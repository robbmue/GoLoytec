package main

import (
	"bufio"
	"fmt"
	"github.com/robbmue/GoLoytec/client"
	"os"
	"strconv"
)

func main() {
	x := client.Init("10.10.160.111",80)
	x.Ping()
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Sunblinds [TOP(0)|UP(1)|DOWN(2)|BOTTOM(3)]")
		input, _, _ := reader.ReadLine()
		i,_ := strconv.Atoi(string(input))
		x.Sunblind(client.Direction(i))
	}

}
