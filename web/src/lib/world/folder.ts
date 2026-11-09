// A Minecraft world folder on the user's computer, as the browser lets the
// page reach it. Chrome and Edge give a folder handle the page can write
// to and keep for the next visit (the File System Access API); other
// browsers only let the page read the files of a folder picked with
// <input webkitdirectory>, as they were when picked. Building into a
// world needs the first; listing and importing structures works with
// either.
import { del, get, set } from 'idb-keyval';

export interface WorldFile {
	/** Path relative to the world folder, with / separators. */
	path: string;
	file: File;
}

export interface WorldFolder {
	/** The folder's name on disk. */
	readonly name: string;
	/** Whether the page can write to it (a folder handle). */
	readonly writable: boolean;
	readonly handle?: FileSystemDirectoryHandle;
	/** The file at path, or undefined when there is none. */
	file(path: string): Promise<File | undefined>;
	/** The files under the folder at path (all of them for ''), with paths relative to the world. */
	walk(path: string): Promise<WorldFile[]>;
	/** Whether a folder exists at path. */
	isDir(path: string): Promise<boolean>;
	/** Writes a file, making its folders; a writable folder only. */
	write(path: string, data: Uint8Array): Promise<void>;
	/** Deletes a file or folder (with what it holds); nothing when there is none. */
	remove(path: string): Promise<void>;
}

const parts = (path: string) => path.split('/').filter(Boolean);
const notFound = (e: unknown) => e instanceof DOMException && (e.name === 'NotFoundError' || e.name === 'TypeMismatchError');

/** Whether this browser can pick a folder to write to. */
export const canWrite = () => typeof window !== 'undefined' && 'showDirectoryPicker' in window;

type PermissionHandle = FileSystemDirectoryHandle & {
	queryPermission(o: { mode: 'readwrite' | 'read' }): Promise<PermissionState>;
	requestPermission(o: { mode: 'readwrite' | 'read' }): Promise<PermissionState>;
	values(): AsyncIterable<FileSystemDirectoryHandle | FileSystemFileHandle>;
};

/** A world folder through a folder handle, which can be written to. */
export class HandleFolder implements WorldFolder {
	readonly writable = true;
	constructor(readonly handle: FileSystemDirectoryHandle) {}

	get name() {
		return this.handle.name;
	}

	private async dir(path: string[], create = false): Promise<FileSystemDirectoryHandle | undefined> {
		let d = this.handle;
		try {
			for (const p of path) d = await d.getDirectoryHandle(p, { create });
		} catch (e) {
			if (notFound(e)) return undefined;
			throw e;
		}
		return d;
	}

	async file(path: string): Promise<File | undefined> {
		const p = parts(path);
		const d = await this.dir(p.slice(0, -1));
		if (!d) return undefined;
		try {
			return await (await d.getFileHandle(p[p.length - 1])).getFile();
		} catch (e) {
			if (notFound(e)) return undefined;
			throw e;
		}
	}

	async isDir(path: string): Promise<boolean> {
		return !!(await this.dir(parts(path)));
	}

	async walk(path: string): Promise<WorldFile[]> {
		const out: WorldFile[] = [];
		const start = await this.dir(parts(path));
		if (!start) return out;
		const visit = async (d: FileSystemDirectoryHandle, at: string) => {
			for await (const e of (d as PermissionHandle).values()) {
				const rel = at ? `${at}/${e.name}` : e.name;
				if (e.kind === 'directory') await visit(e as FileSystemDirectoryHandle, rel);
				else out.push({ path: rel, file: await (e as FileSystemFileHandle).getFile() });
			}
		};
		await visit(start, parts(path).join('/'));
		return out;
	}

	async write(path: string, data: Uint8Array): Promise<void> {
		const p = parts(path);
		const d = (await this.dir(p.slice(0, -1), true))!;
		const w = await (await d.getFileHandle(p[p.length - 1], { create: true })).createWritable();
		// The browser writes to a copy and swaps it in on close, so a
		// failed write never leaves half a file.
		await w.write(data as Uint8Array<ArrayBuffer>);
		await w.close();
	}

	async remove(path: string): Promise<void> {
		const p = parts(path);
		const d = await this.dir(p.slice(0, -1));
		if (!d) return;
		try {
			await d.removeEntry(p[p.length - 1], { recursive: true });
		} catch (e) {
			if (!notFound(e)) throw e;
		}
	}

