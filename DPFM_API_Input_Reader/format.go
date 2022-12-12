package dpfm_api_input_reader

import (
	"data-platform-api-delivery-document-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		DeliveryDocument:              data.DeliveryDocument,
		Buyer:                         data.Buyer,
		Seller:                        data.Seller,
		ReferenceDocument:             data.ReferenceDocument,
		ReferenceDocumentItem:         data.ReferenceDocumentItem,
		OrderID:                       data.OrderID,
		OrderItem:                     data.OrderItem,
		ContractType:                  data.ContractType,
		OrderValidityStartDate:        data.OrderValidityStartDate,
		OrderValidityEndDate:          data.OrderValidityEndDate,
		IssuingPlantTimeZone:          data.IssuingPlantTimeZone,
		ReceivingPlantTimeZone:        data.ReceivingPlantTimeZone,
		DocumentDate:                  data.DocumentDate,
		PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:         data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:       data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:       data.PlannedGoodsReceiptTime,
		BillingDocumentDate:           data.BillingDocumentDate,
		CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
		OverallDeliveryStatus:         data.OverallDeliveryStatus,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		IssuingBlockReason:            data.IssuingBlockReason,
		ReceivingBlockReason:          data.ReceivingBlockReason,
		GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
		HeaderBillingStatus:           data.HeaderBillingStatus,
		HeaderBillingConfStatus:       data.HeaderBillingConfStatus,
		HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
		HeaderGrossWeight:             data.HeaderGrossWeight,
		HeaderNetWeight:               data.HeaderNetWeight,
		HeaderVolume:                  data.HeaderVolume,
		HeaderVolumeUnit:              data.HeaderVolumeUnit,
		HeaderWeightUnit:              data.HeaderWeightUnit,
		Incoterms:                     data.Incoterms,
		IsExportImportDelivery:        data.IsExportImportDelivery,
		LastChangeDate:                data.LastChangeDate,
		IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
		IssuingPlant:                  data.IssuingPlant,
		ReceivingPlant:                data.ReceivingPlant,
		ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
		DeliverToParty:                data.DeliverToParty,
		DeliverFromParty:              data.DeliverFromParty,
		TransactionCurrency:           data.TransactionCurrency,
		OverallDelivReltdBillgStatus:  data.OverallDelivReltdBillgStatus,
	}
}

func (sdc *SDC) ConvertToHeaderPDF() *requests.HeaderPDF {
	dataHeader := sdc.Header
	data := sdc.Header.HeaderPDF
	return &requests.HeaderPDF{
		DeliveryDocument:         dataHeader.DeliveryDocument,
		DocType:                  data.DocType,
		DocVersionID:             data.DocVersionID,
		DocID:                    data.DocID,
		DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		FileName:                 data.FileName,
	}
}

