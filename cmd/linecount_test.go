package cmd_test

import (
	"CLI-fops/cmd"
	"testing"
)

func Test_linecount(t *testing.T)  {
	output, err := executeCommand(cmd.RootCmd, "linecount", "--file", "../myfile.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "4")
}