package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var s string
	fmt.Println("s:" + s + ".")  // s==""
	var _, _, _ int              // error: var i,j,k int,float32,int
	var _, _, _ = true, 1., "sb" // can not use _ as value

	// multi return value case
	var f, err = os.Open("./decalre.go")
	if err != nil {
		log.Println(err)
	}
	f.Close()

	// 简短变量声明
	i, j := 0, 1
	i, j = j, i
	fmt.Println(i, j)
}
