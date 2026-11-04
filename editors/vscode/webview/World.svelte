<script lang="ts">
	// The World view's page: the project, its world, the Build button with
	// the last build's result and what to do in the game next, and the
	// structures saved in the structure world (the build world unless a
	// separate one is chosen). The extension (src/world.ts) owns the
	// state and sends all of it after each change. Actions are buttons,
	// drawn with the theme's button colours, and results are cards without
	// them, so the two never look alike. Each structure row shows at most
	// one button, for the one thing to do next (Import, or Copy call for a
	// structure no pool uses); the rest are in its ⋯ menu, which a
	// right-click also opens, so the list stays mostly text.
	import type { FromWorld, ToWorld, WorldState, WorldStructureRow } from '../src/messages';
	import { group, type Group, type Node } from './groups';

	const vscode = acquireVsCodeApi<WorldState>();
	let s = $state<WorldState | undefined>(vscode.getState());
	// Buttons that just copied something say so for a moment.
	let copied = $state<string | null>(null);
	let copyTimer: ReturnType<typeof setTimeout> | undefined;

	window.addEventListener('message', (e: MessageEvent<ToWorld>) => {
		if (e.data.type === 'state') {
			s = e.data.state;
			vscode.setState(e.data.state);
		}
	});
	const send = (m: FromWorld) => vscode.postMessage(m);
	send({ type: 'ready' });

	function copy(id: string, m: FromWorld) {
		send(m);
		copied = id;
		clearTimeout(copyTimer);
		copyTimer = setTimeout(() => (copied = null), 1500);
	}

	// A row's ⋯ menu (or a world row's): where it opens and what it holds.
	type MenuItem = { label: string; run: () => void; danger?: boolean } | 'sep';
	let menu = $state<{ x: number; y: number; right: boolean; items: MenuItem[] } | null>(null);
	let menuEl = $state<HTMLDivElement>();
	function openMenu(e: MouseEvent, items: MenuItem[]) {
		e.preventDefault();
		e.stopPropagation();
		// From the button: under it, right-aligned. From a right-click: at the pointer.
		const b = e.type === 'contextmenu' ? null : (e.currentTarget as HTMLElement).getBoundingClientRect();
		menu = { x: b ? b.right : e.clientX, y: b ? b.bottom + 2 : e.clientY, right: !!b, items };
	}
	// Keep the menu inside the view, and put the keyboard on its first item.
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

	function rowMenu(r: WorldStructureRow): MenuItem[] {
		return [
			...(r.status !== 'imported' ? [{ label: r.status === 'changed' ? 'Import again' : 'Import', run: () => send({ type: 'import', file: r.file }) }] : []),
			{ label: 'Copy Structure(...) call', run: () => send({ type: 'copyCall', file: r.file }) },
			...(r.status !== 'new' ? [{ label: 'Show file in Explorer', run: () => send({ type: 'revealFile', rel: r.rel }) }] : []),
			'sep',
			{ label: 'Delete from world…', danger: true, run: () => send({ type: 'deleteFromWorld', file: r.file }) }
		];
	}

	// "5 min ago", for when a structure was saved; the full time is in the tooltip.
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

	const waiting = $derived(s?.structures.filter((r) => r.status !== 'imported').length ?? 0);
	const unused = $derived((s?.structures.filter((r) => r.unused).length ?? 0) + (s?.unusedFiles.length ?? 0));
	// The structures as folders, split at / in their ids, so big
	// packs of pieces stay readable. With several namespaces the namespace
	// is the top folder.
	const groups = $derived.by(() => {
		const rows = s?.structures ?? [];
		return group(rows, nameOf);
	});
	// The name the list sorts into folders by, and shows (less its folder).
	const manyNs = $derived(new Set((s?.structures ?? []).map((r) => r.id.split(':')[0])).size > 1);
	const nameOf = (r: WorldStructureRow) => (manyNs ? r.id : r.id.slice(r.id.indexOf(':') + 1));
	// Folders the user opened or closed; the rest are open when they hold
	// something to do (new, changed or not used) and closed otherwise.
	let open = $state<Record<string, boolean>>({});
	const isOpen = (g: Group<WorldStructureRow>) => open[g.key] ?? g.items.some((r) => r.status !== 'imported' || r.unused);
	const sentence = (line: string) => line.charAt(0).toUpperCase() + line.slice(1);
