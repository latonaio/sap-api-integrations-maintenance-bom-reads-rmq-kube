package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	Document   struct {
		DocumentNo                     string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		Quantity                       string      `json:"quantity"`
		PickedQuantity                 string      `json:"picked_quantity"`
		Price                          string      `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string      `json:"api_schema"`
	MaterialCode            string      `json:"material_code"`
	Plant_Supplier          string      `json:"plant/supplier"`
	Stock                   string      `json:"stock"`
	DocumentType            string      `json:"document_type"`
	DocumentNo              string      `json:"document_no"`
	PlannedDate             string      `json:"planned_date"`
	ValidatedDate           string      `json:"validated_date"`
	Deleted                 bool        `json:"deleted"`
}

type SDC struct {
	ConnectionKey             string `json:"connection_key"`
	Result                    bool   `json:"result"`
	RedisKey                  string `json:"redis_key"`
	Filepath                  string `json:"filepath"`
	MaintenanceBillOfMaterial struct {
		BillOfMaterial                string      `json:"BillOfMaterial"`
		BillOfMaterialCategory        string      `json:"BillOfMaterialCategory"`
		BillOfMaterialVariant         string      `json:"BillOfMaterialVariant"`
		BillOfMaterialVersion         string      `json:"BillOfMaterialVersion"`
		TechnicalObject               string      `json:"TechnicalObject"`
		Plant                         string      `json:"Plant"`
		EngineeringChangeDocument     string      `json:"EngineeringChangeDocument"`
		BillOfMaterialVariantUsage    string      `json:"BillOfMaterialVariantUsage"`
		BillOfMaterialHeaderUUID      string      `json:"BillOfMaterialHeaderUUID"`
		IsMultipleBOMAlt              bool        `json:"IsMultipleBOMAlt"`
		BOMHeaderInternalChangeCount  string      `json:"BOMHeaderInternalChangeCount"`
		BOMUsagePriority              string      `json:"BOMUsagePriority"`
		BillOfMaterialAuthsnGrp       string      `json:"BillOfMaterialAuthsnGrp"`
		BOMVersionStatus              string      `json:"BOMVersionStatus"`
		IsVersionBillOfMaterial       bool        `json:"IsVersionBillOfMaterial"`
		IsLatestBOMVersion            bool        `json:"IsLatestBOMVersion"`
		IsConfiguredMaterial          bool        `json:"IsConfiguredMaterial"`
		BOMTechnicalType              string      `json:"BOMTechnicalType"`
		BOMGroup                      string      `json:"BOMGroup"`
		BOMHeaderText                 string      `json:"BOMHeaderText"`
		BOMAlternativeText            string      `json:"BOMAlternativeText"`
		BillOfMaterialStatus          string      `json:"BillOfMaterialStatus"`
		HeaderValidityStartDate       string      `json:"HeaderValidityStartDate"`
		HeaderValidityEndDate         string      `json:"HeaderValidityEndDate"`
		ChgToEngineeringChgDocument   string      `json:"ChgToEngineeringChgDocument"`
		IsMarkedForDeletion           bool        `json:"IsMarkedForDeletion"`
		BOMHeaderBaseUnit             string      `json:"BOMHeaderBaseUnit"`
		BOMHeaderQuantityInBaseUnit   string      `json:"BOMHeaderQuantityInBaseUnit"`
		RecordCreationDate            string      `json:"RecordCreationDate"`
		LastChangeDate                string      `json:"LastChangeDate"`
		BOMIsToBeDeleted              string      `json:"BOMIsToBeDeleted"`
		MaintenanceBillOfMaterialItem struct {
			TechnicalObject               string      `json:"TechnicalObject"`
			Plant                         string      `json:"Plant"`
			EngineeringChangeDocument     string      `json:"EngineeringChangeDocument"`
			BillOfMaterialVariantUsage    string      `json:"BillOfMaterialVariantUsage"`
			BillOfMaterialItemNodeNumber  string      `json:"BillOfMaterialItemNodeNumber"`
			BillOfMaterialItemUUID        string      `json:"BillOfMaterialItemUUID"`
			BOMItemInternalChangeCount    string      `json:"BOMItemInternalChangeCount"`
			ValidityStartDate             string      `json:"ValidityStartDate"`
			ValidityEndDate               string      `json:"ValidityEndDate"`
			ChgToEngineeringChgDocument   string      `json:"ChgToEngineeringChgDocument"`
			InheritedNodeNumberForBOMItem string      `json:"InheritedNodeNumberForBOMItem"`
			BOMItemRecordCreationDate     string      `json:"BOMItemRecordCreationDate"`
			BOMItemLastChangeDate         string      `json:"BOMItemLastChangeDate"`
			BillOfMaterialComponent       string      `json:"BillOfMaterialComponent"`
			BillOfMaterialItemCategory    string      `json:"BillOfMaterialItemCategory"`
			BillOfMaterialItemNumber      string      `json:"BillOfMaterialItemNumber"`
			BillOfMaterialItemUnit        string      `json:"BillOfMaterialItemUnit"`
			BillOfMaterialItemQuantity    string      `json:"BillOfMaterialItemQuantity"`
			IsAssembly                    string      `json:"IsAssembly"`
			IsSubItem                     bool        `json:"IsSubItem"`
			BOMItemSorter                 string      `json:"BOMItemSorter"`
			PurchasingGroup               string      `json:"PurchasingGroup"`
			Currency                      string      `json:"Currency"`
			MaterialComponentPrice        string      `json:"MaterialComponentPrice"`
			IdentifierBOMItem             string      `json:"IdentifierBOMItem"`
			MaterialPriceUnitQty          string      `json:"MaterialPriceUnitQty"`
			ComponentScrapInPercent       string      `json:"ComponentScrapInPercent"`
			OperationScrapInPercent       string      `json:"OperationScrapInPercent"`
			IsNetScrap                    bool        `json:"IsNetScrap"`
			QuantityVariableSizeItem      string      `json:"QuantityVariableSizeItem"`
			FormulaKey                    string      `json:"FormulaKey"`
			ComponentDescription          string      `json:"ComponentDescription"`
			BOMItemDescription            string      `json:"BOMItemDescription"`
			BOMItemText2                  string      `json:"BOMItemText2"`
			MaterialGroup                 string      `json:"MaterialGroup"`
			DocumentType                  string      `json:"DocumentType"`
			DocNumber                     string      `json:"DocNumber"`
			DocumentVersion               string      `json:"DocumentVersion"`
			DocumentPart                  string      `json:"DocumentPart"`
			ClassNumber                   string      `json:"ClassNumber"`
			ClassType                     string      `json:"ClassType"`
			ResultingItemCategory         string      `json:"ResultingItemCategory"`
			DependencyObjectNumber        string      `json:"DependencyObjectNumber"`
			ObjectType                    string      `json:"ObjectType"`
			IsClassificationRelevant      bool        `json:"IsClassificationRelevant"`
			IsBulkMaterial                bool        `json:"IsBulkMaterial"`
			BOMItemIsSparePart            string      `json:"BOMItemIsSparePart"`
			BOMItemIsSalesRelevant        string      `json:"BOMItemIsSalesRelevant"`
			IsProductionRelevant          bool        `json:"IsProductionRelevant"`
			BOMItemIsPlantMaintRelevant   bool        `json:"BOMItemIsPlantMaintRelevant"`
			BOMItemIsCostingRelevant      string      `json:"BOMItemIsCostingRelevant"`
			IsEngineeringRelevant         bool        `json:"IsEngineeringRelevant"`
			SpecialProcurementType        string      `json:"SpecialProcurementType"`
			IsBOMRecursiveAllowed         bool        `json:"IsBOMRecursiveAllowed"`
			OperationLeadTimeOffset       string      `json:"OperationLeadTimeOffset"`
			OpsLeadTimeOffsetUnit         string      `json:"OpsLeadTimeOffsetUnit"`
			IsMaterialProvision           string      `json:"IsMaterialProvision"`
			BOMIsRecursive                bool        `json:"BOMIsRecursive"`
			DocumentIsCreatedByCAD        bool        `json:"DocumentIsCreatedByCAD"`
			DistrKeyCompConsumption       string      `json:"DistrKeyCompConsumption"`
			DeliveryDurationInDays        string      `json:"DeliveryDurationInDays"`
			Creditor                      string      `json:"Creditor"`
			CostElement                   string      `json:"CostElement"`
			Size1                         string      `json:"Size1"`
			Size2                         string      `json:"Size2"`
			Size3                         string      `json:"Size3"`
			UnitOfMeasureForSize1To3      string      `json:"UnitOfMeasureForSize1To3"`
			GoodsReceiptDuration          string      `json:"GoodsReceiptDuration"`
			PurchasingOrganization        string      `json:"PurchasingOrganization"`
			RequiredComponent             bool        `json:"RequiredComponent"`
			MultipleSelectionAllowed      bool        `json:"MultipleSelectionAllowed"`
			ProdOrderIssueLocation        string      `json:"ProdOrderIssueLocation"`
			MaterialIsCoProduct           bool        `json:"MaterialIsCoProduct"`
			ExplosionType                 string      `json:"ExplosionType"`
			AlternativeItemGroup          string      `json:"AlternativeItemGroup"`
			FollowUpGroup                 string      `json:"FollowUpGroup"`
			DiscontinuationGroup          string      `json:"DiscontinuationGroup"`
			IsConfigurableBOM             string      `json:"IsConfigurableBOM"`
			ReferencePoint                string      `json:"ReferencePoint"`
			LeadTimeOffset                string      `json:"LeadTimeOffset"`
			ProductionSupplyArea          string      `json:"ProductionSupplyArea"`
			IsDeleted                     bool        `json:"IsDeleted"`
			BillOfMaterialHeaderUUID      string      `json:"BillOfMaterialHeaderUUID"`
		} `json:"MaintenanceBillOfMaterialItem"`
	} `json:"MaintenanceBillOfMaterial"`
	APISchema       string   `json:"api_schema"`
	Accepter        []string `json:"accepter"`
	TechnicalObject string   `json:"technical_object"`
	Plant           string   `json:"plant"`
	Deleted         bool     `json:"deleted"`
}
