/// <reference lib="webworker" />
import {
	DensityFunction,
	Holder,
	Identifier,
	NoiseGeneratorSettings,
	NoiseRouter,
	RandomState,
	WorldgenRegistries
} from 'deepslate';
import { vanillaBase, type PreviewKind, type PreviewRequest, type PreviewResponse } from './protocol';

// Renders previews with Deepslate. User files are registered under their
// ids; minecraft: ids the pack references but does not define are fetched
// from mcmeta on demand and cached for the session.

const registries = {
	noise: WorldgenRegistries.NOISE,
	density_function: WorldgenRegistries.DENSITY_FUNCTION,
	noise_settings: WorldgenRegistries.NOISE_SETTINGS
} as const;
const folderOf: Record<string, PreviewKind> = {
	'worldgen/noise': 'noise',
	'worldgen/density_function': 'density_function',
	'worldgen/noise_settings': 'noise_settings'
};

let userFiles: { kind: PreviewKind; id: string; json: unknown }[] = [];
const vanilla = new Map<string, unknown | null>(); // "kind|id" → json, null when missing

async function fetchVanilla(kind: PreviewKind, id: Identifier): Promise<boolean> {
	const key = `${kind}|${id}`;
	if (!vanilla.has(key)) {
		let json: unknown | null = null;
		try {
			const res = await fetch(`${vanillaBase}/${id.namespace}/worldgen/${kind}/${id.path}.json`);
			if (res.ok) json = await res.json();
		} catch {
			json = null;
		}
		vanilla.set(key, json);
	}
	return vanilla.get(key) != null;
}

function register(kind: PreviewKind, id: string, json: unknown) {
	const r = registries[kind] as unknown as { register(id: Identifier, v: () => unknown): void; parse(o: unknown): unknown };
	r.register(Identifier.parse(id), () => r.parse(json));
}

function registerAll() {
	for (const r of Object.values(registries)) {
		// Builtin entries (the surface noises) survive; everything else
		// is re-registered from the current files.
		for (const key of r.keys()) {
			if (!(key.equals(Identifier.create('surface')) || key.equals(Identifier.create('surface_secondary')))) r.delete(key);
		}
	}
	for (const [key, json] of vanilla) {
		if (json == null) continue;
		const [kind, id] = key.split('|');
		register(kind as PreviewKind, id, json);
	}
	for (const f of userFiles) {
		register(f.kind, f.id, f.json);
	}
}

/** Runs build, fetching each missing vanilla entry it reports, until it succeeds. */
async function withVanilla<T>(build: () => T): Promise<T> {
	for (let attempt = 0; ; attempt++) {
		registerAll();
		try {
			return build();
		} catch (e) {
			const m = /Missing key in (?:minecraft:)?([a-z_/]+): ([a-z0-9_.-]+:[a-z0-9_./-]+)/.exec(String(e));
			const kind = m && folderOf[m[1]];
			if (!m || !kind || attempt > 60 || !(await fetchVanilla(kind, Identifier.parse(m[2])))) throw e;
		}
	}
}

function seedOf(s: string): bigint {
	try {
		return BigInt(s);
	} catch {
		let h = 0n;
		for (const c of s) h = (h * 31n + BigInt(c.codePointAt(0)!)) & 0xffffffffffffffffn;
		return BigInt.asIntN(64, h);
	}
}

// Diverging palette for density values: blue below zero, orange above.
function diverging(v: number, out: Uint8ClampedArray, i: number) {
	const t = Math.max(-1, Math.min(1, v));
	if (t < 0) {
		const a = -t;
		out[i] = 245 - 200 * a;
		out[i + 1] = 245 - 140 * a;
		out[i + 2] = 245 - 30 * a;
	} else {
		out[i] = 245 - 10 * t;
		out[i + 1] = 245 - 120 * t;
		out[i + 2] = 245 - 200 * t;
	}
	out[i + 3] = 255;
}

