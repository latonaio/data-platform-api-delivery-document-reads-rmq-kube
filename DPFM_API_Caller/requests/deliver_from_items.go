package requests

type DeliverFromItems struct {
	DeliveryDocument                   int     `json:"DeliveryDocument"`
	HeaderDeliveryStatus               *string `json:"HeaderDeliveryStatus"`
	DeliverFromBusinessPartnerFullName *string `json:"DeliverFromBusinessPartnerFullName"`
	DeliverFromBusinessPartnerName     *string `json:"DeliverFromBusinessPartnerName"`
	DeliverToBusinessPartnerFullName   *string `json:"DeliverToBusinessPartnerFullName"`
	DeliverToBusinessPartnerName       *string `json:"DeliverToBusinessPartnerName"`
	ItemBillingStatus                  *string `json:"ItemBillingStatus"`
	ConfirmedDeliveryDate              *string `json:"ConfirmedDeliveryDate"`
}