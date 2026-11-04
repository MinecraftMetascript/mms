<script lang="ts">
	// The function preview: runs the pack's load tags and then the function
	// in cmdlib's command simulator, and shows chat, scores, storage and
	// entities. Ticks step the tick tag and /schedule. Edits rerun it from a
	// fresh world, then replay the same number of ticks.
	import { textColors } from '../book/text';
	import type { OutputFile } from '../types';
	import { Run, runnable, type LogLine, type RunSnapshot } from './run';

	let { output, path }: { output: OutputFile[]; path: string } = $props();

	const id = $derived(runnable(path)?.id ?? '');
	let ticks = $state(0);
	let showAll = $state(false);
	let showTemps = $state(false);

	// Rerun from a fresh world whenever the pack, the function or the tick
	// count changes. The simulator is deterministic, so replaying N ticks is
	// the same as stepping to N, and an edit lands on the same tick.
	const snap = $derived.by((): RunSnapshot => {
		const r = new Run(output);
		r.start(id);
		r.tick(ticks);
		return r.snapshot();
	});

	// The game's own lines are chat, titles, errors and warnings; "data" and
	// "info" are cmdlib's echo of each command's result.
	const quiet = new Set(['data', 'info', 'success']);
	const log = $derived((snap?.log ?? []).filter((l) => showAll || !quiet.has(l.kind)));
	// MMS's own holders: constants (#c10, #cf) and expression temporaries
	// (#fn._t0, #fn._f1). Locals (#fn.name) and globals ($name) stay.
	const temp = /^#(c-?\d+|cf|.*\._[a-z]+\d+)$/;
	const scores = $derived((snap?.scores ?? []).filter((r) => showTemps || !temp.test(r.key)));
	const hidden = $derived((snap?.scores.length ?? 0) - scores.length);
	const problems = $derived((snap?.log ?? []).filter((l) => l.kind === 'error' || l.kind === 'warning').length);
	const kindClass: Record<string, string> = {
		error: 'text-[#ff6b6b]',
		warning: 'text-[#ffcf5c]',
		command: 'text-[#9aa4b2]',
		system: 'text-[#9aa4b2]',
		data: 'text-[#9aa4b2]',
		info: 'text-[#9aa4b2]',
		success: 'text-[#8fd18f]'
	};
	const color = (c?: string) => (c ? (textColors[c] ?? (c.startsWith('#') ? c : undefined)) : undefined);
	const style = (r: LogLine['runs'][number]) =>
		[
			color(r.color) ? `color:${color(r.color)}` : '',
			r.bold ? 'font-weight:700' : '',
			r.italic ? 'font-style:italic' : '',
			r.underlined || r.strikethrough ? `text-decoration:${r.underlined ? 'underline' : ''} ${r.strikethrough ? 'line-through' : ''}` : ''
		]
			.filter(Boolean)
			.join(';');
	const mc = /^(minecraft:)/;
</script>

<div class="flex h-full flex-col text-xs" aria-label="Function preview">
	<div class="flex flex-wrap items-center gap-2 border-b border-line px-3 py-2">
		<span class="font-semibold">Run</span>
		<span class="font-mono text-muted">{id}</span>
		<span class="ml-auto text-muted" title="Game ticks run after the function">tick {snap?.tick ?? 0}</span>
		<button class="rounded border border-line px-2 py-0.5 hover:bg-hover" onclick={() => (ticks += 1)} title="Run the tick tag and due schedules once">+1 tick</button>
		<button class="rounded border border-line px-2 py-0.5 hover:bg-hover" onclick={() => (ticks += 20)} title="One second of ticks">+20</button>
		<button class="rounded border border-line px-2 py-0.5 hover:bg-hover disabled:opacity-50" disabled={ticks === 0} onclick={() => (ticks = 0)} title="Back to just after the function ran">Reset</button>
	</div>
	<div class="min-h-0 flex-1 overflow-auto p-3">
		{#if snap}
			<section class="mb-4">
				<h3 class="mb-1 flex items-center gap-2 font-semibold">
					Chat
					{#if problems}<span class="font-normal text-muted">· {problems} {problems === 1 ? 'problem' : 'problems'}</span>{/if}
					<label class="ml-auto flex items-center gap-1 font-normal text-muted">
						<input type="checkbox" bind:checked={showAll} /> command results
					</label>
				</h3>
				<ol class="max-h-72 overflow-auto rounded bg-[#1d1f24] px-2 py-1.5 font-mono leading-5 text-white" aria-label="Chat">
					{#each log as l, i (i)}
						<li class={kindClass[l.kind] ?? ''}>
							{#each l.runs as r, j (j)}<span style={style(r)}>{r.text}</span>{/each}
						</li>
					{:else}
						<li class="text-[#9aa4b2]">Nothing was said.</li>
					{/each}
				</ol>
			</section>

			{#if snap.scores.length}
				<section class="mb-4">
					<h3 class="mb-1 flex items-center gap-2 font-semibold">
						Scores
						<label class="ml-auto flex items-center gap-1 font-normal text-muted" title="Constants and expression temporaries MMS writes">
							<input type="checkbox" bind:checked={showTemps} /> MMS temporaries{hidden ? ` (${hidden})` : ''}
						</label>
					</h3>
					<div class="overflow-auto">
						<table class="border-collapse font-mono">
							<thead>
								<tr>
									<th class="border-b border-line pr-4 text-left font-normal text-muted">holder</th>
									{#each snap.objectives as o (o)}<th class="border-b border-line px-2 text-right font-normal text-muted">{o}</th>{/each}
								</tr>
							</thead>
							<tbody>
								{#each scores as row (row.key)}
									<tr class="hover:bg-hover">
										<td class="pr-4">{row.holder}</td>
										{#each snap.objectives as o (o)}<td class="px-2 text-right">{row.values[o] ?? ''}</td>{/each}
									</tr>
								{/each}
							</tbody>
						</table>
					</div>
				</section>
			{/if}

			{#if snap.storage.length}
				<section class="mb-4">
					<h3 class="mb-1 font-semibold">Storage</h3>
					{#each snap.storage as s (s.id)}
						<h4 class="mt-1 font-mono text-muted">{s.id}</h4>
						<pre class="overflow-auto rounded border border-line p-2 font-mono">{s.snbt}</pre>
					{/each}
				</section>
			{/if}

			<section class="mb-4">
				<h3 class="mb-1 font-semibold">Entities</h3>
				<ul>
					{#each snap.entities as e, i (i)}
						<li class="flex gap-3 font-mono">
							<span class="w-24 truncate">{e.type}</span>
							<span class="w-24 truncate">{e.name}</span>
							<span class="w-28 text-muted">{e.pos}</span>
							<span class="truncate text-muted">{e.tags.join(', ')}</span>
						</li>
					{/each}
				</ul>
			</section>

			{#if snap.scheduled.length}
				<section class="mb-4">
					<h3 class="mb-1 font-semibold">Scheduled</h3>
					<ul class="font-mono">
						{#each snap.scheduled as s, i (i)}<li>{s.id.replace(mc, '')} <span class="text-muted">at tick {s.due}</span></li>{/each}
					</ul>
				</section>
			{/if}

			<p class="text-muted">
				Simulated with <a class="text-accent hover:underline" href="https://github.com/Gameboy612/cmdlib" target="_blank" rel="noreferrer">cmdlib</a>: the
				load tags run, then this function as the player Dev at 0 64 0. Entities don't move or fight, and commands such as bossbar, particle and
				loot are skipped with a warning.
			</p>
		{/if}
	</div>
</div>
