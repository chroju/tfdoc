package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	var cases = []struct {
		args             []string
		expectedStdout   string
		expectedStderr   string
		expectedExitCode int
	}{
		{strings.Split("tfdoc", " "),
			"", "Usage", 2},
		{strings.Split("tfdoc -h", " "),
			"", "Usage", 0},
		{strings.Split("tfdoc --help", " "),
			"", "Usage", 0},
		{strings.Split("tfdoc -v", " "),
			"", "tfdoc", 0},
		{strings.Split("tfdoc --version", " "),
			"", "tfdoc", 0},
		{strings.Split("tfdoc -l aws", " "),
			"aws_internet_gateway", "", 0},
		{strings.Split("tfdoc --list aws", " "),
			"aws_internet_gateway", "", 0},
		{strings.Split("tfdoc -l", " "),
			"", "Usage", 2},
		{strings.Split("tfdoc aws_instance", " "),
			"Provides an EC2 instance resource.", "", 0},
		{strings.Split("tfdoc aws_instance -s", " "),
			"resource \"aws_instance\" \"sample\"", "", 0},
		{strings.Split("tfdoc aws_instance --snippet", " "),
			"resource \"aws_instance\" \"sample\"", "", 0},
		{strings.Split("tfdoc aws_instance -u", " "),
			"https://www.terraform.io/docs/providers/aws/r/instance.html", "", 0},
		{strings.Split("tfdoc aws_instance --url", " "),
			"https://www.terraform.io/docs/providers/aws/r/instance.html", "", 0},
		// {strings.Split("tfdoc -v", " "),
		// 	version, 0},
		// {strings.Split("tfdoc -l", " "),
		// 	"azurerm_container_group", 0},
	}

	for i, c := range cases {
		outBuffer := new(bytes.Buffer)
		errBuffer := new(bytes.Buffer)
		cli := &CLI{outStream: outBuffer, errStream: errBuffer}
		exitCode := cli.Run(c.args)
		stdout := cli.outStream.(*bytes.Buffer).String()
		// stderr := cli.errStream.(*bytes.Buffer).String()
		if !strings.Contains(stdout, c.expectedStdout) {
			t.Errorf("Case: %d\n Expected: %s\n Result: %s", i, c.expectedStdout, stdout)
		} else if exitCode != c.expectedExitCode {
			t.Errorf("Case: %d\n Expected: %d\n Result: %d", i, c.expectedExitCode, exitCode)
		}
	}
}
