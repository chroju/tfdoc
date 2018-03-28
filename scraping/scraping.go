package scraping

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type TfResourceArg struct {
	Name        string
	Description string
	NestedField []*TfResourceArg
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

func scrapingResourceList(li *goquery.Selection) *TfResourceArg {
	a := &TfResourceArg{}
	a.Name = li.Find("a > code").Text()
	a.Description = strings.TrimSpace(strings.SplitN(li.Text(), "-", 2)[1])
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

	doc.Find("#inner > ul").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			selection.Children().Each(func(_ int, li *goquery.Selection) {
				arg := scrapingResourceList(li)
				ret.Args = append(ret.Args, arg)
			})
		} else {
			fieldName := selection.Prev().Find("code,strong").Text()
			for i, arg := range ret.Args {
				if arg.Name == fieldName {
					selection.Children().Each(func(_ int, li *goquery.Selection) {
						ret.Args[i].NestedField = append(ret.Args[i].NestedField, scrapingResourceList(li))
					})
				}
			}
		}
	})

	return ret
}
