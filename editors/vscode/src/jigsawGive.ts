// The /give command for a jigsaw block set up ahead of time, so a piece
// built in the game gets jigsaw blocks with the right pool, name and
// target without typing them into the jigsaw screen each time. Pure, so
// the Jigsaw view's page and tests can import it.
//
// The block's settings go in the item's block_entity_data component
// (which needs the block entity's id), and the game copies them onto the
// block when it is placed by an operator in creative mode. The way the
// block faces comes from how it is placed, as for any jigsaw block.
import type { JigsawConfig } from './messages';

/** The longest command the chat box takes; a command block takes more. */
export const CHAT_LIMIT = 256;

/**
 * A full id for a pool or jigsaw name: a bare name is in the project's
 * namespace, as in MMS source, except `empty`, which is the game's
 * minecraft:empty. An empty field is minecraft:empty, the game's default.
 */
export function fullId(value: string, ns: string | undefined): string {
	const v = value.trim();
	if (!v || v === 'empty') return 'minecraft:empty';
	if (v.includes(':')) return v;
	return `${ns || 'minecraft'}:${v}`;
}

/** The same id as short as the game reads it: without minecraft:. */
function short(id: string): string {
	return id.startsWith('minecraft:') ? id.slice('minecraft:'.length) : id;
}

/** An SNBT string, quoted. */
function quote(s: string): string {
	return `"${s.replace(/\\/g, '\\\\').replace(/"/g, '\\"')}"`;
}

/**
 * The give command for a configured jigsaw block. Fields at the game's
 * defaults are left out to keep it under the chat limit: minecraft:empty
 * names and pool, air as the final state, and a rollable joint (the
 * default for blocks facing up or down; sideways blocks ignore it).
 */
export function giveCommand(c: JigsawConfig, ns: string | undefined): string {
	const data: string[] = ['id:"jigsaw"'];
	const name = fullId(c.name, ns);
	const target = fullId(c.target, ns);
	const pool = fullId(c.pool, ns);
	if (name !== 'minecraft:empty') data.push(`name:${quote(short(name))}`);
	if (target !== 'minecraft:empty') data.push(`target:${quote(short(target))}`);
	if (pool !== 'minecraft:empty') data.push(`pool:${quote(short(pool))}`);
	const final = c.finalState.trim();
	if (final && final !== 'air' && final !== 'minecraft:air') data.push(`final_state:${quote(final)}`);
	if (c.joint === 'aligned') data.push('joint:"aligned"');
	const components = [`block_entity_data={${data.join(',')}}`];
	if (c.label.trim()) components.push(`item_name=${quote(c.label.trim())}`);
	return `/give @s jigsaw[${components.join(',')}]`;
}

/** Whether two configurations make the same block (the label aside). */
export function sameBlock(a: JigsawConfig, b: JigsawConfig, ns: string | undefined): boolean {
	return (
		fullId(a.name, ns) === fullId(b.name, ns) &&
		fullId(a.target, ns) === fullId(b.target, ns) &&
		fullId(a.pool, ns) === fullId(b.pool, ns) &&
		(a.finalState.trim() || 'minecraft:air').replace(/^minecraft:/, '') === (b.finalState.trim() || 'minecraft:air').replace(/^minecraft:/, '') &&
		a.joint === b.joint
	);
}
