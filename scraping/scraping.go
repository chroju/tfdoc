package scraping

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type TfResourceArg struct {
	Name        string
	Description string
	Field_name  string
	Field       []*TfResourceArg
	Required    bool
}

type TfResource struct {
	Name string
	Args []*TfResourceArg
}

func GetResourceUrl(resource string) string {
	splited_resource := strings.SplitN(resource, "_", 2)
	return "https://www.terraform.io/docs/providers/" +
		splited_resource[0] + "/r/" + splited_resource[1] + ".html"
}

func ScrapingResourceList(li *goquery.Selection) *TfResourceArg {
	a := &TfResourceArg{Field_name: ""}
	a.Name = li.Find("a > code").Text()
	a.Description = "(" + strings.SplitN(li.Text(), "(", 2)[1]
	a.Description = strings.Replace(a.Description, "\n", "", -1)
	if strings.Contains(strings.SplitN(li.Text(), " ", 3)[2], "Required") {
		a.Required = true
	} else {
		a.Required = false
	}
	return a
}

func ScrapingDoc(url string) *TfResource {
	ret := &TfResource{Name: ""}

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
				if strings.Contains(arg.Description, "below for") {
					start_at := strings.Index(arg.Description, "See") + 4
					end_at := strings.LastIndex(arg.Description, "below") - 1

					arg.Field_name = strings.Replace(strings.ToLower(arg.Description[start_at:end_at]), " ", "-", -1)
				}
				ret.Args = append(ret.Args, arg)
			})
		}

		attr, _ := selection.Attr("id")
		if selection.Is("h3") && attr != "example" {
			for i, arg := range ret.Args {
				if arg.Field_name == attr {
					selection.NextAllFiltered("ul").First().Children().Each(func(_ int, li *goquery.Selection) {
						ret.Args[i].Field = append(ret.Args[i].Field, ScrapingResourceList(li))
					})
				}
			}
		}
	})

	return ret
}
