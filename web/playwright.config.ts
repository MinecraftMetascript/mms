import { defineConfig } from '@playwright/test';

// Smoke tests run against the static build served by `vite preview`.
// PW_CHROMIUM points at a preinstalled browser when there is one.
export default defineConfig({
	testDir: 'tests',
	timeout: 60_000,
	use: {
		baseURL: `http://localhost:4173${process.env.BASE_PATH ?? ''}/`,
		viewport: { width: 1400, height: 850 },
		launchOptions: process.env.PW_CHROMIUM ? { executablePath: process.env.PW_CHROMIUM } : {}
	},
	webServer: {
		command: 'npx vite preview --port 4173 --strictPort',
		port: 4173,
		reuseExistingServer: !process.env.CI
	}
});
