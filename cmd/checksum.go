package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/spf13/cobra"
	"io"
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

		checkFileRes, checkFileErr := checkFile(checksumFilePath)
		if !checkFileRes{
			return checkFileErr
		}

		fileContent, readFileErr := ioutil.ReadFile(checksumFilePath)
		if readFileErr != nil{
			return readFileErr
		}

		if md5Flag{
			cmd.Printf(genMd5(string(fileContent)))
		}

		if sha1Flag{
			cmd.Printf(genSha1(string(fileContent)))
		}

		if sha256Flag{
			cmd.Printf(genSha256(string(fileContent)))
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringVarP(&checksumFilePath, "file", "f", "", "File path")
	checksumCmd.Flags().BoolVarP(&md5Flag, "md5", "", false, "Get checksum in hash function-sha256")
	checksumCmd.Flags().BoolVarP(&sha1Flag, "sha1", "", false, "Get checksum in hash function-sha1")
	checksumCmd.Flags().BoolVarP(&sha256Flag, "sha256", "", false, "Get checksum in hash function-sha256")
}

func genMd5(content string) string {
	h := md5.New()
	io.WriteString(h, string(content)) //write str(f) to h
	md5Hash := fmt.Sprintf("%x", h.Sum(nil))
	return md5Hash
}

func genSha1(content string) string {
	h := sha1.New()
	io.WriteString(h, string(content)) //write str(f) to h
	sha1Hash := fmt.Sprintf("%x", h.Sum(nil))
	return sha1Hash
}

func genSha256(content string) string {
	h := sha256.New()
	io.WriteString(h, string(content)) //write str(f) to h
	sha256Hash := fmt.Sprintf("%x", h.Sum(nil))
	return sha256Hash
}
