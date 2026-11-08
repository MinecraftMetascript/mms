import { readFileSync } from 'node:fs';
import { expect, test, type Page } from '@playwright/test';

const status = (page: Page) => page.getByRole('status');

// clickLine clicks the editor line containing text, scrolling down to it:
// CodeMirror renders only the lines near the viewport.
async function clickLine(page: Page, text: string) {
	const line = page.locator('.cm-line', { hasText: text });
	for (let i = 0; i < 50 && (await line.count()) === 0; i++) {
		await page.locator('.cm-scroller').first().evaluate((el) => el.scrollBy(0, 200));
		await page.waitForTimeout(100);
	}
	await line.first().click();
}

async function replaceText(page: Page, text: string) {
	await page.locator('.cm-content').click();
	await page.keyboard.press('ControlOrMeta+a');
	await page.keyboard.press('Delete');
	await page.keyboard.insertText(text);
}

test('the editor loads, checks, completes and previews', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });

	// Diagnostics appear for a typo and clear when it is fixed.
	await replaceText(page, 'DensityFunction x = Clmp(1, 0, 1)\n');
	await expect(status(page)).toHaveText(/1 error/);
	await expect(page.locator('.cm-lintRange-error')).toHaveCount(1);

	// Completion offers the helper.
	await replaceText(page, 'DensityFunction x = Cl');
	await page.keyboard.press('ControlOrMeta+Space');
	await expect(page.locator('.cm-tooltip-autocomplete li[aria-selected]')).toContainText('Clamp');
	await page.waitForTimeout(300);
	await page.keyboard.press('Enter');
	await page.keyboard.insertText('1, 0, 1');
	await expect(page.locator('.cm-content')).toHaveText('DensityFunction x = Clamp(1, 0, 1)');
	await expect(status(page)).toHaveText(/No problems/);

	// A noise renders in the preview, and its JSON is shown.
	await replaceText(page, 'Noise n { base_octave = -5, amplitude_modifiers = [1, 1] }\n');
	await page.locator('.cm-line').first().click();
	await expect(page.getByText('my_pack:n')).toBeVisible();
	await expect(page.locator('canvas')).toBeVisible();
	await expect(page.locator('text=/-?\\d\\.\\d{3} … -?\\d\\.\\d{3}/')).toBeVisible({ timeout: 20_000 });
	await page.getByRole('tab', { name: 'JSON' }).click();
	await expect(page.locator('pre.json')).toContainText('"base_octave": -5');

	expect(errors).toEqual([]);
});

test('projects persist and examples open', async ({ page }) => {
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /surface-rules/ }).click();
	await expect(page.getByRole('button', { name: 'surface.mms' })).toBeVisible();
	await expect(status(page)).toHaveText(/No problems/);
	await page.getByRole('tab', { name: 'Pack' }).click();
	await expect(page.getByText('my_strange_surface.json')).toBeVisible();

	await page.waitForTimeout(800); // saving is debounced
	await page.reload();
	await expect(page.getByRole('button', { name: 'surface.mms' })).toBeVisible();
});

test('an example with several files in folders builds as one project', async ({ page }) => {
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /treasure/ }).click();
	// recipes/crafting.mms uses the #gems tag declared in tags.mms.
	await page.getByRole('button', { name: 'recipes/crafting.mms' }).click();
	await expect(page.locator('.cm-content')).toContainText('Recipe gem_block');
	await expect(status(page)).toHaveText(/No problems/);
	await page.getByRole('tab', { name: 'Pack' }).click();
	await expect(page.getByText('gem_block.json')).toBeVisible();
	await expect(page.getByText('buried.json')).toBeVisible();
});

test('an example with textures and a sound checks and exports both packs', async ({ page }) => {
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /embers/ }).click();
	await expect(page.getByText('art/ember_shard.png', { exact: true })).toBeVisible();
	await expect(page.getByText('sounds/embers.ogg', { exact: true })).toBeVisible();
	await expect(status(page)).toHaveText(/No problems/);

	const { unzipSync } = await import('fflate');
	const { readFileSync } = await import('node:fs');
	const zipped = async (button: string) => {
		const download = page.waitForEvent('download');
		await page.getByRole('button', { name: button }).click();
		const d = await download;
		return { name: d.suggestedFilename(), files: Object.keys(unzipSync(new Uint8Array(readFileSync(await d.path())))) };
	};
	// The resource pack downloads on its own, with the copied files.
	const res = await zipped('Download resources');
	expect(res.name).toBe('embers-resources.zip');
	expect(res.files).toContain('assets/embers/textures/item/ember_shard.png');
	expect(res.files).toContain('assets/embers/textures/painting/hearth.png');
	expect(res.files).toContain('assets/embers/sounds/embers.ogg');
	expect(res.files).toContain('assets/embers/sounds.json');
	expect(res.files).toContain('pack.mcmeta');
	const data = await zipped('Download pack');
	expect(data.name).toBe('embers.zip');
	expect(data.files).toContain('data/embers/painting_variant/hearth.json');
	expect(data.files.some((f) => f.startsWith('assets/'))).toBe(false);

	// Typing a texture path lists the project's images.
	await page.locator('.cm-content').click();
	await page.keyboard.press('ControlOrMeta+End');
	await page.keyboard.press('Enter');
	await page.keyboard.insertText('PaintingVariant extra { width = 1, height = 1, asset_id = Texture(');
	await page.keyboard.type('"art/h');
	const popup = page.locator('.cm-tooltip-autocomplete');
	await expect(popup).toContainText('art/hearth.png');
	await expect(popup).not.toContainText('embers.ogg');
	await page.waitForTimeout(300);
	await page.keyboard.press('Enter');
	// The closing quote was typed for us; the path goes inside it.
	await expect(page.locator('.cm-content')).toContainText('Texture("art/hearth.png"');
	await page.keyboard.press('End');
	await page.keyboard.insertText(') }');
	await expect(status(page)).toHaveText(/No problems/);

	// Removing a texture breaks the code that names it, in both paintings.
	await page.getByRole('listitem').filter({ hasText: 'art/hearth.png' }).hover();
	await page.getByRole('button', { name: 'Delete art/hearth.png' }).click();
	await expect(status(page)).toHaveText(/2 errors/);
});

test('density previews use the picked palette over the range in view', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });

	// A gradient from 0 to 1 spans the whole palette, not a sliver of a -1 to 1 one.
	await replaceText(page, 'DensityFunction grad = Gradient(axis = y, from_coordinate = 0, to_coordinate = 512, from_value = 0, to_value = 1)\n');
	await expect(status(page)).toHaveText(/No problems/);
	await page.locator('.cm-line').first().click();
	const scale = page.getByLabel('Value scale');
	await expect(scale).toContainText(/in view 0\.000 … 0\.\d{3}/, { timeout: 20_000 });
	await expect(scale).toContainText(/^\s*0\.000\s+0\.\d{3}/);

	const pixel = (x: number, y: number) =>
		page.locator('canvas').evaluate((c: HTMLCanvasElement, [fx, fy]) => {
			const d = c.getContext('2d')!.getImageData(Math.floor(c.width * fx), Math.floor(c.height * fy), 1, 1).data;
			return [d[0], d[1], d[2]];
		}, [x, y]);
	// Viridis: dark purple at the bottom (0), yellow at the top.
	await expect(page.getByLabel('Palette')).toHaveValue('viridis');
	expect((await pixel(0.5, 0.98))[2]).toBeGreaterThan((await pixel(0.5, 0.98))[1]);
	expect((await pixel(0.5, 0.01))[0]).toBeGreaterThan(200);

	// Grayscale runs black to white over the same range.
	await page.getByLabel('Palette').selectOption('grayscale');
	await expect.poll(async () => (await pixel(0.5, 0.98))[0]).toBeLessThan(20);
	expect((await pixel(0.5, 0.01))[0]).toBeGreaterThan(200);

	// A fixed -1 to 1 range puts 0 in the middle grey.
	await page.getByLabel('Range').selectOption('unit');
	await expect.poll(async () => (await pixel(0.5, 0.98))[0]).toBeGreaterThan(110);
	await expect(scale).toContainText('-1.00');

	// Looking down, a y gradient is the same everywhere: the value at the camera's y.
	await page.getByRole('button', { name: 'Top', exact: true }).click();
	await page.getByLabel('Camera y').fill('256');
	await expect(scale).toContainText('flat: 0.500 everywhere in view', { timeout: 20_000 });
	await page.getByRole('button', { name: 'Top', exact: true }).click();

	// The palette is remembered.
	await page.reload();
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.locator('.cm-line').first().click();
	await expect(page.getByLabel('Palette')).toHaveValue('grayscale');
	expect(errors).toEqual([]);
});

