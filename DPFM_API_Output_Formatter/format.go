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
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.IsExportImport,
			&pm.DeliverToPlantTimeZone,
			&pm.DeliverFromPlantTimeZone,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.DocumentDate,
			&pm.PlannedGoodsIssueDate,
			&pm.PlannedGoodsIssueTime,
			&pm.PlannedGoodsReceiptDate,
			&pm.PlannedGoodsReceiptTime,
			&pm.InvoiceDocumentDate,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.GoodsIssueOrReceiptSlipNumber,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingConfStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.HeaderGrossWeight,
			&pm.HeaderNetWeight,
			&pm.HeaderWeightUnit,
			&pm.Incoterms,
			&pm.TransactionCurrency,
			&pm.HeaderDeliveryBlockStatus,
			&pm.HeaderIssuingBlockStatus,
			&pm.HeaderReceivingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	header := &Header{
		DeliveryDocument:                       data.DeliveryDocument,
		SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
		SupplyChainRelationshipBillingID:       data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID:       data.SupplyChainRelationshipPaymentID,
		Buyer:                                  data.Buyer,
		Seller:                                 data.Seller,
		DeliverToParty:                         data.DeliverToParty,
		DeliverFromParty:                       data.DeliverFromParty,
		DeliverToPlant:                         data.DeliverToPlant,
		DeliverFromPlant:                       data.DeliverFromPlant,
		BillToParty:                            data.BillToParty,
		BillFromParty:                          data.BillFromParty,
		BillToCountry:                          data.BillToCountry,
		BillFromCountry:                        data.BillFromCountry,
		Payer:                                  data.Payer,
		Payee:                                  data.Payee,
		IsExportImport:                         data.IsExportImport,
		DeliverToPlantTimeZone:                 data.DeliverToPlantTimeZone,
		DeliverFromPlantTimeZone:               data.DeliverFromPlantTimeZone,
		ReferenceDocument:                      data.ReferenceDocument,
		ReferenceDocumentItem:                  data.ReferenceDocumentItem,
		OrderID:                                data.OrderID,
		OrderItem:                              data.OrderItem,
		ContractType:                           data.ContractType,
		OrderValidityStartDate:                 data.OrderValidityStartDate,
		OrderValidityEndDate:                   data.OrderValidityEndDate,
		DocumentDate:                           data.DocumentDate,
		PlannedGoodsIssueDate:                  data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:                  data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:                data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:                data.PlannedGoodsReceiptTime,
		InvoiceDocumentDate:                    data.InvoiceDocumentDate,
		HeaderCompleteDeliveryIsDefined:        data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:                   data.HeaderDeliveryStatus,
		CreationDate:                           data.CreationDate,
		CreationTime:                           data.CreationTime,
		LastChangeDate:                         data.LastChangeDate,
		LastChangeTime:                         data.LastChangeTime,
		GoodsIssueOrReceiptSlipNumber:          data.GoodsIssueOrReceiptSlipNumber,
		HeaderBillingStatus:                    data.HeaderBillingStatus,
		HeaderBillingConfStatus:                data.HeaderBillingConfStatus,
		HeaderBillingBlockStatus:               data.HeaderBillingBlockStatus,
		HeaderGrossWeight:                      data.HeaderGrossWeight,
		HeaderNetWeight:                        data.HeaderNetWeight,
		HeaderWeightUnit:                       data.HeaderWeightUnit,
		Incoterms:                              data.Incoterms,
		TransactionCurrency:                    data.TransactionCurrency,
		HeaderDeliveryBlockStatus:              data.HeaderDeliveryBlockStatus,
		HeaderIssuingBlockStatus:               data.HeaderIssuingBlockStatus,
		HeaderReceivingBlockStatus:             data.HeaderReceivingBlockStatus,
	}
	return header, nil
}

func ConvertToPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Partner, error) {
	var partner []Partner

	for i := 0; true; i++ {
		pm := &requests.Partner{}
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

		data := pm
		partner = append(partner, Partner{
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
		})
	}
	return &partner, nil
}

func ConvertToAddress(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Address, error) {
	var address []Address

	for i := 0; true; i++ {
		pm := &requests.Address{}
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

		data := pm
		address = append(address, Address{
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
		})
	}
	return &address, nil
}

