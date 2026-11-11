// Loads the preview worker headless: host.ts must run before the worker
// module, which reads `self` as it loads, so this module imports them in
// that order and everything else imports the worker through here.
import './host';
import '../preview/preview.worker';

export { render, setup, type Rendered, type VanillaFiles } from './host';
