// Vanilla item icons for the inspectors, drawn the way the game draws an
// inventory slot: Deepslate's ItemRenderer follows the item's model
// definition and models (a block item shows as a lit 3D block) over the
// game's texture atlas. The atlas, models and item model definitions come
// from mcmeta's 26.3 tags, pinned, and load once on the first icon (about
// 3.3 MB, cached by the browser). One small WebGL canvas draws every icon,
// which is kept as a data URL.
import {
	BlockModel,
	Identifier,
	ItemModel,
	ItemRenderer,
	ItemStack,
	NbtString,
	type ItemRendererResources,
	type NbtTag
} from 'deepslate';
import { loadImage, loadJson } from '../book/assets';

const mcmeta = 'https://raw.githubusercontent.com/misode/mcmeta';
/** mcmeta's 26.3-atlas tag: every texture in one image, and where each sits. */
const atlasBase = `${mcmeta}/41de69bb755237ac62cfc4a076fe17aafd41fb07/all`;
/** mcmeta's 26.3-summary tag: block and item models, item model definitions. */
const summaryBase = `${mcmeta}/d96c75fec200c4580dd75e033a76341521165461/assets`;

/** The icon's size in pixels: four per texture pixel, for a crisp 2× tile. */
const size = 64;

type Json = Record<string, unknown>;

let resources: Promise<ItemRendererResources> | null = null;

function loadResources(): Promise<ItemRendererResources> {
	if (!resources) {
		resources = (async () => {
			const [atlas, uvs, models, items] = await Promise.all([
				loadImage(`${atlasBase}/atlas.png`),
				loadJson<Record<string, [number, number, number, number]>>(`${atlasBase}/data.min.json`),
				loadJson<Record<string, Json>>(`${summaryBase}/model/data.min.json`),
				loadJson<Record<string, Json>>(`${summaryBase}/item_definition/data.min.json`)
			]);
			if (!atlas || !uvs || !models || !items) throw new Error('missing vanilla item assets');
			const w = atlas.width;
			const h = atlas.height;
			const blockModels = new Map<string, BlockModel | null>();
			const itemModels = new Map<string, ItemModel | null>();
			const res: ItemRendererResources = {
				getTextureAtlas: () => atlas,
				getPixelSize: () => 1 / w,
				getTextureUV(id) {
					const uv = uvs[id.path] ?? uvs['block/missingno'];
					if (!uv) return [0, 0, 16 / w, 16 / h];
					// An animated texture is a strip of frames: use the first.
					const [x, y, u] = uv;
					return [x / w, y / h, (x + u) / w, (y + u) / h];
				},
				getBlockModel(id) {
					const key = id.namespace === 'minecraft' ? id.path : id.toString();
					if (!blockModels.has(key)) {
						const json = models[key];
						blockModels.set(key, null);
						if (json) {
							const m = BlockModel.fromJson(json);
							m.flatten(res);
							blockModels.set(key, m);
						}
					}
					return blockModels.get(key) ?? null;
				},
				getItemModel(id) {
					const key = id.namespace === 'minecraft' ? id.path : id.toString();
					if (!itemModels.has(key)) {
						const json = items[key];
						itemModels.set(key, json?.model ? ItemModel.fromJson(json.model) : null);
					}
					return itemModels.get(key) ?? null;
				},
				// Only the default item_model matters for drawing an icon.
				getItemComponents: (id) => new Map<string, NbtTag>([['minecraft:item_model', new NbtString(id.toString())]])
			};
			return res;
		})();
		resources.catch(() => (resources = null));
	}
	return resources;
}

let renderer: { canvas: HTMLCanvasElement; r: ItemRenderer } | null = null;
const icons = new Map<string, Promise<string | null>>();

/**
 * A vanilla item's icon as a PNG data URL, or null when it has no model
 * (or WebGL or the assets are unavailable). `id` is an item id such as
 * minecraft:stone; `model` an item_model to draw instead.
 */
export function vanillaIcon(id: string, model?: string): Promise<string | null> {
	const key = model ?? id;
	let p = icons.get(key);
	if (!p) {
		// A failed load (offline, say) is tried again by the next icon.
		p = draw(Identifier.parse(id), model).catch(() => {
			icons.delete(key);
			return null;
		});
		icons.set(key, p);
	}
	return p;
}

async function draw(id: Identifier, model?: string): Promise<string | null> {
	if (typeof document === 'undefined') return null;
	const res = await loadResources();
	const components = new Map<string, NbtTag>();
	if (model) components.set('minecraft:item_model', new NbtString(Identifier.parse(model).toString()));
	if (!res.getItemModel(model ? Identifier.parse(model) : id)) return null;
	const stack = new ItemStack(id, 1, components);
	if (!renderer) {
		const canvas = document.createElement('canvas');
		canvas.width = canvas.height = size;
		// WebGL 2, because the atlas is not a power of two in size and
		// Deepslate mipmaps it.
		const gl = canvas.getContext('webgl2', { preserveDrawingBuffer: true, premultipliedAlpha: false });
		if (!gl) return null;
		renderer = { canvas, r: new ItemRenderer(gl as unknown as WebGLRenderingContext, stack, res) };
	} else renderer.r.setItem(stack);
	const { canvas, r } = renderer;
	const gl = canvas.getContext('webgl2')!;
	gl.clearColor(0, 0, 0, 0);
	gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
	r.setViewport(0, 0, size, size);
	r.drawItem();
	return canvas.toDataURL('image/png');
}