func ConvertToItem(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Item, error) {
	var item []Item

	for i := 0; true; i++ {
		pm := &requests.Item{}
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
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverToPlant,
			&pm.BatchMgmtPolicyInDeliverToPlant,
			&pm.DeliverToPlantBatch,
			&pm.DeliverToPlantBatchValidityStartDate,
			&pm.DliverToPlantBatchValidityEndDate,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverFromPlant,
			&pm.BatchMgmtPolicyInDeliverFromPlant,
			&pm.DeliverFromPlantBatch,
			&pm.DeliverFromPlantBatchValidityStartDate,
			&pm.DliverFromPlantBatchValidityEndDate,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.ProductIsBatchManagedInStockConfirmationPlant,
			&pm.BatchMgmtPolicyInStockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.ProductionPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.DeliveryDocumentItemText,
			&pm.DeliveryDocumentItemTextByBuyer,
			&pm.DeliveryDocumentItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.OriginalQuantityInBaseUnit,
			&pm.DeliveryUnit,
			&pm.ActualGoodsIssueDate,
			&pm.ActualGoodsIssueTime,
			&pm.ActualGoodsReceiptDate,
			&pm.ActualGoodsReceiptTime,
			&pm.ActualGoodsIssueQtyInBaseUnit,
			&pm.ActualGoodsIssueQuantity,
			&pm.ActualGoodsReceiptQtyInBaseUnit,
			&pm.ActualGoodsReceiptQuantity,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.ItemBillingStatus,
			&pm.SalesCostGLAccount,
			&pm.ReceivingGLAccount,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.ItemIsBillingRelevant,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
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
			&pm.ConfirmedDeliveryDate,
			&pm.Project,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.ItemDeliveryBlockStatus,
			&pm.ItemIssuingBlockStatus,
			&pm.ItemReceivingBlockStatus,
			&pm.ItemBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		item = append(item, Item{
			DeliveryDocument:                              data.DeliveryDocument,
			DeliveryDocumentItem:                          data.DeliveryDocumentItem,
			DeliveryDocumentItemCategory:                  data.DeliveryDocumentItemCategory,
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:       data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			SupplyChainRelationshipBillingID:              data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID:              data.SupplyChainRelationshipPaymentID,
			Buyer:                                         data.Buyer,
			Seller:                                        data.Seller,
			DeliverToParty:                                data.DeliverToParty,
			DeliverFromParty:                              data.DeliverFromParty,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverFromPlant:                              data.DeliverFromPlant,
			BillToParty:                                   data.BillToParty,
			BillFromParty:                                 data.BillFromParty,
			BillToCountry:                                 data.BillToCountry,
			BillFromCountry:                               data.BillFromCountry,
			Payer:                                         data.Payer,
			Payee:                                         data.Payee,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductIsBatchManagedInDeliverToPlant:         data.ProductIsBatchManagedInDeliverToPlant,
			BatchMgmtPolicyInDeliverToPlant:               data.BatchMgmtPolicyInDeliverToPlant,
			DeliverToPlantBatch:                           data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:          data.DeliverToPlantBatchValidityStartDate,
			DliverToPlantBatchValidityEndDate:             data.DliverToPlantBatchValidityEndDate,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			ProductIsBatchManagedInDeliverFromPlant:       data.ProductIsBatchManagedInDeliverFromPlant,
			BatchMgmtPolicyInDeliverFromPlant:             data.BatchMgmtPolicyInDeliverFromPlant,
			DeliverFromPlantBatch:                         data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:        data.DeliverFromPlantBatchValidityStartDate,
			DliverFromPlantBatchValidityEndDate:           data.DliverFromPlantBatchValidityEndDate,
			StockConfirmationBusinessPartner:              data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                        data.StockConfirmationPlant,
			ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
			BatchMgmtPolicyInStockConfirmationPlant:       data.BatchMgmtPolicyInStockConfirmationPlant,
			StockConfirmationPlantBatch:                   data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
			ProductionPlantBatch:                          data.ProductionPlantBatch,
			ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
			DeliveryDocumentItemText:                      data.DeliveryDocumentItemText,
			DeliveryDocumentItemTextByBuyer:               data.DeliveryDocumentItemTextByBuyer,
			DeliveryDocumentItemTextBySeller:              data.DeliveryDocumentItemTextBySeller,
			Product:                                       data.Product,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			BaseUnit:                                      data.BaseUnit,
			OriginalQuantityInBaseUnit:                    data.OriginalQuantityInBaseUnit,
			DeliveryUnit:                                  data.DeliveryUnit,
			ActualGoodsIssueDate:                          data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:                          data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:                        data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:                        data.ActualGoodsReceiptTime,
			ActualGoodsIssueQtyInBaseUnit:                 data.ActualGoodsIssueQtyInBaseUnit,
			ActualGoodsIssueQuantity:                      data.ActualGoodsIssueQuantity,
			ActualGoodsReceiptQtyInBaseUnit:               data.ActualGoodsReceiptQtyInBaseUnit,
			ActualGoodsReceiptQuantity:                    data.ActualGoodsReceiptQuantity,
			CreationDate:                                  data.CreationDate,
			CreationTime:                                  data.CreationTime,
			LastChangeDate:                                data.LastChangeDate,
			LastChangeTime:                                data.LastChangeTime,
			ItemBillingStatus:                             data.ItemBillingStatus,
			SalesCostGLAccount:                            data.SalesCostGLAccount,
			ReceivingGLAccount:                            data.ReceivingGLAccount,
			ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
			ItemGrossWeight:                               data.ItemGrossWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
			ItemWeightUnit:                                data.ItemWeightUnit,
			ItemIsBillingRelevant:                         data.ItemIsBillingRelevant,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			OrderID:                                       data.OrderID,
			OrderItem:                                     data.OrderItem,
			OrderType:                                     data.OrderType,
			ContractType:                                  data.ContractType,
			OrderValidityStartDate:                        data.OrderValidityStartDate,
			OrderValidityEndDate:                          data.OrderValidityEndDate,
			PaymentTerms:                                  data.PaymentTerms,
			DueCalculationBaseDate:                        data.DueCalculationBaseDate,
			PaymentDueDate:                                data.PaymentDueDate,
			NetPaymentDays:                                data.NetPaymentDays,
			PaymentMethod:                                 data.PaymentMethod,
			InvoicePeriodStartDate:                        data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:                          data.InvoicePeriodEndDate,
			ConfirmedDeliveryDate:                         data.ConfirmedDeliveryDate,
			Project:                                       data.Project,
			ReferenceDocument:                             data.ReferenceDocument,
			ReferenceDocumentItem:                         data.ReferenceDocumentItem,
			TransactionTaxClassification:                  data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:         data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry:       data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                      data.DefinedTaxClassification,
			AccountAssignmentGroup:                        data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
			TaxCode:                                       data.TaxCode,
			TaxRate:                                       data.TaxRate,
			CountryOfOrigin:                               data.CountryOfOrigin,
			CountryOfOriginLanguage:                       data.CountryOfOriginLanguage,
			ItemDeliveryBlockStatus:                       data.ItemDeliveryBlockStatus,
			ItemIssuingBlockStatus:                        data.ItemIssuingBlockStatus,
			ItemReceivingBlockStatus:                      data.ItemReceivingBlockStatus,
			ItemBillingBlockStatus:                        data.ItemBillingBlockStatus,
		})
	}

	return &item, nil
}

