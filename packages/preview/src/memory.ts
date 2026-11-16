// What each preview remembers per declaration: the mode, camera, cuts,
// threshold and so on, so going back to a file shows it as it was left.
// The host decides where it lives: the web editor keeps it in this
// browser for each project, VS Code in the preview panel's own state.
// The previews reach it through PreviewPane's context.

import { getContext, setContext } from 'svelte';

/** Settings by key, such as `preview:<output path>`. Values are plain JSON. */
export interface PreviewMemory {
	get<T>(key: string): T | undefined;
	set(key: string, value: unknown): void;
	/** Forgets everything, and doesn't save again. */
	clear(): void;
}

/** The most declarations remembered; the ones left longest ago go first. */
export const MEMORY_LIMIT = 200;

type Entries = [string, unknown][];

/**
 * A memory kept in a Map in least recently set order, which hands its
 * entries to save at most every `wait` ms (a camera being dragged changes
 * them every frame) and once more when the page goes away.
 */
export function memory(initial: Entries, save: (entries: Entries) => void, wait = 400): PreviewMemory {
	const map = new Map<string, unknown>();
	for (const e of Array.isArray(initial) ? initial : []) if (Array.isArray(e) && typeof e[0] === 'string') map.set(e[0], e[1]);
	let timer: ReturnType<typeof setTimeout> | undefined;
	const flush = () => {
		clearTimeout(timer);
		timer = undefined;
		try {
			save([...map]);
		} catch {
			// Not remembered past this page; still remembered while it's open.
		}
	};
	let cleared = false;
	const onHide = () => timer !== undefined && flush();
	globalThis.addEventListener?.('pagehide', onHide);
	return {
		clear() {
			cleared = true;
			clearTimeout(timer);
			timer = undefined;
			map.clear();
			globalThis.removeEventListener?.('pagehide', onHide);
		},
		get<T>(key: string) {
			return map.get(key) as T | undefined;
		},
		set(key, value) {
			if (cleared) return;
			const json = JSON.stringify(value);
			if (json === undefined) return;
			const copy = JSON.parse(json);
			if (map.has(key) && JSON.stringify(map.get(key)) === json) return;
			map.delete(key);
			map.set(key, copy);
			while (map.size > MEMORY_LIMIT) map.delete(map.keys().next().value!);
			if (timer === undefined) timer = setTimeout(flush, wait);
		}
	};
}

const stored = new Map<string, PreviewMemory>();
/**
 * The memory kept in this browser's localStorage under `key`. There is one
 * per key for the page's life, so going back to a project finds the
 * settings it hasn't saved yet, and its pagehide listener isn't added twice.
 */
export function storedMemory(key: string): PreviewMemory {
	let m = stored.get(key);
	if (m) return m;
	let initial: Entries = [];
	try {
		initial = JSON.parse(localStorage.getItem(key) ?? '[]');
	} catch {
		// Storage can be unavailable (private windows, some webviews) or hold something else.
	}
	m = memory(initial, (entries) => localStorage.setItem(key, JSON.stringify(entries)));
	stored.set(key, m);
	return m;
}

/** Forgets the memory under `key`, here and in localStorage (its project was deleted). */
export function forgetStoredMemory(key: string) {
	stored.get(key)?.clear();
	stored.delete(key);
	try {
		localStorage.removeItem(key);
	} catch {
		// Nothing was remembered.
	}
}

const contextKey = Symbol('preview memory');

/** Gives the previews under this component a memory; `current` is read each time, so the host can switch it (another project). */
export function provideMemory(source: { readonly current: PreviewMemory }) {
	setContext(contextKey, source);
}

/**
 * The memory PreviewPane provides. A preview used on its own remembers in
 * this browser, under a shared key.
 */
export function useMemory(): { readonly current: PreviewMemory } {
	const source = getContext<{ readonly current: PreviewMemory } | undefined>(contextKey);
	if (source) return source;
	return { current: storedMemory('mms.preview.memory') };
}

/**
 * `base` with what `saved` has for each of its keys, where it has the same
 * type (objects field by field). A setting remembered by an older version
 * that no longer fits keeps the base's value. A key whose base value is
 * null takes any saved value but undefined.
 */
export function restore<T extends object>(base: T, saved: unknown): T {
	if (!saved || typeof saved !== 'object' || Array.isArray(saved)) return base;
	const s = saved as Record<string, unknown>;
	const out = { ...base } as Record<string, unknown>;
	for (const [k, b] of Object.entries(base)) {
		const v = s[k];
		if (v === undefined) continue;
		if (b === null) out[k] = v;
		else if (Array.isArray(b)) {
			if (Array.isArray(v)) out[k] = v;
		} else if (typeof b === 'object') out[k] = restore(b, v);
		else if (typeof v === typeof b && (typeof v !== 'number' || Number.isFinite(v))) out[k] = v;
	}
	return out as T;
}

/** `value` if it's a plain object, else undefined. */
export function record(value: unknown): Record<string, unknown> | undefined {
	return value && typeof value === 'object' && !Array.isArray(value) ? (value as Record<string, unknown>) : undefined;
}