test('dimensions, noise settings and surface rules preview as terrain', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /sky-islands/ }).click();
	await expect(status(page)).toHaveText(/No problems/);
	const failed = page.getByText(/Preview failed/);

	// The dimension opens on a biome map with a legend.
	await page.getByRole('button', { name: 'dimension.mms' }).click();
	await page.locator('.cm-line', { hasText: 'Dimension islands' }).click();
	await expect(page.getByText('sky:islands', { exact: true })).toBeVisible();
	await expect(page.getByLabel('Biomes in view')).toContainText('sky:meadow', { timeout: 20_000 });

	// Blocks mode builds the terrain; hovering names the block.
	await page.getByRole('button', { name: 'Blocks', exact: true }).click();
	const canvas = page.locator('canvas');
	await expect(canvas).toBeVisible({ timeout: 30_000 });
	await expect(failed).toBeHidden();
	const box = (await canvas.boundingBox())!;
	await page.mouse.move(box.x + box.width / 2, box.y + box.height / 2);
	await expect(page.getByText(/x -?\d+, y -?\d+, z 0 · \w+/)).toBeVisible();

	// The camera moves and turns: facing east, the slice runs along z at the camera's x.
	await page.getByLabel('Camera x').fill('100');
	await page.getByRole('button', { name: 'Face east' }).click();
	await expect(page.getByRole('button', { name: 'Face east' })).toHaveAttribute('aria-pressed', 'true');
	await expect(page.getByLabel('Camera facing')).toContainText('Facing east · ← N S →');
	await page.mouse.move(box.x + box.width / 2 + 40, box.y + box.height / 2);
	await expect(page.getByText(/x 100, y -?\d+, z [1-9]\d* · \w+/)).toBeVisible({ timeout: 20_000 });
	await page.getByRole('button', { name: 'Turn right' }).click();
	await expect(page.getByLabel('Camera facing')).toContainText('Facing south · ← E W →');
	await expect(failed).toBeHidden();

	// Top looks straight down from the top of the world; hovering gives the block's own height.
	await page.getByRole('button', { name: 'Top', exact: true }).click();
	await expect(page.getByRole('button', { name: 'Top', exact: true })).toHaveAttribute('aria-pressed', 'true');
	await expect(page.getByLabel('Camera facing')).toContainText('Looking down, top is S · ← E W →');
	await expect(page.getByText('Rendering…')).toBeHidden({ timeout: 60_000 });
	await expect(failed).toBeHidden();
	// Sky islands have gaps, so look for a spot with a block under it.
	const named = page.getByText(/x -?\d+, y -?\d+, z -?\d+ · \w+/);
	await expect
		.poll(async () => {
			for (let i = 0; i < 25 && !(await named.isVisible()); i++) {
				await page.mouse.move(box.x + box.width * (0.1 + 0.2 * (i % 5)), box.y + box.height * (0.1 + 0.2 * Math.floor(i / 5)));
			}
			return named.isVisible();
		})
		.toBe(true);
	// Back to the side view, at the height it had.
	const topY = await page.getByLabel('Camera y').inputValue();
	await page.getByRole('button', { name: 'Top', exact: true }).click();
	await expect(page.getByLabel('Camera facing')).toContainText('Facing south');
	expect(await page.getByLabel('Camera y').inputValue()).not.toBe(topY);

	// A surface rule is drawn on the noise settings that use it.
	await page.getByRole('button', { name: 'terrain.mms' }).click();
	await page.locator('.cm-line', { hasText: 'SurfaceRule surface' }).click();
	await expect(page.getByText('surface rule', { exact: true })).toBeVisible();
	await expect(page.getByRole('combobox', { name: 'Terrain' })).toHaveValue('sky:islands');
	await expect(canvas).toBeVisible({ timeout: 30_000 });
	await expect(failed).toBeHidden();

	// The 3D toggle draws the same rule as an isometric diagram.
	await page.getByRole('button', { name: '3D', exact: true }).click();
	await expect(page.getByRole('button', { name: '3D', exact: true })).toHaveAttribute('aria-pressed', 'true');
	await expect(canvas).toBeVisible({ timeout: 30_000 });
	await expect(failed).toBeHidden();

	// It pans by dragging and zooms with the wheel.
	const iso = (await canvas.boundingBox())!;
	await page.mouse.move(iso.x + iso.width / 2, iso.y + iso.height / 2);
	await page.mouse.down();
	await page.mouse.move(iso.x + iso.width / 2 - 120, iso.y + iso.height / 2 - 60, { steps: 5 });
	await page.mouse.up();
	await page.mouse.wheel(0, -300);
	await expect(page.getByText('Rendering…')).toBeHidden({ timeout: 30_000 });
	await expect(failed).toBeHidden();

	// A rule that tests a biome through a named condition starts on that biome.
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /badlands/ }).click();
	await page.locator('.cm-line', { hasText: 'SurfaceRule badlands' }).click();
	await expect(page.getByRole('combobox', { name: 'Biome' })).toHaveValue('minecraft:badlands');
	expect(errors).toEqual([]);
});

test('an inline noise where an id is needed previews', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	const failed = page.getByText(/Preview failed/);

	// The game and Deepslate seed a noise by its id, so the inline one is
	// written to its own file; the cursor still shows the density function.
	await replaceText(page, 'DensityFunction d = Noise({ base_octave = -5, amplitude_modifiers = [1, 1] }, 1, 1) * 0.5\n');
	await expect(status(page)).toHaveText(/No problems/);
	await page.locator('.cm-line').first().click();
	await expect(page.getByText('my_pack:d', { exact: true })).toBeVisible();
	await expect(page.locator('canvas')).toBeVisible();
	await expect(page.locator('text=/-?\\d\\.\\d{3} … -?\\d\\.\\d{3}/')).toBeVisible({ timeout: 20_000 });
	await expect(failed).toBeHidden();
	await page.getByRole('tab', { name: 'JSON' }).click();
	await expect(page.locator('pre.json')).toContainText('"noise": "my_pack:mms_gen/d/noise"');

	// Noise settings with an inline noise in the router build terrain.
	await replaceText(
		page,
		[
			'Noise simple { base_octave = -4, amplitude_modifiers = [1] }',
			'NoiseSettings custom {',
			'  default_block = stone',
			'  default_fluid = water',
			'  disable_mob_generation = false',
			'  material_rule = [If(NoiseThreshold(simple, 0.5, 1)) { Block(glowstone) }]',
			'  sea_level = 0',
			'  noise = { min_y = -64, height = 256 }',
			'  legacy_random_source = false',
			'  noise_router = {',
			'    chunk_surface_level = 0, temperature = 0, vegetation = 0, continents = 0, erosion = 0, depth = 0, ridges = 0',
			'    final_density = Max(',
			'      left = Gradient(axis = y, from_coordinate = 0, to_coordinate = 256, from_value = 0, to_value = 1),',
			'      right = Noise({ normalize = false, base_octave = 4, base_amplitude = 6 }, 1, 1),',
			'    )',
			'  }',
			'  aquifers = { barrier = 0, fluid_level_floodedness = 0, fluid_level_spread = 0, lava = 0, exclusion = 0, surface_level = 0 }',
			'  spawn_target = []',
			'}',
			''
		].join('\n')
	);
	await expect(status(page)).toHaveText(/No problems/);
	await page.locator('.cm-line', { hasText: 'NoiseSettings custom' }).click();
	await page.getByRole('tab', { name: 'Preview' }).click();
	await expect(page.getByText('my_pack:custom', { exact: true })).toBeVisible();
	await expect(page.locator('canvas')).toBeVisible({ timeout: 30_000 });
	await expect(page.getByText('Rendering…')).toBeHidden({ timeout: 30_000 });
	await expect(failed).toBeHidden();
	expect(errors).toEqual([]);
});

test('a pinned preview stays put while the cursor moves to another declaration', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /sky-islands/ }).click();
	await expect(status(page)).toHaveText(/No problems/);

	// Pin the dimension, then go edit the surface rule in another file.
	await page.getByRole('button', { name: 'dimension.mms' }).click();
	await page.locator('.cm-line', { hasText: 'Dimension islands' }).click();
	await expect(page.getByText('sky:islands', { exact: true })).toBeVisible();
	const pin = page.getByRole('button', { name: /^Pin/ });
	await pin.click();
	await expect(pin).toHaveAttribute('aria-pressed', 'true');
	await expect(pin).toContainText('Pinned sky:islands');

	await page.getByRole('button', { name: 'terrain.mms' }).click();
	await page.locator('.cm-line', { hasText: 'SurfaceRule surface' }).click();
	await expect(page.getByText('sky:islands', { exact: true })).toBeVisible();
	await expect(page.getByText('surface rule', { exact: true })).toBeHidden();
	// The JSON tab still follows the cursor.
	await page.getByRole('tab', { name: 'JSON' }).click();
	await expect(page.locator('pre.json')).not.toContainText('"biome_source"');
	await page.getByRole('tab', { name: 'Preview' }).click();

	// Unpinning goes back to following the cursor.
	await pin.click();
	await expect(pin).toHaveAttribute('aria-pressed', 'false');
	await expect(page.getByText('surface rule', { exact: true })).toBeVisible();
	expect(errors).toEqual([]);
});

