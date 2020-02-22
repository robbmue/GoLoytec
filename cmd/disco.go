package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robbmue/GoLoytec/client"
)

func discoMode(client *client.Client) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		client.Light(0, 0)
		fmt.Println("Exiting disco mode")
		os.Exit(0)
	}()
	for {
		client.Light(1, 100)
		time.Sleep(time.Millisecond * 100)	
		client.Light(0, 0)
		time.Sleep(time.Millisecond * 100)
	}
}
