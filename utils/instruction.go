package utils

import (
	"fmt"
	"os"
)

func Cd(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Can't get in dir: [red]" + err.Error())
	}
}

func Ls(isDir bool) {
	names, _ := os.ReadDir("./")
	var files []string
	for _, file := range names {
		if file.IsDir() == isDir {
			files = append(files, file.Name())
		}
	}
	fmt.Println("Files are listed: ")
	for num, file := range files {
		fmt.Printf("%d. %s\n", num, file)
	}
	fmt.Println("Total file number:", len(files))
	fmt.Println()
}

func FileNum(path string) []string {
	var names []os.DirEntry
	if path != "" && path != "." {
		names, _ = os.ReadDir(path)
	} else {
		names, _ = os.ReadDir("./")
	}
	var files []string
	for _, file := range names {
		if file.IsDir() == false {
			files = append(files, file.Name())
		}
	}
	//fmt.Println(files)
	return files
}

func LsFile() {
	Ls(false)
}

func LsDir() {
	Ls(true)
}

// show Instructions
func listFunc() {
	fmt.Println("Instructions of ExFile are: ")
	for name := range doc {
		fmt.Println("	" + name)
	}
}
