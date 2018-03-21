package printer

import (
	"fmt"
	"strings"

	"github.com/chroju/tfh/scraping"
)

func PrintTfResource(tfr *scraping.TfResource) {
	fmt.Println("resource \"" + tfr.Name + "\" \"sample\" {")
	PrintTfResourceArgs(tfr.Args, 1)
	fmt.Println("}")
}

func PrintTfResourceArgs(args []*scraping.TfResourceArg, indent int) {
	for _, arg := range args {
		fmt.Println("\n" + strings.Repeat("  ", indent) + "// " + arg.Description)
		if len(arg.Field_name) > 0 {
			fmt.Println("\n" + strings.Repeat("  ", indent) + arg.Name + " {")
			PrintTfResourceArgs(arg.Field, indent+1)
			fmt.Println(strings.Repeat("  ", indent) + "}\n")
		} else {
			fmt.Println(strings.Repeat("  ", indent) + arg.Name + " = " + "\"" + arg.Field_name + "\"")
		}
	}
}
