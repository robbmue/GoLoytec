package main

import (
	"fmt"
	"os"
)

func main() {
	config := new(Conf).getConf()
	defaultClient := config.setRoom()
	if defaultClient.Ping() != nil {
		fmt.Println("No Network")
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		textUserInterface(defaultClient)
	} else {
		commandInterface(defaultClient)
	}
}

func exitOnFail() {
	fmt.Println("No valid command")
	os.Exit(0)
}
