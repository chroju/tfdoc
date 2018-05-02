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
	var isVersion bool
	var needlessComment bool
	var requiredOnly bool
	// snippetCommand := flag.NewFlagSet("snippet", flag.ExitOnError)
	flag.BoolVar(&needlessComment, "no-comments", false, "use with -s to output without comments")
	flag.BoolVar(&requiredOnly, "only-required", false, "use with -s to output only required arguments")
	flag.BoolVarP(&isSnippet, "snippet", "s", false, "output snippet format")
	flag.BoolVarP(&isURL, "url", "u", false, "output document url")
	flag.BoolVarP(&isList, "list", "l", false, "list resources about given provider")
	flag.BoolVarP(&isVersion, "version", "v", false, "show version")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [option] RESOURCE or PROVIDER\n", name)
		fmt.Fprintf(os.Stderr, "\n%s\n\n", description)
		flag.PrintDefaults()
	}
	flag.Parse()

	// -v option
	if isVersion {
		return []string{name + " " + version}, ExitCodeError
	}

	if len(flag.Args()) != 1 {
		flag.Usage()
		return []string{""}, ExitCodeError
	}

	target := flag.Args()[0]

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
