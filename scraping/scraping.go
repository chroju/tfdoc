package scraping

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type TfScraper struct {
	Name    string
	DocType string
	Url     string
}

type TfObject interface {
	Doc(...bool) (doc []string)
}

func NewScraper(docType string, name string) (*TfScraper, error) {
	s := TfScraper{Name: name, DocType: docType}

	err := s.convertDocUrl()
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// convert resource or provider name to document url.
func (s *TfScraper) convertDocUrl() error {
	var url string

	switch s.DocType {
	case "provider":
		url = "https://www.terraform.io/docs/providers/" + s.Name + "/index.html"
	case "resource":
		if !strings.Contains(s.Name, "_") {
			return fmt.Errorf("resource name is invalid.")
		}

		splited := strings.SplitN(s.Name, "_", 2)
		url = "https://www.terraform.io/docs/providers/" + splited[0] + "/r/" + splited[1] + ".html"
	}

	_, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Provider error : %s", err)
	}

	s.Url = url
	return nil
}

// scrape from web
func (s *TfScraper) Scrape() (TfObject, error) {
	res, err := http.Get(s.Url)
	defer res.Body.Close()

	if err != nil {
		err = fmt.Errorf("URL Query error : %s", err)
		return nil, err
	}

	if res.StatusCode != 200 {
		err = fmt.Errorf("Status code error : %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	var tfo TfObject
	switch s.DocType {
	case "resource":
		tfo, err = scrapeTfResource(s.Name, res)
		if err != nil {
			err = fmt.Errorf("Scraping error : %s", err)
			return nil, err
		}
	case "provider":
		tfo, err = scrapeTfProvider(s.Name, res)
		if err != nil {
			err = fmt.Errorf("Scraping error : %s", err)
			return nil, err
		}
	}

	return tfo, nil
}

func scrapeTfResource(name string, res *http.Response) (*TfResource, error) {
	var ret = TfResource{Name: name}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		err = fmt.Errorf("HTML Read error: %s", err)
		return nil, err
	}

	ret.Description = strings.Replace(strings.TrimSpace(doc.Find("#inner > p").First().Text()), "\n", "", -1)
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

	return &ret, nil
}

func scrapingResourceList(li *goquery.Selection) *tfResourceArg {
	a := &tfResourceArg{}
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

func scrapeTfProvider(name string, res *http.Response) (*TfProvider, error) {
	var ret = TfProvider{Name: name}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		err = fmt.Errorf("HTML Read error: %s", err)
		return nil, err
	}

	doc.Find(".docs-sidenav > li").Each(func(i int, selection *goquery.Selection) {
		if !(strings.Contains(selection.Text(), "Guides") || strings.Contains(selection.Text(), "Data Sources") || strings.Contains(selection.Text(), "Provider")) {
			selection.Find(".nav-visible > li").Each(func(_ int, li *goquery.Selection) {
				ret.ResourceList = append(ret.ResourceList, strings.TrimSpace(li.Text()))
			})
		}
	})

	return &ret, nil
}
