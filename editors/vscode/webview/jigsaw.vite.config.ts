// Builds the Jigsaw view's page (webview/Jigsaw.svelte) into out/jigsaw: one
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
		outDir: here('../out/jigsaw'),
		emptyOutDir: true,
		target: 'es2022',
		cssCodeSplit: false,
		rollupOptions: {
			input: here('jigsaw.ts'),
			output: {
				entryFileNames: 'jigsaw.js',
				assetFileNames: 'jigsaw[extname]',
				inlineDynamicImports: true
			}
		}
	}
});
