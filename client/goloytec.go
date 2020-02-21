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

func (client *Client) Sunblind(direction Direction) error {

	var buf bytes.Buffer

	httpclient := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	switch direction {
	case Top:
		buf.WriteString(client.parsePayload(5, 0, 0))
	case Up:
		buf.WriteString(client.parsePayload(5, 0, 90))
	case Down:
		buf.WriteString(client.parsePayload(5, 0, 10))
	case Bottom:
		buf.WriteString(client.parsePayload(5, 90, 100))
	default:
		return nil
	}

	req, err := http.NewRequest("POST", client.GetAddress(), &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Basic b3BlcmF0b3I6b3BlcmF0b3I=")
	req.Header.Add("Origin", "http://www.loytec.com")
	req.Header.Add("SOAPAction", "http://opcfoundation.org/webservices/XMLDA/1.0/Write")
	req.Header.Add("Referer", "http://www.loytec.com/lweb802/?project=lstudio%2FSystem.LROC_LROC111.LWEBMobile_Seg04.lweb2&address=10.10.160.111&port=80&https=false")
	fmt.Printf("\nRequest: \n %s \n\n", req)
	resp, err := httpclient.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("Response: \n %s \n\n", resp)

	return nil
}
