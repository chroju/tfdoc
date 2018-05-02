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

func (t *TfResource) Doc(opts ...bool) []string {
	isSnippet := opts[0]
	needlessComment := opts[1]
	requiredOnly := opts[2]
	if isSnippet {
		return t.Snippet(needlessComment, requiredOnly)
	}
	var ret []string
	ret = append(ret, t.Name, t.Description, "", "Argument Reference (= is mandatory):", "")
	ret = append(ret, printTfResourceArgsDoc(t.Args, 0)...)
	return ret
}

func printTfResourceArgsDoc(args []*tfResourceArg, indent int) []string {
	var ret []string

	spaces := strings.Repeat("  ", indent)
	for _, arg := range args {
		ret = append(ret, "")
		var mark string
		if arg.Required {
			mark = "="
		} else {
			mark = "-"
		}

		descln := strings.SplitAfter(arg.Description, ". ")
		for i, v := range descln {
			descln[i] = spaces + "  " + v
		}
		ret = append(ret, spaces+mark+" "+arg.Name)
		ret = append(ret, descln...)
		if len(arg.NestedField) > 0 {
			ret = append(ret, printTfResourceArgsDoc(arg.NestedField, indent+1)...)
		}

	}

	return ret
}

func (t *TfResource) Snippet(needlessComment bool, requiredOnly bool) []string {
	var ret []string
	ret = append(ret, "resource \""+t.Name+"\" \"sample\" {")
	ret = append(ret, printTfResourceArgsSnippet(t.Args, 1, needlessComment, requiredOnly)...)
	ret = append(ret, "}")
	return ret
}

func printTfResourceArgsSnippet(args []*tfResourceArg, indent int, needlessComment bool, requiredOnly bool) []string {
	var ret []string

	spaces := strings.Repeat("  ", indent)
	for _, arg := range args {
		if requiredOnly && !arg.Required {
			continue
		}
		if !needlessComment {
			ret = append(ret, strings.Repeat("  ", indent)+"// "+arg.Description)
		}
		if len(arg.NestedField) > 0 {
			ret = append(ret, spaces+arg.Name+" {")
			ret = append(ret, printTfResourceArgsSnippet(arg.NestedField, indent+1, needlessComment, requiredOnly)...)
			ret = append(ret, spaces+"}")
		} else {
			ret = append(ret, spaces+arg.Name+" = "+"\""+"\"")
		}

		ret = append(ret, "")
	}

	return ret
}
