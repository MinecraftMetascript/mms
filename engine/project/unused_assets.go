package project

import (
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// assetExts are the kinds of file the asset helpers copy into a pack.
var assetExts = []string{".png", ".ogg", ".nbt"}

// UnusedAssets lists the project's .png, .ogg and .nbt files that no
// Texture, Sound or Structure call and no icon uses, so the packs leave
// them out, with the severity of the unused-asset rule. A file has no
// source line to report at, so editors show these on the file itself.
// The list is empty when the rule is off, and while the project has
// errors, since a call in a broken declaration is not counted as a use.
func (r *Result) UnusedAssets() ([]string, syntax.Severity) {
	sev := r.severity(RuleByID("unused-asset"))
	if sev == 0 || r.files == nil || r.HasErrors() {
		return nil, sev
	}
	used := map[string]bool{r.cfg.Icon: true}
	for _, u := range r.assets {
		used[u.Src] = true
	}
	var out []string
	for _, ext := range assetExts {
		for _, p := range r.cfg.AssetFiles(r.files, ext) {
			if !used[p] {
				out = append(out, p)
			}
		}
	}
	return out, sev
}

// UnusedAssetMessage says why a file is reported as unused.
func UnusedAssetMessage(path string) string {
	return "'" + path + "' is not used: no Texture, Sound or Structure call names it, so the pack leaves it out"
}
