package scraping

import (
	"fmt"
	"net/http"
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
	Name        string
	Description string
	Args        []*TfResourceArg
}

func GetResourceUrl(resource string) (string, error) {
	if !strings.Contains(resource, "_") {
		err := fmt.Errorf("resource name is invalid.")
		return "", err
	}

	splited_resource := strings.SplitN(resource, "_", 2)
	return "https://www.terraform.io/docs/providers/" +
		splited_resource[0] + "/r/" + splited_resource[1] + ".html", nil
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

func ScrapingDoc(url string) (*TfResource, error) {
	ret := &TfResource{Name: ""}

	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("URL Query error : %s", err)
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("Status code error : %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		err = fmt.Errorf("URL Query error : %s", err)
		return nil, err
	}

	ret.Description = strings.TrimSpace(doc.Find("#inner > p").First().Text())
	ret.Description = strings.Replace(ret.Description, "\n", "", -1)
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

	return ret, nil
}

func ScrapingResourceList(provider string) (string, error) {
	result := ""
	res, err := http.Get("https://www.terraform.io/docs/providers/" + provider + "/index.html")
	if err != nil {
		err = fmt.Errorf("Provider error : %s", err)
		return "", err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("Status code error : %d %s", res.StatusCode, res.Status)
		return "", err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)

	doc.Find(".nav-visible").Each(func(i int, selection *goquery.Selection) {
		if i > 1 {
			selection.Find("li").Each(func(_ int, li *goquery.Selection) {
				result = result + strings.TrimSpace(li.Text()) + "\n"
			})
		}
	})

	return result, nil
}
