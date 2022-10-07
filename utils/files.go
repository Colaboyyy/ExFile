package utils

import (
	"bufio"
	"fmt"
	"os"
)

var blockSize = 4096

// file is exit or not
func IsExit(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	return false
}

func FileReader(filename string, data chan []byte) bool {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file failed! [red]", err)
		return false
	}
	defer file.Close()
	defer close(data)

	reader := bufio.NewReader(file)
	for {
		tmp := make([]byte, blockSize)
		n, err := reader.Read(tmp)
		if err != nil {
			fmt.Println("File read failed! [red]", err)
			return false
		}
		if n == 0 {
			return true
		}
		data <- tmp
	}
}

func FileWriter(filename string, data chan []byte) bool {
	for IsExit(filename) {
		filename += "-副本"
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Create file failed! [red]", err)
		return false
	}
	defer file.Close()

	for bytes := range data {
		_, err = file.Write(bytes)
		if err != nil {
			fmt.Println("Write file failed! [red]", err)
			return false
		}
	}
	return true
}
