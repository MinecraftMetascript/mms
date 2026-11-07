# Notes for agents working in this repo

- Do not add entries to `docs/decisions.md`. It is frozen (D1 to D94) because parallel pull requests kept conflicting over the next D-number. Explain judgement calls in the pull request description and the commit message instead, and do not cite new D-numbers in code comments.
- The plan is in `docs/plan.md` and the task list is in `docs/tasks.md`.
- The website's pages are Markdown under `site/`, which mirrors the `/docs` URLs (`site/guides/functions.md` is `/docs/guides/functions`). Sidebar groups and their order are in `web/src/lib/docs/site.ts`. The generated reference pages come from `mms reference`.
- Every ```` ```mms ```` block under `site/` is checked by `TestSiteSamples` in `engine/project` as a file of its own. Fence a sample that shows a mistake on purpose as ```` ```mms error ````, a sample that reads asset files as ```` ```mms in=examples/<name> ```` (the folder the files come from), and a fragment that is not a whole file as plain ```` ``` ````.
