package requests

type Item struct {
	DeliveryDocument                       *int     `json:"DeliveryDocument"`
	DeliveryDocumentItem                   *int     `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory           string   `json:"DeliveryDocumentItemCategory"`
	DeliveryDocumentItemText               string   `json:"DeliveryDocumentItemText"`
	Product                                string   `json:"Product"`
	ProductStandardID                      string   `json:"ProductStandardID"`
	ProductGroup                           string   `json:"ProductGroup"`
	BaseUnit                               string   `json:"BaseUnit"`
	OriginalDeliveryQuantity               *float32 `json:"OriginalDeliveryQuantity"`
	DeliveryQuantityUnit                   string   `json:"DeliveryQuantityUnit"`
	ActualGoodsIssueDate                   *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                   *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                 *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                 *string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQtyInBaseUnit          *float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsIssueQuantity               *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsReceiptQtyInBaseUnit        *float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	ActualGoodsReceiptQuantity             *float32 `json:"ActualGoodsReceiptQuantity"`
	CompleteItemDeliveryIsDefined          *bool    `json:"CompleteItemDeliveryIsDefined"`
	StockConfirmationPartnerFunction       string   `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner       *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                 string   `json:"StockConfirmationPlant"`
	StockConfirmationPolicy                string   `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                string   `json:"StockConfirmationStatus"`
	ProductionPlantPartnerFunction         string   `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner         string   `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                        string   `json:"ProductionPlant"`
	ProductionPlantStorageLocation         string   `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction            string   `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner            string   `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                           string   `json:"IssuingPlant"`
	IssuingPlantStorageLocation            string   `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction          string   `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner          string   `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                         string   `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation          string   `json:"ReceivingPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsBatchManagedInIssuingPlant    *bool    `json:"ProductIsBatchManagedInIssuingPlant"`
	ProductIsBatchManagedInReceivingPlant  *bool    `json:"ProductIsBatchManagedInReceivingPlant"`
	BatchMgmtPolicyInProductionPlant       string   `json:"BatchMgmtPolicyInProductionPlant"`
	BatchMgmtPolicyInIssuingPlant          string   `json:"BatchMgmtPolicyInIssuingPlant"`
	BatchMgmtPolicyInReceivingPlant        string   `json:"BatchMgmtPolicyInReceivingPlant"`
	ProductionPlantBatch                   string   `json:"ProductionPlantBatch"`
	IssuingPlantBatch                      string   `json:"IssuingPlantBatch"`
	ReceivingPlantBatch                    string   `json:"ReceivingPlantBatch"`
	ProductionPlantBatchValidityStartDate  *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate    *string  `json:"ProductionPlantBatchValidityEndDate"`
	IssuingPlantBatchValidityStartDate     *string  `json:"IssuingPlantBatchValidityStartDate"`
	IssuingPlantBatchValidityEndDate       *string  `json:"IssuingPlantBatchValidityEndDate"`
	ReceivingPlantBatchValidityStartDate   *string  `json:"ReceivingPlantBatchValidityStartDate"`
	ReceivingPlantBatchValidityEndDate     *string  `json:"ReceivingPlantBatchValidityEndDate"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	ItemBillingStatus                      string   `json:"ItemBillingStatus"`
	ItemBillingConfStatus                  string   `json:"ItemBillingConfStatus"`
	SalesCostGLAccount                     string   `json:"SalesCostGLAccount"`
	ReceivingGLAccount                     string   `json:"ReceivingGLAccount"`
	IssuingGoodsMovementType               string   `json:"IssuingGoodsMovementType"`
	ReceivingGoodsMovementType             string   `json:"ReceivingGoodsMovementType"`
	ItemBillingBlockReason                 *bool    `json:"ItemBillingBlockReason"`
	ItemCompleteDeliveryIsDefined          *bool    `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryIncompletionStatus         string   `json:"ItemDeliveryIncompletionStatus"`
	ItemGrossWeight                        *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                          *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                         string   `json:"ItemWeightUnit"`
	ItemIsBillingRelevant                  *bool    `json:"ItemIsBillingRelevant"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	OrderID                                *int     `json:"OrderID"`
	OrderItem                              *int     `json:"OrderItem"`
	OrderType                              string   `json:"OrderType"`
	ContractType                           string   `json:"ContractType"`
	OrderValidityStartDate                 *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string  `json:"OrderValidityEndDate"`
	PaymentTerms                           string   `json:"PaymentTerms"`
	DueCalculationBaseDate                 *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                         *string  `json:"PaymentDueDate"`
	NetPaymentDays                         *int     `json:"NetPaymentDays"`
	PaymentMethod                          *string  `json:"PaymentMethod"`
	InvoicePeriodStartDate                 *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                   *string  `json:"InvoicePeriodEndDate"`
	ProductAvailabilityDate                *string  `json:"ProductAvailabilityDate"`
	Project                                string   `json:"Project"`
	ReferenceDocument                      *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                  *int     `json:"ReferenceDocumentItem"`
	BPTaxClassification                    string   `json:"BPTaxClassification"`
	ProductTaxClassification               string   `json:"ProductTaxClassification"`
	BPAccountAssignmentGroup               string   `json:"BPAccountAssignmentGroup"`
	ProductAccountAssignmentGroup          string   `json:"ProductAccountAssignmentGroup"`
	TaxCode                                string   `json:"TaxCode"`
	TaxRate                                *float32 `json:"TaxRate"`
	CountryOfOrigin                        string   `json:"CountryOfOrigin"`
}
