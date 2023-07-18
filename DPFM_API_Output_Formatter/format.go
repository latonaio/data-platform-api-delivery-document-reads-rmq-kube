package dpfm_api_output_formatter

import (
	"data-platform-api-delivery-document-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

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
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
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
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
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
			ProductionOrder:                        data.ProductionOrder,
			ProductionOrderItem:                    data.ProductionOrderItem,
			Operations:                             data.Operations,
			OperationsItem:                         data.OperationsItem,
			BillOfMaterial:                         data.BillOfMaterial,
			BillOfMaterialItem:                     data.BillOfMaterialItem,
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
			CreationDate:                           data.CreationDate,
			CreationTime:                           data.CreationTime,
			LastChangeDate:                         data.LastChangeDate,
			LastChangeTime:                         data.LastChangeTime,
			IsCancelled:                            data.IsCancelled,
			IsMarkedForDeletion:                    data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

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
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.OriginalQuantityInBaseUnit,
			&pm.OriginalQuantityInDeliveryUnit,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverToPlant,
			&pm.BatchMgmtPolicyInDeliverToPlant,
			&pm.DeliverToPlantBatch,
			&pm.DeliverToPlantBatchValidityStartDate,
			&pm.DeliverToPlantBatchValidityStartTime,
			&pm.DeliverToPlantBatchValidityEndDate,
			&pm.DeliverToPlantBatchValidityEndTime,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverFromPlant,
			&pm.BatchMgmtPolicyInDeliverFromPlant,
			&pm.DeliverFromPlantBatch,
			&pm.DeliverFromPlantBatchValidityStartDate,
			&pm.DeliverFromPlantBatchValidityStartTime,
			&pm.DeliverFromPlantBatchValidityEndDate,
			&pm.DeliverFromPlantBatchValidityEndTime,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.ProductIsBatchManagedInStockConfirmationPlant,
			&pm.BatchMgmtPolicyInStockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityStartTime,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.StockConfirmationPlantBatchValidityEndTime,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.ProductionPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityStartTime,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.ProductionPlantBatchValidityEndTime,
			&pm.InspectionPlan,
			&pm.InspectionPlant,
			&pm.InspectionOrder,
			&pm.DeliveryDocumentItemText,
			&pm.DeliveryDocumentItemTextByBuyer,
			&pm.DeliveryDocumentItemTextBySeller,
			&pm.PlannedGoodsIssueDate,
			&pm.PlannedGoodsIssueTime,
			&pm.PlannedGoodsReceiptDate,
			&pm.PlannedGoodsReceiptTime,
			&pm.PlannedGoodsIssueQuantity,
			&pm.PlannedGoodsIssueQtyInBaseUnit,
			&pm.PlannedGoodsReceiptQuantity,
			&pm.PlannedGoodsReceiptQtyInBaseUnit,
			&pm.ActualGoodsIssueDate,
			&pm.ActualGoodsIssueTime,
			&pm.ActualGoodsReceiptDate,
			&pm.ActualGoodsReceiptTime,
			&pm.ActualGoodsIssueQuantity,
			&pm.ActualGoodsIssueQtyInBaseUnit,
			&pm.ActualGoodsReceiptQuantity,
			&pm.ActualGoodsReceiptQtyInBaseUnit,
			&pm.QuantityPerPackage,
			&pm.ItemBillingStatus,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.ItemIsBillingRelevant,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.Operations,
			&pm.OperationsItem,
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
			&pm.WBSElement,
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
			&pm.Equipment,
			&pm.ItemDeliveryBlockStatus,
			&pm.ItemIssuingBlockStatus,
			&pm.ItemReceivingBlockStatus,
			&pm.ItemBillingBlockStatus,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
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
			Product:                                       data.Product,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			BaseUnit:                                      data.BaseUnit,
			DeliveryUnit:                                  data.DeliveryUnit,
			OriginalQuantityInBaseUnit:                    data.OriginalQuantityInBaseUnit,
			OriginalQuantityInDeliveryUnit:                data.OriginalQuantityInDeliveryUnit,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductIsBatchManagedInDeliverToPlant:         data.ProductIsBatchManagedInDeliverToPlant,
			BatchMgmtPolicyInDeliverToPlant:               data.BatchMgmtPolicyInDeliverToPlant,
			DeliverToPlantBatch:                           data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:          data.DeliverToPlantBatchValidityStartDate,
			DeliverToPlantBatchValidityStartTime:          data.DeliverToPlantBatchValidityStartTime,
			DeliverToPlantBatchValidityEndDate:            data.DeliverToPlantBatchValidityEndDate,
			DeliverToPlantBatchValidityEndTime:            data.DeliverToPlantBatchValidityEndTime,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			ProductIsBatchManagedInDeliverFromPlant:       data.ProductIsBatchManagedInDeliverFromPlant,
			BatchMgmtPolicyInDeliverFromPlant:             data.BatchMgmtPolicyInDeliverFromPlant,
			DeliverFromPlantBatch:                         data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:        data.DeliverFromPlantBatchValidityStartDate,
			DeliverFromPlantBatchValidityStartTime:        data.DeliverFromPlantBatchValidityStartTime,
			DeliverFromPlantBatchValidityEndDate:          data.DeliverFromPlantBatchValidityEndDate,
			DeliverFromPlantBatchValidityEndTime:          data.DeliverFromPlantBatchValidityEndTime,
			StockConfirmationBusinessPartner:              data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                        data.StockConfirmationPlant,
			ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
			BatchMgmtPolicyInStockConfirmationPlant:       data.BatchMgmtPolicyInStockConfirmationPlant,
			StockConfirmationPlantBatch:                   data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityStartTime:  data.StockConfirmationPlantBatchValidityStartTime,
			StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
			StockConfirmationPlantBatchValidityEndTime:    data.StockConfirmationPlantBatchValidityEndTime,
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
			ProductionPlantBatch:                          data.ProductionPlantBatch,
			ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityStartTime:         data.ProductionPlantBatchValidityStartTime,
			ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
			ProductionPlantBatchValidityEndTime:           data.ProductionPlantBatchValidityEndTime,
			InspectionPlan:           					   data.InspectionPlan,
			InspectionPlant:           					   data.InspectionPlant,
			InspectionOrder:           					   data.InspectionOrder,
			DeliveryDocumentItemText:                      data.DeliveryDocumentItemText,
			DeliveryDocumentItemTextByBuyer:               data.DeliveryDocumentItemTextByBuyer,
			DeliveryDocumentItemTextBySeller:              data.DeliveryDocumentItemTextBySeller,
			PlannedGoodsIssueDate:                         data.PlannedGoodsIssueDate,
			PlannedGoodsIssueTime:                         data.PlannedGoodsIssueTime,
			PlannedGoodsReceiptDate:                       data.PlannedGoodsReceiptDate,
			PlannedGoodsReceiptTime:                       data.PlannedGoodsReceiptTime,
			PlannedGoodsIssueQuantity:                     data.PlannedGoodsIssueQuantity,
			PlannedGoodsIssueQtyInBaseUnit:                data.PlannedGoodsIssueQtyInBaseUnit,
			PlannedGoodsReceiptQuantity:                   data.PlannedGoodsReceiptQuantity,
			PlannedGoodsReceiptQtyInBaseUnit:              data.PlannedGoodsReceiptQtyInBaseUnit,
			ActualGoodsIssueDate:                          data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:                          data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:                        data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:                        data.ActualGoodsReceiptTime,
			ActualGoodsIssueQuantity:                      data.ActualGoodsIssueQuantity,
			ActualGoodsIssueQtyInBaseUnit:                 data.ActualGoodsIssueQtyInBaseUnit,
			ActualGoodsReceiptQuantity:                    data.ActualGoodsReceiptQuantity,
			ActualGoodsReceiptQtyInBaseUnit:               data.ActualGoodsReceiptQtyInBaseUnit,
			QuantityPerPackage:                            data.QuantityPerPackage,
			ItemBillingStatus:                             data.ItemBillingStatus,
			ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
			ItemGrossWeight:                               data.ItemGrossWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
			ItemWeightUnit:                                data.ItemWeightUnit,
			InternalCapacityQuantity:                      data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:                  data.InternalCapacityQuantityUnit,
			ItemIsBillingRelevant:                         data.ItemIsBillingRelevant,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			OrderID:                                       data.OrderID,
			OrderItem:                                     data.OrderItem,
			ProductionOrder:                               data.ProductionOrder,
			ProductionOrderItem:                           data.ProductionOrderItem,
			BillOfMaterial:                                data.BillOfMaterial,
			BillOfMaterialItem:                            data.BillOfMaterialItem,
			Operations:                                    data.Operations,
			OperationsItem:                                data.OperationsItem,
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
			WBSElement:                                    data.WBSElement,
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
			Equipment:                                     data.Equipment,
			ItemDeliveryBlockStatus:                       data.ItemDeliveryBlockStatus,
			ItemIssuingBlockStatus:                        data.ItemIssuingBlockStatus,
			ItemReceivingBlockStatus:                      data.ItemReceivingBlockStatus,
			ItemBillingBlockStatus:                        data.ItemBillingBlockStatus,
			CreationDate:                                  data.CreationDate,
			CreationTime:                                  data.CreationTime,
			LastChangeDate:                                data.LastChangeDate,
			LastChangeTime:                                data.LastChangeTime,
			IsCancelled:                                   data.IsCancelled,
			IsMarkedForDeletion:                           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}

