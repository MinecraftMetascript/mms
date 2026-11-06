<script lang="ts">
	// The Jigsaw view's page: a form for one jigsaw block and the /give
	// command for it, then Recent (blocks copied before) and In structures
	// (the jigsaw blocks of the pack's structure files), both trees like the
	// World view's. Clicking a row loads it into the form; its copy icon
	// copies its command straight away. The extension (src/jigsaw.ts) sends
	// the lists; the form lives here and is kept across reloads.
	import { CHAT_LIMIT, fullId, giveCommand, sameBlock } from '../src/jigsawGive';
	import type { FromJigsaw, JigsawConfig, JigsawExisting, JigsawRecent, JigsawState, ToJigsaw } from '../src/messages';

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

	const ns = $derived(s?.project?.ns);
	const command = $derived(giveCommand(form, ns));
	// An id as the list shows it: the project's namespace and minecraft:empty left out.
	const shown = (id: string) => (id === 'minecraft:empty' ? 'empty' : ns && id.startsWith(`${ns}:`) ? id.slice(ns.length + 1) : id);

	const holes = $derived(new Set(s?.holes.map((h) => h.id) ?? []));
	const pools = $derived(new Set(s?.pools.map((p) => p.id) ?? []));
	// Jigsaw names the pack's pieces use, declared or not.
	const used = $derived([...new Set(s?.existing.map((e) => e.config.name).filter((n) => n !== 'minecraft:empty') ?? [])].sort());
	const pool = $derived(s?.pools.find((p) => p.id === fullId(form.pool, ns)));

	type Option = { value: string; note: string };
	// Name: the declared holes, then names pieces use that are not declared.
	const nameOptions = $derived<Option[]>([
		...(s?.holes ?? []).map((h) => ({ value: h.id, note: h.doc ?? 'JigsawHole' })),
		...used.filter((n) => !holes.has(n)).map((n) => ({ value: n, note: 'in a piece, not a JigsawHole' }))
	]);
	// Target hole: the names the chosen pool's pieces have come first.
	const targetOptions = $derived<Option[]>([
		...(pool?.names ?? []).map((n) => ({ value: n, note: `in ${shown(pool!.id)}` })),
		...nameOptions.filter((o) => !pool?.names.includes(o.value))
	]);
	const poolOptions = $derived<Option[]>([...(s?.pools ?? []).map((p) => ({ value: p.id, note: p.names.length ? p.names.map(shown).join(', ') : 'TemplatePool' })), { value: 'minecraft:empty', note: 'no pool: nothing joins here' }]);

	// What to say under a field whose id the project does not declare.
	function holeNote(value: string): string | null {
		const id = fullId(value, ns);
		if (!value.trim() || id === 'minecraft:empty' || holes.has(id) || !s) return null;
		if (ns && !id.startsWith(`${ns}:`)) return null;
		return id;
	}
	const nameNote = $derived(holeNote(form.name));
	const targetNote = $derived(holeNote(form.target));
	const poolNote = $derived.by(() => {
		const id = fullId(form.pool, ns);
		if (!form.pool.trim() || id === 'minecraft:empty' || pools.has(id) || !s || id.startsWith('minecraft:')) return null;
		return id;
	});
	const targetMissing = $derived(pool && pool.names.length > 0 && fullId(form.target, ns) !== 'minecraft:empty' && !pool.names.includes(fullId(form.target, ns)));

	// Buttons that just copied something say so for a moment.
	let copied = $state<string | null>(null);
	let copyTimer: ReturnType<typeof setTimeout> | undefined;
	function copy(id: string, config: JigsawConfig) {
		send({ type: 'copy', config: $state.snapshot(config), command: giveCommand(config, ns) });
		copied = id;
		clearTimeout(copyTimer);
		copyTimer = setTimeout(() => (copied = null), 1500);
	}
	function load(config: JigsawConfig) {
		form = { ...config };
	}

	// A row's ⋯ menu, as in the World view.
	type MenuItem = { label: string; run: () => void; danger?: boolean } | 'sep';
	let menu = $state<{ x: number; y: number; right: boolean; items: MenuItem[] } | null>(null);
	let menuEl = $state<HTMLDivElement>();
	function openMenu(e: MouseEvent, items: MenuItem[]) {
		e.preventDefault();
		e.stopPropagation();
		const b = e.type === 'contextmenu' ? null : (e.currentTarget as HTMLElement).getBoundingClientRect();
		menu = { x: b ? b.right : e.clientX, y: b ? b.bottom + 2 : e.clientY, right: !!b, items };
	}
	$effect(() => {
		if (!menu || !menuEl) return;
		const r = menuEl.getBoundingClientRect();
		const x = menu.right ? menu.x - r.width : menu.x;
		menuEl.style.left = `${Math.max(4, Math.min(x, window.innerWidth - r.width - 4))}px`;
		menuEl.style.top = `${Math.max(4, menu.y + r.height > window.innerHeight - 4 ? menu.y - r.height - 4 : menu.y)}px`;
		menuEl.querySelector('button')?.focus();
	});
	const closeMenu = () => (menu = null);
	function menuKey(e: KeyboardEvent) {
		if (!menu) return;
		if (e.key === 'Escape') return closeMenu();
		if (e.key !== 'ArrowDown' && e.key !== 'ArrowUp') return;
		e.preventDefault();
		const items = [...(menuEl?.querySelectorAll('button') ?? [])];
		const at = items.indexOf(document.activeElement as HTMLButtonElement);
		items[(at + (e.key === 'ArrowDown' ? 1 : items.length - 1)) % items.length]?.focus();
	}
	function rowKey(e: KeyboardEvent, config: JigsawConfig, items: MenuItem[]) {
		// Keys on the row's own buttons are theirs.
		if (e.target !== e.currentTarget) return;
		if (e.key === 'Enter' || e.key === ' ') {
			e.preventDefault();
			load(config);
		} else if (e.key === 'ContextMenu' || (e.shiftKey && e.key === 'F10')) openMenu(e as unknown as MouseEvent, items);
	}

	const recentMenu = (r: JigsawRecent): MenuItem[] => [
		{ label: 'Load into the form', run: () => load(r.config) },
		{ label: 'Copy give command', run: () => copy(`r:${r.at}`, r.config) },
		'sep',
		{ label: 'Remove from Recent', run: () => send({ type: 'forget', config: $state.snapshot(r.config) }) },
		{ label: 'Clear Recent', danger: true, run: () => send({ type: 'forgetAll' }) }
	];
	// A block placed from a named item keeps the name in the structure
	// file; one without gets the label of a block set up the same way in
	// Recent.
	const labelled = (c: JigsawConfig): JigsawConfig => {
		if (c.label) return c;
		const r = s?.recent.find((r) => r.config.label && sameBlock(r.config, c, ns));
		return r ? { ...c, label: r.config.label } : c;
	};
	// A row of In structures: the blocks of one piece set up the same way,
	// which differ only in where they are.
	type Same = { key: string; config: JigsawConfig; saved: boolean; src: string; at: JigsawExisting[] };
	const existingMenu = (e: Same): MenuItem[] => [
		{ label: 'Load into the form', run: () => load(e.config) },
		{ label: 'Copy give command', run: () => copy(e.key, e.config) },
		'sep',
		{ label: 'Show file in Explorer', run: () => send({ type: 'revealFile', src: e.src }) }
	];

	// In structures: one folder per structure file, duplicates rolled up.
	const pieces = $derived.by(() => {
		const by = new Map<string, Same[]>();
		for (const e of s?.existing ?? []) {
			const rows = by.get(e.id) ?? [];
			const same = rows.find((r) => sameBlock(r.config, e.config, ns));
			if (same) {
				same.at.push(e);
				// A label saved in the file wins over one taken from Recent.
				if (e.config.label && !same.saved) {
					same.config = { ...same.config, label: e.config.label };
					same.saved = true;
				}
			} else rows.push({ key: `${e.src}|${e.pos.join(',')}`, config: { ...labelled(e.config) }, saved: !!e.config.label, src: e.src, at: [e] });
			by.set(e.id, rows);
		}
		return [...by].sort(([a], [b]) => a.localeCompare(b));
	});
	const where = (e: JigsawExisting) => `${e.pos.join(' ')}${e.orientation ? `, facing ${e.orientation.replace('_', ', top ')}` : ''}`;
	let open = $state<Record<string, boolean>>({});

	// "5 min ago", for Recent.
	let now = $state(Date.now());
	setInterval(() => (now = Date.now()), 60_000);
	function ago(ms: number): string {
		const min = Math.round((now - ms) / 60_000);
		if (min < 1) return 'just now';
		if (min < 60) return `${min} min ago`;
		const h = Math.round(min / 60);
		if (h < 24) return `${h} h ago`;
		const d = Math.round(h / 24);
		return d === 1 ? 'yesterday' : `${d} days ago`;
	}
	// A block in one line: its name, then where it leads.
	const leads = (c: JigsawConfig) => (fullId(c.pool, ns) === 'minecraft:empty' ? 'no pool' : `→ ${shown(fullId(c.pool, ns))} · ${shown(fullId(c.target, ns))}`);
	const tip = (c: JigsawConfig) =>
		[`Name: ${fullId(c.name, ns)}`, `Target pool: ${fullId(c.pool, ns)}`, `Target name: ${fullId(c.target, ns)}`, `Turns into: ${c.finalState || 'minecraft:air'}`, `Joint: ${c.joint}`].join('\n');
