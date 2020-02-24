package client

import (
	"bytes"
	"fmt"
	"net/http"
)

func (client *Client) parseLightPayload(lampZoneState, lampZoneValue int) string {
	payload := `<?xml version="1.0" encoding="utf-8"?>
					<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
						<soap:Body>
							<Write xmlns="http://opcfoundation.org/webservices/XMLDA/1.0/" ReturnValuesOnReply="true">
								<Options ReturnErrorText="true" ReturnItemTime="true" ClientRequestHandle="Write-65"/>
								<ItemList>
									<Items ItemPath="User Registers.Seg04.Lights1.HmiLWeb" ItemName="lampZoneState" ClientItemHandle="gen0x01200091-2">
										<Value xsi:type="xsd:int">%v</Value>
									</Items>
									<Items ItemPath="User Registers.Seg04.Lights1.HmiLWeb" ItemName="lampZoneValue" ClientItemHandle="gen0x01200092-2">
										<Value xsi:type="xsd:double">%v</Value>
									</Items>
								</ItemList>
							</Write>
						</soap:Body>
					</soap:Envelope> `
	return fmt.Sprintf(payload, lampZoneState, lampZoneValue)
}

func (client *Client) Light(state, value int) error {
	var buf bytes.Buffer

	buf.WriteString(client.parseLightPayload(state, value))

	err := sendRequest(buf)
	if err != nil {
		return err
	}

	return nil
}
