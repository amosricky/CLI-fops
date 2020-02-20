package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"os"
)

var linecountFilePath string

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Count line for file.",
	Long: "Count line for file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		for{
			checkFileRes, checkFileErr := checkFile(linecountFilePath)
			if !checkFileRes{
				err = checkFileErr
				break
			}

			checkTextRes, checkTextErr := checkText(linecountFilePath)
			if !checkTextRes{
				err = checkTextErr
				break
			}

			f, openFileErr := os.Open(linecountFilePath)
			if openFileErr != nil{
				err = openFileErr
				break
			}
			defer f.Close()

			var reader io.Reader
			reader = f

			countResult, countErr := lineCounter(reader)
			if countErr != nil{
				err = countErr
				break
			}else {
				cmd.Printf("%v", countResult)
			}
			break
		}
		return err
	},
}

func init() {
	RootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringVarP(&linecountFilePath, "file", "f", "", "File path")
}
