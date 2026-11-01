import adapter from '@sveltejs/adapter-static';
import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [
		tailwindcss(),
		sveltekit({
			compilerOptions: {
				runes: ({ filename }) =>
					filename.split(/[/\\]/).includes('node_modules') ? undefined : true
			},
			// A fully static site for GitHub Pages; BASE_PATH is the repository
			// path (/mms) when deployed there.
			adapter: adapter({ fallback: '404.html' }),
			paths: { base: (process.env.BASE_PATH ?? '') as '' | `/${string}` }
		})
	],
	// The previews live in packages/preview (shared with the VS Code
	// extension); the host's own Svelte and Deepslate serve it.
	resolve: { dedupe: ['svelte', 'deepslate'] },
	server: { fs: { allow: ['../packages'] } },
	worker: { format: 'es' }
});
