<script lang="ts">
	// The Jigsaw tab: the Jigsaw tool VS Code shows in its side bar
	// (packages/preview), fed from the web engine's pack. Copied blocks are
	// kept per project in this browser; the form stays while the tab is
	// open.
	import JigsawTool from '@mms/preview/jigsaw/JigsawTool.svelte';
	import { sameBlock, type JigsawConfig } from '@mms/preview/jigsaw/give';
	import { holePlacement, jigsawToolState, RECENT_MAX, type JigsawAction, type JigsawRecent } from '@mms/preview/jigsaw/tool';
	import type { Engine } from '../engine/engine.svelte';

	let {
		engine,
		projectId,
		projectName,
		ns,
		onreveal,
		onpiece,
		oninsert
	}: {
		engine: Engine;
		projectId: string;
		projectName: string;
		ns: string;
		/** Shows a place in the project's source. */
		onreveal: (uri: string, range: { start: { line: number; character: number }; end: { line: number; character: number } }) => void;
		/** Shows a structure file of the project, by its path in the pack. */
		onpiece: (packPath: string) => void;
		/** Inserts text before a 0-based line of a source file (or at its end), and selects `select` in it. */
		oninsert: (uri: string, line: number | undefined, text: string, select: string) => void;
	} = $props();

	const key = $derived(`mms-jigsaw-recent:${projectId}`);
	let recent = $state<JigsawRecent[]>([]);
	$effect(() => {
		try {
			recent = JSON.parse(localStorage.getItem(key) ?? '[]');
		} catch {
			recent = [];
		}
	});
	function setRecent(list: JigsawRecent[]) {
		recent = list.slice(0, RECENT_MAX);
		try {
			localStorage.setItem(key, JSON.stringify(recent));
		} catch {
			/* storage unavailable: Recent lasts for this visit */
		}
	}

	let form = $state<JigsawConfig>({ label: '', pool: '', target: '', name: '', finalState: 'minecraft:air', joint: 'rollable' });
	const toolState = $derived(
		jigsawToolState({
			project: { name: projectName, ns },
			holes: engine.outputHoles,
			files: engine.output,
			assets: engine.outputAssets,
			recent,
			hasErrors: engine.hasErrors
		})
	);

	let note = $state('');
	async function send(a: JigsawAction) {
		switch (a.type) {
			case 'copy':
				try {
					await navigator.clipboard.writeText(a.command);
					note = 'Copied the give command; paste it in the game (creative, as an operator).';
				} catch {
					note = 'The browser did not allow copying; select the command and copy it.';
				}
				setRecent([{ config: a.config, at: Date.now() }, ...recent.filter((r) => !sameBlock(r.config, a.config, ns))]);
				return;
			case 'forget':
				return setRecent(recent.filter((r) => !sameBlock(r.config, a.config, ns) || r.config.label !== a.config.label));
			case 'forgetAll':
				return setRecent([]);
			case 'declare': {
				const at = holePlacement(a.id, ns, engine.outputHoles, engine.output);
				if (!at) {
					note = `Add \`JigsawHole ${a.id}\` to one of the project's .mms files to declare it.`;
					return;
				}
				oninsert(at.uri, at.line, at.text, at.name);
				return;
			}
			case 'revealHole': {
				const h = engine.outputHoles.find((h) => h.id === a.id);
				if (h) onreveal(h.uri, h.range);
				return;
			}
			case 'revealFile': {
				const asset = engine.outputAssets.find((x) => x.src === a.src);
				if (asset) onpiece(asset.path);
				return;
			}
		}
	}
</script>

<div class="h-full overflow-auto text-sm" aria-label="Jigsaw">
	<JigsawTool state={toolState} bind:form {send} revealFileLabel="Preview the piece" />
	{#if note}<p class="px-3 py-2 text-xs text-muted" role="status">{note}</p>{/if}
</div>
