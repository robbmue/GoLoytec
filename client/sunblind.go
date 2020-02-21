package client

import "fmt"

const(
	payload = `<?xml version="1.0" encoding="utf-8"?>
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
)

func (client *Client) parsePayload(funcZone, rotZone, posZone int) string {
	return fmt.Sprintf(payload, funcZone, rotZone, posZone)
}