test('data packs and pasted JSON import as MMS', async ({ page }) => {
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });

	// Paste JSON appends a declaration to the open file.
	await page.getByRole('button', { name: /Paste JSON/ }).click();
	await page.getByLabel('Kind', { exact: true }).selectOption('LootTable');
	await page.getByLabel('Name', { exact: true }).fill('blocks/ruby_ore');
	await page.getByLabel('JSON', { exact: true }).fill('{"type":"minecraft:block","pools":[{"rolls":1,"entries":[{"type":"minecraft:item","name":"minecraft:diamond"}]}]}');
	await page.getByRole('button', { name: 'Convert' }).click();
	await expect(page.locator('.cm-content')).toContainText('LootTable blocks/ruby_ore {');
	await expect(page.locator('.cm-content')).toContainText('entries = [diamond]');
	await expect(status(page)).toHaveText(/No problems/);
	await page.getByRole('tab', { name: 'Pack' }).click();
	await expect(page.getByRole('button', { name: 'ruby_ore.json' })).toBeVisible();

	// A data pack zip becomes a project.
	const { zipSync, strToU8 } = await import('fflate');
	const zip = zipSync({
		'gems/pack.mcmeta': strToU8('{"pack":{"pack_format":121,"description":"Gems"}}'),
		'gems/data/gems/tags/item/shiny.json': strToU8('{"values":["minecraft:diamond"]}'),
		'gems/data/gems/function/hello.mcfunction': strToU8('say hello\n')
	});
	await page.locator('input[type=file]').setInputFiles({ name: 'gems.zip', mimeType: 'application/zip', buffer: Buffer.from(zip) });
	await page.getByRole('tab', { name: 'Sources' }).click();
	await expect(page.getByRole('button', { name: 'tags_item.mms' })).toBeVisible();
	await expect(page.getByRole('button', { name: 'function.mms' })).toBeVisible();
	await expect(status(page)).toHaveText(/No problems/);
	await page.getByRole('tab', { name: 'Pack' }).click();
	await expect(page.getByRole('button', { name: 'shiny.json' })).toBeVisible();
	await expect(page.getByRole('button', { name: 'hello.mcfunction' })).toBeVisible();
});

test('project zips keep book and text files, and their paths stay inside', async ({ page }) => {
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	const { zipSync, strToU8 } = await import('fflate');
	const zip = zipSync({
		'shelf/mms.json': strToU8('{"namespace":"shelf","description":"Shelf","output":"build"}'),
		'shelf/main.mms': strToU8('Function hello {\n  /say hi\n}\n'),
		'shelf/guide.mmsbook': strToU8('Page one\n---\nPage two\n'),
		'shelf/notes.mmstext': strToU8('Some *notes*\n'),
		'shelf/../escape.mms': strToU8('Function escape {\n  /say out\n}\n')
	});
	await page.locator('input[type=file][accept=".zip"]').setInputFiles({ name: 'shelf.zip', mimeType: 'application/zip', buffer: Buffer.from(zip) });
	await page.getByRole('tab', { name: 'Sources' }).click();
	await expect(page.getByRole('button', { name: 'guide.mmsbook' })).toBeVisible();
	await expect(page.getByRole('button', { name: 'notes.mmstext' })).toBeVisible();
	await expect(page.getByRole('button', { name: 'escape.mms', exact: true })).toBeVisible();
	await expect(page.getByRole('button', { name: /\.\./ })).toHaveCount(0);
});

test('lint fixes apply from the editor', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'ItemTag gems = ["minecraft:diamond", minecraft:emerald]\nDensityFunction d = 2 * 0.25\n');
	await expect(page.locator('.cm-lintRange-hint')).toHaveCount(4); // two names, the constant, and d is unused
	await page.keyboard.press('ControlOrMeta+Alt+f');
	await expect(page.locator('.cm-content')).toContainText('ItemTag gems = [diamond, emerald]DensityFunction d = 0.5');
	await expect(page.locator('.cm-lintRange-hint')).toHaveCount(1);
	await expect(page.locator('.cm-unnecessary')).toHaveText('d');
});

test('a custom item in a recipe ingredient shows its base item', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'CustomItem ember { base = amethyst_shard, name = "Ember" }\nRecipe r = Shapeless([ember, stick], torch)\n');
	await expect(page.locator('.cm-inlay-hint')).toHaveText('→ amethyst_shard');
	await expect(page.locator('.cm-lintRange-hint')).toHaveCount(1);
	await page.screenshot({ path: 'test-results/custom-item-base.png' });
});

test('a recipe making a custom item shows its name and texture', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /embers/ }).click();
	await expect(status(page)).toHaveText(/No problems/);
	await page.getByRole('button', { name: 'items.mms' }).click();
	await clickLine(page, 'Recipe ember_shard');
	const inspector = page.getByLabel('Inspector');
	await expect(inspector.getByLabel('Result')).toHaveText('2 × Ember Shard');
	const result = inspector.getByRole('img', { name: 'Ember Shard' });
	await expect(result).toHaveAttribute('title', /Custom item embers:ember_shard\nBase: Amethyst Shard \(minecraft:amethyst_shard\)/);
	await expect(result.locator('img')).toHaveAttribute('src', /^data:image\/png;base64,/);
	// The ingredient is the base id, which the custom item also fits.
	const shard = page.getByLabel('Crafting grid').getByRole('img', { name: 'Amethyst Shard' });
	await expect(shard).toHaveAttribute('title', /including the custom item Ember Shard/);
	await expect(page.getByLabel('Crafting grid').getByRole('img', { name: 'Blaze Powder' }).locator('img')).toHaveAttribute('src', /^data:image\/png/, { timeout: 30_000 });
	await page.screenshot({ path: 'test-results/recipe-custom-item.png' });

	// A custom item that keeps its base's look shows its name as text.
	await clickLine(page, 'Recipe ember_disc');
	await expect(inspector.getByLabel('Result')).toHaveText('Music Disc');
	await expect(inspector.getByRole('img', { name: 'Music Disc' }).locator('img')).toHaveCount(1);
	expect(errors).toEqual([]);
});

test('loot tables, recipes and biomes show in the inspector', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(
		page,
		[
			'LootTable chest { pools = [Pool(2, [Item(name = diamond).Weight(3), Item(name = emerald)])] }',
			'Recipe ruby_block = Shaped(["RRR", "R R", "RRR"], { R = emerald }, emerald_block)',
			'Biome warm { temperature = 0.8, downfall = 0.4, has_precipitation = true, effects = { water_color = "#3f76e4" }, features = [], carvers = [] }',
			''
		].join('\n')
	);
	await expect(status(page)).toHaveText(/No problems/);
	const inspector = page.getByLabel('Inspector');

	await page.locator('.cm-line').nth(0).click();
	await expect(inspector).toContainText('Pool 1');
	await expect(inspector.getByRole('row', { name: /Diamond/ })).toContainText('75%');
	await page.getByRole('button', { name: 'Roll' }).click();
	await expect(page.getByLabel('Roll result')).toContainText('×');

	await page.locator('.cm-line').nth(1).click();
	await expect(page.getByLabel('Crafting grid').getByRole('img', { name: 'Emerald' })).toHaveCount(8);
	await expect(page.getByLabel('Result')).toHaveText('Emerald Block');
	// Vanilla items draw their game icons: a flat item and a 3D block.
	await expect(page.getByLabel('Crafting grid').getByRole('img', { name: 'Emerald' }).first().locator('img')).toHaveAttribute('src', /^data:image\/png/, { timeout: 30_000 });
	await expect(inspector.getByRole('img', { name: 'Emerald Block' }).locator('img')).toHaveAttribute('src', /^data:image\/png/);
	await page.screenshot({ path: 'test-results/recipe-vanilla-icons.png' });

	await page.locator('.cm-line').nth(2).click();
	await expect(inspector).toContainText('Temperature');
	await expect(inspector).toContainText('#3f76e4');
	expect(errors).toEqual([]);
});

