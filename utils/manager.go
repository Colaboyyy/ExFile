package utils

import "fmt"

type GUI struct {
	FileList []string
	LocalIps []string
	funcList map[string]func()
}

var filePort string = ":12345"

var doc = map[string]string{
	"ls_file": "显示当前目录文件",
	"ls_dir":  "显示当前目录文件夹",
	"accept":  "进入文件接收模式，在这个模式下等待对方通过你的 ID 向你传输文件",
	"send":    "发送文件",
	"list":    "显示指令",
	"help":    "显示指令说明",
}

func handleErr(err error, msg string) bool {
	if err != nil {
		fmt.Println("[red]" + msg + err.Error())
		return false
	}
	return true
}

func (gui *GUI) Run() {
	gui.FileList = FileNum(".")
	gui.funcList = map[string]func(){
		"ls_file": LsFile,
		"ls_dir":  LsDir,
		"list":    listFunc,
		"accept":  gui.Receive,
		"send":    gui.fileTransport,
		"help":    gui.helpInfo,
	}
	for {
		var cmd string
		_, err := fmt.Scanln(&cmd)
		if cmd == "quit" {
			return
		} else if err != nil {
			fmt.Println("[blue]please enter 'help' for more information")
		}
		if fun, ok := gui.funcList[cmd]; ok {
			fun()
		} else {
			fmt.Println("No this instruction!")
		}
	}
}

// show help information
func (gui *GUI) helpInfo() {
	fmt.Println("If you want to send files, you should turn the receiver into 'receiver' mode and send files by specific ID")
	for cmd, info := range doc {
		fmt.Println(cmd + ":\t" + info)
	}
}

// receiver
func (gui *GUI) Receive() {
	Receive()
}

// fileTransport
func (gui *GUI) fileTransport() {
	LsFile()
	//fmt.Println(gui.FileList)
	fmt.Println("Please enter index to send file..")
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if !handleErr(err, "enter error:") {
		return
	}
	if n < 0 || n > len(gui.FileList) {
		fmt.Println("[red]Select file error!")
		return
	} else if n == len(gui.FileList) {
		return
	}
	fmt.Println("Please enter receiver ID:")
	var ip string
	_, _ = fmt.Scanln(&ip)
	Send(gui.FileList[n], ip)
}
