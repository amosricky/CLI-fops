package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var checksumFilePath string
var md5Flag bool
var sha1Flag bool
var sha256Flag bool

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Get checksum",
	Long: "Get checksum",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		for{
			checkFileRes, checkFileErr := checkFile(checksumFilePath)
			if !checkFileRes{
				err = checkFileErr
				break
			}

			fileContent, readFileErr := ioutil.ReadFile(checksumFilePath)
			if readFileErr != nil{
				err = readFileErr
				break
			}

			if md5Flag{
				fmt.Println(genMd5(string(fileContent)))
			}

			if sha1Flag{
				fmt.Println(genSha1(string(fileContent)))
			}

			if sha256Flag{
				fmt.Println(genSha256(string(fileContent)))
			}
			break
		}
		return err
	},
}

func init() {
	RootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringVarP(&checksumFilePath, "file", "f", "", "File path")
	checksumCmd.Flags().BoolVarP(&md5Flag, "md5", "", false, "Get checksum in hash function-sha256")
	checksumCmd.Flags().BoolVarP(&sha1Flag, "sha1", "", false, "Get checksum in hash function-sha1")
	checksumCmd.Flags().BoolVarP(&sha256Flag, "sha256", "", false, "Get checksum in hash function-sha256")
}