func (sdc *SDC) ConvertToHeaderPartner() *requests.HeaderPartner {
	dataHeader := sdc.Header
	data := sdc.Header.HeaderPartner
	return &requests.HeaderPartner{
		DeliveryDocument:        dataHeader.DeliveryDocument,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
}

func (sdc *SDC) ConvertToHeaderPartnerContact() *requests.HeaderPartnerContact {
	dataHeader := sdc.Header
	dataHeaderPartner := sdc.Header.HeaderPartner
	data := sdc.Header.HeaderPartner.HeaderPartnerContact
	return &requests.HeaderPartnerContact{
		DeliveryDocument:  dataHeader.DeliveryDocument,
		PartnerFunction:   dataHeaderPartner.PartnerFunction,
		ContactID:         data.ContactID,
		ContactPersonName: data.ContactPersonName,
		EmailAddress:      data.EmailAddress,
		PhoneNumber:       data.PhoneNumber,
		MobilePhoneNumber: data.MobilePhoneNumber,
		FaxNumber:         data.FaxNumber,
		ContactTag1:       data.ContactTag1,
		ContactTag2:       data.ContactTag2,
		ContactTag3:       data.ContactTag3,
		ContactTag4:       data.ContactTag4,
	}
}

func (sdc *SDC) ConvertToHeaderPartnerPlant() *requests.HeaderPartnerPlant {
	dataHeader := sdc.Header
	dataHeaderPartner := sdc.Header.HeaderPartner
	data := sdc.Header.HeaderPartner.HeaderPartnerPlant
	return &requests.HeaderPartnerPlant{
		DeliveryDocument: dataHeader.DeliveryDocument,
		PartnerFunction:  dataHeaderPartner.PartnerFunction,
		BusinessPartner:  dataHeaderPartner.BusinessPartner,
		Plant:            data.Plant,
	}
}

func (sdc *SDC) ConvertToAddress() *requests.Address {
	dataHeader := sdc.Header
	data := sdc.Header.Address
	return &requests.Address{
		DeliveryDocument: dataHeader.DeliveryDocument,
		AddressID:        data.AddressID,
		PostalCode:       data.PostalCode,
		LocalRegion:      data.LocalRegion,
		Country:          data.Country,
		District:         data.District,
		StreetName:       data.StreetName,
		CityName:         data.CityName,
		Building:         data.Building,
		Floor:            data.Floor,
		Room:             data.Room,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataHeader := sdc.Header
	data := sdc.Header.Item
	return &requests.Item{
		DeliveryDocument:                       dataHeader.DeliveryDocument,
		DeliveryDocumentItem:                   data.DeliveryDocumentItem,
		DeliveryDocumentItemCategory:           data.DeliveryDocumentItemCategory,
		DeliveryDocumentItemText:               data.DeliveryDocumentItemText,
		Product:                                data.Product,
		ProductStandardID:                      data.ProductStandardID,
		ProductGroup:                           data.ProductGroup,
		BaseUnit:                               data.BaseUnit,
		OriginalDeliveryQuantity:               data.OriginalDeliveryQuantity,
		DeliveryQuantityUnit:                   data.DeliveryQuantityUnit,
		ActualGoodsIssueDate:                   data.ActualGoodsIssueDate,
		ActualGoodsIssueTime:                   data.ActualGoodsIssueTime,
		ActualGoodsReceiptDate:                 data.ActualGoodsReceiptDate,
		ActualGoodsReceiptTime:                 data.ActualGoodsReceiptTime,
		ActualGoodsIssueQtyInBaseUnit:          data.ActualGoodsIssueQtyInBaseUnit,
		ActualGoodsIssueQuantity:               data.ActualGoodsIssueQuantity,
		ActualGoodsReceiptQtyInBaseUnit:        data.ActualGoodsReceiptQtyInBaseUnit,
		ActualGoodsReceiptQuantity:             data.ActualGoodsReceiptQuantity,
		CompleteItemDeliveryIsDefined:          data.CompleteItemDeliveryIsDefined,
		StockConfirmationPartnerFunction:       data.StockConfirmationPartnerFunction,
		StockConfirmationBusinessPartner:       data.StockConfirmationBusinessPartner,
		StockConfirmationPlant:                 data.StockConfirmationPlant,
		StockConfirmationPolicy:                data.StockConfirmationPolicy,
		StockConfirmationStatus:                data.StockConfirmationStatus,
		ProductionPlantPartnerFunction:         data.ProductionPlantPartnerFunction,
		ProductionPlantBusinessPartner:         data.ProductionPlantBusinessPartner,
		ProductionPlant:                        data.ProductionPlant,
		ProductionPlantStorageLocation:         data.ProductionPlantStorageLocation,
		IssuingPlantPartnerFunction:            data.IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner:            data.IssuingPlantBusinessPartner,
		IssuingPlant:                           data.IssuingPlant,
		IssuingPlantStorageLocation:            data.IssuingPlantStorageLocation,
		ReceivingPlantPartnerFunction:          data.ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner:          data.ReceivingPlantBusinessPartner,
		ReceivingPlant:                         data.ReceivingPlant,
		ReceivingPlantStorageLocation:          data.ReceivingPlantStorageLocation,
		ProductIsBatchManagedInProductionPlant: data.ProductIsBatchManagedInProductionPlant,
		ProductIsBatchManagedInIssuingPlant:    data.ProductIsBatchManagedInIssuingPlant,
		ProductIsBatchManagedInReceivingPlant:  data.ProductIsBatchManagedInReceivingPlant,
		BatchMgmtPolicyInProductionPlant:       data.BatchMgmtPolicyInProductionPlant,
		BatchMgmtPolicyInIssuingPlant:          data.BatchMgmtPolicyInIssuingPlant,
		BatchMgmtPolicyInReceivingPlant:        data.BatchMgmtPolicyInReceivingPlant,
		ProductionPlantBatch:                   data.ProductionPlantBatch,
		IssuingPlantBatch:                      data.IssuingPlantBatch,
		ReceivingPlantBatch:                    data.ReceivingPlantBatch,
		ProductionPlantBatchValidityStartDate:  data.ProductionPlantBatchValidityStartDate,
		ProductionPlantBatchValidityEndDate:    data.ProductionPlantBatchValidityEndDate,
		IssuingPlantBatchValidityStartDate:     data.IssuingPlantBatchValidityStartDate,
		IssuingPlantBatchValidityEndDate:       data.IssuingPlantBatchValidityEndDate,
		ReceivingPlantBatchValidityStartDate:   data.ReceivingPlantBatchValidityStartDate,
		ReceivingPlantBatchValidityEndDate:     data.ReceivingPlantBatchValidityEndDate,
		CreationDate:                           data.CreationDate,
		CreationTime:                           data.CreationTime,
		ItemBillingStatus:                      data.ItemBillingStatus,
		ItemBillingConfStatus:                  data.ItemBillingConfStatus,
		SalesCostGLAccount:                     data.SalesCostGLAccount,
		ReceivingGLAccount:                     data.ReceivingGLAccount,
		IssuingGoodsMovementType:               data.IssuingGoodsMovementType,
		ReceivingGoodsMovementType:             data.ReceivingGoodsMovementType,
		ItemBillingBlockReason:                 data.ItemBillingBlockReason,
		ItemCompleteDeliveryIsDefined:          data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryIncompletionStatus:         data.ItemDeliveryIncompletionStatus,
		ItemGrossWeight:                        data.ItemGrossWeight,
		ItemNetWeight:                          data.ItemNetWeight,
		ItemWeightUnit:                         data.ItemWeightUnit,
		ItemIsBillingRelevant:                  data.ItemIsBillingRelevant,
		LastChangeDate:                         data.LastChangeDate,
		OrderID:                                data.OrderID,
		OrderItem:                              data.OrderItem,
		OrderType:                              data.OrderType,
		ContractType:                           data.ContractType,
		OrderValidityStartDate:                 data.OrderValidityStartDate,
		OrderValidityEndDate:                   data.OrderValidityEndDate,
		PaymentTerms:                           data.PaymentTerms,
		DueCalculationBaseDate:                 data.DueCalculationBaseDate,
		PaymentDueDate:                         data.PaymentDueDate,
		NetPaymentDays:                         data.NetPaymentDays,
		PaymentMethod:                          data.PaymentMethod,
		InvoicePeriodStartDate:                 data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:                   data.InvoicePeriodEndDate,
		ProductAvailabilityDate:                data.ProductAvailabilityDate,
		Project:                                data.Project,
		ReferenceDocument:                      data.ReferenceDocument,
		ReferenceDocumentItem:                  data.ReferenceDocumentItem,
		BPTaxClassification:                    data.BPTaxClassification,
		ProductTaxClassification:               data.ProductTaxClassification,
		BPAccountAssignmentGroup:               data.BPAccountAssignmentGroup,
		ProductAccountAssignmentGroup:          data.ProductAccountAssignmentGroup,
		TaxCode:                                data.TaxCode,
		TaxRate:                                data.TaxRate,
		CountryOfOrigin:                        data.CountryOfOrigin,
	}
}

func (sdc *SDC) ConvertToItemPartner() *requests.ItemPartner {
	dataHeader := sdc.Header
	dataItem := sdc.Header.Item
	data := sdc.Header.Item.ItemPartner
	return &requests.ItemPartner{
		DeliveryDocument:     dataHeader.DeliveryDocument,
		DeliveryDocumentItem: dataItem.DeliveryDocumentItem,
		PartnerFunction:      data.PartnerFunction,
		BusinessPartner:      data.BusinessPartner,
	}
}

func (sdc *SDC) ConvertToItemPartnerPlant() *requests.ItemPartnerPlant {
	dataHeader := sdc.Header
	dataItem := sdc.Header.Item
	dataItemPartner := sdc.Header.Item.ItemPartner
	data := sdc.Header.Item.ItemPartner.ItemPartnerPlant
	return &requests.ItemPartnerPlant{
		DeliveryDocument:     dataHeader.DeliveryDocument,
		DeliveryDocumentItem: dataItem.DeliveryDocumentItem,
		PartnerFunction:      dataItemPartner.PartnerFunction,
		BusinessPartner:      dataItemPartner.BusinessPartner,
		Plant:                data.Plant,
	}
}
