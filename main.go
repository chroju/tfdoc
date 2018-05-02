package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chroju/tfdoc/scraping"
	flag "github.com/spf13/pflag"
)

const (
	name        = "tfdoc"
	description = "output Terraform documents and snippets in command line."
	version     = "0.1.0"
)

const (
	// ExitCodeOK normal code
	ExitCodeOK = iota
	// ExitCodeError error code
	ExitCodeError
)

func main() {
	result, exitCode := run(os.Args)
	fmt.Println(strings.Join(result, "\n"))
	os.Exit(exitCode)
}

func run(args []string) ([]string, int) {
	os.Args = args

	// parse flags
	var isSnippet bool
	var isURL bool
	var isList bool
	var needlessComment bool
	var requiredOnly bool
	flag.BoolVarP(&isSnippet, "snippet", "s", false, "output snippet")
	flag.BoolVarP(&isURL, "url", "u", false, "output document url")
	flag.BoolVarP(&isList, "list", "l", false, "output resource list")
	flag.BoolVar(&needlessComment, "no-comments", false, "output snippets with no comment")
	flag.BoolVar(&requiredOnly, "required-only", false, "output only required arguments")
	flag.Parse()

	if len(flag.Args()) != 1 {
		return []string{""}, ExitCodeError
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
		return []string{err.Error()}, ExitCodeError
	}

	// --url option
	if isURL {
		return []string{s.Url}, ExitCodeOK
	}

	tfobject, err := s.Scrape()
	if err != nil {
		return []string{err.Error()}, ExitCodeError
	}

	return tfobject.Doc(isSnippet, needlessComment, requiredOnly), ExitCodeOK
}
