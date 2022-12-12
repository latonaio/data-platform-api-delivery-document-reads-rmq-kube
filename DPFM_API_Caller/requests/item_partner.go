package requests

type ItemPartner struct {
	DeliveryDocument     *int   `json:"DeliveryDocument"`
	DeliveryDocumentItem *int   `json:"DeliveryDocumentItem"`
	PartnerFunction      string `json:"PartnerFunction"`
	BusinessPartner      *int   `json:"BusinessPartner"`
}