test('the docs render guides and the reference', async ({ page }) => {
	await page.goto('./');
	await page.getByRole('link', { name: 'Docs' }).click();
	await expect(page.getByRole('heading', { level: 1 })).toHaveText('MMS docs');

	// A guide from the grouped sidebar, and a link from one guide to a
	// section of another.
	const nav = page.getByRole('navigation', { name: 'Docs' });
	await expect(nav.getByRole('heading', { name: 'Concepts' })).toBeVisible();
	await nav.getByRole('link', { name: 'Tutorial: a custom dimension' }).click();
	await expect(page.getByRole('heading', { level: 1 })).toHaveText('Tutorial: a custom dimension');
	await expect(page.getByRole('navigation', { name: 'On this page' }).getByRole('link').first()).toBeVisible();
	await page.getByRole('article').getByRole('link', { name: 'Writing resources' }).click();
	await expect(page).toHaveURL(/\/docs\/guides\/resources#density-functions$/);
	await expect(page.locator('#density-functions')).toBeVisible();
	await expect(page.getByRole('navigation', { name: 'Previous and next' }).getByRole('link', { name: /Functions/ })).toBeVisible();

	// The Reference tab: helpers by group, one page each.
	await page.getByRole('navigation', { name: 'Sections' }).getByRole('link', { name: 'Reference' }).click();
	await expect(page.getByRole('heading', { level: 1 })).toHaveText('Reference');
	const refNav = page.getByRole('navigation', { name: 'Reference' });
	await refNav.getByRole('link', { name: 'Declaration kinds' }).click();
	await expect(page.getByRole('cell', { name: 'DensityFunction', exact: true })).toBeVisible();
	await refNav.getByRole('link', { name: 'Helpers' }).click();
	await page.getByRole('main').getByRole('link', { name: /^Density function/ }).click();
	await expect(page).toHaveURL(/\/docs\/reference\/helpers\/density_function$/);
	await expect(page.getByText(/^Gradient\(axis, /)).toBeVisible();
	await expect(page.locator('#loot_function-SetCount')).toHaveCount(0);

	// A type in a signature links to its entry on the Types page.
	await page.locator('#density_function-Gradient').getByRole('link', { name: 'Axis', exact: true }).click();
	await expect(page).toHaveURL(/\/docs\/reference\/types#type-util-direction-Axis$/);
	await expect(page.locator('#type-util-direction-Axis')).toBeInViewport();

	// A lint rule has a page of its own.
	await page.goto('./docs/reference/lint');
	await page.getByRole('table').getByRole('link', { name: 'unused', exact: true }).click();
	await expect(page.getByRole('heading', { level: 1 })).toHaveText('unused');
	await expect(page.getByText('hint by default')).toBeVisible();

	// Kinds written in MMS have a page each, from the prelude.
	await page.goto('./docs/reference/kinds');
	await page.getByRole('table').getByRole('link', { name: 'CustomItem', exact: true }).click();
	await expect(page.getByRole('heading', { level: 1 })).toHaveText('CustomItem');
	await expect(page.locator('article dt').filter({ hasText: /^base$/ })).toBeVisible();
	await expect(page.getByRole('heading', { name: 'What its name stands for' })).toBeVisible();

	// The language reference, and the forms of expression by type.
	await page.goto('./docs/reference/language/blueprints');
	await expect(page.getByRole('heading', { level: 1 })).toHaveText('Blueprints');
	await page.goto('./docs/reference/prelude');
	await expect(page.locator('#prelude-Range')).toContainText('IntNumberProviderRef');

	// Search finds pages, helpers and rules from anywhere.
	await page.keyboard.press('Control+k');
	const search = page.getByRole('dialog', { name: 'Search the docs' });
	const box = search.getByRole('searchbox');
	// Words of a name in any order, a typo, and a guide's sub-section.
	await box.fill('count set');
	await expect(search.getByRole('option').first()).toContainText('SetCount');
	await box.fill('jigsaw hole');
	await expect(search.getByRole('option', { name: /Jigsaw holes.*Writing resources/ })).toBeVisible();
	await box.fill('setcuont');
	await expect(search.getByRole('option', { name: /SetCount/ })).toBeVisible();
	await box.fill('setcount');
	await expect(search.getByRole('option').first()).toContainText('SetCount');
	await page.keyboard.press('Enter');
	await expect(page).toHaveURL(/\/docs\/reference\/helpers\/loot_function#loot_function-SetCount$/);

	// The reference for AI tools.
	const llms = await (await page.request.get('./docs/llms.txt')).text();
	expect(llms).toMatch(/^# MMS\n\n> /);
	expect(llms).toContain('/docs/guides/functions');
	const yaml = await (await page.request.get('./docs/reference.yaml')).text();
	expect(yaml).toContain('lint_rules:');
	expect(yaml).toContain('- name: "SetCount"');
	expect(yaml).toContain('- name: "CustomItem"');
});

test('the Project menu closes on outside click and Escape, and the editor header formats and fixes', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	const projectButton = page.getByRole('button', { name: /Project/ });
	const newProject = page.getByRole('button', { name: 'New project' });

	await projectButton.click();
	await expect(newProject).toBeVisible();
	await page.locator('.cm-content').click();
	await expect(newProject).toBeHidden();

	await projectButton.click();
	await expect(newProject).toBeVisible();
	await page.keyboard.press('Escape');
	await expect(newProject).toBeHidden();
	await expect(projectButton).toBeFocused();

	await replaceText(page, 'DensityFunction   x=Clamp( 1,0 ,1 )\n');
	await expect(status(page)).toHaveText(/No problems/);
	await page.getByRole('button', { name: 'Format', exact: true }).click();
	await expect(page.locator('.cm-content')).toHaveText('DensityFunction x = Clamp(1, 0, 1)');

	// The Project menu holds only project actions; editing actions live in the editor header.
	await projectButton.click();
	await expect(page.getByRole('button', { name: /Fix all in this file/ })).toHaveCount(0);
	await expect(page.getByRole('button', { name: /badlands/ })).toHaveCount(0);
	await page.keyboard.press('Escape');

	await replaceText(page, 'DensityFunction a = 2 * 0.25\n');
	await expect(page.locator('.cm-lintRange-hint').first()).toBeVisible();
	await page.getByRole('button', { name: 'Fix all', exact: true }).click();
	await expect(page.locator('.cm-content')).toContainText('DensityFunction a = 0.5');
});

test('doc comments show on hover over a use', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(
		page,
		'/**\n * Kills needed for each reward.\n * @deprecated use reward_every\n */\nGlobal every: int = 5\n\nFunction f {\n  /say ${every}\n}\n'
	);
	await expect(page.locator('.cm-deprecated')).toHaveCount(1);
	await page.locator('.cm-deprecated').hover();
	const tip = page.locator('.cm-tooltip-hover, .cm-tooltip-lint').filter({ hasText: 'Kills needed' });
	await expect(tip.first()).toContainText('Deprecated.');
});

test('doc comments cannot run script in a hover', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(
		page,
		'/**\n * Pwned <img src=x onerror="window.__pwned=1"> [link](javascript:window.__pwned=2)\n * @deprecated use reward_every\n */\nGlobal every: int = 5\n\nFunction f {\n  /say ${every}\n}\n'
	);
	await expect(page.locator('.cm-deprecated')).toHaveCount(1);
	const tip = page.locator('.cm-lsp-hover-tooltip').filter({ hasText: 'Pwned' });
	// The first hover can land while the line is redrawn; hover again.
	await expect(async () => {
		await page.mouse.move(0, 0);
		await page.locator('.cm-deprecated').hover();
		await expect(tip).toBeVisible({ timeout: 1500 });
	}).toPass({ timeout: 15_000 });
	await expect(tip.locator('img')).toHaveCount(0);
	await expect(tip.locator('a', { hasText: 'link' })).not.toHaveAttribute('href');
	// Highlighting classes on the signature survive.
	await expect(tip.locator('code.language-mms')).toHaveCount(1);
	expect(await page.evaluate(() => (window as { __pwned?: number }).__pwned)).toBeUndefined();
});

test('a problem in text a Blueprint wrote links to the line that wrote it', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(
		page,
		'Blueprint Thing {\n  n: int\n  emit {\n    Noise ${self.path} { base_octave = n, amplitude_modifiers = "oops" }\n  }\n}\nThing a { n = 1 }\n'
	);
	const range = page.locator('.cm-lintRange-error').last();
	await expect(range).toBeVisible();
	await range.hover();
	const go = page.locator('.cm-tooltip-lint').getByRole('button', { name: 'Go to Blueprint Thing' });
	await expect(go).toBeVisible();
	await go.click();
	await expect(page.locator('.cm-activeLine')).toContainText('amplitude_modifiers = "oops"');
});

test('right-click opens editor actions with quick fixes', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'DensityFunction a = 2 * 0.25\nDensityFunction b = a * 2\n');
	await expect(page.locator('.cm-lintRange-hint').first()).toBeVisible();
	const menu = page.getByRole('menu', { name: 'Editor actions' });

	// Go to definition from a reference.
	await page.locator('.cm-line').nth(1).getByText('a', { exact: true }).click({ button: 'right' });
	await expect(menu).toBeVisible();
	await expect(menu.getByRole('menuitem', { name: /Find references/ })).toContainText('Shift+F12');
	await menu.getByRole('menuitem', { name: /Go to definition/ }).click();
	await expect(menu).toBeHidden();
	await expect(page.locator('.cm-activeLine')).toContainText('DensityFunction a');

	// Escape and outside clicks close it.
	await page.locator('.cm-line').first().click({ button: 'right' });
	await expect(menu).toBeVisible();
	await page.keyboard.press('Escape');
	await expect(menu).toBeHidden();
	await page.locator('.cm-line').first().click({ button: 'right' });
	await page.getByRole('tab', { name: 'JSON' }).click();
	await expect(menu).toBeHidden();

	// A quick fix for the problem under the cursor.
	await page.locator('.cm-lintRange-hint', { hasText: '2 * 0.25' }).click({ button: 'right' });
	await expect(menu.getByRole('menuitem', { name: /mms-ignore|Ignore/i })).toBeVisible();
	const fix = menu.getByRole('menuitem').filter({ hasNotText: /Fix all|ignore/i }).last();
	await expect(fix).not.toContainText('Format');
	await fix.click();
	await expect(page.locator('.cm-content')).toContainText('DensityFunction a = 0.5');

	// Ctrl+. lists the fixes at the cursor, and Enter applies the first.
	await page.locator('.cm-line', { hasText: 'b = a * 2' }).click();
	await page.keyboard.press('ControlOrMeta+.');
	await expect(menu).toBeVisible();
	await expect(menu.getByRole('menuitem', { name: /Go to definition/ })).toHaveCount(0);
	await expect(menu.getByRole('menuitem', { name: /Fix all in this file/ })).toBeVisible();
	await page.keyboard.press('Escape');
	await expect(menu).toBeHidden();
	await page.locator('.cm-lintRange-hint').first().click();
	await page.keyboard.press('ControlOrMeta+.');
	await expect(menu.getByRole('menuitem').first()).toBeFocused();
	await expect(menu.getByRole('menuitem').first()).not.toContainText('Fix all');
});

