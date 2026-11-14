<script lang="ts">
	// A labelled slider whose value is also a text box: type a number and
	// press Enter (Esc puts it back, ↑/↓ steps, Shift for ten steps), or
	// double-click the slider to reset it.

	interface Props {
		label: string;
		value: number;
		min: number;
		max: number;
		step: number;
		/** What a double-click puts back. */
		reset: number;
		/** Decimals shown. */
		digits?: number;
		/** Shown before and after the number, as units. */
		prefix?: string;
		suffix?: string;
		title?: string;
		onchange: (v: number) => void;
	}
	let { label, value, min, max, step, reset, digits = 0, prefix = '', suffix = '', title, onchange }: Props = $props();

	const id = `slider-${Math.random().toString(36).slice(2, 9)}`;
	let editing = $state(false);
	let text = $state('');
	const shown = $derived(value.toFixed(digits));

	const clamp = (v: number) => Math.min(max, Math.max(min, v));
	function commit() {
		const n = parseFloat(text.replace(',', '.'));
		if (Number.isFinite(n)) onchange(clamp(n));
		text = clamp(Number.isFinite(n) ? n : value).toFixed(digits);
	}
	function key(e: KeyboardEvent & { currentTarget: HTMLInputElement }) {
		if (e.key === 'Enter') {
			commit();
			e.currentTarget.select();
		} else if (e.key === 'Escape') {
			text = shown;
			e.currentTarget.blur();
		} else if (e.key === 'ArrowUp' || e.key === 'ArrowDown') {
			e.preventDefault();
			const v = clamp(value + (e.key === 'ArrowUp' ? 1 : -1) * step * (e.shiftKey ? 10 : 1));
			onchange(v);
			text = v.toFixed(digits);
		}
	}
</script>

<div class="flex flex-col gap-0.5" {title}>
	<div class="flex items-center justify-between gap-2">
		<label for={id}>{label}</label>
		<span class="flex items-center gap-1 font-mono text-[11px] text-muted">
			{#if prefix}<span>{prefix}</span>{/if}
			<input
				type="text"
				inputmode="decimal"
				spellcheck="false"
				autocomplete="off"
				aria-label="{label} value"
				class="w-[8ch] rounded border border-transparent bg-transparent px-1 py-px text-right tabular-nums hover:border-line focus:border-accent focus:bg-panel focus:text-ink focus:outline-none"
				value={editing ? text : shown}
				onfocus={(e) => {
					editing = true;
					text = shown;
					e.currentTarget.select();
				}}
				oninput={(e) => (text = e.currentTarget.value)}
				onkeydown={key}
				onblur={() => {
					commit();
					editing = false;
				}}
			/>
			{#if suffix}<span>{suffix}</span>{/if}
		</span>
	</div>
	<input
		{id}
		type="range"
		class="w-full accent-[var(--c-accent,#2563eb)]"
		{min}
		{max}
		{step}
		{value}
		title="Double-click to reset"
		oninput={(e) => onchange(parseFloat(e.currentTarget.value))}
		ondblclick={() => onchange(reset)}
	/>
</div>
