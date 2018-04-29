package main

import (
	"fmt"
	"os"

	"github.com/chroju/tfdoc/scraping"
	flag "github.com/spf13/pflag"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

func main() {
	result, exitCode := run(os.Args)
	fmt.Println(result)
	os.Exit(exitCode)
}

func run(args []string) (string, int) {
	os.Args = args

	// parse flags
	var isSnippet bool
	var isURL bool
	var isList bool
	flag.BoolVarP(&isSnippet, "snippet", "s", false, "output snippet")
	flag.BoolVarP(&isURL, "url", "u", false, "output document url")
	flag.BoolVarP(&isList, "list", "l", false, "output resource list")
	flag.Parse()

	if len(flag.Args()) != 1 {
		return "", ExitCodeError
	}

	target := flag.Args()[0]

	// -v and -h option

	// --list option
	var tfResourceType string
	if isList {
		tfResourceType = "provider"
	} else {
		tfResourceType = "resource"
	}

	s, err := scraping.NewScraper(tfResourceType, target)
	if err != nil {
		return err.Error(), ExitCodeError
	}

	// --url option
	if isURL {
		return s.Url, ExitCodeOK
	}

	tfobject, err := s.Scrape()
	if err != nil {
		return err.Error(), ExitCodeError
	}

	return tfobject.Doc(isSnippet), ExitCodeOK
}
