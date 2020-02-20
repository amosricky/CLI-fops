package cmd

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/tools/godoc/util"
	"io"
	"io/ioutil"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "fops",
	Short: "File Ops",
	Long: "File Ops",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize()

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkFile(path string) (bool, error){
	checkRes := false
	var checkErr error
	for{
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			errorStr := fmt.Sprintf("No such file %v", path)
			checkErr = errors.New(errorStr)
			break
		}else if info.Mode().IsDir(){
			errorStr := fmt.Sprintf("Expected file got directory %v", path)
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
		errorStr := fmt.Sprintf("Cannot do linecount for binary file %v", path)
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
