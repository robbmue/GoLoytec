package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Address string
	Port    int
}

func Init(addr string, port int) *Client {
	return &Client{Address: addr, Port: port}
}

func (client *Client) Ping() error {
	resp, err := http.Get(client.GetAddress())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(body)
	return nil
}

func (client *Client) GetAddress() string {
	return fmt.Sprintf("http://%s:%v/DA", client.Address, client.Port)
}




