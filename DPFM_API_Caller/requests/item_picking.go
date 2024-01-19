package requests

type ItemPicking struct {
	DeliveryDocument                                 int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                             int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemPickingID                    int      `json:"DeliveryDocumentItemPickingID"`
	SupplyChainRelationshipID                        int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID                int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID           int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                            int      `json:"Buyer"`
	Seller                                           int      `json:"Seller"`
	Product                                          string   `json:"Product"`
	DeliverToParty                                   int      `json:"DeliverToParty"`
	DeliverToPlant                                   string   `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                    string   `json:"DeliverToPlantStorageLocation"`
	DeliverToPlantStorageBin                         *string  `json:"DeliverToPlantStorageBin"`
	DeliverToPlantKanbanContainer                    *int     `json:"DeliverToPlantKanbanContainer"`
	DeliverFromParty                                 int      `json:"DeliverFromParty"`
	DeliverFromPlant                                 string   `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation                  string   `json:"DeliverFromPlantStorageLocation"`
	DeliverFromPlantStorageBin                       *string  `json:"DeliverFromPlantStorageBin"`
	DeliverFromPlantKanbanContainer                  *int	  `json:"DeliverFromPlantKanbanContainer"`
	DeliverToPlantPlannedPickingQuantityInBaseUnit   float32  `json:"DeliverToPlantPlannedPickingQuantityInBaseUnit"`
	DeliverFromPlantPlannedPickingQuantityInBaseUnit float32  `json:"DeliverFromPlantPlannedPickingQuantityInBaseUnit"`
	DeliverToPlantPlannedPickingDate                 string   `json:"DeliverToPlantPlannedPickingDate"`
	DeliverToPlantPlannedPickingTime                 string   `json:"DeliverToPlantPlannedPickingTime"`
	DeliverFromPlantPlannedPickingDate               string   `json:"DeliverFromPlantPlannedPickingDate"`
	DeliverFromPlantPlannedPickingTime               string   `json:"DeliverFromPlantPlannedPickingTime"`
	DeliverToPlantActualPickingQuantityInBaseUnit    *float32 `json:"DeliverToPlantActualPickingQuantityInBaseUnit"`
	DeliverToPlantActualPickingDate                  *string  `json:"DeliverToPlantActualPickingDate"`
	DeliverToPlantActualPickingTime                  *string  `json:"DeliverToPlantActualPickingTime"`
	DeliverFromPlantActualPickingQuantityInBaseUnit  *float32 `json:"DeliverFromPlantActualPickingQuantityInBaseUnit"`
	DeliverFromPlantActualPickingDate                *string  `json:"DeliverFromPlantActualPickingDate"`
	DeliverFromPlantActualPickingTime                *string  `json:"DeliverFromPlantActualPickingTime"`
	ExternalReferenceDocument        			  	 *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem      			  	 *string  `json:"ExternalReferenceDocumentItem"`
	ExternalReferenceDocumentItemPickingID		  	 *string  `json:"ExternalReferenceDocumentItemPickingID"`
	CreationDate                                     string   `json:"CreationDate"`
	CreationTime                                     string   `json:"CreationTime"`
	LastChangeDate                                   string   `json:"LastChangeDate"`
	LastChangeTime                                   string   `json:"LastChangeTime"`
	IsCancelled                                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                              *bool    `json:"IsMarkedForDeletion"`
}
