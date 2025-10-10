package spec

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
)

func GenerateSpecString(blocks *BlockSpecList) string {
	out := []string{
		"# MMS Language Specification",
		"This document describes the language specification for MMS. It is generated using the `mms spec` command.\n",
		"## Blocks",
	}

	for name, specs := range blocks.All() {
		out = append(out,
			fmt.Sprintf("### %s  \n", name),
			fmt.Sprintf("#### Allowed Values:  \n"),
			fmt.Sprintf("---"),
		)

		for _, s := range specs.AllowedValues.All() {
			out = append(out, mkSpecTxt(s)...)
		}
	}

	return strings.Join(out, "\n")
}
func mkSpecTxt(s ValueSpec) []string {
	out := []string{}
	includeUsage := true
	usage := strings.Trim(s.UsageStr(), "\n")
	if l, ok := s.(*ValueSpecList); ok {
		for _, v := range l.All() {
			out = append(out, mkSpecTxt(v)...)
		}
		return out
	}

	if fn, ok := s.(FunctionSpec); ok {
		out = append(out, fmt.Sprintf("##### Function: %s  ", fn.Name))
		usage = strings.Join(fn.Usage(), "\n")
	}
	if _, ok := s.(*NumberSpec); ok {
		out = append(out, fmt.Sprintf("##### Number  "))
		includeUsage = false
	}
	if r, ok := s.(*ReferenceSpec); ok {
		out = append(out, fmt.Sprintf("##### Reference: %s  ", r.Kind))
		usage = "namespace:name"
	}
	if l, ok := s.(*ListSpec); ok {
		out = append(out, fmt.Sprintf("##### List: %s  ", l.Kind))
	}

	if c, ok := s.(*ConditionalSpec); ok {
		out = append(out, fmt.Sprintf("##### Conditional: %s  ", c.Kind))
	}

	if h, ok := s.(ast.Helpful); ok {
		help := h.GetHelp()
		if help != "" {
			out = append(out,
				fmt.Sprintf("###### About  "),
				fmt.Sprintf("%s  ", help),
			)
		}
	}

	if includeUsage {
		if usage != "" {
			out = append(out,
				fmt.Sprintf("###### Usage  "),
				fmt.Sprintf("```"),
			)
			for _, line := range strings.Split(usage, "\n") {
				out = append(out, "    "+line)
			}
			out = append(out, "```")
		}
	}
	out = append(out, "---")
	return out
}
