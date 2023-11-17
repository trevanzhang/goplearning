package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[1])
	txt, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(txt))
	for _, line := range strings.Split(string(txt), "\r\n") {
		fmt.Println(line)
	}
}
