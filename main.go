package main

import (
	"fmt"
	"main/utils"
)

func main() {
	fmt.Println("===***** ExFile *****===")
	LocalIps := utils.GetIntranetIp()
	for _, LocalIps := range LocalIps {
		fmt.Println(LocalIps)
	}
	fmt.Println("please enter 'help' for more information")

	gui := utils.GUI{}
	gui.LocalIps = LocalIps
	gui.Run()
}
