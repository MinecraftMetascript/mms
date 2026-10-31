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
	await expect(page.locator('.cm-content')).toContainText('name = diamond');
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
