import { inspectable } from './inspect/common';
import { previewable } from './preview/protocol';

/** What the preview pane shows for an output path: a preview, an inspector, or nothing. */
export const shows = (path: string | null) => previewable(path) ?? inspectable(path);
