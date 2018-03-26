package main

import (
	"github.com/chroju/tfdoc/printer"
	"github.com/chroju/tfdoc/scraping"
	flag "github.com/spf13/pflag"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

func main() {
	flag.Parse()
	resource_name := flag.Args()[0]

	url := scraping.GetResourceUrl(resource_name)
	resource := scraping.ScrapingDoc(url)
	resource.Name = resource_name
	printer.PrintTfResource(resource)
}
