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
	Auth    string
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
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) GetAddress() string {
	return fmt.Sprintf("http://%s:%v/DA", client.Address, client.Port)
}

func (client *Client) SendRequest(buf bytes.Buffer) error {
	httpclient := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	req, err := http.NewRequest("POST", client.GetAddress(), &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", client.Auth)
	req.Header.Add("Origin", "http://www.loytec.com")
	req.Header.Add("SOAPAction", "http://opcfoundation.org/webservices/XMLDA/1.0/Write")
	req.Header.Add("Referer", fmt.Sprintf("http://www.loytec.com/lweb802/?project=lstudio%2FSystem.LROC_LROC111.LWEBMobile_Seg04.lweb2&address=%v&port=%s&https=false", client.Address, client.Port))
	_, err = httpclient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
