package main

import (
	"CLI-fops/cli"
	"CLI-fops/setting"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	setting.Setup()
	fmt.Println("$ You could use [help] to get some instruction or [exit] to leave the terminal.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdString = strings.TrimSuffix(cmdString, "\n")
		arrCommandStr := strings.Fields(cmdString)

		newOsArgs := []string{os.Args[0]}
		newOsArgs = append(newOsArgs, arrCommandStr...)
		os.Args = newOsArgs

		if len(os.Args) == 1{
			continue
		}

		switch os.Args[1] {
		case "help":
			os.Args = []string{os.Args[0], "-h"}
			cli.Execute()
		case "exit":
			cli.Exit()
		default:
			cli.Execute()
		}
	}
}