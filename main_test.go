package main

import (
	"strings"
	"testing"
)

const (
	commandName = "tfdoc"
)

func TestCLI(t *testing.T) {
	var cases = []struct {
		args             []string
		expectedResult   string
		expectedExitCode int
	}{
		{[]string{commandName, "aws_instance"},
			"Provides an EC2 instance resource.", 0},
		{[]string{commandName, "aws_instance", "-s"},
			"resource \"aws_instance\" \"sample\"", 0},
	}

	for _, c := range cases {
		result, exitCode := run(c.args)
		joinedResult := strings.Join(result, "\n")
		if !strings.Contains(joinedResult, c.expectedResult) {
			t.Errorf("Expected: %s\n Result: %s", c.expectedResult, joinedResult)
		} else if exitCode != c.expectedExitCode {
			t.Errorf("Expected: %d\n Result: %d", c.expectedExitCode, exitCode)
		}
	}
}
