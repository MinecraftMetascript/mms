# Decision log

Judgement calls made while building the rewrite without stopping for input. Each entry says what was chosen and why, so any of them can be revisited. Newest last.

Decisions Brian made directly are recorded in the plan ([docs/plan.md](plan.md)) and are not repeated here.

| # | Date | Area | Decision | Why |
| --- | --- | --- | --- | --- |
| D1 | 2026-10-31 | Repo | The old implementation was deleted from the tree in one commit, not moved to a folder. | The old code is reference only and stays in git history (`git show 9748807:<path>`). Keeping it in the tree would confuse tooling and search. |
| D2 | 2026-10-31 | Tracking | Milestones are sub-issues of a tracking issue (#22) and carry a checklist; task sub-issues are filed per milestone when that milestone starts, not all up front. | Filing every task early would freeze a breakdown that will change as the earlier milestones teach us things. |
| D3 | 2026-10-31 | Spec pack | The spec pack is the vanilla-mcdoc symbol table filtered to 26.3, in the same node shape, plus registries, block states and the command tree, stored as gzipped JSON. It is not converted to a new normalised schema model. | Keeping mcdoc's own shape is lossless (templates, dispatchers, spreads, attributes survive), so no information is thrown away before the checker needs it. One loader in `spec` interprets it. A custom binary format can come later if size or load time need it; the generator reports both. |
| D4 | 2026-10-31 | Spec pack | Version gates: `#[since=v]` keeps a node when 26.3 >= v; `#[until=v]` keeps it when 26.3 < v. Versions compare as dotted integers. | This is how Spyglass reads the attributes. Checked against `density_function::Clamp.input`, which switches from inline-only to reference-or-inline at 26.2. |
| D5 | 2026-10-31 | Spec pack | Only resource kinds whose type lives under `::java::data::` are data-pack kinds. Asset kinds (models, atlases, fonts, lang) are dropped from the pack. | MMS targets data packs. The `minecraft:resource` dispatcher mixes both. |
| D6 | 2026-10-31 | Spec pack | Generator inputs are fetched over HTTPS from pinned commits and the generated pack is committed to the repo. | Builds, tests and the WASM bundle then need no network, and every spec change shows up as a reviewable diff. |
| D7 | 2026-10-31 | Toolchain | Go 1.24, the version installed in the dev environment and current stable. | No feature in the plan needs a newer release. |
