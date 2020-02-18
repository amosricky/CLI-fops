package cli

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/tools/godoc/util"
	"io"
	"io/ioutil"
	"os"
	"crypto/md5"
)

var filePath string
var md5Flag bool
var sha1Flag bool
var sha256Flag bool

var rootCmd = &cobra.Command{Use: "",}

var fopsCmd = &cobra.Command{
	Use: "fops",
	Short: "File Ops",
	Long: "File Ops",}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Get system version.",
	Long: "Get system version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("versionCmd 123")
	},
}

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Get checksum",
	Long: "Get checksum",
	Run: func(cmd *cobra.Command, args []string) {

		for{
			checkFileRes, checkFileErr := checkFile(filePath)
			if !checkFileRes{
				fmt.Println(checkFileErr)
				break
			}

			fileContent, err := ioutil.ReadFile(filePath)
			if err != nil{
				fmt.Println(err)
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
		refresh()
	},
}

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Count line for file.",
	Long: "Count line for file.",
	Run: func(cmd *cobra.Command, args []string) {

		for{
			checkFileRes, checkFileErr := checkFile(filePath)
			if !checkFileRes{
				fmt.Println(checkFileErr)
				break
			}

			checkTextRes, checkTextErr := checkText(filePath)
			if !checkTextRes{
				fmt.Println(checkTextErr)
				break
			}

			f, _ := os.Open(filePath)
			defer f.Close()

			var reader io.Reader
			reader = f

			countResult, countErr := lineCounter(reader)
			if countErr != nil{
				fmt.Println(countErr)
			}else {
				fmt.Println(countResult)
			}
			break
		}
		refresh()
	},
}

func checkFile(path string) (bool, error){
	checkRes := false
	var checkErr error
	for{
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			errorStr := fmt.Sprintf("error: No such file %v", path)
			checkErr = errors.New(errorStr)
			break
		}else if info.Mode().IsDir(){
			errorStr := fmt.Sprintf("error: Expected file got directory %v", path)
			checkErr = errors.New(errorStr)
			break
		}
		checkRes = true
		break
	}
	return checkRes, checkErr
}

func checkText(path string) (bool, error)  {
	checkRes := false
	var checkErr error
	f, _ := ioutil.ReadFile(path)
	if isText := util.IsText(f); !isText{
		errorStr := fmt.Sprintf("error: Cannot do linecount for binary file %v", path)
		checkErr = errors.New(errorStr)
	}else {
		checkRes = true
	}
	return checkRes, checkErr
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		if c != 0 {
			count += bytes.Count(buf[:c], lineSep)+1
		}
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
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

func refresh()  {
	filePath = ""
	md5Flag = false
	sha1Flag = false
	sha256Flag = false
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Can't execute salve node：%v", err.Error())
	}
}

func Exit()  {
	os.Exit(1)
}

func init() {
	// Init Cli
	cobra.OnInitialize()
	linecountCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
	checksumCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
	checksumCmd.Flags().BoolVarP(&md5Flag, "md5", "", false, "Get checksum in hash function-sha256")
	checksumCmd.Flags().BoolVarP(&sha1Flag, "sha1", "", false, "Get checksum in hash function-sha1")
	checksumCmd.Flags().BoolVarP(&sha256Flag, "sha256", "", false, "Get checksum in hash function-sha256")
	rootCmd.AddCommand(fopsCmd)
	fopsCmd.AddCommand(versionCmd, checksumCmd, linecountCmd)
}