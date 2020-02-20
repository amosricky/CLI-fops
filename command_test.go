package main

import (
	"CLI-fops/cmd"
	"bytes"
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)
	c, err = root.ExecuteC()
	return c, buf.String(), err
}

func checkStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}

func Test_version(t *testing.T)  {
	output, err := executeCommand(cmd.RootCmd, "version")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "v0.0.1")
}

func Test_checksum(t *testing.T)  {
	output, err := executeCommand(cmd.RootCmd, "checksum", "--file", "./myfile.txt", "--md5")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(output) != 32{
		t.Error("Not correct")
	}
}

func Test_linecount(t *testing.T)  {
	output, err := executeCommand(cmd.RootCmd, "linecount", "--file", "./myfile.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "4")
}
