/**
 * Shortcut labels shown in menus and tooltips. Keep them in step with the
 * keymap in Editor.svelte.
 *
 * On macOS, Option+letter types a character (Option+Shift+F is "Ï"), and
 * CodeMirror leaves such keys to typing, so Option-only bindings never fire
 * there. Format uses Cmd+Shift+I on macOS instead (Shift+Alt+F elsewhere,
 * since Ctrl+Shift+I opens the browser's developer tools).
 */
export const mac = typeof navigator !== 'undefined' && /Mac|iP(hone|ad)/.test(navigator.platform);
export const mod = mac ? '⌘' : 'Ctrl';
export const alt = mac ? '⌥' : 'Alt';

export const formatKeys = mac ? '⌘+Shift+I' : 'Shift+Alt+F';
export const fixAllKeys = `${mod}+${alt}+F`;
export const quickFixKeys = `${mod}+.`;
