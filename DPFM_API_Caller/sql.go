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
	var itemPicking *[]dpfm_api_output_formatter.ItemPicking
	var partner *[]dpfm_api_output_formatter.Partner
	var address *[]dpfm_api_output_formatter.Address
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
			//		case "Headers":
			//			func() {
			//				header = c.Headers(mtx, input, output, errs, log)
			//			}()
		case "HeadersByDeliverToParty":
			func() {
				header = c.HeadersByDeliverToParty(mtx, input, output, errs, log)
			}()
		case "HeadersByDeliverFromParty":
			func() {
				header = c.HeadersByDeliverFromParty(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()
		case "ItemPicking":
			func() {
				itemPicking = c.ItemPicking(mtx, input, output, errs, log)
			}()
		case "ItemPickings":
			func() {
				itemPicking = c.ItemPickings(mtx, input, output, errs, log)
			}()
		case "Partner":
			func() {
				partner = c.Partner(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:      header,
		Item:        item,
		ItemPicking: itemPicking,
		Partner:     partner,
		Address:     address,
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
		` + where + ` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocument ASC;`,
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

func (c *DPFMAPICaller) HeadersByDeliverToParty(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.DeliverToParty != nil {
		where = fmt.Sprintf("%s\nAND DeliverToParty = %v", where, *input.Header.DeliverToParty)
	}
	if input.Header.HeaderCompleteDeliveryIsDefined != nil {
		where = fmt.Sprintf("%s\nAND HeaderCompleteDeliveryIsDefined = %t", where, *input.Header.HeaderCompleteDeliveryIsDefined)
	}
	if input.Header.HeaderDeliveryBlockStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderDeliveryBlockStatus = %t", where, *input.Header.HeaderDeliveryBlockStatus)
	}
	if input.Header.HeaderDeliveryStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderDeliveryStatus != '%s'", where, *input.Header.HeaderDeliveryStatus)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %t", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		` + where + `;`,
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

func (c *DPFMAPICaller) HeadersByDeliverFromParty(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.DeliverFromParty != nil {
		where = fmt.Sprintf("%s\nAND DeliverFromParty = %v", where, *input.Header.DeliverFromParty)
	}
	if input.Header.HeaderCompleteDeliveryIsDefined != nil {
		where = fmt.Sprintf("%s\nAND HeaderCompleteDeliveryIsDefined = %t", where, *input.Header.HeaderCompleteDeliveryIsDefined)
	}
	if input.Header.HeaderDeliveryBlockStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderDeliveryBlockStatus = %t", where, *input.Header.HeaderDeliveryBlockStatus)
	}
	if input.Header.HeaderDeliveryStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderDeliveryStatus != '%s'", where, *input.Header.HeaderDeliveryStatus)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %t", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		` + where + `;`,
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

//func (c *DPFMAPICaller) Headers(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *[]dpfm_api_output_formatter.Header {
//	where := "WHERE 1 = 1"
//	if input.Header.HeaderCompleteDeliveryIsDefined != nil {
//		where = fmt.Sprintf("%s\nAND HeaderCompleteDeliveryIsDefined = %v", where, *input.Header.HeaderCompleteDeliveryIsDefined)
//	}
//	if input.Header.HeaderDeliveryBlockStatus != nil {
//		where = fmt.Sprintf("%s\nAND HeaderDeliveryBlockStatus = %v", where, *input.Header.HeaderDeliveryBlockStatus)
//	}
//	if input.Header.HeaderDeliveryStatus != nil {
//		where = fmt.Sprintf("%s\nAND HeaderDeliveryStatus = '%s'", where, *input.Header.HeaderDeliveryStatus)
//	}
//	if input.Header.IsCancelled != nil {
//		where = fmt.Sprintf("%s\nAND IsCancelled = %v", where, *input.Header.IsCancelled)
//	}
//	if input.Header.IsMarkedForDeletion != nil {
//		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
//	}
// if input.Header.HeaderBillingStatusException != nil {
// 	where = fmt.Sprintf("%s\nAND HeaderBillingStatus != '%v' ", where, *input.Header.HeaderBillingStatusException)
// }
//
//	idWhere := ""
//	if input.Header.DeliverFromParty != nil && input.Header.DeliverToParty != nil {
//		idWhere = fmt.Sprintf("\nAND ( DeliverFromParty = %d OR DeliverToParty = %d ) ", *input.Header.DeliverFromParty, *input.Header.DeliverToParty)
//	} else if input.Header.DeliverFromParty != nil {
//		idWhere = fmt.Sprintf("\nAND DeliverFromParty = %d ", *input.Header.DeliverFromParty)
//	} else if input.Header.DeliverToParty != nil {
//		idWhere = fmt.Sprintf("\nAND DeliverToParty = %d ", *input.Header.DeliverToParty)
//	}
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
//		` + where + idWhere + ` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocument ASC;`,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}

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
		ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocument ASC, DeliveryDocumentItem ASC;`, args...,
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
	item := &dpfm_api_input_reader.Item{}

	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	where := "WHERE 1 = 1"

	deliveryDocument := input.Header.DeliveryDocument

	if deliveryDocument != 0 {
		where = fmt.Sprintf("%s\nAND item.DeliveryDocument = %d", where, deliveryDocument)
	}

	if item != nil {
		if item.ItemCompleteDeliveryIsDefined != nil {
			where = fmt.Sprintf("%s\nAND item.ItemCompleteDeliveryIsDefined = %v", where, *item.ItemCompleteDeliveryIsDefined)
		}
		if item.DeliveryDocumentItem != 0 {
			where = fmt.Sprintf("%s\nAND item.DeliveryDocumentItem = '%v'", where, item.DeliveryDocumentItem)
		}
		if item.ItemDeliveryBlockStatus != nil {
			where = fmt.Sprintf("%s\nAND item.ItemDeliveryBlockStatus = %v", where, *item.ItemDeliveryBlockStatus)
		}
		if item.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND item.IsCancelled = %v", where, *item.IsCancelled)
		}
		if item.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND item.IsMarkedForDeletion = %v", where, *item.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data as item
		` + where + ` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocumentItem ASC;`,
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

func (c *DPFMAPICaller) ItemPicking(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPicking {
	var args []interface{}
	deliveryDocument := input.Header.DeliveryDocument
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemPicking := v.ItemPicking
		for _, w := range itemPicking {
			args = append(args, deliveryDocument, v.DeliveryDocumentItem, w.DeliveryDocumentItemPickingID)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data
		WHERE (DeliveryDocument, DeliveryDocumentItem, DeliveryDocumentItemPickingID) IN ( `+repeat+` )
		ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, DeliveryDocument ASC, DeliveryDocumentItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPicking(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPickings(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPicking {
	item := &dpfm_api_input_reader.Item{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	where := fmt.Sprintf("WHERE deliveryDocument = %v", input.Header.DeliveryDocument)
	if item != nil {
		if item.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *item.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_item_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPicking(rows)
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
		ORDER BY DeliveryDocument ASC, BusinessPartner ASC, AddressID ASC;`, args...,
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
		ORDER BY DeliveryDocument ASC, AddressID ASC;`, args...,
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
