<script lang="ts">
	// The World view's page: the project, its world, the Build button with
	// the last build's result and what to do in the game next, and the
	// structures saved in the structure world (the build world unless a
	// separate one is chosen). The extension (src/world.ts) owns the
	// state and sends all of it after each change. Actions are buttons,
	// drawn with the theme's button colours, and results are cards without
	// them, so the two never look alike.
	import type { FromWorld, ToWorld, WorldState, WorldStructureRow } from '../src/messages';

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

	const label: Record<WorldStructureRow['status'], string> = { new: 'New', imported: 'Imported', changed: 'Changed in world' };
	const waiting = $derived(s?.structures.filter((r) => r.status !== 'imported').length ?? 0);
	const unused = $derived((s?.structures.filter((r) => r.unused).length ?? 0) + (s?.unusedFiles.length ?? 0));
	const sentence = (line: string) => line.charAt(0).toUpperCase() + line.slice(1);
</script>

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
				<button class="secondary small-btn" title="Open the world's datapacks folder" onclick={() => send({ type: 'run', action: 'revealWorld' })}>Open folder</button>
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
					<button class="secondary small-btn" title="Open the world's generated folder, where structure blocks save" onclick={() => send({ type: 'run', action: 'revealStructureWorld' })}>Open folder</button>
				</div>
				<p class="muted small">
					Builds never write to this world.
					<button class="link" onclick={() => send({ type: 'run', action: 'sameStructureWorld' })}>{s.world ? 'Use the build world instead' : 'Forget this world'}</button>
				</p>
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
					{[s.structures.length ? (waiting ? `${waiting} to import` : 'All imported') : '', unused ? `${unused} not used` : ''].filter(Boolean).join(' · ')}
					{#if s.unusedPaused}· the not-used check waits until the project has no errors{/if}
				</p>
				<ul>
					{#each s.structures as r (r.file)}
						<li class="st" class:unused={r.unused} class:changed={r.status === 'changed'}>
							<div class="st-head">
								<span class="name grow" title={`Saved ${r.saved} in ${r.file}`}>{r.id}</span>
								{#if r.unused}
									<span class="badge warn-badge">Not used</span>
								{:else}
									<span class="badge" class:changed-badge={r.status === 'changed'}>{label[r.status]}</span>
								{/if}
							</div>
							{#if r.unused}
								<div class="small">No <code>Structure(...)</code> call uses <code>{r.rel}</code>, so the pack leaves it out. Add it to a template pool:</div>
								<code class="call">{r.call}</code>
							{:else if r.status === 'new'}
								<div class="muted small">Import copies it to {r.rel}.</div>
							{:else}
								<div class="muted small">{r.rel}</div>
							{/if}
							{#if r.packId}<div class="muted small">Its id in the pack is {r.packId}, in the project's namespace.</div>{/if}
							<div class="actions">
								{#if r.status !== 'imported'}
									<button class="primary" onclick={() => send({ type: 'import', file: r.file })}>{r.status === 'changed' ? 'Import again' : 'Import'}</button>
								{/if}
								<button class="secondary" title={r.call} onclick={() => copy(r.file, { type: 'copyCall', file: r.file })}>{copied === r.file ? 'Copied' : 'Copy call'}</button>
								{#if r.status !== 'new'}
									<button class="secondary" onclick={() => send({ type: 'revealFile', rel: r.rel })}>Show file</button>
								{/if}
								<button class="secondary danger" title="Delete this save from the world, for one saved under the wrong name. Asks first." onclick={() => send({ type: 'deleteFromWorld', file: r.file })}>Delete from world</button>
							</div>
						</li>
					{/each}
					{#each s.unusedFiles as rel (rel)}
						<li class="st unused">
							<div class="st-head">
								<span class="name grow">{rel}</span>
								<span class="badge warn-badge">Not used</span>
							</div>
							<div class="small">In the project but not saved in {from.name}, and no <code>Structure(...)</code> call uses it, so the pack leaves it out.</div>
							<div class="actions">
								<button class="secondary" onclick={() => send({ type: 'revealFile', rel })}>Show file</button>
							</div>
						</li>
					{/each}
				</ul>
			{/if}
		</section>
	{/if}
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
	.card-head,
	.st-head {
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

	/* Structures */
	.st {
		padding: 6px 8px;
		border: 1px solid var(--vscode-widget-border, var(--vscode-panel-border, transparent));
		border-left-width: 3px;
		border-radius: 2px;
	}
	.st.changed {
		border-left-color: var(--vscode-gitDecoration-modifiedResourceForeground);
	}
	.st.unused {
		border-left-color: var(--vscode-editorWarning-foreground);
		background: color-mix(in srgb, var(--vscode-editorWarning-foreground) 8%, transparent);
	}
	.badge {
		font-size: 11px;
		padding: 0 6px;
		border-radius: 9px;
		background: var(--vscode-badge-background);
		color: var(--vscode-badge-foreground);
		white-space: nowrap;
	}
	.changed-badge {
		background: var(--vscode-gitDecoration-modifiedResourceForeground);
		color: var(--vscode-editor-background);
	}
	.warn-badge {
		background: var(--vscode-editorWarning-foreground);
		color: var(--vscode-editor-background);
		font-weight: 600;
	}
	.call {
		display: block;
		margin-top: 4px;
		padding: 2px 6px;
		background: var(--vscode-textCodeBlock-background);
		border-radius: 2px;
		overflow-wrap: anywhere;
	}
</style>
