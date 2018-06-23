package main

import (
	"fmt"
	"io"
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
	// ExitCodeHelp error code
	ExitCodeHelp
)

type CLI struct {
	outStream, errStream io.Writer
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	exitCode := cli.Run(os.Args)
	os.Exit(exitCode)
}

func (c *CLI) Run(args []string) int {

	// parse flags
	var isSnippet bool
	var isURL bool
	var isList bool
	var isVersion bool
	var needlessComment bool
	var requiredOnly bool

	f := flag.NewFlagSet(args[0], flag.ContinueOnError)
	f.SetOutput(c.errStream)
	f.BoolVar(&needlessComment, "no-comments", false, "use with -s to output without comments")
	f.BoolVar(&requiredOnly, "only-required", false, "use with -s to output only required arguments")
	f.BoolVarP(&isSnippet, "snippet", "s", false, "output snippet format")
	f.BoolVarP(&isURL, "url", "u", false, "output document url")
	f.BoolVarP(&isList, "list", "l", false, "list resources about given provider")
	f.BoolVarP(&isVersion, "version", "v", false, "show version")
	f.Usage = func() {
		fmt.Fprintf(c.errStream, "Usage: %s [option] RESOURCE or PROVIDER\n", name)
		fmt.Fprintf(c.errStream, "\n%s\n\n", description)
		f.PrintDefaults()
	}

	// parse args
	if err := f.Parse(args[1:]); err != nil {
		return ExitCodeOK
	} else if len(args) < 2 {
		f.Usage()
		return ExitCodeHelp
	}

	// -v option
	if isVersion {
		fmt.Fprint(c.outStream, name+" "+version)
		return ExitCodeOK
	}

	if len(f.Args()) < 1 {
		f.Usage()
		return ExitCodeHelp
	}
	target := f.Args()[0]

	// --list option
	var tfResourceType string
	if isList {
		tfResourceType = "provider"
	} else {
		tfResourceType = "resource"
	}

	s, err := scraping.NewScraper(tfResourceType, target)
	if err != nil {
		fmt.Fprintf(c.errStream, "%sn", err.Error())
		return ExitCodeError
	}

	// --url option
	if isURL {
		fmt.Fprint(c.outStream, s.Url)
		return ExitCodeOK
	}

	tfobject, err := s.Scrape()
	if err != nil {
		fmt.Fprintf(c.errStream, "%sn", err.Error())
		return ExitCodeError
	}

	fmt.Fprint(c.outStream, strings.Join(tfobject.Doc(isSnippet, needlessComment, requiredOnly), "\n"))
	return ExitCodeOK
}
