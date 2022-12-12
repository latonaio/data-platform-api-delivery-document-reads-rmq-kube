package requests

type HeaderPartnerPlant struct {
	DeliveryDocument int    `json:"DeliveryDocument"`
	PartnerFunction  string `json:"PartnerFunction"`
	BusinessPartner  int    `json:"BusinessPartner"`
	Plant            string `json:"Plant"`
}
