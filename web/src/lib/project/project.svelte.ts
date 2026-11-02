import * as store from './store';

export interface SourceFile {
	path: string;
	text: string;
}

/** A binary project file, such as a texture or sound, as base64. */
export interface AssetFile {
	path: string;
	data: string;
}

/** The binary files asset helpers read: Texture, Sound and Structure. */
export const assetExtensions = ['.png', '.ogg', '.nbt'];
export const isAssetPath = (p: string) => assetExtensions.some((e) => p.toLowerCase().endsWith(e));

export interface Config {
	namespace: string;
	description: string;
}

export interface ProjectData {
	name: string;
	config: Config;
	files: SourceFile[];
	/** Textures, sounds and structures; older saves have none. */
	assets?: AssetFile[];
	active: string;
}

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
 * The open project: its files and config. It is one of the projects saved
 * in IndexedDB (see store.ts), so a reload picks up where the user left off.
 */
export class Project {
	/** The saved project's key; empty until load(). */
	id = $state('');
	name = $state(starter.name);
	config = $state<Config>({ ...starter.config });
	files = $state<SourceFile[]>([]);
	assets = $state<AssetFile[]>([]);
	active = $state('');
	/** Every saved project, most recent first, for the switcher. */
	saved = $state<store.ProjectMeta[]>([]);
	private saveTimer: ReturnType<typeof setTimeout> | null = null;

	activeFile = $derived(this.files.find((f) => f.path === this.active) ?? null);

	/** Opens this tab's project, or the last one, or a new starter project. */
	async load(): Promise<void> {
		try {
			await this.refresh();
			const id = await store.currentId();
			const data = id ? await store.loadProject(id) : undefined;
			if (id && data?.files?.length) {
				this.show(id, data);
				await store.setCurrentId(id);
				return;
			}
			const last = this.saved[0];
			const lastData = last ? await store.loadProject(last.id) : undefined;
			if (last && lastData) {
				this.show(last.id, lastData);
				await store.setCurrentId(last.id);
				return;
			}
		} catch {
			/* storage can be unavailable (private windows); the session still works */
		}
		await this.create(starter);
	}

	/** Re-reads the list of saved projects. */
	async refresh(): Promise<void> {
		try {
			this.saved = await store.listProjects();
		} catch {
			/* storage unavailable */
		}
	}

	data(): ProjectData {
		return {
			name: this.name,
			config: $state.snapshot(this.config),
			files: $state.snapshot(this.files),
			assets: $state.snapshot(this.assets),
			active: this.active
		};
	}

	private show(id: string, data: ProjectData) {
		this.id = id;
		this.name = data.name;
		this.config = { ...data.config };
		this.files = data.files.map((f) => ({ ...f }));
		this.assets = (data.assets ?? []).map((f) => ({ ...f }));
		this.active = this.files.some((f) => f.path === data.active) ? data.active : (this.files[0]?.path ?? '');
	}

	/**
	 * Saves the open project, then opens `data` as a new saved project. Its
	 * name is made unique among the saved ones.
	 */
	async create(data: ProjectData): Promise<void> {
		await this.flush();
		const name = store.uniqueName(data.name || 'project', this.saved.map((m) => m.name));
		this.show(store.newId(), { ...data, name });
		try {
			await store.saveProject(this.id, this.data());
			await store.setCurrentId(this.id);
		} catch {
			/* storage unavailable */
		}
		await this.refresh();
	}

	/** Saves the open project and opens a saved one. */
	async open(id: string): Promise<boolean> {
		if (id === this.id) return true;
		await this.flush();
		const data = await store.loadProject(id);
		if (!data) {
			await this.refresh();
			return false;
		}
		this.show(id, data);
		await store.setCurrentId(id);
		await this.refresh();
		return true;
	}

	/** Renames a saved project, or the open one. */
	async renameProject(id: string, name: string): Promise<void> {
		name = name.trim();
		if (!name) return;
		if (id === this.id) {
			this.name = name;
			await this.flush(true);
		} else {
			const data = await store.loadProject(id);
			if (data) await store.saveProject(id, { ...data, name });
		}
		await this.refresh();
	}

	/** Copies a saved project (or the open one) and opens the copy. */
	async duplicate(id: string): Promise<void> {
		await this.flush();
		const data = id === this.id ? this.data() : await store.loadProject(id);
		if (data) await this.create(structuredClone(data));
	}

	/** Deletes a saved project; deleting the open one opens the next. */
	async deleteProject(id: string): Promise<void> {
		if (id === this.id && this.saveTimer) {
			clearTimeout(this.saveTimer);
			this.saveTimer = null;
		}
		await store.deleteProject(id);
		await this.refresh();
		if (id !== this.id) return;
		const next = this.saved[0];
		if (!next || !(await this.open(next.id))) await this.create(starter);
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

	/** Adds or replaces a binary file. */
	setAsset(path: string, data: string) {
		path = path.trim().replace(/\\/g, '/').replace(/^\/+/, '');
		const f = this.assets.find((f) => f.path === path);
		if (f) f.data = data;
		else {
			this.assets.push({ path, data });
			this.assets.sort((a, b) => a.path.localeCompare(b.path));
		}
		this.save();
	}

	removeAsset(path: string) {
		this.assets = this.assets.filter((f) => f.path !== path);
		this.save();
	}

	/** Saves soon; edits come in bursts while typing. */
	save() {
		if (this.saveTimer) clearTimeout(this.saveTimer);
		this.saveTimer = setTimeout(() => void this.flush(true), 300);
	}

	/** Writes a pending save now (or always, with `force`). */
	async flush(force = false): Promise<void> {
		const pending = this.saveTimer !== null;
		if (this.saveTimer) clearTimeout(this.saveTimer);
		this.saveTimer = null;
		if (!this.id || (!pending && !force)) return;
		try {
			await store.saveProject(this.id, this.data());
		} catch {
			/* storage can be unavailable (private windows); the session still works */
		}
		// Keep the switcher's order and names current without re-reading.
		const meta = { id: this.id, name: this.name, updated: Date.now(), files: this.files.length };
		this.saved = [meta, ...this.saved.filter((m) => m.id !== this.id)];
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
