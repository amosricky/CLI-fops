package cmd_test

import (
	"CLI-fops/cmd"
	"testing"
)

func Test_checksum(t *testing.T)  {
	output, err := executeCommand(cmd.RootCmd, "checksum", "--file", "../myfile.txt", "--md5")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(output) != 32{
		t.Error("Not correct")
	}
}