test('Tab indents, Shift+Tab unindents, and Tab accepts a completion', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'Function f {\n/say hi\n}\n');
	const line = page.locator('.cm-line').nth(1);
	await line.click();
	await page.keyboard.press('Tab');
	await expect(line).toHaveText('  /say hi');
	await expect(page.locator('.cm-content')).toBeFocused();
	await page.keyboard.press('Shift+Tab');
	await expect(line).toHaveText('/say hi');

	await replaceText(page, 'DensityFunction x = Cl');
	await page.keyboard.press('ControlOrMeta+Space');
	await expect(page.locator('.cm-tooltip-autocomplete li[aria-selected]')).toContainText('Clamp');
	await page.waitForTimeout(300);
	await page.keyboard.press('Tab');
	await expect(page.locator('.cm-content')).toContainText('DensityFunction x = Clamp');
});

test('completion opens on request with Ctrl+Space, Cmd/Ctrl+I and Alt+Space', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	const popup = page.locator('.cm-tooltip-autocomplete');
	// A blank line in a function offers execute helpers and keywords.
	await replaceText(page, 'Function f {\n  /say hi\n  \n}\n');
	await page.locator('.cm-line').nth(2).click();
	await page.keyboard.press('End');
	for (const key of ['Control+Space', 'ControlOrMeta+i', 'Alt+Space']) {
		await page.keyboard.press(key);
		await expect(popup).toContainText('Positioned');
		await expect(popup).toContainText('let');
		await page.keyboard.press('Escape');
		await expect(popup).toHaveCount(0);
	}
	await expect(page.locator('.cm-line').nth(2)).toHaveText('  ');
	// After an operator with nothing typed yet.
	await replaceText(page, 'DensityFunction x = 1 / ');
	await page.keyboard.press('Control+Space');
	await expect(popup).toContainText('Clamp');
});

test('F2 renames inline with a preview of occurrences', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'DensityFunction a = 0.5\nDensityFunction b = a * a\n');
	await expect(status(page)).toHaveText(/No problems|warning/);
	await page.locator('.cm-line').first().getByText('a', { exact: true }).click();

	// Escape cancels without changes.
	await page.keyboard.press('F2');
	const box = page.getByRole('textbox', { name: 'New name' });
	await expect(box).toBeFocused();
	await expect(box).toHaveValue('a');
	await expect(page.locator('.cm-rename-hint')).toContainText('3 occurrences');
	await expect(page.locator('.cm-rename-occurrence')).toHaveCount(3);
	await page.keyboard.press('Escape');
	await expect(box).toBeHidden();
	await expect(page.locator('.cm-rename-occurrence')).toHaveCount(0);
	await expect(page.locator('.cm-content')).toBeFocused();

	// Typing replaces the selected name; Enter applies it everywhere.
	await page.keyboard.press('F2');
	await expect(box).toBeFocused();
	await page.keyboard.type('base');
	await page.keyboard.press('Enter');
	await expect(box).toBeHidden();
	await expect(page.locator('.cm-content')).toContainText('DensityFunction base = 0.5DensityFunction b = base * base');

	// The right-click menu opens the same box.
	await page.locator('.cm-line').first().getByText('base', { exact: true }).click({ button: 'right' });
	await page.getByRole('menuitem', { name: /Rename symbol/ }).click();
	await expect(box).toHaveValue('base');
	await page.keyboard.press('Escape');
});

test('built-in helpers explain themselves on hover and while typing', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'SurfaceRule r = If(StoneDepth(Floor)) { Block(grass_block) }\n');
	await expect(status(page)).toHaveText(/No problems/);

	// Hover: what it does, each parameter, an example and a reference link.
	const at = await page.evaluate(() => {
		const walk = document.createTreeWalker(document.querySelector('.cm-content')!, NodeFilter.SHOW_TEXT);
		for (let n = walk.nextNode(); n; n = walk.nextNode()) {
			const i = n.textContent!.indexOf('StoneDepth');
			if (i < 0) continue;
			const r = document.createRange();
			r.setStart(n, i + 2);
			r.setEnd(n, i + 3);
			const b = r.getBoundingClientRect();
			return { x: b.x + b.width / 2, y: b.y + b.height / 2 };
		}
		return { x: 0, y: 0 };
	});
	// "No problems" already showed before the edit, so the server may not
	// have the new text yet; hover again until it answers.
	const hover = page.locator('.cm-lsp-hover-tooltip');
	await expect(async () => {
		await page.mouse.move(0, 0);
		await page.mouse.move(at.x, at.y);
		await expect(hover).toContainText('True for blocks close to a surface', { timeout: 1500 });
	}).toPass({ timeout: 15_000 });
	await expect(hover).toContainText('secondary_depth_range');
	await expect(hover.locator('a[href*="reference/helpers/material_condition#material_condition-StoneDepth"]')).toHaveCount(1);
	await page.mouse.move(0, 0);

	// Signature help shows the active parameter's meaning first.
	await page.locator('.cm-content').click();
	await page.keyboard.press('ControlOrMeta+End');
	// Typed key by key, so brackets close as they would for a person.
	await page.keyboard.type('SurfaceRule s = If(StoneDepth(Floor,');
	const sig = page.locator('.cm-lsp-signature-tooltip');
	await expect(sig.locator('.cm-lsp-active-parameter')).toHaveText(/^offset/);
	await expect(sig).toContainText('Extra depth that still passes');
});

