// The MMS language server in a Web Worker. The page posts JSON-RPC message
// strings and receives the server's replies the same way; the Go side
// reads globalThis.mmsSend and calls globalThis.mmsReceive (see
// engine/cmd/mms-wasm).
importScripts('wasm_exec.js');

const queue = [];
self.mmsReceive = (text) => self.postMessage({ type: 'message', text });
self.mmsReady = (version) => {
	for (const text of queue.splice(0)) self.mmsSend(text);
	self.postMessage({ type: 'ready', version });
};
self.onmessage = (e) => {
	if (typeof self.mmsSend === 'function') self.mmsSend(e.data);
	else queue.push(e.data);
};

const go = new Go();
WebAssembly.instantiateStreaming(fetch('mms.wasm'), go.importObject)
	.then((result) => go.run(result.instance))
	.then(() => self.postMessage({ type: 'exit' }))
	.catch((err) => self.postMessage({ type: 'error', message: String(err) }));
