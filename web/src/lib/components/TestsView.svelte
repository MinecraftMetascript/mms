<script lang="ts">
	import type { Diagnostic, TestOutcome } from '../engine/engine.svelte';

	interface Props {
		tests: TestOutcome[];
		error: string;
		onsource: (uri: string, range: Diagnostic['range']) => void;
	}
	let { tests, error, onsource }: Props = $props();
	const failed = $derived(tests.filter((t) => !t.pass).length);
</script>

<div class="flex h-full min-h-0 flex-col text-sm">
	<p class="border-b border-line px-3 py-1.5 text-xs text-muted" aria-live="polite" data-testid="tests-summary">
		{#if error}
			{error}
		{:else if tests.length === 0}
			No tests ran.
		{:else if failed}
			{failed} of {tests.length} {tests.length === 1 ? 'test' : 'tests'} failed
		{:else}
			{tests.length} {tests.length === 1 ? 'test' : 'tests'} passed
		{/if}
	</p>
	<ul class="min-h-0 flex-1 overflow-auto" aria-label="Tests">
		{#each tests as t (t.uri + ' ' + t.name)}
			<li class="border-b border-line px-3 py-1.5">
				<button
					class="flex w-full items-center gap-2 text-left hover:text-accent"
					onclick={() => onsource(t.uri, t.expectRange ?? t.range)}
					aria-label="{t.name}: {t.pass ? 'passed' : 'failed'}"
				>
					<span class={t.pass ? 'text-green-600' : 'text-red-600'} aria-hidden="true">{t.pass ? '✓' : '✗'}</span>
					<span class="font-mono">{t.name}</span>
				</button>
				{#if t.message}
					<p class="mt-0.5 pl-6 text-xs text-ink">{t.message}</p>
				{/if}
				{#each t.said as line, i (i)}
					<p class="pl-6 font-mono text-xs text-muted">{line}</p>
				{/each}
			</li>
		{/each}
	</ul>
</div>
