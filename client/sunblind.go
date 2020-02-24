package client

import (
	"bytes"
	"fmt"
	"net/http"
)

func (client *Client) parseSunblindPayload(funcZone, rotZone, posZone int) string {
	payload := `<?xml version="1.0" encoding="utf-8"?>
				<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Body>
						<Write xmlns="http://opcfoundation.org/webservices/XMLDA/1.0/" ReturnValuesOnReply="true">
							<Options ReturnErrorText="true" ReturnItemTime="true" ClientRequestHandle="Write-49"/>
							<ItemList>
								<Items ItemPath="User Registers.Seg04.Sunblind1.HmiLWeb" ItemName="functionZone" ClientItemHandle="gen0x012000ae-49">
									<Value xsi:type="xsd:int">%v</Value>
								</Items>
								<Items ItemPath="User Registers.Seg04.Sunblind1.HmiLWeb" ItemName="rotationZone" ClientItemHandle="gen0x012000ad-25">
									<Value xsi:type="xsd:double">%v</Value>
								</Items>
								<Items ItemPath="User Registers.Seg04.Sunblind1.HmiLWeb" ItemName="positionZone" ClientItemHandle="gen0x012000ac-25">
									<Value xsi:type="xsd:double">%v</Value>
								</Items>
							</ItemList>
						</Write>
					</soap:Body>
				</soap:Envelope>`
	return fmt.Sprintf(payload, funcZone, rotZone, posZone)
}

func (client *Client) Sunblind(direction Direction, custom ...int) error {

	var buf bytes.Buffer

	httpclient := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	switch direction {
	case Top:
		buf.WriteString(client.parseSunblindPayload(5, 0, 0))
	case Up:
		//TODO
		fmt.Println("TODO")
	case Down:
		//TODO
		fmt.Println("TODO")
	case Bottom:
		buf.WriteString(client.parseSunblindPayload(5, 90, 100))
	case Custom:
		if len(custom) != 3 {
			return nil
		}
		buf.WriteString(client.parseSunblindPayload(custom[0], custom[1], custom[2]))

	default:
		return nil
	}

	err := sendRequest(buf)
	if err != nil {
		return err
	}

	return nil
}
