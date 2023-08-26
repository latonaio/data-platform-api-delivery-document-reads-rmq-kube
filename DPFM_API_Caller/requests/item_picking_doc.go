package requests

type ItemPickingDoc struct {
	DeliveryDocument              int     `json:"DeliveryDocument"`
	DeliveryDocumentItem          int     `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemPickingID int     `json:"DeliveryDocumentItemPickingID"`
	DocType                       string  `json:"DocType"`
	DocVersionID                  int     `json:"DocVersionID"`
	DocID                         string  `json:"DocID"`
	FileExtension                 string  `json:"FileExtension"`
	FileName                      *string `json:"FileName"`
	FilePath                      *string `json:"FilePath"`
	DocIssuerBusinessPartner      *int    `json:"DocIssuerBusinessPartner"`
}
