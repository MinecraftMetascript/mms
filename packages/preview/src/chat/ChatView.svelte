<script lang="ts">
	// A message a text command shows: chat for tell, the big centred text
	// of a title and its subtitle, or the line above the hotbar for an
	// action bar. It reads the component the way the book preview does
	// (book/text.ts), with sample values for scores, storage and selectors,
	// which the game fills in when the message is shown.
	import { loadLang } from '../book/assets';
	import { flatten, type Json, type Piece, type Sample, type Style } from '../book/text';
	import type { OutputFile, OutputText } from '../types';

	let { text, output }: { text: OutputText; output: OutputFile[] } = $props();

	const white: Style = {
		color: '#FFFFFF',
		bold: false,
		italic: false,
		underlined: false,
		strikethrough: false,
		obfuscated: false,
		font: 'minecraft:default'
	};

	const parse = (s: string | undefined): Json => {
		if (s === undefined) return null;
		try {
			return JSON.parse(s) as Json;
		} catch {
			return s;
		}
	};

	// The vanilla language file plus the project's en_us lines, for
	// translated text and key names.
	let lang = $state<Record<string, string>>({});
	const packLang = $derived(
		output.filter((f) => /^assets\/[^/]+\/lang\/en_us\.json$/.test(f.path)).map((f) => {
			try {
				return JSON.parse(f.text) as Record<string, string>;
			} catch {
				return {};
			}
		})
	);
	$effect(() => {
		const own = packLang;
		let stale = false;
		loadLang()
			.then((vanilla) => {
				if (!stale) lang = Object.assign({ ...vanilla }, ...own);
			})
			.catch(() => {
				if (!stale) lang = Object.assign({}, ...own);
			});
		return () => (stale = true);
	});

	// Sample values, by sample key; kept while the panel is open.
	let samples = $state<Record<string, string>>({});
	const sampleOf = (s: Sample) => samples[s.key] ?? (s.kind === 'score' ? '0' : s.kind === 'selector' ? 'Steve' : '');

	/** Pieces split into lines at line breaks. */
	function lines(c: Json, used: Map<string, Sample>): Piece[][] {
		const vals = { ...samples };
		const pieces = flatten(c, white, { lang, sample: (s) => vals[s.key] ?? sampleOf(s), used });
		const out: Piece[][] = [[]];
		for (const p of pieces) {
			if (!('text' in p)) {
				out[out.length - 1].push(p);
				continue;
			}
			p.text.split('\n').forEach((part, i) => {
				if (i > 0) out.push([]);
				if (part) out[out.length - 1].push({ style: p.style, text: part });
			});
		}
		return out;
	}

	const laid = $derived.by(() => {
		const used = new Map<string, Sample>();
		const message = lines(parse(text.text), used);
		const subtitle = text.subtitle !== undefined ? lines(parse(text.subtitle), used) : null;
		return { message, subtitle, used: [...used.values()] };
	});

	// The game draws text with a shadow a quarter as bright, one pixel down
	// and to the right.
	function shadow(color: string): string {
		const m = /^#([0-9a-f]{6})$/i.exec(color);
		if (!m) return '#3f3f3f';
		const n = parseInt(m[1], 16);
		const q = (v: number) => ((v >> 2) & 0xff).toString(16).padStart(2, '0');
		return `#${q((n >> 16) & 0xff)}${q((n >> 8) & 0xff)}${q(n & 0xff)}`;
	}
	function style(s: Style): string {
		const deco = [s.underlined ? 'underline' : '', s.strikethrough ? 'line-through' : ''].filter(Boolean).join(' ');
		return [
			`color:${s.color}`,
			`text-shadow:0.08em 0.08em 0 ${shadow(s.color)}`,
			s.bold ? 'font-weight:700' : '',
			s.italic ? 'font-style:italic' : '',
			deco ? `text-decoration:${deco}` : ''
		]
			.filter(Boolean)
			.join(';');
	}
	const scramble = (t: string) => t.replace(/\S/g, () => String.fromCharCode(0x41 + Math.floor(Math.random() * 26)));
	function hoverText(h: Json | undefined): string {
		if (h === undefined) return '';
		return flatten(h, white, { lang, sample: sampleOf, used: new Map() })
			.map((p) => ('text' in p ? p.text : ''))
			.join('');
	}
	function what(s: Style): string {
		const parts: string[] = [];
		const c = s.click;
		if (c) {
			const v = String(c.command ?? c.value ?? c.url ?? c.page ?? '');
			const act: Record<string, string> = {
				run_command: 'Click: runs ',
				suggest_command: 'Click: puts in chat ',
				copy_to_clipboard: 'Click: copies ',
				open_url: 'Click: opens ',
				change_page: 'Click: turns to page '
			};
			parts.push((act[String(c.action)] ?? 'Click: ') + v);
		}
		const h = hoverText(s.hover);
		if (h) parts.push('Hover: ' + h);
		return parts.join('\n');
	}
	const kindLabel = { chat: 'Chat', title: 'Title', actionbar: 'Action bar' } as const;
