<script lang="ts">
	// The Jigsaw view's page: the shared Jigsaw tool (packages/preview),
	// fed by the extension (src/jigsaw.ts), which carries out what it asks.
	// The form and the last state are kept across reloads of the view.
	import JigsawTool from '@mms/preview/jigsaw/JigsawTool.svelte';
	import type { JigsawConfig } from '@mms/preview/jigsaw/give';
	import type { FromJigsaw, JigsawState, ToJigsaw } from '../src/messages';

	type Saved = { state?: JigsawState; form?: JigsawConfig };
	const vscode = acquireVsCodeApi<Saved>();
	const saved = vscode.getState();
	const blank = (): JigsawConfig => ({ label: '', pool: '', target: '', name: '', finalState: 'minecraft:air', joint: 'rollable' });
	let s = $state<JigsawState | undefined>(saved?.state);
	let form = $state<JigsawConfig>(saved?.form ?? blank());
	$effect(() => vscode.setState({ state: $state.snapshot(s), form: $state.snapshot(form) }));

	window.addEventListener('message', (e: MessageEvent<ToJigsaw>) => {
		if (e.data.type === 'state') s = e.data.state;
	});
	const send = (m: FromJigsaw) => vscode.postMessage(m);
	send({ type: 'ready' });
</script>

<JigsawTool state={s} bind:form {send} revealFileLabel="Show file in Explorer" />

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