func ConvertToItemPicking(rows *sql.Rows) (*[]ItemPicking, error) {
	defer rows.Close()
	itemPicking := make([]ItemPicking, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemPicking{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.DeliveryDocumentItemPickingID,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.Product,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverToPlantStorageBin,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantStorageLocation,
			&pm.DeliverFromPlantStorageBin,
			&pm.DeliverToPlantPlannedPickingQuantityInBaseUnit,
			&pm.DeliverFromPlantPlannedPickingQuantityInBaseUnit,
			&pm.DeliverToPlantPlannedPickingDate,
			&pm.DeliverToPlantPlannedPickingTime,
			&pm.DeliverFromPlantPlannedPickingDate,
			&pm.DeliverFromPlantPlannedPickingTime,
			&pm.DeliverToPlantActualPickingQuantityInBaseUnit,
			&pm.DeliverToPlantActualPickingDate,
			&pm.DeliverToPlantActualPickingTime,
			&pm.DeliverFromPlantActualPickingQuantityInBaseUnit,
			&pm.DeliverFromPlantActualPickingDate,
			&pm.DeliverFromPlantActualPickingTime,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemPicking, err
		}

		data := pm
		itemPicking = append(itemPicking, ItemPicking{
			DeliveryDocument:                       data.DeliveryDocument,
			DeliveryDocumentItem:                   data.DeliveryDocumentItem,
			DeliveryDocumentItemPickingID:          data.DeliveryDocumentItemPickingID,
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			Product:                                data.Product,
			DeliverToParty:                         data.DeliverToParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverToPlantStorageLocation:          data.DeliverToPlantStorageLocation,
			DeliverToPlantStorageBin:               data.DeliverToPlantStorageBin,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverFromPlant:                       data.DeliverFromPlant,
			DeliverFromPlantStorageLocation:        data.DeliverFromPlantStorageLocation,
			DeliverFromPlantStorageBin:             data.DeliverFromPlantStorageBin,
			DeliverToPlantPlannedPickingQuantityInBaseUnit:   data.DeliverToPlantPlannedPickingQuantityInBaseUnit,
			DeliverFromPlantPlannedPickingQuantityInBaseUnit: data.DeliverFromPlantPlannedPickingQuantityInBaseUnit,
			DeliverToPlantPlannedPickingDate:                 data.DeliverToPlantPlannedPickingDate,
			DeliverToPlantPlannedPickingTime:                 data.DeliverToPlantPlannedPickingTime,
			DeliverFromPlantPlannedPickingDate:               data.DeliverFromPlantPlannedPickingDate,
			DeliverFromPlantPlannedPickingTime:               data.DeliverFromPlantPlannedPickingTime,
			DeliverToPlantActualPickingQuantityInBaseUnit:    data.DeliverToPlantActualPickingQuantityInBaseUnit,
			DeliverToPlantActualPickingDate:                  data.DeliverToPlantActualPickingDate,
			DeliverToPlantActualPickingTime:                  data.DeliverToPlantActualPickingTime,
			DeliverFromPlantActualPickingQuantityInBaseUnit:  data.DeliverFromPlantActualPickingQuantityInBaseUnit,
			DeliverFromPlantActualPickingDate:                data.DeliverFromPlantActualPickingDate,
			DeliverFromPlantActualPickingTime:                data.DeliverFromPlantActualPickingTime,
			CreationDate:                                     data.CreationDate,
			CreationTime:                                     data.CreationTime,
			LastChangeDate:                                   data.LastChangeDate,
			LastChangeTime:                                   data.LastChangeTime,
			IsCancelled:                                      data.IsCancelled,
			IsMarkedForDeletion:                              data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemPicking, nil
	}

	return &itemPicking, nil
}

func ConvertToAddress(rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	address := make([]Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

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
			return &address, err
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
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &address, nil
	}

	return &address, nil
}

func ConvertToPartner(rows *sql.Rows) (*[]Partner, error) {
	defer rows.Close()
	partner := make([]Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Partner{}

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
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &partner, nil
	}

	return &partner, nil
}