</script>

{#snippet icon(name: 'import' | 'copy' | 'more' | 'check')}
	<svg viewBox="0 0 16 16" width="16" height="16" aria-hidden="true" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round">
		{#if name === 'import'}<path d="M8 2.5v7M5 7l3 3 3-3M3 11v2.5h10V11" />
		{:else if name === 'copy'}<rect x="5.5" y="5.5" width="8" height="8" rx="1" /><path d="M10.5 5.5V3.5a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2" />
		{:else if name === 'check'}<path d="M3 8.5l3 3 7-7" />
		{:else}<circle cx="3.5" cy="8" r=".6" fill="currentColor" /><circle cx="8" cy="8" r=".6" fill="currentColor" /><circle cx="12.5" cy="8" r=".6" fill="currentColor" />{/if}
	</svg>
{/snippet}

{#snippet row(r: WorldStructureRow, prefix: string, depth: number)}
	{@const tip = [
		r.id,
		r.status === 'new' ? `New: not in the project yet. Import copies it to ${r.rel}.` : r.status === 'changed' ? `Changed: differs from the project copy ${r.rel}. Import again overwrites it.` : r.rel,
		r.unused ? `Not used: no Structure(...) call uses it, so the pack leaves it out. Add it to a template pool:\n${r.call}` : '',
		r.packId ? `In the pack as ${r.packId}, the project's namespace.` : '',
		`Saved ${r.saved}`
	].filter(Boolean).join('\n')}
	<li>
		<div class="item" style:--depth={depth} title={tip} tabindex="0" role="treeitem" aria-selected="false" oncontextmenu={(e) => openMenu(e, rowMenu(r))} onkeydown={(e) => { if (e.key === 'ContextMenu' || (e.shiftKey && e.key === 'F10')) openMenu(e as unknown as MouseEvent, rowMenu(r)); }}>
			<span class="twistie"></span>
			<span class="label" class:new-text={r.status === 'new'} class:changed-text={r.status === 'changed'} class:warn-text={r.unused}>{nameOf(r).slice(prefix.length)}</span>
			<span class="desc">
				{#if r.unused}<span class="warn-text">not used</span>{' · '}{:else if r.status === 'new'}<span class="new-text">new</span>{' · '}{:else if r.status === 'changed'}<span class="changed-text">changed</span>{' · '}{/if}{ago(r.savedMs)}
			</span>
			<span class="actions-inline">
				{#if r.status !== 'imported'}
					<button class="icon-btn" title={r.status === 'changed' ? 'Import again (overwrites the project copy)' : 'Import into the project'} aria-label="Import {r.id}" onclick={() => send({ type: 'import', file: r.file })}>{@render icon('import')}</button>
				{/if}
				<button class="icon-btn" title="Copy {r.call}" aria-label="Copy the Structure call for {r.id}" onclick={() => copy(r.file, { type: 'copyCall', file: r.file })}>{@render icon(copied === r.file ? 'check' : 'copy')}</button>
				<button class="icon-btn" title="More actions" aria-label="More actions for {r.id}" aria-haspopup="menu" onclick={(e) => openMenu(e, rowMenu(r))}>{@render icon('more')}</button>
			</span>
		</div>
	</li>
{/snippet}

{#snippet tree(nodes: Node<WorldStructureRow>[], prefix: string, depth: number)}
	{#each nodes as n (n.kind === 'group' ? `g:${n.key}` : n.item.file)}
		{#if n.kind === 'leaf'}
			{@render row(n.item, prefix, depth)}
		{:else}
			{@const todo = n.items.filter((r) => r.status !== 'imported')}
			{@const notUsed = n.items.filter((r) => r.unused).length}
			<li>
				<details open={isOpen(n)} ontoggle={(e) => (open[n.key] = e.currentTarget.open)}>
					<summary class="item" style:--depth={depth} title={n.key}>
						<span class="twistie" aria-hidden="true"></span>
						<span class="label folder">{n.label}</span>
						<span class="desc">
							{n.items.length}{#if todo.length}{' · '}<span class="new-text">{todo.length} to import</span>{/if}{#if notUsed}{' · '}<span class="warn-text">{notUsed} not used</span>{/if}
						</span>
						{#if todo.length}
							<span class="actions-inline">
								<button class="icon-btn" title="Import the {todo.length} new and changed structures in {n.label}" aria-label="Import {todo.length} in {n.label}" onclick={(e) => { e.preventDefault(); send({ type: 'importFiles', files: todo.map((r) => r.file) }); }}>{@render icon('import')}</button>
							</span>
						{/if}
					</summary>
					<ul class="tree">{@render tree(n.children, n.key, depth + 1)}</ul>
				</details>
			</li>
		{/if}
	{/each}
{/snippet}

{#if !s}
	<p class="muted pad">Loading…</p>
{:else if !s.project}
	<p class="pad">No MMS project in this workspace. Open a folder with an <code>mms.json</code>, or create one with <code>mms init</code>.</p>
{:else}
	<section>
		<div class="row">
			<span class="icon" aria-hidden="true">▣</span>
			<div class="grow">
				<div class="name" title={s.project.path}>{s.project.name}</div>
				{#if s.project.ns}<div class="muted small">namespace {s.project.ns}</div>{/if}
			</div>
			{#if s.project.canChoose}
				<button class="secondary small-btn" onclick={() => send({ type: 'run', action: 'chooseProject' })}>Change</button>
			{/if}
		</div>
		{#if s.world}
			<div class="row">
				<span class="icon" aria-hidden="true">◍</span>
				<div class="grow">
					<div class="muted small">Builds into</div>
					<div class="name" title={s.world.path}>{s.world.name}</div>
					{#if s.world.folder}<div class="muted small">{s.world.folder}</div>{/if}
				</div>
				<button class="secondary small-btn" title="Choose another world to build into" onclick={() => send({ type: 'run', action: 'chooseWorld' })}>Change</button>
				<button class="icon-btn more" title="More actions" aria-label="More actions for the build world" aria-haspopup="menu" onclick={(e) => openMenu(e, [{ label: 'Open datapacks folder', run: () => send({ type: 'run', action: 'revealWorld' }) }])}>⋯</button>
			</div>
		{/if}
	</section>

	{#if !s.world}
		<section>
			<p>Choose a Minecraft world to build this project's pack into. The data pack goes in the world's datapacks folder, and a resource pack, if the project has one, in resourcepacks/resources.zip.</p>
			<button class="primary wide" onclick={() => send({ type: 'run', action: 'chooseWorld' })}>Choose world…</button>
			<p class="muted small">Worlds come from the Minecraft launcher's and Prism Launcher's saves folders, or browse to any world folder.</p>
			{#if !s.structureWorld}
				<button class="secondary small-btn" onclick={() => send({ type: 'run', action: 'chooseStructureWorld' })}>Only import structures from a world…</button>
			{/if}
		</section>
	{:else}
		<section>
			<button class="primary wide" disabled={s.building} onclick={() => send({ type: 'run', action: 'build' })}>
				{#if s.building}<span class="spin" aria-hidden="true">◌</span> Building…{:else}▶ Build into world{/if}
			</button>
			<label class="check small">
				<input type="checkbox" checked={s.buildOnSave} onchange={(e) => send({ type: 'buildOnSave', on: e.currentTarget.checked })} />
				Build on save
			</label>

			{#if s.result && !s.building}
				{@const r = s.result}
				<div class="card" class:error={!r.ok} role="status">
					<div class="card-head">
						<span class={r.ok ? 'ok' : 'bad'} aria-hidden="true">{r.ok ? '✓' : '✕'}</span>
						<span class="grow">{r.ok ? `Built ${r.files} ${r.files === 1 ? 'file' : 'files'}` : r.message}</span>
						<span class="muted small">{r.at}</span>
						<button class="icon-btn" title="Dismiss" aria-label="Dismiss" onclick={() => send({ type: 'run', action: 'dismiss' })}>✕</button>
					</div>
					{#if r.ok}
						{#if r.next?.action === 'reload'}
							<div class="next">
								<div class="next-title">Next, in the game: run <code>/reload</code></div>
								{#if r.next.changed}<div class="muted small">{sentence(r.next.changed)} changed.</div>{/if}
								<button class="secondary" onclick={() => copy('reload', { type: 'run', action: 'copyReload' })}>{copied === 'reload' ? 'Copied' : 'Copy /reload'}</button>
							</div>
						{:else if r.next?.action === 'rejoin'}
							<div class="next warn">
								<div class="next-title">Next, in the game: exit and rejoin the world</div>
								<div class="muted small">{r.next.changed ? `${sentence(r.next.changed)} changed. ` : ''}Minecraft only reads these when the world is opened, so /reload is not enough.</div>
							</div>
						{:else if r.next?.action === 'none'}
							<div class="muted small">Nothing changed since the last build.</div>
						{:else}
							<div class="muted small">Run <code>/reload</code> in the game to load the data pack.{r.resources ? ' Rejoin the world to load the resource pack.' : ''}</div>
						{/if}
					{:else}
						<div class="actions">
							{#if r.problems}
								<button class="secondary" onclick={() => send({ type: 'run', action: 'showProblems' })}>Show problems</button>
							{/if}
							<button class="secondary" onclick={() => send({ type: 'run', action: 'showOutput' })}>Show output</button>
						</div>
						{#if r.detail}
							<details>
								<summary class="small">Build output</summary>
								<pre>{r.detail}</pre>
							</details>
						{/if}
					{/if}
				</div>
			{/if}
		</section>
	{/if}

	{#if s.structureWorld}
		{@const from = s.structureWorld}
		<section>
			<div class="head">
				<h2 class="grow">Structures</h2>
				{#if waiting}
					<button class="secondary small-btn" title="Import the new and changed structures into the project" onclick={() => send({ type: 'run', action: 'importWaiting' })}>Import {waiting}</button>
				{/if}
				<button class="icon-btn" title="Refresh" aria-label="Refresh" onclick={() => send({ type: 'run', action: 'refresh' })}>↻</button>
			</div>
			{#if from.separate}
				<div class="row">
					<span class="icon" aria-hidden="true">◍</span>
					<div class="grow">
						<div class="muted small">From</div>
						<div class="name" title={from.path}>{from.name}</div>
						{#if from.folder}<div class="muted small">{from.folder}</div>{/if}
					</div>
					<button class="secondary small-btn" title="Choose another world to import structures from" onclick={() => send({ type: 'run', action: 'chooseStructureWorld' })}>Change</button>
					<button class="icon-btn more" title="More actions" aria-label="More actions for the structure world" aria-haspopup="menu" onclick={(e) => openMenu(e, [
						{ label: 'Open generated folder', run: () => send({ type: 'run', action: 'revealStructureWorld' }) },
						{ label: s?.world ? 'Use the build world instead' : 'Forget this world', run: () => send({ type: 'run', action: 'sameStructureWorld' }) }
					])}>⋯</button>
				</div>
				<p class="muted small">Builds never write to this world.</p>
			{:else}
				<p class="muted small">
					From the build world.
					<button class="link" title="Save structures in one world and build the pack into another" onclick={() => send({ type: 'run', action: 'chooseStructureWorld' })}>Use a separate world…</button>
				</p>
			{/if}
			{#if !s.structures.length && !s.unusedFiles.length}
				<p class="muted small">None saved yet. Save one with a structure block in SAVE mode; it shows up here to import into the project for a template pool.</p>
			{:else}
				<p class="muted small">
					{s.structures.length ? (waiting ? `${waiting} to import` : 'All imported') : ''}
					{#if unused}
						{s.structures.length ? ' · ' : ''}<span class="warn-text">{unused} not used</span>: no <code>Structure(...)</code> call uses {unused === 1 ? 'it' : 'them'}, so the pack leaves {unused === 1 ? 'it' : 'them'} out. Copy a row's call into a template pool.
					{/if}
					{#if s.unusedPaused}(The not-used check waits until the project has no errors.){/if}
				</p>
				<ul class="tree" role="tree">
					{@render tree(groups, '', 0)}
					{#each s.unusedFiles as rel (rel)}
						{@const items: MenuItem[] = [{ label: 'Show file in Explorer', run: () => send({ type: 'revealFile', rel }) }]}
						<li>
							<div class="item" style:--depth={0} tabindex="0" role="treeitem" aria-selected="false" title={`${rel}\nNot saved in ${from.name}, and no Structure(...) call uses it, so the pack leaves it out.`} oncontextmenu={(e) => openMenu(e, items)}>
								<span class="twistie"></span>
								<span class="label warn-text">{rel}</span>
								<span class="desc"><span class="warn-text">not used</span> · not in world</span>
								<span class="actions-inline">
									<button class="icon-btn" title="More actions" aria-label="More actions for {rel}" aria-haspopup="menu" onclick={(e) => openMenu(e, items)}>{@render icon('more')}</button>
								</span>
							</div>
						</li>
					{/each}
				</ul>
			{/if}
		</section>
	{/if}
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
		margin: 0;
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
		gap: 6px;
	}
	code {
		font-family: var(--vscode-editor-font-family);
		font-size: 0.95em;
	}
	.row,
	.head,
	.card-head {
		display: flex;
		align-items: center;
		gap: 6px;
	}
	.row {
		padding: 3px 0;
	}
	.grow {
		flex: 1;
		min-width: 0;
	}
	.name {
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}
	.icon {
		width: 16px;
		text-align: center;
		color: var(--vscode-icon-foreground);
	}
	.muted {
		color: var(--vscode-descriptionForeground);
	}
	.small {
		font-size: 12px;
		line-height: 1.4;
	}

	/* Buttons: the theme's button colours, a border and a pointer, so
	   they never read as a line of text. */
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
	button:disabled {
		opacity: 0.6;
		cursor: default;
	}
	.primary {
		background: var(--vscode-button-background);
		color: var(--vscode-button-foreground);
	}
	.primary:hover:not(:disabled) {
		background: var(--vscode-button-hoverBackground);
	}
	.secondary {
		background: var(--vscode-button-secondaryBackground);
		color: var(--vscode-button-secondaryForeground);
		border-color: var(--vscode-button-border, var(--vscode-contrastBorder, transparent));
	}
	.secondary:hover:not(:disabled) {
		background: var(--vscode-button-secondaryHoverBackground);
	}
	.danger:hover:not(:disabled) {
		color: var(--vscode-errorForeground);
	}
	.small-btn {
		padding: 1px 8px;
		font-size: 12px;
	}
	.wide {
		width: 100%;
		padding: 6px 10px;
	}
	/* An inline action inside a sentence; still a button, but it reads as a link. */
	.link {
		background: none;
		border: 0;
		padding: 0;
		color: var(--vscode-textLink-foreground);
		text-decoration: underline;
		font-size: inherit;
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
	.check {
		display: flex;
		align-items: center;
		gap: 6px;
		margin: 6px 0 2px;
		cursor: pointer;
	}
	.actions {
		display: flex;
		flex-wrap: wrap;
		gap: 6px;
		margin-top: 6px;
	}
	.spin {
		display: inline-block;
		animation: spin 1s linear infinite;
	}
	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}

	/* Results: a card with no button colours, and a coloured edge. */
	.card {
		margin-top: 10px;
		padding: 8px 10px;
		background: var(--vscode-textBlockQuote-background, var(--vscode-editorWidget-background));
		border-left: 3px solid var(--vscode-testing-iconPassed, #388a34);
		border-radius: 2px;
		display: flex;
		flex-direction: column;
		gap: 6px;
	}
	.card.error {
		border-left-color: var(--vscode-errorForeground);
	}
	.ok {
		color: var(--vscode-testing-iconPassed, #388a34);
	}
	.bad {
		color: var(--vscode-errorForeground);
	}
	.next {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		gap: 4px;
	}
	.next-title {
		font-weight: 600;
	}
	.next.warn .next-title {
		color: var(--vscode-editorWarning-foreground);
	}
	details pre {
		white-space: pre-wrap;
		word-break: break-word;
		font-family: var(--vscode-editor-font-family);
		font-size: 11px;
		max-height: 200px;
		overflow: auto;
		margin: 4px 0 0;
	}
	summary {
		cursor: pointer;
	}

	/* Structures: a tree like the Explorer's. One 22px line per item, the
	   status as coloured words in the description (like git decorations),
	   and the actions as icons that show on hover or keyboard focus. The
	   full explanation is in each item's tooltip. */
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
		cursor: default;
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
		cursor: pointer;
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
		gap: 0;
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
	.new-text {
		color: var(--vscode-gitDecoration-untrackedResourceForeground, #73c991);
	}
	.changed-text {
		color: var(--vscode-gitDecoration-modifiedResourceForeground);
	}
	.warn-text {
		color: var(--vscode-editorWarning-foreground);
	}
	.more {
		opacity: 0.45;
		font-weight: 700;
		letter-spacing: 1px;
	}
	.row:hover .more,
	.more:focus-visible {
		opacity: 1;
	}

	/* The ⋯ menu, drawn like VS Code's own context menus. */
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
