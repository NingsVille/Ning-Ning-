package main

//https://github.com/zanshin/interpreter

import (
	"fmt"
	"logic/config"
)

func main() {
	fmt.Println("_______________________________")
	fmt.Println("\n-$^ Programming Language v0.8")
	fmt.Println("_______________________________")
	fmt.Println()
	line1 := config.MainAuthor
	fmt.Printf("Main author is %s\n", line1)
}
