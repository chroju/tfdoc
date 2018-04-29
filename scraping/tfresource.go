package scraping

import (
	"strings"
)

type tfResourceArg struct {
	Name        string
	Description string
	NestedField []*tfResourceArg
	Required    bool
}

type TfResource struct {
	Name        string
	Description string
	Args        []*tfResourceArg
}

func (t *TfResource) Doc(isSnippet bool) string {
	if isSnippet {
		return t.Snippet()
	}
	var ret string
	ret = t.Name + "\n" + t.Description + "\n\nArgument Reference (= is mandatory):\n"
	ret += printTfResourceArgsDoc(t.Args, 0)
	return ret
}

func printTfResourceArgsDoc(args []*tfResourceArg, indent int) string {
	ret := ""

	spaces := strings.Repeat("  ", indent)
	for _, arg := range args {
		var mark string
		if arg.Required {
			mark = "="
		} else {
			mark = "-"
		}

		descln := strings.Join(strings.SplitAfter(arg.Description, ". "), "\n"+spaces+"  ")
		ret += "\n" + spaces + mark + " " + arg.Name + "\n" + spaces + "  " + descln + "\n"
		if len(arg.NestedField) > 0 {
			ret += printTfResourceArgsDoc(arg.NestedField, indent+1)
		}
	}

	return ret
}

func (t *TfResource) Snippet() string {
	var ret string
	ret = "resource \"" + t.Name + "\" \"sample\" {"
	ret += printTfResourceArgsSnippet(t.Args, 1)
	ret += "}"
	return ret
}

func printTfResourceArgsSnippet(args []*tfResourceArg, indent int) string {
	ret := ""

	spaces := strings.Repeat("  ", indent)
	for _, arg := range args {
		ret += "\n" + strings.Repeat("  ", indent) + "// " + arg.Description + "\n"
		if len(arg.NestedField) > 0 {
			ret += spaces + arg.Name + " {"
			ret += printTfResourceArgsSnippet(arg.NestedField, indent+1)
			ret += spaces + "}\n"
		} else {
			ret += spaces + arg.Name + " = " + "\"" + "\"" + "\n"
		}
	}

	return ret
}
