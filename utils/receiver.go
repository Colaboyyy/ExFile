package utils

import (
	"fmt"
	"net"
	"strconv"
)

func Receive() bool {
	host, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0"+filePort)
	fmt.Println("Listening: ", host.IP, host.Port)
	listener, err := net.ListenTCP("tcp", host)
	if err != nil {
		fmt.Println("Listen failed! [red]" + err.Error())
		return false
	}
	conn, err := listener.AcceptTCP()
	if err != nil {
		fmt.Println("Receive failed! [red]" + err.Error())
		return false
	}
	defer conn.Close()

	// connect success
	filename := receiveName(conn)
	if len(filename) == 0 {
		fmt.Println("[red]Receive fileName err!")
		return false
	}
	size := receiveSize(conn)
	if size == 0 {
		fmt.Println("[red]Receive fileSize err!")
		return false
	}

	data := make(chan []byte, blockSize)
	writerResult := make(chan bool)
	receiveResult := make(chan bool)
	counter := make(chan int64)
	go func() {
		writerResult <- FileWriter(filename, data)
	}()
	go func() {
		receiveResult <- Receiver(conn, data, true, counter)
	}()

	go DisplayCounter(size, counter)

	if <-writerResult && <-receiveResult {
		fmt.Println("[yellow] Receive file success!")
	} else {
		fmt.Println("[red] Receive file failed!")
		return false
	}
	return true
}

func receiveName(conn *net.TCPConn) string {
	tmp := make([]byte, 200)
	n, err := conn.Read(tmp)
	if err != nil {
		fmt.Println("Receive fileName failed! [red]" + err.Error())
		tmp = []byte("fail")
		_, _ = conn.Write(tmp)
		return ""
	}
	res := string(tmp[:n])
	tmp = []byte("success")
	_, _ = conn.Write(tmp)
	return res
}

func receiveSize(conn *net.TCPConn) int64 {
	tmp := make([]byte, 200)
	n, err := conn.Read(tmp)
	if err != nil {
		fmt.Println("Receive fileSize failed! [red]" + err.Error())
		tmp = []byte("fail")
		_, _ = conn.Write(tmp)
		return 0
	}
	res, _ := strconv.ParseInt(string(tmp[:n]), 10, 64)
	tmp = []byte("success")
	_, _ = conn.Write(tmp)
	return res
}

// DisplayCounter show progress
func DisplayCounter(size int64, counter chan int64) {
	now := int64(0)
	for tmp := range counter {
		now += tmp
		fmt.Printf("Progress:%f%%\r[green]", float64(now)/float64(size)*100)
	}
	fmt.Println()
}
