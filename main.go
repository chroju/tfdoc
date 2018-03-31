package main

import (
	"fmt"
	"os"

	"github.com/chroju/tfdoc/printer"
	"github.com/chroju/tfdoc/scraping"
	flag "github.com/spf13/pflag"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

func main() {
	var isSnippet bool
	var isUrl bool
	flag.BoolVarP(&isSnippet, "snippet", "s", false, "output snippet")
	flag.BoolVarP(&isUrl, "url", "u", false, "output snippet")
	flag.Parse()

	if len(flag.Args()) != 1 {
		os.Exit(ExitCodeError)
	}

	resourceName := flag.Args()[0]

	url, err := scraping.GetResourceUrl(resourceName)
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitCodeError)
	}

	if isUrl {
		fmt.Println(url)
		os.Exit(ExitCodeOK)
	}

	resource, err := scraping.ScrapingDoc(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitCodeError)
	}
	resource.Name = resourceName

	if isSnippet {
		printer.PrintTfResourceSnippet(resource)
	} else if isUrl {
	} else {
		printer.PrintTfResourceDoc(resource)
	}

	os.Exit(ExitCodeOK)
}
