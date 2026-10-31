import { get, set } from 'idb-keyval';

export interface SourceFile {
	path: string;
	text: string;
}

export interface Config {
	namespace: string;
	description: string;
}

export interface ProjectData {
	name: string;
	config: Config;
	files: SourceFile[];
	active: string;
}

const storageKey = 'mms-project';
const uriPrefix = 'file:///project/';

export const uriOf = (path: string) => uriPrefix + path.split('/').map(encodeURIComponent).join('/');
export const pathOf = (uri: string) =>
	uri.startsWith(uriPrefix) ? uri.slice(uriPrefix.length).split('/').map(decodeURIComponent).join('/') : uri;

export const starter: ProjectData = {
	name: 'my-pack',
	config: { namespace: 'my_pack', description: 'Built with MMS' },
	active: 'main.mms',
	files: [
		{
			path: 'main.mms',
			text: `// Welcome to MMS. Declarations here become data pack files; the
// preview on the right renders noises and density functions.

Noise hills = {
  base_octave = -6
  amplitude_modifiers = [1, 1, 0.5]
}

DensityFunction terrain = Clamp(Noise(hills, 1, 0.5) * 2 - Abs(Noise(hills, 0.25, 0)), -1, 1)
`
		}
	]
};

/**
 * The open project: its files and config, kept in IndexedDB so a reload
 * picks up where the user left off.
 */
export class Project {
	name = $state(starter.name);
	config = $state<Config>({ ...starter.config });
	files = $state<SourceFile[]>([]);
	active = $state('');
	private saveTimer: ReturnType<typeof setTimeout> | null = null;

	activeFile = $derived(this.files.find((f) => f.path === this.active) ?? null);

	async load(): Promise<void> {
		let data: ProjectData | undefined;
		try {
			data = await get<ProjectData>(storageKey);
		} catch {
			data = undefined;
		}
		this.replace(data && data.files?.length ? data : starter, false);
	}

	data(): ProjectData {
		return {
			name: this.name,
			config: $state.snapshot(this.config),
			files: $state.snapshot(this.files),
			active: this.active
		};
	}

	replace(data: ProjectData, save = true) {
		this.name = data.name;
		this.config = { ...data.config };
		this.files = data.files.map((f) => ({ ...f }));
		this.active = this.files.some((f) => f.path === data.active) ? data.active : (this.files[0]?.path ?? '');
		if (save) this.save();
	}

	setText(path: string, text: string) {
		const f = this.files.find((f) => f.path === path);
		if (f && f.text !== text) {
			f.text = text;
			this.save();
		}
	}

	/** Adds a file and returns its path, made unique. */
	add(path: string, text = ''): string {
		path = normalizePath(path);
		let unique = path;
		for (let i = 2; this.files.some((f) => f.path === unique); i++) unique = path.replace(/(\.mms)?$/, `_${i}.mms`);
		this.files.push({ path: unique, text });
		this.files.sort((a, b) => a.path.localeCompare(b.path));
		this.active = unique;
		this.save();
		return unique;
	}

	rename(from: string, to: string): boolean {
		to = normalizePath(to);
		if (from === to || this.files.some((f) => f.path === to)) return false;
		const f = this.files.find((f) => f.path === from);
		if (!f) return false;
		f.path = to;
		this.files.sort((a, b) => a.path.localeCompare(b.path));
		if (this.active === from) this.active = to;
		this.save();
		return true;
	}

	remove(path: string) {
		this.files = this.files.filter((f) => f.path !== path);
		if (this.active === path) this.active = this.files[0]?.path ?? '';
		this.save();
	}

	save() {
		if (this.saveTimer) clearTimeout(this.saveTimer);
		this.saveTimer = setTimeout(() => {
			set(storageKey, this.data()).catch(() => {
				/* storage can be unavailable (private windows); the session still works */
			});
		}, 300);
	}
}

export function normalizePath(p: string): string {
	p = p
		.trim()
		.replace(/\\/g, '/')
		.replace(/^\/+/, '')
		.replace(/\/{2,}/g, '/');
	if (!p) p = 'untitled';
	if (!p.endsWith('.mms')) p += '.mms';
	return p;
}
