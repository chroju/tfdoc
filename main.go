package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	// flag "github.com/spf13/pflag"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

type TfResourceArg struct {
	name        string
	description string
	field_name  string
	field       []*TfResourceArg
	required    bool
}

type TfResource struct {
	name string
	args []*TfResourceArg
}

func GetResourceUrl(resource string) string {
	splited_resource := strings.SplitN(resource, "_", 2)
	return "https://www.terraform.io/docs/providers/" +
		splited_resource[0] + "/r/" + splited_resource[1] + ".html"
}

func ScrapingResourceList(li *goquery.Selection) *TfResourceArg {
	a := &TfResourceArg{field_name: ""}
	a.name = li.Find("a > code").Text()
	a.description = "(" + strings.SplitN(li.Text(), "(", 2)[1]
	a.description = strings.Replace(a.description, "\n", "", -1)
	if strings.Contains(strings.SplitN(li.Text(), " ", 3)[2], "Required") {
		a.required = true
	} else {
		a.required = false
	}
	return a
}

func ScrapingDoc(url string) *TfResource {
	ret := &TfResource{}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		// return fmt.Errorf("error: " + url)
		// return "error: " + url
	}

	inner := doc.Find("#inner").Children()

	inner.Each(func(_ int, selection *goquery.Selection) {
		if strings.Contains(selection.Text(), "The following arguments") {
			selection.Next().Children().Each(func(_ int, li *goquery.Selection) {
				arg := ScrapingResourceList(li)
				if strings.Contains(arg.description, "below for") {
					start_at := strings.Index(arg.description, "See") + 4
					end_at := strings.LastIndex(arg.description, "below") - 1

					arg.field_name = strings.Replace(strings.ToLower(arg.description[start_at:end_at]), " ", "-", -1)
				}
				ret.args = append(ret.args, arg)
			})
		}

		attr, _ := selection.Attr("id")
		if selection.Is("h3") && attr != "example" {
			for i, arg := range ret.args {
				if arg.field_name == attr {
					selection.NextAllFiltered("ul").First().Children().Each(func(_ int, li *goquery.Selection) {
						ret.args[i].field = append(ret.args[i].field, ScrapingResourceList(li))
					})
				}
			}
		}
	})

	return ret
}

func PrintTfResource(tfr *TfResource) {
	fmt.Println("resource \"" + tfr.name + "\" \"sample\" {")
	PrintTfResourceArgs(tfr.args, 1)
	fmt.Println("}")
}

func PrintTfResourceArgs(args []*TfResourceArg, indent int) {
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

func main() {
	resource_name := "aws_route53_record"
	url := GetResourceUrl(resource_name)
	s := ScrapingDoc(url)
	s.name = resource_name
	PrintTfResource(s)
}
