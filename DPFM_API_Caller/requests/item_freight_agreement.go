package requests

type ItemFreightAgreement struct {
	DeliveryDocument                        int     `json:"DeliveryDocument"`
	DeliveryDocumentItem                    int     `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemFreightAgreement    int     `json:"DeliveryDocumentItemFreightAgreement"`
	FreightAgreement                        int     `json:"FreightAgreement"`
	FreightAgreementItem                    int     `json:"FreightAgreementItem"`
	SupplyChainRelationshipID               int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID        int     `json:"SupplyChainRelationshipFreightID"`
	SupplyChainRelationshipFreightBillingID *int    `json:"SupplyChainRelationshipFreightBillingID"`
	SupplyChainRelationshipFreightPaymentID *int    `json:"SupplyChainRelationshipFreightPaymentID"`
	FreightPartner                          int     `json:"FreightPartner"`
	FreightBillToParty                      *int    `json:"FreightBillToParty"`
	FreightBillFromParty                    *int    `json:"FreightBillFromParty"`
	FreightBillToCountry                    *string `json:"FreightBillToCountry"`
	FreightBillFromCountry                  *string `json:"FreightBillFromCountry"`
	Product                                 *string `json:"Product"`
	Incoterms                               *string `json:"Incoterms"`
	Project                                 *int    `json:"Project"`
	WBSElement                              *int    `json:"WBSElement"`
	IsCancelled                             *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool   `json:"IsMarkedForDeletion"`
}
