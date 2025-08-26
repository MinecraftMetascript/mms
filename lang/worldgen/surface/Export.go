package surface

import (
	"encoding/json"
	"errors"

	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_conditions"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_rules"
	"github.com/minecraftmetascript/mms/lib"
)

func exportRules(ns *lib.Namespace, rootDir *lib.FileTreeLike) error {
	rulesDir := rootDir.MkDir("surface_rules", nil)

	for name, rule := range lib.AllOf[surface_rules.SurfaceRule](ns) {
		str, err := json.MarshalIndent(rule.Value, "", "  ")
		if err != nil {
			return err
		}
		rulesDir.MkFile(
			name+".json",
			string(str),
			rule,
		)
	}

	return nil
}

func exportConditions(ns *lib.Namespace, rootDir *lib.FileTreeLike) error {
	conditionsDir := rootDir.MkDir("surface_conditions", nil)

	for name, condition := range lib.AllOf[surface_conditions.SurfaceCondition](ns) {
		str, err := json.MarshalIndent(condition.Value, "", "  ")
		if err != nil {
			return err
		}
		conditionsDir.MkFile(
			name+".json",
			string(str),
			condition,
		)
	}

	return nil
}

func Export(ns *lib.Namespace, rootDir *lib.FileTreeLike) error {
	if !rootDir.IsDir {
		return errors.New("root must be a directory")
	}

	debugDir := rootDir.MkDir("_debug", nil)
	exportRules(ns, debugDir)
	exportConditions(ns, debugDir)

	return nil

}
