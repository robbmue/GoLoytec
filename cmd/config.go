package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/robbmue/GoLoytec/client"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Def  string `yaml:"default"`
	Room []Room `yaml:"room"`
}

type Room struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func (c *Conf) getConf() *Conf {
	var address string
	if _, err := os.Stat("./config.yml"); os.IsNotExist(err) {
		if _, err := os.Stat("/etc/goloytec/config.yml"); os.IsNotExist(err) {
			os.Exit(1)
		} else {
			address = "/etc/goloytec/config.yml"
		}
	} else {
		address = "./config.yml"
	}

	yamlFile, err := ioutil.ReadFile(address)
	if err != nil {
		log.Printf("yaml File.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func (config *Conf) setRoom(context ...string) *client.Client {
	var room Room
	if len(context) == 0 {
		for _, v := range config.Room {
			if v.Name == config.Def {
				room = v
			}
		}
	} else {
		for _, v := range config.Room {
			if v.Name == context[0] {
				room = v
			}
		}
	}
	return client.Init(room.Address, room.Port)
}