</script>

{#snippet icon(name: 'copy' | 'more' | 'check')}
	<svg viewBox="0 0 16 16" width="16" height="16" aria-hidden="true" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round">
		{#if name === 'copy'}<rect x="5.5" y="5.5" width="8" height="8" rx="1" /><path d="M10.5 5.5V3.5a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2" />
		{:else if name === 'check'}<path d="M3 8.5l3 3 7-7" />
		{:else}<circle cx="3.5" cy="8" r=".6" fill="currentColor" /><circle cx="8" cy="8" r=".6" fill="currentColor" /><circle cx="12.5" cy="8" r=".6" fill="currentColor" />{/if}
	</svg>
{/snippet}

{#snippet datalist(id: string, options: Option[])}
	<datalist {id}>
		{#each options as o (o.value)}<option value={o.value} label={o.note}></option>{/each}
	</datalist>
{/snippet}

{#snippet undeclared(id: string)}
	<div class="note small">
		<span class="warn-text">Not a JigsawHole.</span>
		<button class="link" title="Add JigsawHole {shown(id)} to the project, next to its other holes" onclick={() => send({ type: 'declare', id })}>Declare it</button>
	</div>
{/snippet}

{#if !s}
	<p class="muted pad">Loading…</p>
{:else if !s.project}
	<p class="pad">No MMS project in this workspace. Open a folder with an <code>mms.json</code>, or create one with <code>mms init</code>.</p>
{:else}
	<section class="form">
		<label class="field">
			<span class="field-label">Label</span>
			<input type="text" bind:value={form.label} placeholder="Name on the item, optional" spellcheck="false" />
		</label>
		<label class="field">
			<span class="field-label">Target pool</span>
			<input type="text" list="jigsaw-pools" bind:value={form.pool} placeholder="minecraft:empty" spellcheck="false" autocomplete="off" />
		</label>
		{#if poolNote}<div class="note small warn-text">No TemplatePool declares {shown(poolNote)}.</div>{/if}
		<label class="field">
			<span class="field-label">Target hole</span>
			<input type="text" list="jigsaw-targets" bind:value={form.target} placeholder="minecraft:empty" spellcheck="false" autocomplete="off" />
		</label>
		{#if targetNote}{@render undeclared(targetNote)}{:else if targetMissing}<div class="note small warn-text">No piece of {shown(pool!.id)} has this hole.</div>{/if}
		<label class="field">
			<span class="field-label">Name</span>
			<input type="text" list="jigsaw-names" bind:value={form.name} placeholder="minecraft:empty" spellcheck="false" autocomplete="off" />
		</label>
		{#if nameNote}{@render undeclared(nameNote)}{/if}
		<label class="field">
			<span class="field-label">Replace with</span>
			<input type="text" bind:value={form.finalState} placeholder="minecraft:air" spellcheck="false" />
		</label>
		<div class="field" role="radiogroup" aria-label="Alignment">
			<span class="field-label">Alignment</span>
			<span class="toggle">
				<button role="radio" aria-checked={form.joint === 'rollable'} class:on={form.joint === 'rollable'} title="Rollable: a piece joined here may turn about the joint. Only used when the block faces up or down." onclick={() => (form.joint = 'rollable')}>Rollable</button>
				<button role="radio" aria-checked={form.joint === 'aligned'} class:on={form.joint === 'aligned'} title="Aligned: a piece joined here keeps the turn of this block's top. Only used when the block faces up or down." onclick={() => (form.joint = 'aligned')}>Aligned</button>
			</span>
		</div>
		<button class="primary wide" title="Copy the command, and keep this block in Recent" onclick={() => copy('form', form)}>
			{copied === 'form' ? '✓ Copied' : 'Copy give command'}
		</button>
		<code class="command" title="Paste it in the game in creative mode, as an operator. The block faces the way it is placed.">{command}</code>
		{#if command.length > CHAT_LIMIT}
			<div class="note small warn-text">{command.length} characters: too long for the chat box ({CHAT_LIMIT}). Shorten the label, or run it from a command block.</div>
		{/if}
		{#if s.hasErrors && !s.holes.length}
			<div class="note small muted">The lists fill in once the project has no errors.</div>
		{/if}
		{@render datalist('jigsaw-pools', poolOptions)}
		{@render datalist('jigsaw-targets', targetOptions)}
		{@render datalist('jigsaw-names', nameOptions)}
	</section>

	<section>
		<div class="head">
			<h2 class="grow">Recent</h2>
		</div>
		{#if !s.recent.length}
			<p class="muted small">Blocks you copy show up here, to copy again.</p>
		{:else}
			<ul class="tree" role="tree">
				{#each s.recent as r (r.at)}
					{@const items = recentMenu(r)}
					<li>
						<div class="item" style:--depth={0} title={`${r.config.label ? `${r.config.label}\n` : ''}${tip(r.config)}\nCopied ${new Date(r.at).toLocaleString()}`} tabindex="0" role="treeitem" aria-selected="false" onclick={() => load(r.config)} onkeydown={(e) => rowKey(e, r.config, items)} oncontextmenu={(e) => openMenu(e, items)}>
							<span class="twistie"></span>
							<span class="label">{r.config.label || shown(fullId(r.config.name, ns))}</span>
							<span class="desc">{r.config.label ? `${shown(fullId(r.config.name, ns))} ` : ''}{leads(r.config)} · {ago(r.at)}</span>
							<span class="actions-inline">
								<button class="icon-btn" title="Copy give command" aria-label="Copy the give command for {r.config.label || r.config.name}" onclick={(e) => { e.stopPropagation(); copy(`r:${r.at}`, r.config); }}>{@render icon(copied === `r:${r.at}` ? 'check' : 'copy')}</button>
								<button class="icon-btn" title="More actions" aria-label="More actions" aria-haspopup="menu" onclick={(e) => openMenu(e, items)}>{@render icon('more')}</button>
							</span>
						</div>
					</li>
				{/each}
			</ul>
		{/if}
	</section>

	<section>
		<div class="head">
			<h2 class="grow">In structures</h2>
		</div>
		{#if !pieces.length}
			<p class="muted small">No jigsaw blocks in the pack's structure files yet. Pieces that StructureFile(...) calls use are listed here.</p>
		{:else}
			<ul class="tree" role="tree">
				{#each pieces as [id, list] (id)}
					<li>
						<details open={open[id] ?? false} ontoggle={(e) => (open[id] = e.currentTarget.open)}>
							<summary class="item" style:--depth={0} title={`${id}\n${list[0].src}`}>
								<span class="twistie" aria-hidden="true"></span>
								<span class="label folder">{shown(id)}</span>
								<span class="desc" title="{list.reduce((n, r) => n + r.at.length, 0)} jigsaw blocks">⧉{list.reduce((n, r) => n + r.at.length, 0)}</span>
							</summary>
							<ul class="tree">
								{#each list as e (e.key)}
									{@const items = existingMenu(e)}
									<li>
										<div class="item" style:--depth={1} title={`${e.config.label ? `${e.config.label}\n` : ''}${tip(e.config)}\n${e.at.length === 1 ? `At ${where(e.at[0])}` : `${e.at.length} blocks set up this way, at:\n${e.at.map((a) => `  ${where(a)}`).join('\n')}`}\nClick to load it into the form.`} tabindex="0" role="treeitem" aria-selected="false" onclick={() => load(e.config)} onkeydown={(ev) => rowKey(ev, e.config, items)} oncontextmenu={(ev) => openMenu(ev, items)}>
											<span class="twistie"></span>
											<span class="label">{e.config.label || shown(fullId(e.config.name, ns))}</span>
											<span class="desc">{e.config.label ? `${shown(fullId(e.config.name, ns))} ` : ''}{leads(e.config)}</span>
											{#if e.at.length > 1}<span class="count" title="{e.at.length} jigsaw blocks set up this way">⧉{e.at.length}</span>{/if}
											<span class="actions-inline">
												<button class="icon-btn" title="Copy give command" aria-label="Copy the give command for this block" onclick={(ev) => { ev.stopPropagation(); copy(e.key, e.config); }}>{@render icon(copied === e.key ? 'check' : 'copy')}</button>
												<button class="icon-btn" title="More actions" aria-label="More actions" aria-haspopup="menu" onclick={(ev) => openMenu(ev, items)}>{@render icon('more')}</button>
											</span>
										</div>
									</li>
								{/each}
							</ul>
						</details>
					</li>
				{/each}
			</ul>
		{/if}
	</section>
{/if}

<svelte:window onclick={closeMenu} onblur={closeMenu} onresize={closeMenu} onkeydown={menuKey} />
{#if menu}
	<div class="menu" role="menu" tabindex="-1" bind:this={menuEl} onclick={(e) => e.stopPropagation()} onkeydown={(e) => { e.stopPropagation(); menuKey(e); }}>
		{#each menu.items as it, i (i)}
			{#if it === 'sep'}
				<div class="sep" role="separator"></div>
			{:else}
				<button role="menuitem" class:danger={it.danger} onclick={() => { closeMenu(); it.run(); }}>{it.label}</button>
			{/if}
		{/each}
	</div>
{/if}

<style>
	:global(body) {
		margin: 0;
		padding: 0 0 12px;
		background: var(--vscode-sideBar-background);
		color: var(--vscode-foreground);
		font-family: var(--vscode-font-family);
		font-size: var(--vscode-font-size);
	}
	section {
		padding: 8px 12px;
		border-bottom: 1px solid var(--vscode-sideBarSectionHeader-border, var(--vscode-widget-border, transparent));
	}
	section:last-child {
		border-bottom: 0;
	}
	p {
		margin: 4px 0 8px;
		line-height: 1.4;
	}
	.pad {
		padding: 8px 12px;
	}
	h2 {
		margin: 0 0 4px;
		font-size: 11px;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 0.04em;
	}
	ul {
		list-style: none;
		margin: 0;
		padding: 0;
		display: flex;
		flex-direction: column;
	}
	code {
		font-family: var(--vscode-editor-font-family);
		font-size: 0.95em;
	}
	.head {
		display: flex;
		align-items: center;
		gap: 6px;
	}
	.grow {
		flex: 1;
		min-width: 0;
	}
	.muted {
		color: var(--vscode-descriptionForeground);
	}
	.small {
		font-size: 12px;
		line-height: 1.4;
	}

	/* The form: label column and input, like VS Code's settings rows. */
	.form {
		display: flex;
		flex-direction: column;
		gap: 4px;
	}
	.field {
		display: flex;
		align-items: center;
		gap: 8px;
	}
	.field-label {
		flex: 0 0 84px;
		color: var(--vscode-foreground);
	}
	.field input {
		flex: 1;
		min-width: 0;
		font: inherit;
		padding: 2px 6px;
		height: 24px;
		box-sizing: border-box;
		color: var(--vscode-input-foreground);
		background: var(--vscode-input-background);
		border: 1px solid var(--vscode-input-border, transparent);
		border-radius: 2px;
	}
	.field input:focus {
		outline: 1px solid var(--vscode-focusBorder);
		outline-offset: -1px;
	}
	.field input::placeholder {
		color: var(--vscode-input-placeholderForeground);
	}
	.note {
		margin: -1px 0 2px 92px;
	}
	.toggle {
		display: inline-flex;
		border: 1px solid var(--vscode-input-border, var(--vscode-widget-border, transparent));
		border-radius: 2px;
		overflow: hidden;
	}
	.toggle button {
		border: 0;
		border-radius: 0;
		padding: 1px 10px;
		background: var(--vscode-input-background);
		color: var(--vscode-foreground);
	}
	.toggle button.on {
		background: var(--vscode-button-secondaryBackground);
		color: var(--vscode-button-secondaryForeground);
		font-weight: 600;
	}
	.command {
		display: block;
		margin-top: 2px;
		padding: 4px 6px;
		font-size: 11px;
		word-break: break-all;
		color: var(--vscode-descriptionForeground);
		background: var(--vscode-textCodeBlock-background, var(--vscode-editorWidget-background));
		border-radius: 2px;
		user-select: all;
	}

	button {
		font: inherit;
		cursor: pointer;
		border-radius: 2px;
		padding: 4px 10px;
		border: 1px solid var(--vscode-button-border, transparent);
		line-height: 18px;
	}
	button:focus-visible {
		outline: 1px solid var(--vscode-focusBorder);
		outline-offset: 2px;
	}
	.primary {
		background: var(--vscode-button-background);
		color: var(--vscode-button-foreground);
	}
	.primary:hover {
		background: var(--vscode-button-hoverBackground);
	}
	.wide {
		width: 100%;
		margin-top: 6px;
		padding: 4px 10px;
	}
	.link {
		background: none;
		border: 0;
		padding: 0;
		color: var(--vscode-textLink-foreground);
		text-decoration: underline;
		font-size: inherit;
		line-height: inherit;
	}
	.link:hover {
		color: var(--vscode-textLink-activeForeground);
	}
	.icon-btn {
		background: none;
		border: 0;
		padding: 0 4px;
		color: var(--vscode-icon-foreground);
	}
	.icon-btn:hover {
		background: var(--vscode-toolbar-hoverBackground);
	}

	/* Trees like the Explorer's, as in the World view. */
	ul.tree {
		gap: 0;
		margin: 0 -12px;
	}
	ul.tree ul.tree {
		margin: 0;
	}
	.item {
		display: flex;
		align-items: center;
		gap: 6px;
		height: 22px;
		padding: 0 4px 0 calc(8px + var(--depth, 0) * 12px);
		cursor: pointer;
		white-space: nowrap;
		outline: none;
	}
	.item:hover {
		background: var(--vscode-list-hoverBackground);
	}
	.item:focus-visible {
		outline: 1px solid var(--vscode-list-focusOutline, var(--vscode-focusBorder));
		outline-offset: -1px;
	}
	.twistie {
		flex: 0 0 12px;
		text-align: center;
		color: var(--vscode-icon-foreground);
	}
	summary.item {
		list-style: none;
	}
	summary.item::-webkit-details-marker {
		display: none;
	}
	summary.item .twistie::before {
		content: '›';
		display: inline-block;
		font-size: 14px;
		line-height: 1;
		transition: transform 0.1s;
	}
	details[open] > summary.item .twistie::before {
		transform: rotate(90deg);
	}
	.label {
		flex: 0 1 auto;
		min-width: 0;
		overflow: hidden;
		text-overflow: ellipsis;
	}
	.desc {
		flex: 1 1 0;
		min-width: 0;
		overflow: hidden;
		text-overflow: ellipsis;
		font-size: 0.92em;
		color: var(--vscode-descriptionForeground);
	}
	.actions-inline {
		display: none;
		flex: 0 0 auto;
	}
	.item:hover .actions-inline,
	.item:focus-within .actions-inline {
		display: flex;
	}
	.actions-inline .icon-btn {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 22px;
		height: 20px;
		padding: 0;
		border-radius: 4px;
	}
	/* How many blocks a rolled-up row stands for, as the jigsaw preview's piece cards show it. */
	.count {
		flex: 0 0 auto;
		font-size: 10px;
		color: var(--vscode-descriptionForeground);
	}
	.item:hover .count,
	.item:focus-within .count {
		display: none;
	}
	.warn-text {
		color: var(--vscode-editorWarning-foreground);
	}

	.menu {
		position: fixed;
		z-index: 10;
		min-width: 170px;
		padding: 4px 0;
		background: var(--vscode-menu-background, var(--vscode-editorWidget-background));
		color: var(--vscode-menu-foreground, var(--vscode-foreground));
		border: 1px solid var(--vscode-menu-border, var(--vscode-widget-border, transparent));
		border-radius: 4px;
		box-shadow: 0 2px 8px var(--vscode-widget-shadow, rgba(0, 0, 0, 0.36));
		display: flex;
		flex-direction: column;
	}
	.menu button {
		background: none;
		border: 0;
		border-radius: 0;
		text-align: left;
		padding: 3px 14px;
		color: inherit;
		white-space: nowrap;
	}
	.menu button:hover,
	.menu button:focus-visible {
		background: var(--vscode-menu-selectionBackground, var(--vscode-list-activeSelectionBackground));
		color: var(--vscode-menu-selectionForeground, inherit);
		outline: none;
	}
	.menu .danger {
		color: var(--vscode-errorForeground);
	}
	.menu .sep {
		height: 1px;
		margin: 4px 0;
		background: var(--vscode-menu-separatorBackground, var(--vscode-widget-border, transparent));
	}
</style>