test.describe('on a phone', () => {
	test.use({ viewport: { width: 390, height: 664 }, hasTouch: true, isMobile: true });

	test('the 3D preview pans with a finger and zooms with a pinch about the pinch', async ({ page, context }) => {
		page.on('dialog', (d) => d.accept());
		await page.goto('./');
		await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
		// Nothing is wider than the phone. If something were, the browser would widen
		// the page and pan the screen over it, and the touches below would land off
		// from where the canvas's box says they are.
		expect(await page.evaluate(() => [innerWidth, document.documentElement.scrollWidth])).toEqual([390, 390]);
		await page.getByRole('button', { name: /Examples/ }).click();
		await page.getByRole('button', { name: /badlands/ }).click();
		const rule = page.locator('.cm-line', { hasText: 'SurfaceRule badlands' });
		// The editor only renders lines near the viewport; scroll until the rule is there.
		for (let i = 0; i < 20 && (await rule.count()) === 0; i++) {
			await page.locator('.cm-scroller').evaluate((s) => (s.scrollTop += 300));
		}
		await rule.tap();
		await page.getByRole('button', { name: '3D', exact: true }).tap();
		const canvas = page.locator('canvas');
		await expect(canvas).toBeVisible({ timeout: 30_000 });
		await expect(canvas).toHaveAttribute('data-view', /"scale":[1-9]/);

		const box = (await canvas.boundingBox())!;
		const x = box.x + box.width * 0.65;
		const y = box.y + Math.min(box.height / 2, 50);
		const view = async () => JSON.parse((await canvas.getAttribute('data-view'))!);
		/** The world x/z under a client point, by the view's projection. */
		const worldAt = async () => {
			const v = await view();
			return canvas.evaluate((c: HTMLCanvasElement, [x, y, v]) => {
				const r = c.getBoundingClientRect();
				const dx = ((x - r.left) / r.width) * c.width - c.width / 2;
				const dy = ((y - r.top) / r.height) * c.height - c.height / 2;
				return [v.x + (dx / v.scale + (2 * dy) / v.scale) / 2, v.y + ((2 * dy) / v.scale - dx / v.scale) / 2];
			}, [x, y, v] as const);
		};
		const cdp = await context.newCDPSession(page);
		const touch = (type: string, points: [number, number][]) =>
			cdp.send('Input.dispatchTouchEvent', { type: type as 'touchStart', touchPoints: points.map(([x, y], id) => ({ x, y, id })) });

		// Spreading two fingers zooms in and keeps the ground under them in place.
		const before = await view();
		const ground = await worldAt();
		await touch('touchStart', [[x - 20, y], [x + 20, y]]);
		for (let i = 1; i <= 10; i++) await touch('touchMove', [[x - 20 - i * 6, y], [x + 20 + i * 6, y]]);
		await touch('touchEnd', []);
		await expect.poll(async () => (await view()).scale).toBeCloseTo(before.scale * 4, 5);
		const after = await worldAt();
		expect(after[0]).toBeCloseTo(ground[0], 5);
		expect(after[1]).toBeCloseTo(ground[1], 5);

		// One finger pans; the page itself neither zooms nor scrolls.
		const zoomed = await view();
		await touch('touchStart', [[x, y]]);
		for (let i = 1; i <= 8; i++) await touch('touchMove', [[x - i * 10, y]]);
		await touch('touchEnd', []);
		await expect.poll(async () => (await view()).x).toBeGreaterThan(zoomed.x);
		expect(await page.evaluate(() => [visualViewport!.scale, scrollX, scrollY])).toEqual([1, 0, 0]);
		await expect(page.getByText(/Preview failed/)).toBeHidden();
	});
});

test('several projects are saved and can be switched', async ({ page }) => {
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'Noise first { base_octave = -4, amplitude_modifiers = [1] }\n');

	// An example opens as a second project; the first stays saved.
	await page.getByRole('button', { name: 'Examples ▾' }).click();
	await page.getByRole('button', { name: /surface-rules/ }).click();
	await expect(page.getByRole('button', { name: 'surface.mms' })).toBeVisible();

	await page.getByRole('button', { name: 'surface-rules (switch project)' }).click();
	const list = page.getByRole('list', { name: 'Saved projects' });
	await expect(list.getByRole('listitem')).toHaveCount(2);
	await list.getByRole('button', { name: /^my-pack/ }).click();
	await expect(page.locator('.cm-content')).toContainText('Noise first');

	// Rename it, and the name survives a reload along with both projects.
	await page.getByRole('button', { name: 'my-pack (switch project)' }).click();
	await list.getByRole('button', { name: 'Rename my-pack' }).click();
	await page.getByRole('textbox', { name: 'Project name' }).fill('hills');
	await page.getByRole('textbox', { name: 'Project name' }).press('Enter');
	await expect(list.getByRole('button', { name: /^hills/ })).toBeVisible();
	await page.getByRole('button', { name: 'Close' }).click();

	await page.reload();
	await expect(page.getByRole('button', { name: 'hills (switch project)' })).toBeVisible();
	await expect(page.locator('.cm-content')).toContainText('Noise first');

	// Deleting the open project opens the other one.
	await page.getByRole('button', { name: 'hills (switch project)' }).click();
	await list.getByRole('button', { name: 'Delete hills' }).click();
	await expect(page.getByRole('button', { name: 'surface.mms' })).toBeVisible();
	await expect(list.getByRole('listitem')).toHaveCount(1);
});

test('Format is Shift+Alt+F, or Cmd+Shift+I on macOS', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'DensityFunction   x=Clamp( 1,0 ,1 )\n');
	await page.keyboard.press('Shift+Alt+f');
	await expect(page.locator('.cm-content')).toHaveText('DensityFunction x = Clamp(1, 0, 1)');
	await expect(page.getByRole('button', { name: 'Format', exact: true })).toHaveAttribute('title', /Shift\+Alt\+F/);
});

test.describe('on macOS', () => {
	test.beforeEach(async ({ page }) => {
		await page.addInitScript(() => Object.defineProperty(Navigator.prototype, 'platform', { get: () => 'MacIntel' }));
	});

	test('Cmd+Shift+I formats, and Cmd+Option+F fixes all', async ({ page }) => {
		await page.goto('./');
		await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
		await page.locator('.cm-content').click();
		await page.keyboard.press('Meta+a');
		await page.keyboard.press('Delete');
		await page.keyboard.insertText('DensityFunction   x=Clamp( 1,0 ,1 )\n');
		await page.keyboard.press('Meta+Shift+I');
		await expect(page.locator('.cm-content')).toHaveText('DensityFunction x = Clamp(1, 0, 1)');
		await expect(page.getByRole('button', { name: 'Format', exact: true })).toHaveAttribute('title', /⌘\+Shift\+I/);

		await page.keyboard.press('Meta+a');
		await page.keyboard.press('Delete');
		await page.keyboard.insertText('DensityFunction d = 2 * 0.25\n');
		await expect(page.locator('.cm-lintRange-hint').first()).toBeVisible();
		await page.keyboard.press('Meta+Alt+f');
		await expect(page.locator('.cm-content')).toContainText('DensityFunction d = 0.5');

		await page.locator('.cm-line').first().click({ button: 'right' });
		const menu = page.getByRole('menu', { name: 'Editor actions' });
		await expect(menu.getByRole('menuitem', { name: /Format document/ })).toContainText('⌘+Shift+I');
		await expect(menu.getByRole('menuitem', { name: /Fix all/ })).toContainText('⌘+⌥+F');
	});
});

test('a book file in an example is checked and completed in place', async ({ page }) => {
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /library/ }).click();
	await page.getByRole('button', { name: 'books/day_12.mmsbook' }).click();
	await expect(page.locator('.cm-content')).toContainText('The lamps went out');
	await expect(status(page)).toHaveText(/No problems/);
	// A bad style is reported in the book file itself.
	await page.locator('.cm-line').last().click();
	await page.keyboard.press('End');
	await page.keyboard.insertText('\n{purpl}x{/}');
	await expect(status(page)).toHaveText(/1 error/);
	await expect(page.locator('.cm-lintRange-error')).toHaveCount(1);
	// Completion after { offers colors.
	await page.keyboard.press('Backspace');
	for (let i = 0; i < 'purpl}x{/}'.length - 1; i++) await page.keyboard.press('Backspace');
	await page.keyboard.insertText('go');
	await page.keyboard.press('ControlOrMeta+Space');
	await expect(page.locator('.cm-tooltip-autocomplete')).toContainText('gold');
});

test('the Pack tab lists custom items and splits the two packs', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'CustomItem ruby {\n  base = stick\n  components = { max_stack_size = 1 }\n}\nCustomItem charm {\n  base = torch\n  name = "Charm"\n}\n');
	await expect(status(page)).toHaveText(/No problems/);
	await page.getByRole('tab', { name: 'Pack' }).click();
	// ruby writes no file, but is listed with the stack its name stands for.
	await expect(page.getByText('Custom items')).toBeVisible();
	await page.getByRole('button', { name: 'my_pack:ruby' }).click();
	await page.getByRole('tab', { name: 'JSON' }).click();
	await expect(page.locator('pre.json')).toContainText('"max_stack_size": 1');
	await expect(page.getByText('stick[max_stack_size=1]')).toBeVisible();
	// Both packs have a pack.mcmeta; each one opens its own.
	await expect(page.getByText('Resource pack')).toBeVisible();
	await page.getByRole('button', { name: 'pack.mcmeta' }).last().click();
	await expect(page.locator('pre.json')).toContainText('"min_format": 97');
});

