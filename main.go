package main

import (
	"fmt"

	"github.com/chroju/tfh/printer"
	"github.com/chroju/tfh/scraping"
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
	resource.name = resource_name
	fmt.Println(resource.name)
	fmt.Println(resource_name)
	printer.PrintTfResource(resource)
}
