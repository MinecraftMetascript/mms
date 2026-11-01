package project

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// SchemaURL is where the docs site serves the mms.json JSON Schema.
// `mms init` writes it as "$schema" so editors without the VS Code
// extension (Zed, JetBrains, ...) still validate and complete mms.json.
const SchemaURL = "https://minecraftmetascript.github.io/mms/mms.schema.json"

// schemaNode is the subset of JSON Schema that ConfigSchema uses. VS Code
// reads markdownDescription for hovers and enumDescriptions for
// completion details; other editors fall back to description.
type schemaNode struct {
	Schema               string                 `json:"$schema,omitempty"`
	ID                   string                 `json:"$id,omitempty"`
	Title                string                 `json:"title,omitempty"`
	Description          string                 `json:"description,omitempty"`
	MarkdownDescription  string                 `json:"markdownDescription,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	Format               string                 `json:"format,omitempty"`
	Pattern              string                 `json:"pattern,omitempty"`
	PatternErrorMessage  string                 `json:"patternErrorMessage,omitempty"`
	Default              any                    `json:"default,omitempty"`
	Examples             []any                  `json:"examples,omitempty"`
	Enum                 []string               `json:"enum,omitempty"`
	EnumDescriptions     []string               `json:"enumDescriptions,omitempty"`
	Items                *schemaNode            `json:"items,omitempty"`
	Properties           map[string]*schemaNode `json:"properties,omitempty"`
	AdditionalProperties *bool                  `json:"additionalProperties,omitempty"`
	Ref                  string                 `json:"$ref,omitempty"`
	Defs                 map[string]*schemaNode `json:"definitions,omitempty"`
}

// md fills description and markdownDescription from one Markdown text.
func md(n *schemaNode, text string) *schemaNode {
	n.Description = text
	n.MarkdownDescription = text
	return n
}

// ConfigSchema returns the JSON Schema for mms.json, generated from Config,
// DefaultConfig and the lint rules so it cannot drift from the loader.
// TestConfigSchemaFiles keeps the committed copies in sync.
func ConfigSchema() []byte {
	def := DefaultConfig()
	no := false

	severityNames := map[syntax.Severity]string{
		syntax.SeverityError: "error", syntax.SeverityWarning: "warning",
		syntax.SeverityInfo: "info", syntax.SeverityHint: "hint",
	}
	lint := map[string]*schemaNode{}
	for _, r := range Rules {
		lint[r.ID] = md(&schemaNode{Ref: "#/definitions/severity"},
			fmt.Sprintf("%s\n\nDefault: `%s`.", r.Doc, severityNames[r.Severity]))
	}

	s := &schemaNode{
		Schema: "http://json-schema.org/draft-07/schema#",
		ID:     SchemaURL,
		Title:  "MMS project (mms.json)",
		Type:   "object",
		Properties: map[string]*schemaNode{
			"$schema": {Type: "string", Description: "The JSON Schema for this file.", Default: SchemaURL},
			"namespace": md(&schemaNode{
				Type: "string", Pattern: nsPattern.String(), Default: def.Namespace,
				PatternErrorMessage: "A namespace is lowercase letters, digits, '_', '-' or '.'.",
			}, "The default namespace for files without a `Namespace` declaration."),
			"description": md(&schemaNode{Type: "string", Default: def.Description},
				"The data pack's description, shown in the game's pack list (`pack.mcmeta`)."),
			"output": md(&schemaNode{Type: "string", Default: def.Output},
				"The folder `mms build` writes the data pack to, relative to the project."),
			"sources": md(&schemaNode{
				Type:     "array",
				Items:    md(&schemaNode{Type: "string", Examples: []any{"src/**/*.mms"}}, "A glob relative to the project: `*` matches within a folder, `**` any number of folders and `?` one character."),
				Examples: []any{[]string{"src/**/*.mms"}},
			}, "Glob patterns for the project's .mms files, relative to the project; `**` matches any number of folders. The default is every .mms file (`**/*.mms`), skipping the output folders, hidden folders and `node_modules`."),
			"lint": md(&schemaNode{
				Type:                 "object",
				Properties:           lint,
				AdditionalProperties: &no,
			}, "Lint rule severities by rule id. Each rule can be `error`, `warning`, `info`, `hint` or `off`."),
			"icon": md(&schemaNode{
				Type: "string", Pattern: `\.[pP][nN][gG]$`, Examples: []any{"icon.png"},
				PatternErrorMessage: "The icon must be a .png file.",
			}, "A PNG in the project that becomes `pack.png` in both the data pack and the resource pack."),
			"resourcePack": md(&schemaNode{
				Type: "object",
				Properties: map[string]*schemaNode{
					"output": md(&schemaNode{Type: "string"},
						"The folder `mms build` writes the resource pack to, relative to the project. The default is the data pack's `output` with `-resources` appended."),
					"description": md(&schemaNode{Type: "string"},
						"The resource pack's description. The default is the data pack's `description`."),
				},
				AdditionalProperties: &no,
			}, "Settings for the resource pack `mms build` writes when the project uses `Texture`, `Sound` or other resource pack declarations."),
		},
		AdditionalProperties: &no,
		Defs: map[string]*schemaNode{
			"severity": {
				Type:        "string",
				Description: "A lint severity.",
				Enum:        []string{"error", "warning", "info", "hint", "off", "warn"},
				EnumDescriptions: []string{
					"Report as an error; `mms build` fails.",
					"Report as a warning.",
					"Report as information.",
					"Report as a hint (shown faintly in editors).",
					"Turn the rule off.",
					"Same as `warning`.",
				},
			},
		},
	}
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return append(b, '\n')
}
