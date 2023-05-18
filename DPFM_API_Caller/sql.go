package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-delivery-document-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-delivery-document-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
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
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var partner *[]dpfm_api_output_formatter.Partner
	var address *[]dpfm_api_output_formatter.Address
	var deliverFromItems *[]dpfm_api_output_formatter.DeliverFromItems
	var deliverToItems *[]dpfm_api_output_formatter.DeliverToItems
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "Headers":
			func() {
				header = c.Headers(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()
		case "Partner":
			func() {
				partner = c.Partner(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "DeliverFromItems":
			func() {
				deliverFromItems = c.DeliverFromItems(mtx, input, output, errs, log)
			}()
		case "DeliverToItems":
			func() {
				deliverToItems = c.DeliverToItems(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:           header,
		Item:             item,
		Partner:          partner,
		Address:          address,
		DeliverFromItems: deliverFromItems,
		DeliverToItems:   deliverToItems,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE DeliveryDocument = %d ", input.Header.DeliveryDocument)
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %v", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocument DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Headers(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.HeaderCompleteDeliveryIsDefined != nil {
		where = fmt.Sprintf("%s\nAND HeaderCompleteDeliveryIsDefined = %v", where, *input.Header.HeaderCompleteDeliveryIsDefined)
	}
	if input.Header.HeaderDeliveryBlockStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderDeliveryBlockStatus = %v", where, *input.Header.HeaderDeliveryBlockStatus)
	}
	if input.Header.HeaderDeliveryStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderDeliveryStatus = '%s'", where, *input.Header.HeaderDeliveryStatus)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %v", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	idWhere := ""
	if input.Header.DeliverFromParty != nil && input.Header.DeliverToParty != nil {
		idWhere = fmt.Sprintf("\nAND ( DeliverFromParty = %d OR DeliverToParty = %d ) ", *input.Header.DeliverFromParty, *input.Header.DeliverToParty)
	} else if input.Header.DeliverFromParty != nil {
		idWhere = fmt.Sprintf("\nAND DeliverFromParty = %d ", *input.Header.DeliverFromParty)
	} else if input.Header.DeliverToParty != nil {
		idWhere = fmt.Sprintf("\nAND DeliverToParty = %d ", *input.Header.DeliverToParty)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		` + where + idWhere + ` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocument DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
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
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	deliveryDocument := input.Header.DeliveryDocument
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		args = append(args, deliveryDocument, v.DeliveryDocumentItem)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data
		WHERE (DeliveryDocument, DeliveryDocumentItem) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocument DESC, DeliveryDocumentItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	header := input.Header
	where := fmt.Sprintf("where DeliveryDocument = %d", header.DeliveryDocument)

	if header.DeliverFromParty != nil && header.DeliverToParty != nil {
		where = fmt.Sprintf("%s\nAND ( DeliverFromParty = %d OR DeliverToParty = %d ) ", where, *input.Header.DeliverFromParty, *input.Header.DeliverToParty)
	} else if header.DeliverFromParty != nil {
		where = fmt.Sprintf("%s\nAND DeliverFromParty = %d ", where, *header.DeliverFromParty)
	} else if header.DeliverToParty != nil {
		where = fmt.Sprintf("%s\nAND DeliverToParty = %d ", where, *header.DeliverToParty)
	}

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %v", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocument DESC, DeliveryDocumentItem ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	deliveryDocument := input.Header.DeliveryDocument
	partner := input.Header.Partner

	cnt := 0
	for _, v := range partner {
		args = append(args, deliveryDocument, v.PartnerFunction, v.BusinessPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_partner_data
		WHERE (DeliveryDocument, PartnerFunction, BusinessPartner) IN ( `+repeat+` ) 
		ORDER BY DeliveryDocument DESC, BusinessPartner DESC, AddressID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
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
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	deliveryDocument := input.Header.DeliveryDocument
	address := input.Header.Address

	cnt := 0
	for _, v := range address {
		args = append(args, deliveryDocument, v.AddressID)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_address_data
		WHERE (DeliveryDocument, AddressID) IN ( `+repeat+` )
		ORDER BY DeliveryDocument DESC, AddressID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliverFromItems(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliverFromItems {
	deliverFromParty := input.Header.DeliverFromParty

	rows, err := c.db.Query(
		`SELECT DeliveryDocumentHeader.DeliveryDocument,
			DeliveryDocumentHeader.HeaderDeliveryStatus,
			BusinessPartnerGeneral.BusinessPartnerFullName as DeliverFromBusinessPartnerFullName, 
			BusinessPartnerGeneral.BusinessPartnerName as DeliverFromBusinessPartnerName,
			DeliverToBusinessPartnerGeneral.BusinessPartnerFullName as DeliverToBusinessPartnerFullName,
			DeliverToBusinessPartnerGeneral.BusinessPartnerName as DeliverToBusinessPartnerName,
			DeliveryDocumentItem.ItemBillingStatus,
			DeliveryDocumentItem.ConfirmedDeliveryDate
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data as DeliveryDocumentHeader
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as BusinessPartnerGeneral
		ON DeliveryDocumentHeader.DeliverFromParty = BusinessPartnerGeneral.BusinessPartner
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data as DeliveryDocumentItem
		ON DeliveryDocumentHeader.DeliveryDocument = DeliveryDocumentItem.DeliveryDocument
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as DeliverToBusinessPartnerGeneral
		ON DeliveryDocumentHeader.DeliverToParty = DeliverToBusinessPartnerGeneral.BusinessPartner
		WHERE (DeliveryDocumentHeader.DeliverFromParty) = (?) 
		ORDER BY DeliveryDocumentHeader.IsMarkedForDeletion ASC, DeliveryDocumentHeader.IsCancelled ASC, DeliveryDocumentHeader.DeliveryDocument DESC;`, deliverFromParty,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToDeliverFromItems(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliverToItems(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliverToItems {
	deliverToParty := input.Header.DeliverToParty

	rows, err := c.db.Query(
		`SELECT DeliveryDocumentHeader.DeliveryDocument,
			DeliveryDocumentHeader.HeaderDeliveryStatus,
			BusinessPartnerGeneral.BusinessPartnerFullName as DeliverToBusinessPartnerFullName, 
			BusinessPartnerGeneral.BusinessPartnerName as DeliverToBusinessPartnerName,
			DeliverToBusinessPartnerGeneral.BusinessPartnerFullName as DeliverFromBusinessPartnerFullName,
			DeliverToBusinessPartnerGeneral.BusinessPartnerName as DeliverFromBusinessPartnerName,
			DeliveryDocumentItem.ItemBillingStatus,
			DeliveryDocumentItem.ConfirmedDeliveryDate
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data as DeliveryDocumentHeader
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as BusinessPartnerGeneral
		ON DeliveryDocumentHeader.DeliverToParty = BusinessPartnerGeneral.BusinessPartner
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data as DeliveryDocumentItem
		ON DeliveryDocumentHeader.DeliveryDocument = DeliveryDocumentItem.DeliveryDocument
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as DeliverToBusinessPartnerGeneral
		ON DeliveryDocumentHeader.DeliverFromParty = DeliverToBusinessPartnerGeneral.BusinessPartner
		WHERE (DeliveryDocumentHeader.DeliverToParty) = (?) 
		ORDER BY DeliveryDocumentHeader.IsMarkedForDeletion ASC, DeliveryDocumentHeader.IsCancelled ASC, DeliveryDocumentHeader.DeliveryDocument DESC;`, deliverToParty,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToDeliverToItems(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