	/** Whether the page may write to the folder now, asking when it may ask (a click). */
	async allowed(ask: boolean): Promise<boolean> {
		const h = this.handle as PermissionHandle;
		if (!h.queryPermission) return true;
		if ((await h.queryPermission({ mode: 'readwrite' })) === 'granted') return true;
		return ask && (await h.requestPermission({ mode: 'readwrite' })) === 'granted';
	}
}

/** A world folder picked with <input webkitdirectory>: its files as they were then, read only. */
export class FilesFolder implements WorldFolder {
	readonly writable = false;
	private files = new Map<string, File>();
	readonly name: string;

	constructor(list: FileList | File[]) {
		const all = [...list];
		// Each file's path starts with the picked folder's name.
		this.name = all[0]?.webkitRelativePath.split('/')[0] ?? 'world';
		for (const f of all) this.files.set(f.webkitRelativePath.split('/').slice(1).join('/'), f);
	}

	async file(path: string) {
		return this.files.get(parts(path).join('/'));
	}

	async isDir(path: string) {
		const prefix = parts(path).join('/') + '/';
		return [...this.files.keys()].some((k) => k.startsWith(prefix));
	}

	async walk(path: string) {
		const prefix = parts(path).length ? parts(path).join('/') + '/' : '';
		return [...this.files].filter(([k]) => k.startsWith(prefix)).map(([path, file]) => ({ path, file }));
	}

	async write(): Promise<void> {
		throw new Error('this browser cannot write to a folder');
	}

	async remove(): Promise<void> {
		throw new Error('this browser cannot write to a folder');
	}
}

/** Asks the user for a folder the page can write to; undefined when they cancel. */
export async function pickFolder(id: string): Promise<HandleFolder | undefined> {
	try {
		const pick = (window as unknown as { showDirectoryPicker(o: object): Promise<FileSystemDirectoryHandle> }).showDirectoryPicker;
		return new HandleFolder(await pick({ id, mode: 'readwrite' }));
	} catch (e) {
		if (e instanceof DOMException && e.name === 'AbortError') return undefined;
		throw e;
	}
}

/** Whether a folder is a world: it has a level.dat. */
export async function isWorld(f: WorldFolder): Promise<boolean> {
	return !!(await f.file('level.dat'));
}

/**
 * The world's name in the game, from level.dat, a gzipped NBT file. Rather
 * than parse the whole tree, this finds the LevelName string tag (type 8,
 * name length 9), which only Data holds.
 */
export async function levelName(f: WorldFolder): Promise<string | undefined> {
	try {
		const file = await f.file('level.dat');
		if (!file) return undefined;
		const nbt = new Uint8Array(await new Response(file.stream().pipeThrough(new DecompressionStream('gzip'))).arrayBuffer());
		const tag = [8, 0, 9, ...new TextEncoder().encode('LevelName')];
		const at = indexOf(nbt, tag);
		if (at < 0) return undefined;
		const start = at + tag.length;
		const len = (nbt[start] << 8) | nbt[start + 1];
		return new TextDecoder().decode(nbt.subarray(start + 2, start + 2 + len)).trim() || undefined;
	} catch {
		return undefined;
	}
}

function indexOf(hay: Uint8Array, needle: number[]): number {
	outer: for (let i = 0; i + needle.length <= hay.length; i++) {
		for (let j = 0; j < needle.length; j++) if (hay[i + j] !== needle[j]) continue outer;
		return i;
	}
	return -1;
}

// A project's worlds are kept by project and purpose, so the next visit
// builds into the same one. Handles keep in IndexedDB; the browser asks
// again for permission to write on the first build of a visit.
const key = (project: string, purpose: 'build' | 'structures') => `mms-world:${project}:${purpose}`;

export async function savedFolder(project: string, purpose: 'build' | 'structures'): Promise<HandleFolder | undefined> {
	try {
		const h = await get<FileSystemDirectoryHandle>(key(project, purpose));
		return h ? new HandleFolder(h) : undefined;
	} catch {
		return undefined;
	}
}

export async function saveFolder(project: string, purpose: 'build' | 'structures', f: WorldFolder | undefined): Promise<void> {
	try {
		if (f?.handle) await set(key(project, purpose), f.handle);
		else await del(key(project, purpose));
	} catch {
		/* storage unavailable: the world lasts for this visit */
	}
}
