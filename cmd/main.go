package main

import (
	"fmt"
	"os"

	"github.com/robbmue/GoLoytec/client"
)

func main() {
	x := client.Init("10.10.160.111", 80)
	x.Ping()
	if len(os.Args) == 1 {
		textUserInterface(x)
	} else {
		commandInterface(x)
	}
}

func exitOnFail() {
	fmt.Println("No valid command")
	os.Exit(0)
}
