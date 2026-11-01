/// <reference types="svelte" />
/// <reference types="vite/client" />

interface VsCodeApi<S> {
	postMessage(message: unknown): void;
	getState(): S | undefined;
	setState(state: S): void;
}
declare function acquireVsCodeApi<S = unknown>(): VsCodeApi<S>;
