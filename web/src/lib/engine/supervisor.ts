import type { Transport } from '@codemirror/lsp-client';
import { SignatureShim } from './signature';

export type EngineStatus = 'starting' | 'ready' | 'restarting' | 'failed';

export interface SupervisorOptions {
	/** URL of the classic worker script (static/mms-worker.js). */
	workerURL: string;
	/** Called with every status change. */
	onStatus?: (status: EngineStatus, detail?: string) => void;
	/** Called after a restart, once the new worker is up, to replay state. */
	onRestart?: () => void;
	/** How often to ping the server, in ms. */
	heartbeat?: number;
	/** How long a ping may take before the worker counts as hung, in ms. */
	hangTimeout?: number;
}

const maxRestarts = 5;
const restartWindow = 60_000;

/**
 * Supervisor runs the WASM language server in a worker and presents it to
 * the LSP client as a Transport. When the worker crashes, exits or stops
 * answering pings, it starts a fresh one and asks the owner to replay
 * the open documents, so an engine bug costs a moment, not the session.
 */
export class Supervisor {
	private worker: Worker | null = null;
	private handlers: ((value: string) => void)[] = [];
	private pending: string[] = [];
	private shim = new SignatureShim();
	private ready = false;
	private restarts: number[] = [];
	private pingTimer: ReturnType<typeof setInterval> | null = null;
	private pingID = 0;
	private pingSent = 0;
	private disposed = false;
	status: EngineStatus = 'starting';
	version = '';

	constructor(private opts: SupervisorOptions) {
		this.start();
	}

	/** The transport for LSPClient.connect. */
	readonly transport: Transport = {
		send: (message: string) => {
			this.shim.sent(message);
			if (this.ready && this.worker) this.worker.postMessage(message);
			else this.pending.push(message);
		},
		subscribe: (handler) => {
			this.handlers.push(handler);
		},
		unsubscribe: (handler) => {
			this.handlers = this.handlers.filter((h) => h !== handler);
		}
	};

	private setStatus(s: EngineStatus, detail?: string) {
		this.status = s;
		this.opts.onStatus?.(s, detail);
	}

	private start() {
		const w = new Worker(this.opts.workerURL);
		this.worker = w;
		this.ready = false;
		w.onmessage = (e: MessageEvent) => {
			if (w !== this.worker) return;
			const m = e.data as { type: string; text?: string; version?: string; message?: string };
			switch (m.type) {
				case 'ready':
					this.ready = true;
					this.version = m.version ?? '';
					for (const msg of this.pending.splice(0)) w.postMessage(msg);
					this.setStatus('ready');
					this.startHeartbeat();
					break;
				case 'message':
					if (this.isPong(m.text!)) return;
					for (const h of this.handlers) h(this.shim.received(m.text!));
					break;
				case 'exit':
					this.crashed('the engine stopped');
					break;
				case 'error':
					this.crashed(m.message ?? 'the engine failed to load');
					break;
			}
		};
		w.onerror = (e) => {
			e.preventDefault();
			this.crashed(e.message || 'the engine crashed');
		};
	}

	private startHeartbeat() {
		this.stopHeartbeat();
		const every = this.opts.heartbeat ?? 5000;
		const hang = this.opts.hangTimeout ?? 15000;
		this.pingTimer = setInterval(() => {
			if (this.pingSent && Date.now() - this.pingSent > hang) {
				this.crashed('the engine stopped responding');
				return;
			}
			if (!this.pingSent) {
				this.pingSent = Date.now();
				this.worker?.postMessage(JSON.stringify({ jsonrpc: '2.0', id: `ping-${++this.pingID}`, method: 'mms/ping' }));
			}
		}, every);
	}

	private stopHeartbeat() {
		if (this.pingTimer) clearInterval(this.pingTimer);
		this.pingTimer = null;
		this.pingSent = 0;
	}

	private isPong(text: string): boolean {
		if (!text.includes('"ping-')) return false;
		try {
			const m = JSON.parse(text);
			if (typeof m.id === 'string' && m.id.startsWith('ping-')) {
				this.pingSent = 0;
				return true;
			}
		} catch {
			/* not ours */
		}
		return false;
	}

	private crashed(reason: string) {
		if (this.disposed) return;
		this.stopHeartbeat();
		this.worker?.terminate();
		this.worker = null;
		this.ready = false;
		this.pending = [];
		const now = Date.now();
		this.restarts = this.restarts.filter((t) => now - t < restartWindow);
		if (this.restarts.length >= maxRestarts) {
			this.setStatus('failed', reason);
			return;
		}
		this.restarts.push(now);
		this.setStatus('restarting', reason);
		this.start();
		this.opts.onRestart?.();
	}

	/** Restart by hand, for example after a failure. */
	restart() {
		this.restarts = [];
		this.crashed('restarted');
	}

	dispose() {
		this.disposed = true;
		this.stopHeartbeat();
		this.worker?.terminate();
		this.worker = null;
	}
}
