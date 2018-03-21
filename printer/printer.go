package printer

import (
	"fmt"
	"strings"

	"github.com/chroju/tfh/scraping"
)

func PrintTfResource(tfr *scraping.TfResource) {
	fmt.Println("resource \"" + tfr.name + "\" \"sample\" {")
	PrintTfResourceArgs(tfr.args, 1)
	fmt.Println("}")
}

func PrintTfResourceArgs(args []*scraping.TfResourceArg, indent int) {
	for _, arg := range args {
		fmt.Println("\n" + strings.Repeat("  ", indent) + "// " + arg.description)
		if len(arg.field_name) > 0 {
			fmt.Println("\n" + strings.Repeat("  ", indent) + arg.name + " {")
			PrintTfResourceArgs(arg.field, indent+1)
			fmt.Println(strings.Repeat("  ", indent) + "}\n")
		} else {
			fmt.Println(strings.Repeat("  ", indent) + arg.name + " = " + "\"" + arg.field_name + "\"")
		}
	}
}