test('a book previews its pages and opens on the page at the cursor', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /library/ }).click();
	await page.getByRole('button', { name: 'books.mms' }).click();
	await expect(status(page)).toHaveText(/No problems/);
	// The cursor in the guide's text shows the book on that page.
	await page.locator('.cm-line', { hasText: '# Contents' }).click();
	const book = page.getByLabel('Book preview');
	await expect(book.getByText('Library Guide')).toBeVisible();
	await expect(book.getByText('Page 2 of 3')).toBeVisible();
	await expect(book.getByText("Loading the game's book")).toBeHidden({ timeout: 20_000 });
	// The score on that page gets a sample value to try.
	await expect(book.getByRole('textbox', { name: 'deaths[@s]' })).toBeVisible();
	await book.getByRole('button', { name: 'Next ›' }).click();
	await expect(book.getByText('Page 3 of 3')).toBeVisible();
	// A .mmsbook file shows the book that reads it.
	await page.getByRole('button', { name: 'books/day_12.mmsbook' }).click();
	await page.locator('.cm-line', { hasText: 'Something knocks' }).click();
	await expect(book.getByText('Diary, Day 12')).toBeVisible();
	await expect(book.getByText('Page 2 of 2')).toBeVisible();
	expect(errors).toEqual([]);
});

test('Ctrl+. extracts an inline noise to a declaration and renames it', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'DensityFunction hills = Noise({ base_octave = -6, amplitude_modifiers = [1] }, 1, 1)\n');
	await expect(status(page)).toHaveText(/No problems|warning/);
	await page.locator('.cm-line').first().getByText('base_octave').click();
	await page.keyboard.press('ControlOrMeta+.');
	const menu = page.getByRole('menu', { name: 'Editor actions' });
	const extract = menu.getByRole('menuitem', { name: 'Extract { ... } to a Noise declaration' });
	// The menu opens once the language server answers, which waits behind
	// the analysis of the new text; on a loaded machine that takes longer
	// than the default 5 seconds.
	await expect(extract).toBeVisible({ timeout: 15_000 });
	await extract.click();
	await expect(page.locator('.cm-content')).toContainText(
		'Noise hills_noise { base_octave = -6, amplitude_modifiers = [1] }DensityFunction hills = Noise(hills_noise, 1, 1)'
	);
	// The rename box opens on the new name.
	const box = page.getByRole('textbox', { name: 'New name' });
	await expect(box).toBeFocused();
	await expect(box).toHaveValue('hills_noise');
	await page.keyboard.type('hill_shape');
	await page.keyboard.press('Enter');
	await expect(page.locator('.cm-content')).toContainText('Noise hill_shape {');
	await expect(page.locator('.cm-content')).toContainText('Noise(hill_shape, 1, 1)');
	await expect(status(page)).toHaveText(/No problems|warning/);
});

test('Deepslate is credited in About, on previews, and in the shipped licenses', async ({ page, request }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });

	await page.getByRole('button', { name: 'About and credits' }).click();
	const credits = page.getByRole('list', { name: 'Credits' });
	await expect(credits.getByRole('link', { name: 'Deepslate' })).toHaveAttribute('href', 'https://github.com/misode/deepslate');
	await expect(credits.getByRole('link', { name: 'Misode' }).first()).toHaveAttribute('href', 'https://github.com/misode');
	await page.getByRole('button', { name: 'Close' }).click();

	// A rendered preview carries a small credit in its corner.
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /sky-islands/ }).click();
	await page.getByRole('button', { name: 'dimension.mms' }).click();
	await page.locator('.cm-line', { hasText: 'Dimension islands' }).click();
	const credit = page.locator('[data-credit="deepslate"]');
	await expect(credit).toContainText('Rendered with Deepslate by Misode', { timeout: 30_000 });
	await expect(credit.getByRole('link', { name: 'Deepslate' })).toHaveAttribute('href', 'https://github.com/misode/deepslate');

	const licenses = await request.get('third-party-licenses.txt');
	expect(licenses.ok()).toBe(true);
	const text = await licenses.text();
	expect(text).toContain('deepslate 0.27');
	expect(text).toContain('Copyright (c) 2021 Misode');
	expect(text).toContain('Copyright (c) 2022 SpyglassMC');
});

test('a jigsaw structure simulates a layout in 3D, with a seed, blocks and a slice', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /jigsaw/ }).click();
	await expect(status(page)).toHaveText(/No problems/);
	await page.locator('.cm-line', { hasText: 'JigsawStructure dungeon {' }).click();
	await page.getByRole('tab', { name: 'Simulation' }).click();
	const sim = page.getByLabel('Jigsaw simulation');
	const summary = sim.getByLabel('Simulation summary');
	// The dungeon is six pieces deep, and caps close the doors past that.
	await expect(summary).toContainText(/\d+ pieces · 7 deep \(size 6\)/);
	const first = await summary.textContent();
	// The same seed gives the same layout; another seed another one.
	await sim.getByLabel('Layout seed').fill('42');
	await expect(summary).toContainText('46 pieces');
	await sim.getByLabel('Layout seed').fill('0');
	await expect(summary).toHaveText(first!);

	// Clicking the middle of the view picks a piece and shows what joined it.
	const canvas = sim.getByLabel('Simulated structure');
	await page.getByRole('button', { name: 'Fit' }).click();
	const box = (await canvas.boundingBox())!;
	await page.mouse.click(box.x + box.width / 2, box.y + box.height / 2);
	const piece = sim.getByLabel('Simulated piece');
	await expect(piece).toContainText(/start|halls|rooms|caps/);
	await page.screenshot({ path: 'test-results/jigsaw-boxes.png' });

	// Blocks reads the structure files; a slice cuts through them.
	await sim.getByRole('button', { name: 'Blocks' }).click();
	await expect(sim).not.toContainText('Blocks not read');
	await sim.getByRole('group', { name: 'Slice' }).getByRole('button', { name: 'y' }).click();
	await expect(sim.getByLabel('Slice position')).toBeVisible();
	await sim.getByLabel('Slice position').fill('1');
	await expect(sim).toContainText('y ≤ 1');
	await page.screenshot({ path: 'test-results/jigsaw-blocks.png' });
	expect(errors).toEqual([]);
});

test('a jigsaw structure previews which pieces join which, and marks dead ends', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /jigsaw/ }).click();
	await expect(status(page)).toHaveText(/No problems/);
	await page.locator('.cm-line', { hasText: 'JigsawStructure dungeon {' }).click();
	const graph = page.getByLabel('Jigsaw connections');
	await expect(graph).toContainText('Jigsaw structure');
	await expect(graph).toContainText('size 6');
	await expect(graph.getByLabel('Problems')).toHaveText('No problems found');
	for (const pool of ['start', 'halls', 'rooms', 'caps']) await expect(graph.getByLabel(`Pool dungeon:${pool}`).first()).toBeVisible();
	await expect(graph.getByLabel('Pool dungeon:halls')).toContainText('75%');
	// Rooms lead back to halls: each step gets a column until growth repeats.
	await expect(graph.getByLabel('Pool dungeon:caps')).toHaveCount(2);
	await expect(graph).toContainText('then as depth 1 again');

	// A piece lists its jigsaw blocks and what joins each.
	await graph.getByRole('button', { name: /^start/ }).click();
	const details = page.getByLabel('Piece details');
	await expect(details.getByRole('row')).toHaveCount(5);
	await expect(details.getByRole('row').nth(1)).toContainText('halls/hall (Long hall), halls/short_hall (Short hall)');
	// Each jigsaw block copies a /give for another one set up the same way.
	await page.context().grantPermissions(['clipboard-read', 'clipboard-write']);
	await details.getByRole('row').nth(1).getByRole('button', { name: 'Copy /give' }).click();
	expect(await page.evaluate(() => navigator.clipboard.readText())).toMatch(/^\/give @s jigsaw\[block_entity_data=\{id:"jigsaw",name:"dungeon:room_exit",target:"dungeon:hall_entry",pool:"dungeon:halls"/);
	// A doc comment above a pool element labels its piece.
	await expect(graph.getByLabel('Pool dungeon:halls').first()).toContainText('Long hall');

	// Renaming a pool leaves the halls' exits pointing at nothing. The
	// preview is pinned, since the edit moves the cursor away.
	await page.getByRole('button', { name: /^Pin/ }).click();
	const src = readFileSync(new URL('../../examples/jigsaw/dungeon.mms', import.meta.url), 'utf8');
	await replaceText(page, src.replace('rooms = { fallback', 'rooms2 = { fallback'));
	await expect(graph.getByLabel('Problems')).toContainText('2 dead ends');
	await expect(graph.getByLabel('Pool dungeon:rooms')).toContainText('No TemplatePool declares this');

	// The graph copies as Markdown, dead ends and all, to share it.
	await page.context().grantPermissions(['clipboard-read', 'clipboard-write']);
	await graph.getByRole('button', { name: 'Copy as Markdown' }).click();
	await expect(graph.getByRole('button', { name: 'Copied' })).toBeVisible();
	const md = await page.evaluate(() => navigator.clipboard.readText());
	expect(md).toContain('# Jigsaw connections: dungeon:dungeon');
	expect(md).toContain('- Dead end: `dungeon:halls/hall` "Long hall" in `dungeon:halls`, jigsaw block at 2 1 8 facing south (name `dungeon:hall_exit`, pool `dungeon:rooms`, target `dungeon:room_entry`): no TemplatePool declares dungeon:rooms.');
	expect(md).toContain('- `dungeon:halls/hall` "Long hall" · weight 3 (75%) · size 5×5×9');
	expect(md).toContain('No TemplatePool declares this pool.');
	expect(md).toContain('```mermaid\nflowchart LR');
	await graph.getByRole('button', { name: 'Copy Mermaid' }).click();
	expect(await page.evaluate(() => navigator.clipboard.readText())).toMatch(/^flowchart LR\n {2}subgraph P0\["start · start"\]/);
	expect(errors).toEqual([]);
});

