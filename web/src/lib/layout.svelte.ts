/**
 * The editor page's layout: whether the sidebar is shown, how wide it is,
 * and how the rest is shared between the text and the preview. It is kept
 * in this browser for every project, since it is about the screen, not the
 * pack.
 */
const key = 'mms-layout';

export const sidebarMin = 160;
export const sidebarMax = 480;
export const sidebarDefault = 208; // 13rem
/** The editor's share of the room next to the sidebar, in percent. */
export const editorMin = 15;
export const editorMax = 85;
export const editorDefault = 50;

const clamp = (n: number, lo: number, hi: number) => Math.min(hi, Math.max(lo, n));

export class Layout {
	sidebarOpen = $state(true);
	sidebarWidth = $state(sidebarDefault);
	editorShare = $state(editorDefault);
	#timer: ReturnType<typeof setTimeout> | undefined;

	load() {
		try {
			const raw = localStorage.getItem(key);
			if (!raw) return;
			const v = JSON.parse(raw);
			if (typeof v.sidebarOpen === 'boolean') this.sidebarOpen = v.sidebarOpen;
			if (typeof v.sidebarWidth === 'number' && isFinite(v.sidebarWidth)) this.sidebarWidth = clamp(Math.round(v.sidebarWidth), sidebarMin, sidebarMax);
			if (typeof v.editorShare === 'number' && isFinite(v.editorShare)) this.editorShare = clamp(v.editorShare, editorMin, editorMax);
		} catch {
			/* storage unavailable or bad data: keep the defaults */
		}
	}

	/** Saves soon, so a drag writes once when it settles. */
	save() {
		clearTimeout(this.#timer);
		this.#timer = setTimeout(() => {
			try {
				localStorage.setItem(key, JSON.stringify({ sidebarOpen: this.sidebarOpen, sidebarWidth: this.sidebarWidth, editorShare: this.editorShare }));
			} catch {
				/* ignore */
			}
		}, 200);
	}

	toggleSidebar() {
		this.sidebarOpen = !this.sidebarOpen;
		this.save();
	}
	setSidebarWidth(px: number) {
		this.sidebarWidth = clamp(Math.round(px), sidebarMin, sidebarMax);
		this.sidebarOpen = true;
		this.save();
	}
	setEditorShare(pct: number) {
		this.editorShare = Math.round(clamp(pct, editorMin, editorMax) * 10) / 10;
		this.save();
	}
}
