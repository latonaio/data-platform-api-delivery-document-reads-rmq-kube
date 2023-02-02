package requests

type DeliverToItems struct {
	DeliveryDocument                   int     `json:"DeliveryDocument"`
	HeaderDeliveryStatus               *string `json:"HeaderDeliveryStatus"`
	DeliverToBusinessPartnerFullName   *string `json:"DeliverToBusinessPartnerFullName"`
	DeliverToBusinessPartnerName       *string `json:"DeliverToBusinessPartnerName"`
	DeliverFromBusinessPartnerFullName *string `json:"DeliverFromBusinessPartnerFullName"`
	DeliverFromBusinessPartnerName     *string `json:"DeliverFromBusinessPartnerName"`
	ItemBillingStatus                  *string `json:"ItemBillingStatus"`
	ConfirmedDeliveryDate              *string `json:"ConfirmedDeliveryDate"`
}