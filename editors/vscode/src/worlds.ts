// Finding Minecraft worlds: the saves folders of the vanilla launcher and
// of Prism Launcher (and MultiMC) instances, plus any the user adds with
// the mms.savesFolders setting. A world is a folder with a level.dat.
import * as fs from 'fs';
import * as os from 'os';
import * as path from 'path';
import * as zlib from 'zlib';

export interface World {
	/** The world folder. */
	dir: string;
	/** Its name in the game (level.dat's LevelName), else the folder name. */
	name: string;
	/** Where it was found, such as "Minecraft" or "Prism: Fabric 26.3". */
	source: string;
	/** When it was last saved (level.dat's modification time). */
	saved: Date;
}

export interface SavesFolder {
	dir: string;
	label: string;
}

/** The saves folders that exist on this computer, the vanilla one first. */
export function savesFolders(extra: string[] = []): SavesFolder[] {
	const home = os.homedir();
	const out: SavesFolder[] = [];
	const add = (dir: string, label: string) => {
		if (isDir(dir) && !out.some((s) => same(s.dir, dir))) out.push({ dir, label });
	};
	let launchers: string[];
	if (process.platform === 'win32') {
		const appData = process.env.APPDATA ?? path.join(home, 'AppData', 'Roaming');
		add(path.join(appData, '.minecraft', 'saves'), 'Minecraft');
		launchers = [path.join(appData, 'PrismLauncher'), path.join(appData, 'MultiMC')];
	} else if (process.platform === 'darwin') {
		const support = path.join(home, 'Library', 'Application Support');
		add(path.join(support, 'minecraft', 'saves'), 'Minecraft');
		launchers = [path.join(support, 'PrismLauncher'), path.join(support, 'MultiMC')];
	} else {
		const data = process.env.XDG_DATA_HOME || path.join(home, '.local', 'share');
		add(path.join(home, '.minecraft', 'saves'), 'Minecraft');
		// The Flatpak builds keep their files under ~/.var/app.
		add(path.join(home, '.var', 'app', 'com.mojang.Minecraft', '.minecraft', 'saves'), 'Minecraft (Flatpak)');
		launchers = [path.join(data, 'PrismLauncher'), path.join(home, '.var', 'app', 'org.prismlauncher.PrismLauncher', 'data', 'PrismLauncher'), path.join(data, 'multimc')];
	}
	for (const launcher of launchers) {
		const name = /prism/i.test(launcher) ? 'Prism' : 'MultiMC';
		for (const instance of list(path.join(launcher, 'instances'))) {
			// Newer Prism instances use minecraft/, older ones .minecraft/.
			for (const game of ['minecraft', '.minecraft']) {
				add(path.join(launcher, 'instances', instance, game, 'saves'), `${name}: ${instance}`);
			}
		}
	}
	for (const dir of extra) add(dir, dir.replace(home, '~'));
	return out;
}

/** The worlds in the saves folders, most recently saved first. */
export function findWorlds(folders: SavesFolder[]): World[] {
	const worlds: World[] = [];
	for (const f of folders) {
		for (const entry of list(f.dir)) {
			const w = readWorld(path.join(f.dir, entry), f.label);
			if (w) worlds.push(w);
		}
	}
	return worlds.sort((a, b) => b.saved.getTime() - a.saved.getTime());
}

/** The world in a folder, or undefined if it has no level.dat. */
export function readWorld(dir: string, source = ''): World | undefined {
	const level = path.join(dir, 'level.dat');
	let saved: Date;
	try {
		saved = fs.statSync(level).mtime;
	} catch {
		return undefined;
	}
	return { dir, name: levelName(level) ?? path.basename(dir), source, saved };
}

/**
 * The world's name from level.dat, a gzipped NBT file. Rather than parse
 * the whole tree, this finds the LevelName string tag (type 8, name length
 * 9), which only Data holds.
 */
export function levelName(file: string): string | undefined {
	try {
		const nbt = zlib.gunzipSync(fs.readFileSync(file));
		const tag = Buffer.from([8, 0, 9, ...Buffer.from('LevelName')]);
		const at = nbt.indexOf(tag);
		if (at < 0) return undefined;
		const start = at + tag.length;
		const len = nbt.readUInt16BE(start);
		const name = nbt.subarray(start + 2, start + 2 + len).toString('utf8').trim();
		return name || undefined;
	} catch {
		return undefined;
	}
}

function list(dir: string): string[] {
	try {
		return fs
			.readdirSync(dir, { withFileTypes: true })
			.filter((e) => e.isDirectory())
			.map((e) => e.name);
	} catch {
		return [];
	}
}

function isDir(p: string): boolean {
	try {
		return fs.statSync(p).isDirectory();
	} catch {
		return false;
	}
}

function same(a: string, b: string): boolean {
	return process.platform === 'win32' ? a.toLowerCase() === b.toLowerCase() : a === b;
}
