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

	httpclient := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	buf.WriteString(client.parseLightPayload(state,value))


	req, err := http.NewRequest("POST", client.GetAddress(), &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Basic b3BlcmF0b3I6b3BlcmF0b3I=")
	req.Header.Add("Origin", "http://www.loytec.com")
	req.Header.Add("SOAPAction", "http://opcfoundation.org/webservices/XMLDA/1.0/Write")
	req.Header.Add("Referer", "http://www.loytec.com/lweb802/?project=lstudio%2FSystem.LROC_LROC111.LWEBMobile_Seg04.lweb2&address=10.10.160.111&port=80&https=false")
	_, err = httpclient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

