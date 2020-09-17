package main

import (
	//	"sort"

	//	"reflect"
	//	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	//	"strings"
)

//type FileSome interface {
//	Show(int) error
//	}

//func Show

const DirResult = 
`
├───project
├───static
│	├───a_lorem
│	│	└───ipsum
│	├───css
│	├───html
│	├───js
│	└───z_lorem
│		└───ipsum
└───zline
	└───lorem
		└───ipsum
`

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for idx, file := range files {
		if idx > 0 {
			findPath(files, file, "├───")
		}
	}
	return nil
}

func findPath(fileList []string, currLine string, deep string) error {
	var newDeep = "│   " + deep
	var subList []string
	var currLineForm = currLine + "/"
	if strings.Contains(currLine, "/") != true {

		fmt.Println(deep + currLine)

		for _, line := range fileList {
			if strings.Contains(line, currLine) && line != currLine {

				line = strings.Replace(line, currLineForm, "", -1)
				subList = append(subList, line)
			}
		}
		if len(subList) > 0 {

			for _, line := range subList {
				findPath(subList, line, newDeep)
			}
		}
	}
	return nil
}
