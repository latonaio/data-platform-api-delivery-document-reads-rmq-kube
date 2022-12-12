package dpfm_api_output_formatter

import (
	"data-platform-api-delivery-document-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-delivery-document-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToHeader(sdc *api_input_reader.SDC, rows *sql.Rows) (*Header, error) {
	pm := &requests.Header{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.Buyer,
			&pm.Seller,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.IssuingPlantTimeZone,
			&pm.ReceivingPlantTimeZone,
			&pm.DocumentDate,
			&pm.PlannedGoodsIssueDate,
			&pm.PlannedGoodsIssueTime,
			&pm.PlannedGoodsReceiptDate,
			&pm.PlannedGoodsReceiptTime,
			&pm.BillingDocumentDate,
			&pm.CompleteDeliveryIsDefined,
			&pm.OverallDeliveryStatus,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.IssuingBlockReason,
			&pm.ReceivingBlockReason,
			&pm.GoodsIssueOrReceiptSlipNumber,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingConfStatus,
			&pm.HeaderBillingBlockReason,
			&pm.HeaderGrossWeight,
			&pm.HeaderNetWeight,
			&pm.HeaderVolume,
			&pm.HeaderVolumeUnit,
			&pm.HeaderWeightUnit,
			&pm.Incoterms,
			&pm.IsExportImportDelivery,
			&pm.LastChangeDate,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantBusinessPartner,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.TransactionCurrency,
			&pm.OverallDelivReltdBillgStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	header := &Header{
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
	return header, nil
}

func ConvertToHeaderPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartner, error) {
	pm := &requests.HeaderPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartner := &HeaderPartner{
		DeliveryDocument:        data.DeliveryDocument,
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
	return headerPartner, nil
}

func ConvertToHeaderPartnerContact(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartnerContact, error) {
	pm := &requests.HeaderPartnerContact{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.PartnerFunction,
			&pm.ContactID,
			&pm.ContactPersonName,
			&pm.EmailAddress,
			&pm.PhoneNumber,
			&pm.MobilePhoneNumber,
			&pm.FaxNumber,
			&pm.ContactTag1,
			&pm.ContactTag2,
			&pm.ContactTag3,
			&pm.ContactTag4,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartnerContact := &HeaderPartnerContact{
		DeliveryDocument:  data.DeliveryDocument,
		PartnerFunction:   data.PartnerFunction,
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
	return headerPartnerContact, nil
}

func ConvertToHeaderPartnerPlant(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartnerPlant, error) {
	pm := &requests.HeaderPartnerPlant{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.Plant,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartnerPlant := &HeaderPartnerPlant{
		DeliveryDocument: data.DeliveryDocument,
		PartnerFunction:  data.PartnerFunction,
		BusinessPartner:  data.BusinessPartner,
		Plant:            data.Plant,
	}
	return headerPartnerPlant, nil
}

func ConvertToAddress(sdc *api_input_reader.SDC, rows *sql.Rows) (*Address, error) {
	pm := &requests.Address{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	address := &Address{
		DeliveryDocument: data.DeliveryDocument,
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
	return address, nil
}

func ConvertToItem(sdc *api_input_reader.SDC, rows *sql.Rows) (*Item, error) {
	pm := &requests.Item{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.DeliveryDocumentItemCategory,
			&pm.DeliveryDocumentItemText,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.OriginalDeliveryQuantity,
			&pm.DeliveryQuantityUnit,
			&pm.ActualGoodsIssueDate,
			&pm.ActualGoodsIssueTime,
			&pm.ActualGoodsReceiptDate,
			&pm.ActualGoodsReceiptTime,
			&pm.ActualGoodsIssueQtyInBaseUnit,
			&pm.ActualGoodsIssueQuantity,
			&pm.ActualGoodsReceiptQtyInBaseUnit,
			&pm.ActualGoodsReceiptQuantity,
			&pm.CompleteItemDeliveryIsDefined,
			&pm.StockConfirmationPartnerFunction,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ProductionPlantPartnerFunction,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.IssuingPlantPartnerFunction,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ReceivingPlantBusinessPartner,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.ProductIsBatchManagedInIssuingPlant,
			&pm.ProductIsBatchManagedInReceivingPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.BatchMgmtPolicyInIssuingPlant,
			&pm.BatchMgmtPolicyInReceivingPlant,
			&pm.ProductionPlantBatch,
			&pm.IssuingPlantBatch,
			&pm.ReceivingPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.IssuingPlantBatchValidityStartDate,
			&pm.IssuingPlantBatchValidityEndDate,
			&pm.ReceivingPlantBatchValidityStartDate,
			&pm.ReceivingPlantBatchValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.ItemBillingStatus,
			&pm.ItemBillingConfStatus,
			&pm.SalesCostGLAccount,
			&pm.ReceivingGLAccount,
			&pm.IssuingGoodsMovementType,
			&pm.ReceivingGoodsMovementType,
			&pm.ItemBillingBlockReason,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryIncompletionStatus,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.ItemIsBillingRelevant,
			&pm.LastChangeDate,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderType,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.ProductAvailabilityDate,
			&pm.Project,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.BPTaxClassification,
			&pm.ProductTaxClassification,
			&pm.BPAccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	item := &Item{
		DeliveryDocument:                       data.DeliveryDocument,
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
	return item, nil
}

func ConvertToItemPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*ItemPartner, error) {
	pm := &requests.ItemPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	itemPartner := &ItemPartner{
		DeliveryDocument:     data.DeliveryDocument,
		DeliveryDocumentItem: data.DeliveryDocumentItem,
		PartnerFunction:      data.PartnerFunction,
		BusinessPartner:      data.BusinessPartner,
	}
	return itemPartner, nil
}

func ConvertToItemPartnerPlant(sdc *api_input_reader.SDC, rows *sql.Rows) (*ItemPartnerPlant, error) {
	pm := &requests.ItemPartnerPlant{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.Plant,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	itemPartnerPlant := &ItemPartnerPlant{
		DeliveryDocument:     data.DeliveryDocument,
		DeliveryDocumentItem: data.DeliveryDocumentItem,
		PartnerFunction:      data.PartnerFunction,
		BusinessPartner:      data.BusinessPartner,
		Plant:                data.Plant,
	}
	return itemPartnerPlant, nil
}
