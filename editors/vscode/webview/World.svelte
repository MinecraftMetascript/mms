<script lang="ts">
	// The page of two side bar views, chosen by the body's data-page: World
	// and Structures, both drawn by the shared WorldTool. The extension
	// (src/world.ts) owns the state and sends all of it to both after each
	// change.
	import WorldTool from '@mms/preview/world/WorldTool.svelte';
	import type { WorldFeatures } from '@mms/preview/world/tool';
	import type { FromWorld, ToWorld, WorldState } from '../src/messages';

	const page = document.body.dataset.page === 'structures' ? 'structures' : 'world';
	const vscode = acquireVsCodeApi<WorldState>();
	let s = $state<WorldState | undefined>(vscode.getState());

	window.addEventListener('message', (e: MessageEvent<ToWorld>) => {
		if (e.data.type === 'state') {
			s = e.data.state;
			vscode.setState(e.data.state);
		}
	});
	const send = (m: FromWorld) => vscode.postMessage(m);
	send({ type: 'ready' });

	const features: WorldFeatures = {
		revealFile: 'Show file in Explorer',
		revealWorld: 'Open datapacks folder',
		revealStructureWorld: 'Open generated folder',
		buildOnSave: 'Build on save',
		showOutput: true,
		deleteFromWorld: true,
		chooseWorldHint: "Worlds come from the Minecraft launcher's and Prism Launcher's saves folders, or browse to any world folder."
	};
</script>

<WorldTool {page} state={s} {send} {features} />

<style>
	:global(body) {
		margin: 0;
		padding: 0 0 12px;
		background: var(--vscode-sideBar-background);
		color: var(--vscode-foreground);
		font-family: var(--vscode-font-family);
		font-size: var(--vscode-font-size);
	}
</style>
