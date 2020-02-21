package main

import (
	"bufio"
	"fmt"
	"github.com/robbmue/GoLoytec/client"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := client.Init("10.10.160.111",80)
	x.Ping()
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Sunblinds [TOP(0)|UP(1)|DOWN(2)|BOTTOM(3)]")
		input, _, _ := reader.ReadLine()
		params := make([]int, 3)
		i,_ := strconv.Atoi(string(input))
		if i == 4 {
			fmt.Println("Insert custom params (3):")
			para, _, _ := reader.ReadLine()
			for i , x := range strings.Split(string(para), " "){
				params[i],_ = strconv.Atoi(x)
			}

		}
		x.Sunblind(client.Direction(i), params[0], params[1], params[2])
	}

}