async function render(req: Extract<PreviewRequest, { type: 'render' }>): Promise<PreviewResponse> {
	const { width, height, view } = req;
	const seed = seedOf(req.seed);
	const id = Identifier.parse(req.id);
	const values = new Float32Array(width * height);
	const pixels = new Uint8ClampedArray(width * height * 4);
	const x0 = view.x - (width / 2) * view.scale;
	const yTop = view.y + (height / 2) * view.scale;
	let min = Infinity;
	let max = -Infinity;
	const track = (v: number) => {
		if (v < min) min = v;
		if (v > max) max = v;
	};

	if (req.kind === 'noise') {
		// Top-down: x across, z down, at y = 0.
		const noise = await withVanilla(() => {
			const rs = new RandomState(NoiseGeneratorSettings.create({}), seed);
			return rs.getOrCreateNoise(id);
		});
		const z0 = view.y - (height / 2) * view.scale;
		for (let py = 0; py < height; py++) {
			for (let px = 0; px < width; px++) {
				const v = noise.get3D(x0 + px * view.scale, 0, z0 + py * view.scale);
				values[py * width + px] = v;
				track(v);
			}
		}
		for (let i = 0; i < values.length; i++) {
			const g = Math.max(0, Math.min(255, (values[i] + 1) * 127.5));
			pixels[i * 4] = pixels[i * 4 + 1] = pixels[i * 4 + 2] = g;
			pixels[i * 4 + 3] = 255;
		}
		return { type: 'rendered', seq: req.seq, width, height, pixels, values, min, max };
	}

	if (req.kind === 'density_function') {
		// Side view: x across, y up, at z = 0.
		const fn = await withVanilla(() => {
			const rs = new RandomState(NoiseGeneratorSettings.create({}), seed);
			// A reference holder, so the visitor binds noises and caches by id.
			const df = new DensityFunction.HolderHolder(Holder.reference(WorldgenRegistries.DENSITY_FUNCTION, id));
			return df.mapAll(rs.createVisitor(false));
		});
		for (let py = 0; py < height; py++) {
			const y = yTop - py * view.scale;
			for (let px = 0; px < width; px++) {
				const v = fn.compute(DensityFunction.context(x0 + px * view.scale, y, 0));
				values[py * width + px] = v;
				track(v);
			}
		}
		const range = Math.max(Math.abs(min), Math.abs(max), 1e-9);
		for (let i = 0; i < values.length; i++) diverging(values[i] / range, pixels, i * 4);
		return { type: 'rendered', seq: req.seq, width, height, pixels, values, min, max };
	}

	// Noise settings: the terrain's final density, side view at z = 0.
	const { settings, router } = await withVanilla(() => {
		const settings = WorldgenRegistries.NOISE_SETTINGS.getOrThrow(id);
		const rs = new RandomState(settings, seed);
		return { settings, router: rs.router as NoiseRouter };
	});
	const minY = settings.noise.minY;
	const maxY = minY + settings.noise.height;
	const sea = settings.seaLevel;
	for (let py = 0; py < height; py++) {
		const y = yTop - py * view.scale;
		for (let px = 0; px < width; px++) {
			const i = py * width + px;
			const o = i * 4;
			let v = NaN;
			if (y >= minY && y < maxY) {
				v = router.finalDensity.compute(DensityFunction.context(x0 + px * view.scale, y, 0));
				track(v);
			}
			values[i] = v;
			if (Number.isNaN(v)) {
				pixels[o] = 30; pixels[o + 1] = 30; pixels[o + 2] = 36;
			} else if (v > 0) {
				const shade = 120 + Math.min(1, v) * 40;
				pixels[o] = shade; pixels[o + 1] = shade; pixels[o + 2] = shade + 6;
			} else if (y < sea) {
				pixels[o] = 52; pixels[o + 1] = 104; pixels[o + 2] = 196;
			} else {
				pixels[o] = 186; pixels[o + 1] = 218; pixels[o + 2] = 247;
			}
			pixels[o + 3] = 255;
		}
	}
	return { type: 'rendered', seq: req.seq, width, height, pixels, values, min, max, seaLevel: sea, minY, maxY };
}

let latest = 0;
self.onmessage = async (e: MessageEvent<PreviewRequest>) => {
	const req = e.data;
	if (req.type === 'setup') {
		userFiles = req.files;
		return;
	}
	latest = req.seq;
	try {
		const res = await render(req);
		if (res.seq !== latest) return; // superseded while fetching
		if (res.type === 'rendered') {
			(self as unknown as Worker).postMessage(res, [res.pixels.buffer, res.values.buffer]);
		}
	} catch (err) {
		(self as unknown as Worker).postMessage({ type: 'error', seq: req.seq, message: String(err instanceof Error ? err.message : err) } satisfies PreviewResponse);
	}
};
