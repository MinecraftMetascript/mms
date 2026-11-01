import { expect, test, type Page } from '@playwright/test';

const status = (page: Page) => page.getByRole('status');

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
	await replaceText(page, 'Noise n = { base_octave = -5, amplitude_modifiers = [1, 1] }\n');
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
	await page.getByRole('button', { name: /Project/ }).click();
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
	await page.getByRole('button', { name: /Project/ }).click();
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
	await page.getByRole('button', { name: /Project/ }).click();
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

	// Removing a texture breaks the code that names it.
	await page.getByRole('listitem').filter({ hasText: 'art/hearth.png' }).hover();
	await page.getByRole('button', { name: 'Delete art/hearth.png' }).click();
	await expect(status(page)).toHaveText(/1 error/);
});

test('dimensions, noise settings and surface rules preview as terrain', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Project/ }).click();
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
	await expect(page.getByText(/x -?\d+, y -?\d+ · \w+/)).toBeVisible();

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
	await page.getByRole('button', { name: /Project/ }).click();
	await page.getByRole('button', { name: /badlands/ }).click();
	await page.locator('.cm-line', { hasText: 'SurfaceRule badlands' }).click();
	await expect(page.getByRole('combobox', { name: 'Biome' })).toHaveValue('minecraft:badlands');
	expect(errors).toEqual([]);
});

test('a pinned preview stays put while the cursor moves to another declaration', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	page.on('dialog', (d) => d.accept());
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await page.getByRole('button', { name: /Project/ }).click();
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
	await page.getByRole('button', { name: /Project/ }).click();
	await page.getByRole('button', { name: /Paste JSON/ }).click();
	await page.getByLabel('Kind', { exact: true }).selectOption('LootTable');
	await page.getByLabel('Name', { exact: true }).fill('blocks/ruby_ore');
	await page.getByLabel('JSON', { exact: true }).fill('{"type":"minecraft:block","pools":[{"rolls":1,"entries":[{"type":"minecraft:item","name":"minecraft:diamond"}]}]}');
	await page.getByRole('button', { name: 'Convert' }).click();
	await expect(page.locator('.cm-content')).toContainText('LootTable blocks/ruby_ore = {');
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

test('loot tables, recipes and biomes show in the inspector', async ({ page }) => {
	const errors: string[] = [];
	page.on('pageerror', (e) => errors.push(e.message));
	await page.goto('./');
	await expect(status(page)).toHaveText(/No problems/, { timeout: 30_000 });
	await replaceText(
		page,
		[
			'LootTable chest = { pools = [Pool(2, [Item(name = diamond).Weight(3), Item(name = emerald)])] }',
			'Recipe ruby_block = Shaped(["RRR", "R R", "RRR"], { R = emerald }, emerald_block)',
			'Biome warm = { temperature = 0.8, downfall = 0.4, has_precipitation = true, effects = { water_color = "#3f76e4" }, features = [], carvers = [] }',
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

	await page.locator('.cm-line').nth(2).click();
	await expect(inspector).toContainText('Temperature');
	await expect(inspector).toContainText('#3f76e4');
	expect(errors).toEqual([]);
});

test('the docs render guides and the reference', async ({ page }) => {
	await page.goto('./');
	await page.getByRole('link', { name: 'Docs' }).click();
	await expect(page.getByRole('heading', { level: 1 })).toHaveText('Getting started');

	// A guide, and a link from one guide to a section of another.
	const nav = page.getByRole('navigation', { name: 'Docs' });
	await nav.getByRole('link', { name: 'Tutorial: a custom dimension' }).click();
	await expect(page.getByRole('heading', { level: 1 })).toHaveText('Tutorial: a custom dimension');
	await expect(page.getByRole('navigation', { name: 'On this page' }).getByRole('link').first()).toBeVisible();
	await page.getByRole('main').getByRole('link', { name: 'Writing resources' }).click();
	await expect(page).toHaveURL(/\/docs\/resources#density-functions$/);
	await expect(page.locator('#density-functions')).toBeVisible();

	// The generated reference, filtered.
	await nav.getByRole('link', { name: 'Reference' }).click();
	await expect(page.getByRole('cell', { name: 'DensityFunction', exact: true })).toBeVisible();
	const filter = page.getByRole('searchbox', { name: 'Filter helpers' });
	const toc = page.getByRole('navigation', { name: 'On this page' });
	await expect(toc.getByRole('link', { name: 'Loot function', exact: true })).toBeVisible();
	await filter.fill('gradient');
	// The page sections follow the filter.
	await expect(toc.getByRole('link', { name: 'Loot function', exact: true })).toHaveCount(0);
	await expect(toc.getByRole('link', { name: 'Density function', exact: true })).toBeVisible();
	await expect(page.getByText(/^Gradient\(axis, /)).toBeVisible();
	await expect(page.getByRole('region', { name: 'Loot function' })).toHaveCount(0);

	// A type in a signature links to its entry, even from a filtered list.
	const gradient = page.locator('#density_function-Gradient');
	await gradient.getByRole('link', { name: 'Axis', exact: true }).click();
	await expect(page).toHaveURL(/#type-util-direction-Axis$/);
	await expect(filter).toHaveValue('');
	await expect(page.locator('#type-util-direction-Axis')).toBeInViewport();
});

test('the Project menu closes on outside click and Escape, and Format formats', async ({ page }) => {
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
	await expect(hover.locator('a[href*="reference#material_condition-StoneDepth"]')).toHaveCount(1);
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
		await page.getByRole('button', { name: /Project/ }).click();
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
