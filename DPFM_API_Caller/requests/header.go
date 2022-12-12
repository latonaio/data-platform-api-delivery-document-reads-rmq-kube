package requests

type Header struct {
	DeliveryDocument              int      `json:"DeliveryDocument"`
	Buyer                         *int     `json:"Buyer"`
	Seller                        *int     `json:"Seller"`
	ReferenceDocument             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem         *int     `json:"ReferenceDocumentItem"`
	OrderID                       *int     `json:"OrderID"`
	OrderItem                     *int     `json:"OrderItem"`
	ContractType                  *string  `json:"ContractType"`
	OrderValidityStartDate        *string  `json:"OrderVaridityStartDate"`
	OrderValidityEndDate          *string  `json:"OrderValidityEndDate"`
	IssuingPlantTimeZone          *string  `json:"IssuingPlantTimeZone"`
	ReceivingPlantTimeZone        *string  `json:"ReceivingPlantTimeZone"`
	DocumentDate                  *string  `json:"DocumentDate"`
	PlannedGoodsIssueDate         *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime         *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate       *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime       *string  `json:"PlannedGoodsReceiptTime"`
	BillingDocumentDate           *string  `json:"BillingDocumentDate"`
	CompleteDeliveryIsDefined     *bool    `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus         *string  `json:"OverallDeliveryStatus"`
	CreationDate                  *string  `json:"CreationDate"`
	CreationTime                  *string  `json:"CreationTime"`
	IssuingBlockReason            *bool    `json:"IssuingBlockReason"`
	ReceivingBlockReason          *bool    `json:"ReceivingBlockReason"`
	GoodsIssueOrReceiptSlipNumber *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus           *string  `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus       *string  `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockReason      *bool    `json:"HeaderBillingBlockReason"`
	HeaderGrossWeight             *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight               *float32 `json:"HeaderNetWeight"`
	HeaderVolume                  *string  `json:"HeaderVolume"`
	HeaderVolumeUnit              *string  `json:"HeaderVolumeUnit"`
	HeaderWeightUnit              *string  `json:"HeaderWeightUnit"`
	Incoterms                     *string  `json:"Incoterms"`
	IsExportImportDelivery        *bool    `json:"IsExportImportDelivery"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	IssuingPlantBusinessPartner   *string  `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                  *string  `json:"IssuingPlant"`
	ReceivingPlant                *string  `json:"ReceivingPlant"`
	ReceivingPlantBusinessPartner *string  `json:"ReceivingPlantBusinessPartner"`
	DeliverToParty                *int     `json:"DeliverToParty"`
	DeliverFromParty              *int     `json:"DeliverFromParty"`
	TransactionCurrency           *string  `json:"TransactionCurrency"`
	OverallDelivReltdBillgStatus  *string  `json:"OverallDelivReltdBillgStatus"`
}
