package printer

import (
	"fmt"
	"strings"

	"github.com/chroju/tfdoc/scraping"
)

func PrintTfResourceDoc(tfr *scraping.TfResource) {
	fmt.Println(tfr.Name + "\n" + tfr.Description + "\n\nArgument Reference (= is mandatory):\n")
	printTfResourceArgsDoc(tfr.Args, 0)
}

func printTfResourceArgsDoc(args []*scraping.TfResourceArg, indent int) {
	spaces := strings.Repeat("  ", indent)
	for _, arg := range args {

		var mark string
		if arg.Required {
			mark = "="
		} else {
			mark = "-"
		}

		descln := strings.Join(strings.SplitAfter(arg.Description, ". "), "\n"+spaces+"  ")
		fmt.Println("\n" + spaces + mark + " " + arg.Name + "\n" + spaces + "  " + descln)
		if len(arg.NestedField) > 0 {
			printTfResourceArgsDoc(arg.NestedField, indent+1)
		}
	}
}

func PrintTfResourceSnippet(tfr *scraping.TfResource) {
	fmt.Println("resource \"" + tfr.Name + "\" \"sample\" {")
	printTfResourceArgsSnippet(tfr.Args, 1)
	fmt.Println("}")
}

func printTfResourceArgsSnippet(args []*scraping.TfResourceArg, indent int) {
	spaces := strings.Repeat("  ", indent)
	for _, arg := range args {
		fmt.Println("\n" + strings.Repeat("  ", indent) + "// " + arg.Description)
		if len(arg.NestedField) > 0 {
			fmt.Println(spaces + arg.Name + " {")
			printTfResourceArgsSnippet(arg.NestedField, indent+1)
			fmt.Println(spaces + "}\n")
		} else {
			fmt.Println(spaces + arg.Name + " = " + "\"" + "\"")
		}
	}
}
