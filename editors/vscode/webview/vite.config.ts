// Builds the preview panel's page (webview/) into out/webview: one script
// and one stylesheet, with the preview worker inlined, since a webview can
// only start workers from blob URLs.
import { svelte } from '@sveltejs/vite-plugin-svelte';
import tailwindcss from '@tailwindcss/vite';
import { fileURLToPath } from 'node:url';
import { defineConfig } from 'vite';

const here = (p: string) => fileURLToPath(new URL(p, import.meta.url));

export default defineConfig({
	root: here('.'),
	plugins: [tailwindcss(), svelte()],
	// packages/preview is linked in; the host's own Svelte and Deepslate serve it.
	resolve: { dedupe: ['svelte', 'deepslate'] },
	worker: { format: 'es' },
	build: {
		outDir: here('../out/webview'),
		emptyOutDir: true,
		target: 'es2022',
		cssCodeSplit: false,
		rollupOptions: {
			input: here('main.ts'),
			output: {
				entryFileNames: 'preview.js',
				assetFileNames: 'preview[extname]',
				inlineDynamicImports: true
			}
		}
	}
});
