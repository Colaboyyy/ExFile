package utils

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func Send(filename string, ip string) bool {
	host, _ := net.ResolveTCPAddr("tcp4", ip+filePort)
	client, err := net.DialTCP("tcp", nil, host)
	if err != nil {
		fmt.Println("Connect failed! [red]" + err.Error())
		return false
	}
	defer client.Close()

	// success
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file for sending failed! [red]" + err.Error())
		return false
	}
	info, _ := file.Stat()
	size := info.Size()
	if !sendName(filename, client) {
		return false
	}
	if !sendSize(size, client) {
		return false
	}
	file.Close()

	readerResult := make(chan bool)
	senderResult := make(chan bool)
	counter := make(chan int64)
	data := make(chan []byte, 1024)
	go func() {
		readerResult <- FileReader(filename, data)
	}()
	go func() {
		senderResult <- Sender(client, data, true, counter)
	}()

	go DisplayCounter(size, counter)

	if <-readerResult && <-senderResult {
		fmt.Println("[yellow]Send file success!")
	} else {
		fmt.Println("[red]Send file failed!")
		return false
	}
	return true
}

func sendName(filename string, client *net.TCPConn) bool {
	tmp := []byte(filename)
	_, err := client.Write(tmp)
	if err != nil {
		fmt.Println("Send fileName failed! [red]" + err.Error())
		return false
	}
	n, _ := client.Read(tmp)
	if string(tmp[:n]) != "success" {
		fmt.Println("Receive fileName failed! [red]" + err.Error())
		return false
	}
	return true
}

func sendSize(size int64, client *net.TCPConn) bool {
	tmp := []byte(strconv.FormatInt(size, 10))
	_, err := client.Write(tmp)
	if err != nil {
		fmt.Println("Send fileSize failed! [red]" + err.Error())
		return false
	}
	n, _ := client.Read(tmp)
	if string(tmp[:n]) != "success" {
		fmt.Println("Receive fileSize failed! [red]" + err.Error())
		return false
	}
	return true
}