</script>

{#snippet line(pieces: Piece[])}
	{#each pieces as p, i (i)}
		{#if 'text' in p}
			<span style={style(p.style)} title={what(p.style) || undefined} class:cursor-help={!!what(p.style)}
				>{p.style.obfuscated ? scramble(p.text) : p.text}</span
			>
		{:else}
			<span class="text-[#AAAAAA]" title={p.object === 'player' ? `${p.id}'s head` : p.id}>▪</span>
		{/if}
	{/each}
{/snippet}

<div class="h-full overflow-auto p-3 text-xs" aria-label="Text preview">
	<div class="mb-2 flex items-center gap-2">
		<span class="font-semibold">{kindLabel[text.kind] ?? 'Text'}</span>
		<span class="truncate font-mono text-muted">{text.uri.replace(/^.*\//, '')}</span>
	</div>
	<div
		class="relative overflow-hidden rounded bg-[#4a6b3a] bg-gradient-to-b from-[#6b8fb5] to-[#4a6b3a] font-mono leading-snug"
		style="min-height: {text.kind === 'chat' ? 6 : 9}rem"
	>
		{#if text.kind === 'chat'}
			<ol class="absolute bottom-2 left-2 right-2 bg-black/50 px-1.5 py-1 text-[13px]" aria-label="Chat">
				{#each laid.message as l, i (i)}
					<li class="min-h-[1.2em] whitespace-pre-wrap break-words">{@render line(l)}</li>
				{/each}
			</ol>
		{:else if text.kind === 'title'}
			<div class="absolute inset-0 flex flex-col items-center justify-center px-2 text-center">
				<div class="text-[30px] leading-tight" aria-label="Title">
					{#each laid.message as l, i (i)}<div class="whitespace-pre">{@render line(l)}</div>{/each}
				</div>
				{#if laid.subtitle}
					<div class="mt-1 text-[15px]" aria-label="Subtitle">
						{#each laid.subtitle as l, i (i)}<div class="whitespace-pre">{@render line(l)}</div>{/each}
					</div>
				{/if}
			</div>
		{:else}
			<div class="absolute bottom-9 left-0 right-0 text-center text-[13px]" aria-label="Action bar">
				{#each laid.message as l, i (i)}<div class="whitespace-pre">{@render line(l)}</div>{/each}
			</div>
			<div class="absolute bottom-2 left-1/2 flex -translate-x-1/2 gap-px" aria-hidden="true">
				{#each Array(9) as _, i (i)}<div class="h-5 w-5 border border-[#8b8b8b] bg-black/40"></div>{/each}
			</div>
		{/if}
	</div>
	{#if text.kind === 'title'}
		<p class="mt-2 text-muted">A title stays for 3.5 seconds by default, fading in and out.</p>
	{/if}
	{#if laid.used.length}
		<fieldset class="mt-3">
			<legend class="mb-1 text-muted">Filled in when the message is shown. Try values:</legend>
			{#each laid.used as s (s.key)}
				<label class="mb-1 flex items-center gap-2">
					<span class="w-40 truncate font-mono" title={s.label}>{s.label}</span>
					<input
						class="min-w-0 flex-1 rounded border border-line bg-bg px-1.5 py-0.5"
						value={sampleOf(s)}
						oninput={(e) => (samples = { ...samples, [s.key]: e.currentTarget.value })}
					/>
				</label>
			{/each}
		</fieldset>
	{/if}
	<p class="mt-3 text-muted">Point at a link to see what clicking it does and its hover text.</p>
</div>