func ConvertToDeliverFromItems(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]DeliverFromItems, error) {
	var deliverFromItems []DeliverFromItems

	for i := 0; true; i++ {
		pm := &requests.DeliverFromItems{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_delivery_document_header_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.HeaderDeliveryStatus,
			&pm.DeliverFromBusinessPartnerFullName,
			&pm.DeliverFromBusinessPartnerName,
			&pm.DeliverToBusinessPartnerName,
			&pm.DeliverToBusinessPartnerFullName,
			&pm.ItemBillingStatus,
			&pm.ConfirmedDeliveryDate,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliverFromItems = append(deliverFromItems, DeliverFromItems{
			DeliveryDocument:                   data.DeliveryDocument,
			HeaderDeliveryStatus:               data.HeaderDeliveryStatus,
			DeliverFromBusinessPartnerFullName: data.DeliverFromBusinessPartnerFullName,
			DeliverFromBusinessPartnerName:     data.DeliverFromBusinessPartnerName,
			DeliverToBusinessPartnerName:       data.DeliverToBusinessPartnerName,
			DeliverToBusinessPartnerFullName:   data.DeliverToBusinessPartnerFullName,
			ItemBillingStatus:                  data.ItemBillingStatus,
			ConfirmedDeliveryDate:              data.ConfirmedDeliveryDate,
		})
	}

	return &deliverFromItems, nil
}

func ConvertToDeliverToItems(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]DeliverToItems, error) {
	var deliverToItems []DeliverToItems

	for i := 0; true; i++ {
		pm := &requests.DeliverToItems{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_delivery_document_header_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.HeaderDeliveryStatus,
			&pm.DeliverToBusinessPartnerFullName,
			&pm.DeliverToBusinessPartnerName,
			&pm.DeliverFromBusinessPartnerFullName,
			&pm.DeliverFromBusinessPartnerName,
			&pm.ItemBillingStatus,
			&pm.ConfirmedDeliveryDate,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliverToItems = append(deliverToItems, DeliverToItems{
			DeliveryDocument:                   data.DeliveryDocument,
			HeaderDeliveryStatus:               data.HeaderDeliveryStatus,
			DeliverToBusinessPartnerFullName:   data.DeliverToBusinessPartnerFullName,
			DeliverToBusinessPartnerName:       data.DeliverToBusinessPartnerName,
			DeliverFromBusinessPartnerFullName: data.DeliverFromBusinessPartnerFullName,
			DeliverFromBusinessPartnerName:     data.DeliverFromBusinessPartnerName,
			ItemBillingStatus:                  data.ItemBillingStatus,
			ConfirmedDeliveryDate:              data.ConfirmedDeliveryDate,
		})
	}

	return &deliverToItems, nil
}
