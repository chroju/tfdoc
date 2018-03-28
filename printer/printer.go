package printer

import (
	"fmt"
	"strings"

	"github.com/chroju/tfdoc/scraping"
)

func PrintTfResource(tfr *scraping.TfResource) {
	fmt.Println("resource \"" + tfr.Name + "\" \"sample\" {")
	printTfResourceArgs(tfr.Args, 1)
	fmt.Println("}")
}

func printTfResourceArgs(args []*scraping.TfResourceArg, indent int) {
	for _, arg := range args {
		fmt.Println("\n" + strings.Repeat("  ", indent) + "// " + arg.Description)
		if len(arg.NestedField) > 0 {
			fmt.Println(strings.Repeat("  ", indent) + arg.Name + " {")
			printTfResourceArgs(arg.NestedField, indent+1)
			fmt.Println(strings.Repeat("  ", indent) + "}\n")
		} else {
			fmt.Println(strings.Repeat("  ", indent) + arg.Name + " = " + "\"" + "\"")
		}
	}
}
