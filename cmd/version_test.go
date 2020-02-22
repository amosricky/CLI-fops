package cmd_test

import (
	"CLI-fops/cmd"
	"testing"
)

func Test_version(t *testing.T)  {
	output, err := executeCommand(cmd.RootCmd, "version")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "v0.0.1")
}