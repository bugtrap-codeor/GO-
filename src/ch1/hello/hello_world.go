package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args)
	if len(os.Args)>1 {
		fmt.Println("Hello World",os.Args[1])	
	}
	//os.Exit(-1)返回状态
}

//go run hello_world.go bugtrap