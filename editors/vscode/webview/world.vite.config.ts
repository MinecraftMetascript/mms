// Builds the World view's page (webview/World.svelte) into out/world: one
// script and one stylesheet. It is a separate build from the preview panel
// (vite.config.ts) because each page inlines its imports into one file.
import { svelte } from '@sveltejs/vite-plugin-svelte';
import { fileURLToPath } from 'node:url';
import { defineConfig } from 'vite';

const here = (p: string) => fileURLToPath(new URL(p, import.meta.url));

export default defineConfig({
	root: here('.'),
	plugins: [svelte()],
	build: {
		outDir: here('../out/world'),
		emptyOutDir: true,
		target: 'es2022',
		cssCodeSplit: false,
		rollupOptions: {
			input: here('world.ts'),
			output: {
				entryFileNames: 'world.js',
				assetFileNames: 'world[extname]',
				inlineDynamicImports: true
			}
		}
	}
});
