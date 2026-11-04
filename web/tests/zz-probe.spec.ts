import { test, type Page } from '@playwright/test';
async function replaceText(page: Page, text: string) {
	await page.locator('.cm-content').click();
	await page.keyboard.press('ControlOrMeta+a');
	await page.keyboard.press('Delete');
	await page.keyboard.insertText(text);
}
const popup = (page: Page) => page.locator('.cm-tooltip-autocomplete');
async function probe(page: Page, name: string, pre: string, typed: string) {
	await replaceText(page, pre);
	await page.keyboard.type(typed, { delay: 60 });
	await page.waitForTimeout(1200);
	const txt = (await popup(page).count()) ? await popup(page).innerText() : '(no popup)';
	console.log(`### ${name}: ${JSON.stringify(typed)} -> ${txt.replace(/\n/g, ' | ').slice(0, 200)}`);
	await page.keyboard.press('Escape');
}
test('probe', async ({ page }) => {
	await page.goto('./');
	await page.getByRole('status').filter({ hasText: /No problems|error|warning/ }).waitFor({ timeout: 30000 });
	await probe(page, 'field', 'Book b = {\n  ', 'ti');
	await probe(page, 'field2', 'Book b = {\n  title = "T"\n  ', 'au');
	await probe(page, 'gen', 'Book b = {\n  generation = ', 't');
	await probe(page, 'kind', '', 'Boo');
	await probe(page, 'text tag', 'Book b = {\n  text = [\n    ', '"Hello {go');
	await probe(page, 'text tag2', 'Book b = {\n  text = [\n    ', '"Hello {gold bo');
	await probe(page, 'text link', 'Book b = {\n  text = [\n    ', '"[x](ru');
	await probe(page, 'text expr', 'Global wave = 1\nBook b = {\n  text = [\n    ', '"${wa');
	await probe(page, 'pages', 'Book b = {\n  pages = [\n    { ', 'hov');
	await probe(page, 'action', 'Book b = {\n  pages = [\n    { text = "a", hover_event = { action = ', 'sh');
	await probe(page, 'booktext', 'Book b = {\n  text = ', 'BookT');
});
