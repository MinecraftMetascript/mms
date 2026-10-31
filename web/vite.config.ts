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
	worker: { format: 'es' }
});
