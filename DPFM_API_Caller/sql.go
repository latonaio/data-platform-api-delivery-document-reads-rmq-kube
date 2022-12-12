package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-delivery-document-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-delivery-document-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var headerPartner *dpfm_api_output_formatter.HeaderPartner
	var headerPartnerContact *dpfm_api_output_formatter.HeaderPartnerContact
	var headerPartnerPlant *dpfm_api_output_formatter.HeaderPartnerPlant
	var address *dpfm_api_output_formatter.Address
	var item *dpfm_api_output_formatter.Item
	var itemPartner *dpfm_api_output_formatter.ItemPartner
	var itemPartnerPlant *dpfm_api_output_formatter.ItemPartnerPlant
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeaderPartner":
			func() {
				headerPartner = c.HeaderPartner(mtx, input, output, errs, log)
			}()
		case "HeaderPartnerContact":
			func() {
				headerPartnerContact = c.HeaderPartnerContact(mtx, input, output, errs, log)
			}()
		case "HeaderPartnerPlant":
			func() {
				headerPartnerPlant = c.HeaderPartnerPlant(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "ItemPartner":
			func() {
				itemPartner = c.ItemPartner(mtx, input, output, errs, log)
			}()
		case "ItemPartnerPlant":
			func() {
				itemPartnerPlant = c.ItemPartnerPlant(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:               header,
		HeaderPartner:        headerPartner,
		HeaderPartnerContact: headerPartnerContact,
		HeaderPartnerPlant:   headerPartnerPlant,
		Address:              address,
		Item:                 item,
		ItemPartner:          itemPartner,
		ItemPartnerPlant:     itemPartnerPlant,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	deliveryDocument := input.Header.DeliveryDocument

	rows, err := c.db.Query(
		`SELECT DeliveryDocument, Buyer, Seller, ReferenceDocument, ReferenceDocumentItem, OrderID, 
		OrderItem, ContractType, OrderValidityStartDate, OrderValidityEndDate, IssuingPlantTimeZone, 
		ReceivingPlantTimeZone, DocumentDate, PlannedGoodsIssueDate, PlannedGoodsIssueTime, PlannedGoodsReceiptDate, 
		PlannedGoodsReceiptTime, BillingDocumentDate, CompleteDeliveryIsDefined, OverallDeliveryStatus, 
		CreationDate, CreationTime, IssuingBlockReason, ReceivingBlockReason, GoodsIssueOrReceiptSlipNumber, 
		HeaderBillingStatus, HeaderBillingConfStatus, HeaderBillingBlockReason, HeaderGrossWeight, 
		HeaderNetWeight, HeaderVolume, HeaderVolumeUnit, HeaderWeightUnit, Incoterms, IsExportImportDelivery, 
		LastChangeDate, IssuingPlantBusinessPartner, IssuingPlant, ReceivingPlant, ReceivingPlantBusinessPartner, 
		DeliverToParty, DeliverFromParty, TransactionCurrency, OverallDelivReltdBillgStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE DeliveryDocument = ?;`, deliveryDocument,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartner {
	deliveryDocument := input.Header.DeliveryDocument
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT DeliveryDocument, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, 
		Organization, Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_partner_data
		WHERE (DeliveryDocument, PartnerFunction, BusinessPartner) = (?, ?, ?);`, deliveryDocument, partnerFunction, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartnerContact(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartnerContact {
	deliveryDocument := input.Header.DeliveryDocument
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	contactID := input.Header.HeaderPartner.HeaderPartnerContact.ContactID

	rows, err := c.db.Query(
		`SELECT DeliveryDocument, PartnerFunction, ContactID, ContactPersonName, EmailAddress, 
		PhoneNumber, MobilePhoneNumber, FaxNumber, ContactTag1, ContactTag2, ContactTag3, ContactTag4
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_partner_contact_data
		WHERE (DeliveryDocument, PartnerFunction, ContactID) = (?, ?, ?);`, deliveryDocument, partnerFunction, contactID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartnerContact(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartnerPlant(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartnerPlant {
	deliveryDocument := input.Header.DeliveryDocument
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner
	plant := input.Header.HeaderPartner.HeaderPartnerPlant.Plant

	rows, err := c.db.Query(
		`SELECT DeliveryDocument, PartnerFunction, BusinessPartner, Plant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_partner_plant_data
		WHERE (DeliveryDocument, PartnerFunction, BusinessPartner, Plant) = (?, ?, ?, ?);`, deliveryDocument, partnerFunction, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartnerPlant(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Address {
	deliveryDocument := input.Header.DeliveryDocument
	addressID := input.Header.Address.AddressID

	rows, err := c.db.Query(
		`SELECT DeliveryDocument, AddressID, PostalCode, LocalRegion, Country, District, StreetName, 
		CityName, Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_address_data
		WHERE (DeliveryDocument, AddressID) = (?, ?);`, deliveryDocument, addressID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAddress(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Item {
	deliveryDocument := input.Header.DeliveryDocument
	deliveryDocumentItem := input.Header.Item.DeliveryDocumentItem

	rows, err := c.db.Query(
		`SELECT DeliveryDocument, DeliveryDocumentItem, DeliveryDocumentItemCategory, DeliveryDocumentItemText, 
		Product, ProductStandardID, ProductGroup, BaseUnit, OriginalDeliveryQuantity, DeliveryQuantityUnit, 
		ActualGoodsIssueDate, ActualGoodsIssueTime, ActualGoodsReceiptDate, ActualGoodsReceiptTime, 
		ActualGoodsIssueQtyInBaseUnit, ActualGoodsIssueQuantity, ActualGoodsReceiptQtyInBaseUnit, ActualGoodsReceiptQuantity, 
		CompleteItemDeliveryIsDefined, StockConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlant, 
		StockConfirmationPolicy, StockConfirmationStatus, ProductionPlantPartnerFunction, ProductionPlantBusinessPartner, 
		ProductionPlant, ProductionPlantStorageLocation, IssuingPlantPartnerFunction, IssuingPlantBusinessPartner, IssuingPlant, 
		IssuingPlantStorageLocation, ReceivingPlantPartnerFunction, ReceivingPlantBusinessPartner, ReceivingPlant, 
		ReceivingPlantStorageLocation, ProductIsBatchManagedInProductionPlant, ProductIsBatchManagedInIssuingPlant, 
		ProductIsBatchManagedInReceivingPlant, BatchMgmtPolicyInProductionPlant, BatchMgmtPolicyInIssuingPlant, 
		BatchMgmtPolicyInReceivingPlant, ProductionPlantBatch, IssuingPlantBatch, ReceivingPlantBatch, ProductionPlantBatchValidityStartDate, 
		ProductionPlantBatchValidityEndDate, IssuingPlantBatchValidityStartDate, IssuingPlantBatchValidityEndDate, ReceivingPlantBatchValidityStartDate, 
		ReceivingPlantBatchValidityEndDate, CreationDate, CreationTime, ItemBillingStatus, ItemBillingConfStatus, SalesCostGLAccount, 
		ReceivingGLAccount, IssuingGoodsMovementType, ReceivingGoodsMovementType, ItemBillingBlockReason, ItemCompleteDeliveryIsDefined, 
		ItemDeliveryIncompletionStatus, ItemGrossWeight, ItemNetWeight, ItemWeightUnit, ItemIsBillingRelevant, LastChangeDate, 
		OrderID, OrderItem, OrderType, ContractType, OrderValidityStartDate, OrderValidityEndDate, PaymentTerms, 
		DueCalculationBaseDate, PaymentDueDate, NetPaymentDays, PaymentMethod, InvoicePeriodStartDate, InvoicePeriodEndDate, 
		ProductAvailabilityDate, Project, ReferenceDocument, ReferenceDocumentItem, BPTaxClassification, ProductTaxClassification, 
		BPAccountAssignmentGroup, ProductAccountAssignmentGroup, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data
		WHERE (DeliveryDocument, DeliveryDocumentItem) = (?, ?, ?);`, deliveryDocument, deliveryDocumentItem,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItem(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPartner {
	deliveryDocument := input.Header.DeliveryDocument
	deliveryDocumentItem := input.Header.Item.DeliveryDocumentItem
	partnerFunction := input.Header.Item.ItemPartner.PartnerFunction
	businessPartner := input.Header.Item.ItemPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT DeliveryDocument, DeliveryDocumentItem, PartnerFunction, BusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_partner_data
		WHERE (DeliveryDocument, DeliveryDocumentItem, PartnerFunction, BusinessPartner) = (?, ?, ?, ?);`, deliveryDocument, deliveryDocumentItem, partnerFunction, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPartnerPlant(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPartnerPlant {
	deliveryDocument := input.Header.DeliveryDocument
	deliveryDocumentItem := input.Header.Item.DeliveryDocumentItem
	partnerFunction := input.Header.Item.ItemPartner.PartnerFunction
	businessPartner := input.Header.Item.ItemPartner.BusinessPartner
	plant := input.Header.Item.ItemPartner.ItemPartnerPlant.Plant

	rows, err := c.db.Query(
		`SELECT DeliveryDocument, DeliveryDocumentItem, PartnerFunction, BusinessPartner, Plant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_partner_plant_data
		WHERE (DeliveryDocument, DeliveryDocumentItem, PartnerFunction, BusinessPartner, Plant) = (?, ?, ?, ?, ?);`, deliveryDocument, deliveryDocumentItem, partnerFunction, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPartnerPlant(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
