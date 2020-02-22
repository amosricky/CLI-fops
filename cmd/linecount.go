package cmd

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/tools/godoc/util"
	"io"
	"io/ioutil"
	"os"
)

var linecountFilePath string

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Count line for file.",
	Long: "Count line for file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		
		checkFileRes, checkFileErr := checkFile(linecountFilePath)
		if !checkFileRes{
			return checkFileErr
		}

		checkTextRes, checkTextErr := checkText(linecountFilePath)
		if !checkTextRes{
			return checkTextErr
		}

		f, openFileErr := os.Open(linecountFilePath)
		if openFileErr != nil{
			return openFileErr
		}
		defer f.Close()

		var reader io.Reader
		reader = f

		countResult, countErr := lineCounter(reader)
		if countErr != nil{
			return countErr
		}else {
			cmd.Printf("%v", countResult)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringVarP(&linecountFilePath, "file", "f", "", "File path")
}

func checkFile(path string) (bool, error){
	checkRes := false
	var checkErr error

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		errorStr := fmt.Sprintf("No such file %v", path)
		checkErr = errors.New(errorStr)
		return checkRes, checkErr
	}else if info.Mode().IsDir(){
		errorStr := fmt.Sprintf("Expected file got directory %v", path)
		checkErr = errors.New(errorStr)
		return checkRes, checkErr
	}
	checkRes = true
	return checkRes, checkErr
}

func checkText(path string) (bool, error)  {
	checkRes := false
	var checkErr error

	f, readErr := ioutil.ReadFile(path)
	if readErr != nil{
		return checkRes, readErr
	}
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