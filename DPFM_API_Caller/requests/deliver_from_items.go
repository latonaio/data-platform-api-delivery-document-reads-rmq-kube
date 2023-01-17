package requests

type DeliverFromItems struct {
	DeliveryDocument                   int     `json:"DeliveryDocument"`
	DeliverFromBusinessPartnerFullName *string `json:"DeliverFromBusinessPartnerFullName"`
	DeliverFromBusinessPartnerName     *string `json:"DeliverFromBusinessPartnerName"`
	DeliverToBusinessPartnerFullName   *string `json:"DeliverToBusinessPartnerFullName"`
	DeliverToBusinessPartnerName       *string `json:"DeliverToBusinessPartnerName"`
	HeaderDeliveryStatus               *string `json:"HeaderDeliveryStatus"`
	ItemBillingStatus                  *string `json:"ItemBillingStatus"`
	ConfirmedDeliveryDate              *string `json:"ConfirmedDeliveryDate"`
}
