package main

import(
	"os"
	"strconv"

	"github.com/robbmue/GoLoytec/client"
)

func commandInterface(clientInstance *client.Client) {
	switch os.Args[1] {
	case "light":
		switch os.Args[2] {
		case "on":
			clientInstance.Light(1, 100)
		case "off":
			clientInstance.Light(0, 0)
		case "set":
			if !(len(os.Args) == 4) {
				exitOnFail()
			}
			intensity, err := strconv.Atoi(os.Args[3])
			if err != nil {
				exitOnFail()
			}
			clientInstance.Light(1, intensity)
		default:
			exitOnFail()
		}
	case "blinds":
		switch os.Args[2] {
		case "top":
			clientInstance.Sunblind(client.Top)
		case "bottom":
			clientInstance.Sunblind(client.Bottom)
		case "set":
			rotation := 0
			if !(len(os.Args) >= 4) {
				exitOnFail()
			}
			intensity, err := strconv.Atoi(os.Args[3])
			if err != nil {
				exitOnFail()
			}
			if len(os.Args) == 6 && os.Args[4] == "rotate" {
				rotation, err = strconv.Atoi(os.Args[5])
				if err != nil {
					exitOnFail()
				}
			}
			clientInstance.Sunblind(client.Custom, 5, rotation, intensity)
		default:
			exitOnFail()
		}
	case "disco":
		discoMode(clientInstance)
	default:
		exitOnFail()
	}
}