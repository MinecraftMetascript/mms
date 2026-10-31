<script lang="ts">
	import { fmtRange, label, pct, rng } from './common';
	import { rollLoot, type LootPoolInfo } from './loot';

	let { pools }: { pools: LootPoolInfo[] } = $props();
	let seed = $state(1);
	let roll = $state<{ name: string; count: number }[] | null>(null);
	function doRoll() {
		roll = rollLoot(pools, rng(seed));
		seed++;
	}
	$effect(() => {
		void pools;
		roll = null;
	});
	const name = (n: string, kind: string) => (kind === 'item' || kind === 'tag' || kind === 'loot_table' ? label(n) : n);
</script>

<div class="flex flex-col gap-4 text-xs">
	{#if !pools.length}
		<p class="text-muted">This loot table has no pools, so it drops nothing.</p>
	{/if}
	{#each pools as pool, i (i)}
		<section aria-label="Pool {i + 1}">
			<h3 class="mb-1 font-semibold">
				Pool {i + 1}
				<span class="font-normal text-muted">· {fmtRange(pool.rolls)} {pool.rolls?.max === 1 ? 'roll' : 'rolls'}{pool.bonus ? ' + luck' : ''}{pool.conditional ? ' · conditional' : ''}{pool.functions.length ? ` · ${pool.functions.join(', ')}` : ''}</span>
			</h3>
			<table class="w-full border-collapse">
				<thead class="text-left text-muted">
					<tr><th class="py-0.5 font-normal">Entry</th><th class="font-normal">Chance</th><th class="font-normal">Count</th><th class="font-normal" title="Average items per table roll">Avg.</th></tr>
				</thead>
				<tbody>
					{#each pool.rows as r, j (j)}
						<tr class="border-t border-line">
							<td class="py-1 pr-2">
								<span class={r.kind === 'tag' ? 'text-accent' : r.kind === 'empty' ? 'text-muted' : ''} title={r.name}>{name(r.name, r.kind)}</span>
								{#if r.conditional}<span class="ml-1 rounded bg-hover px-1 text-[10px] text-muted" title="Only drops when its conditions pass">if</span>{/if}
								{#if r.via}<span class="ml-1 text-[10px] text-muted">({r.via})</span>{/if}
								{#if r.functions.length}<div class="text-[10px] text-muted">{r.functions.join(', ')}</div>{/if}
							</td>
							<td class="pr-2">{pct(r.chance)}</td>
							<td class="pr-2">{fmtRange(r.count)}</td>
							<td>{r.expected === null ? '?' : r.expected.toFixed(2)}</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</section>
	{/each}
	{#if pools.length}
		<div class="flex flex-col gap-2 border-t border-line pt-3">
			<div class="flex items-center gap-2">
				<button class="rounded bg-accent px-3 py-1 font-medium text-white" onclick={doRoll}>Roll</button>
				<span class="text-muted">Conditions are assumed to pass.</span>
			</div>
			{#if roll}
				<ul aria-label="Roll result" class="flex flex-wrap gap-1">
					{#each roll as d (d.name)}<li class="rounded border border-line px-2 py-0.5">{d.count} × {label(d.name)}</li>{:else}<li class="text-muted">Nothing dropped.</li>{/each}
				</ul>
			{/if}
		</div>
	{/if}
</div>