test('a project file nothing uses is marked unused in the asset list', async ({ page }) => {
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /jigsaw/ }).click();
	await expect(status(page)).toHaveText(/No problems/);
	const vault = page.getByRole('listitem').filter({ hasText: 'structures/rooms/vault.nbt' });
	await expect(vault).not.toContainText('unused');
	const src = readFileSync(new URL('../../examples/jigsaw/dungeon.mms', import.meta.url), 'utf8');
	await replaceText(page, src.replace('Piece("rooms/vault")', 'Piece("rooms/room")'));
	await expect(vault).toContainText('unused');
	await expect(vault).toHaveAttribute('title', /Not used/);
});

test('a tell or title previews its message as the game shows it', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(
		page,
		'Objective kills: dummy\nFunction f {\n  tell(@a, "{gold}**Wave** ${kills[@s]} begins")\n  title(@a, "Victory", subtitle = "{gray}*well held*")\n}\n'
	);
	const text = page.getByLabel('Text preview');
	const chat = text.getByLabel('Chat');
	// The status still reads "No problems" from before the edit until the
	// new analysis lands, so wait for the preview, not the status.
	await expect(async () => {
		await page.locator('.cm-line', { hasText: 'tell(@a' }).click();
		await expect(chat).toContainText('Wave 0 begins', { timeout: 2_000 });
	}).toPass({ timeout: 30_000 });
	await expect(chat.getByText('Wave', { exact: true })).toHaveCSS('font-weight', '700');
	// The score is filled in when the message shows; try a value.
	await text.getByRole('textbox', { name: 'kills[@s]' }).fill('7');
	await expect(chat).toContainText('Wave 7 begins');
	await page.locator('.cm-line', { hasText: 'title(@a' }).click();
	await expect(text.getByLabel('Title', { exact: true })).toHaveText('Victory');
	await expect(text.getByLabel('Subtitle')).toHaveText('well held');
	expect(errors).toEqual([]);
});

test('Test declarations run in the Tests tab and follow edits', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'Global n: int = 0\nFunction bump {\n  n += 1\n  /say bumped\n}\nTest bumps {\n  run bump()\n  expect n == 2\n}\n');
	const tab = page.getByRole('tab', { name: 'Tests (1 failed)' });
	await expect(tab).toBeVisible({ timeout: 15_000 });
	await tab.click();
	await expect(page.getByTestId('tests-summary')).toHaveText('1 of 1 test failed');
	await expect(page.getByText('expected n == 2, but n is 1')).toBeVisible();
	await expect(page.getByText('bumped', { exact: true })).toBeVisible();
	await page.getByRole('button', { name: 'bumps: failed' }).click();
	await expect(page.locator('.cm-activeLine')).toContainText('expect n == 2');
	await replaceText(page, 'Global n: int = 0\nFunction bump {\n  n += 2\n}\nTest bumps {\n  run bump()\n  expect n == 2\n}\n');
	await expect(page.getByTestId('tests-summary')).toHaveText('1 test passed', { timeout: 15_000 });
	await expect(page.getByRole('tab', { name: 'Tests', exact: true })).toBeVisible();
});

test('the Expanded tab shows what a template call or blueprint writes', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'Template ring(p: Id<particle>, points: int = 2): Commands {\n  for i in 0..points - 1 {\n    /particle ${p} ^${i} ^1 ^0\n  }\n}\nFunction f {\n  let n = 0\n  ring(flame)\n}\n');
	// The call gets an `inlined` hint and the let its inferred type.
	await expect(page.locator('.cm-inlay-hint', { hasText: 'inlined' })).toBeVisible({ timeout: 15_000 });
	await expect(page.locator('.cm-inlay-type')).toHaveText(': int');
	await page.getByRole('tab', { name: 'Expanded' }).click();
	await page.locator('.cm-line', { hasText: 'ring(flame)' }).click();
	await expect(page.getByTestId('expanded-text')).toHaveText('/particle flame ^0 ^1 ^0\n/particle flame ^1 ^1 ^0');
	await expect(page.getByLabel('Expanded')).toContainText('Template ring');
});

test('go to definition on a built-in opens the prelude read-only', async ({ page }) => {
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(page, 'CustomItem ruby { base = diamond }\n');
	await expect(status(page)).toHaveText(/No problems|problem/, { timeout: 15_000 });
	await page.locator('.cm-line', { hasText: 'CustomItem ruby' }).click({ position: { x: 20, y: 5 } });
	await expect(async () => {
		await page.keyboard.press('F12');
		await expect(page.getByTestId('prelude-view')).toBeVisible({ timeout: 2_000 });
	}).toPass({ timeout: 15_000 });
	await expect(page.locator('#prelude-title')).toContainText('items.mms');
	// The definition is selected and scrolled into view.
	await expect(page.getByTestId('prelude-view').locator('.cm-activeLine')).toContainText('Blueprint CustomItem');
	await page.getByRole('button', { name: 'Close' }).click();
	await expect(page.getByTestId('prelude-view')).toBeHidden();
});

test('the Jigsaw tab copies a jigsaw block, keeps it in Recent, and declares a new hole', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	// A project with no jigsaw has no Jigsaw tab.
	await expect(page.getByRole('tab', { name: 'Jigsaw' })).toHaveCount(0);
	await page.getByRole('button', { name: /Examples/ }).click();
	await page.getByRole('button', { name: /jigsaw/ }).click();
	await expect(status(page)).toHaveText(/No problems/);
	await page.getByRole('tab', { name: 'Jigsaw' }).click();
	const tool = page.getByLabel('Jigsaw', { exact: true });

	// The pools and holes the project declares are offered.
	await expect(tool.locator('#jigsaw-pools option[value="dungeon:halls"]')).toHaveCount(1);
	await expect(tool.locator('#jigsaw-names option[value="dungeon:room_exit"]')).toHaveCount(1);
	// The structures' own jigsaw blocks are listed.
	await expect(tool.getByRole('heading', { name: 'In structures' })).toBeVisible();

	await page.context().grantPermissions(['clipboard-read', 'clipboard-write']);
	await tool.getByLabel('Target pool').fill('dungeon:halls');
	await tool.getByLabel('Target hole').fill('dungeon:hall_entry');
	await tool.getByLabel('Name', { exact: true }).fill('dungeon:room_exit');
	await tool.getByRole('button', { name: 'Copy give command' }).click();
	expect(await page.evaluate(() => navigator.clipboard.readText())).toMatch(/^\/give @s jigsaw\[block_entity_data=\{id:"jigsaw",name:"dungeon:room_exit",target:"dungeon:hall_entry",pool:"dungeon:halls"/);
	await expect(tool.getByRole('heading', { name: 'Recent' })).toBeVisible();
	await expect(tool.getByLabel('Copy the give command for dungeon:room_exit')).toHaveCount(1);

	// Recent survives a reload.
	await page.reload();
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('tab', { name: 'Jigsaw' }).click();
	await expect(tool.getByLabel('Copy the give command for dungeon:room_exit')).toHaveCount(1);

	// A hole no JigsawHole declares is declared next to the others.
	await tool.getByLabel('Name', { exact: true }).fill('dungeon:side_door');
	await expect(tool.getByText('Not a JigsawHole.')).toBeVisible();
	await tool.getByRole('button', { name: 'Declare it' }).click();
	await expect(page.locator('.cm-line', { hasText: 'JigsawHole side_door' })).toBeVisible();
	await expect(status(page)).toHaveText(/No problems/);
	await expect(tool.getByText('Not a JigsawHole.')).toBeHidden();
	expect(errors).toEqual([]);
});
