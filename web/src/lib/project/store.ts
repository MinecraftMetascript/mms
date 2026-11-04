import { del, get, set, update } from 'idb-keyval';
import type { ProjectData } from './project.svelte';

/**
 * Saved projects, in IndexedDB. Each project is stored under its own key so
 * a save writes one project, and an index lists them for the switcher.
 *
 * IndexedDB rather than localStorage: localStorage holds about 5 MB of
 * strings per origin, shared by every project, while IndexedDB gets a share
 * of the disk (hundreds of MB or more) and takes binary assets. Both are
 * cleared together when the browser evicts the origin, so persistence comes
 * from navigator.storage.persist(), which we ask for once projects exist.
 */

export interface ProjectMeta {
	id: string;
	name: string;
	/** Milliseconds since the epoch of the last save. */
	updated: number;
	files: number;
}

const indexKey = 'mms-projects';
const lastKey = 'mms-current';
const legacyKey = 'mms-project';
const dataKey = (id: string) => `mms-project:${id}`;
// This tab's project, so a reload reopens it even if another tab switched.
const tabKey = 'mms-tab-project';

// Other tabs hear about saves and deletes, to refresh their project list.
const channel = typeof BroadcastChannel !== 'undefined' ? new BroadcastChannel('mms-projects') : null;
export function onProjectsChanged(fn: () => void): () => void {
	if (!channel) return () => {};
	const h = () => fn();
	channel.addEventListener('message', h);
	return () => channel.removeEventListener('message', h);
}
const changed = () => channel?.postMessage('changed');

export const newId = () =>
	typeof crypto !== 'undefined' && 'randomUUID' in crypto
		? crypto.randomUUID()
		: Date.now().toString(36) + Math.random().toString(36).slice(2);

/** All saved projects, most recently saved first. Moves a pre-index save over. */
export async function listProjects(): Promise<ProjectMeta[]> {
	let list = await get<ProjectMeta[]>(indexKey);
	if (!list) {
		list = [];
		const legacy = await get<ProjectData>(legacyKey);
		if (legacy?.files?.length) {
			const id = newId();
			await saveProject(id, legacy);
			await set(lastKey, id);
			await del(legacyKey);
			list = (await get<ProjectMeta[]>(indexKey)) ?? [];
		}
	}
	return [...list].sort((a, b) => b.updated - a.updated);
}

export const loadProject = (id: string) => get<ProjectData>(dataKey(id));

export async function saveProject(id: string, data: ProjectData): Promise<void> {
	await set(dataKey(id), data);
	const meta: ProjectMeta = { id, name: data.name, updated: Date.now(), files: data.files.length };
	await update<ProjectMeta[]>(indexKey, (list = []) => [...list.filter((m) => m.id !== id), meta]);
	changed();
}

export async function deleteProject(id: string): Promise<void> {
	await del(dataKey(id));
	await update<ProjectMeta[]>(indexKey, (list = []) => list.filter((m) => m.id !== id));
	changed();
}

/** The project this tab had open, else the last one opened anywhere. */
export async function currentId(): Promise<string | undefined> {
	try {
		const id = sessionStorage.getItem(tabKey);
		if (id) return id;
	} catch {
		/* storage unavailable */
	}
	return get<string>(lastKey);
}

export async function setCurrentId(id: string): Promise<void> {
	try {
		sessionStorage.setItem(tabKey, id);
	} catch {
		/* storage unavailable */
	}
	await set(lastKey, id);
}

/** A name not used by another project: "badlands", then "badlands 2". */
export function uniqueName(name: string, taken: Iterable<string>): string {
	const used = new Set(taken);
	if (!used.has(name)) return name;
	const base = name.replace(/ \d+$/, '');
	for (let i = 2; ; i++) if (!used.has(`${base} ${i}`)) return `${base} ${i}`;
}

export interface StorageState {
	/** True when the browser will not evict the projects to free space. */
	persisted: boolean;
	usage?: number;
	quota?: number;
}

/** Asks the browser to keep this site's storage; returns what it reports. */
export async function requestPersistence(ask = true): Promise<StorageState> {
	const s = typeof navigator !== 'undefined' ? navigator.storage : undefined;
	if (!s) return { persisted: false };
	let persisted = false;
	try {
		persisted = (await s.persisted?.()) ?? false;
		if (!persisted && ask) persisted = (await s.persist?.()) ?? false;
	} catch {
		/* unsupported */
	}
	let usage: number | undefined;
	let quota: number | undefined;
	try {
		({ usage, quota } = (await s.estimate?.()) ?? {});
	} catch {
		/* unsupported */
	}
	return { persisted, usage, quota };
}

