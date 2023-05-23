package dpfm_api_output_formatter

type OrdersSDC struct {
	ConnectionKey       string             `json:"connection_key"`
	Result              bool               `json:"result"`
	RedisKey            string             `json:"redis_key"`
	Filepath            string             `json:"filepath"`
	APIStatusCode       int                `json:"api_status_code"`
	RuntimeSessionID    string             `json:"runtime_session_id"`
	BusinessPartnerID   *int               `json:"business_partner"`
	ServiceLabel        string             `json:"service_label"`
	APIType             string             `json:"api_type"`
	DataConcatenation   *DataConcatenation `json:"DataConcatenation"`
	APISchema           string             `json:"api_schema"`
	Accepter            []string           `json:"accepter"`
	Deleted             bool               `json:"deleted"`
	APIProcessingResult bool               `json:"api_processing_result"`
	APIProcessingError  *string            `json:"api_processing_error"`
}

type DataConcatenation struct {
	Header OrdersHeader `json:"OrdersHeader"`
	Item   []OrdersItem `json:"OrdersItem"`
}

type OrdersHeader struct {
	ExchangedOrdersDocumentIdentifier                        string   `json:"ExchangedOrdersDocumentIdentifier" csv:"1"`
	OrdersDocument                                           *string  `json:"OrdersDocument" csv:"2"`
	ExchangedDocumentContextSpecifiedTransactionIdentifier   *string  `json:"ExchangedDocumentContextSpecifiedTransactionIdentifier" csv:"3"`
	ExchangedOrdersDocumentName                              *string  `json:"ExchangedOrdersDocumentName" csv:"4"`
	ExchangedOrdersDocumentTypeCode                          *string  `json:"ExchangedOrdersDocumentTypeCode" csv:"5"`
	ExchangedOrdersDocumentIssueDate                         *string  `json:"ExchangedOrdersDocumentIssueDate" csv:"6"`
	ExchangedOrdersDocumentPurposeCode                       *string  `json:"ExchangedOrdersDocumentPurposeCode" csv:"7"`
	ExchangedOrdersDocumentRevisionDate                      *string  `json:"ExchangedOrdersDocumentRevisionDate" csv:"8"`
	ExchangedOrdersDocumentRevisionIdentifier                *string  `json:"ExchangedOrdersDocumentRevisionIdentifier" csv:"9"`
	ExchangedOrdersDocumentStatusCode                        *string  `json:"ExchangedOrdersDocumentStatusCode" csv:"10"`
	ExchangedOrdersDocumentSubtypeCode                       *string  `json:"ExchangedOrdersDocumentSubtypeCode" csv:"11"`
	NoteOrdersSubjectText                                    *string  `json:"NoteOrdersSubjectText" csv:"12"`
	NoteOrdersContentText                                    *string  `json:"NoteOrdersContentText" csv:"13"`
	NoteOrdersIdentifier                                     *string  `json:"NoteOrdersIdentifier" csv:"14"`
	SpecifiedBinaryFileIdentifier                            *string  `json:"SpecifiedBinaryFileIdentifier" csv:"15"`
	SpecifiedBinaryFileVersionIdentifier                     *string  `json:"SpecifiedBinaryFileVersionIdentifier" csv:"16"`
	SpecifiedBinaryFileNameText                              *string  `json:"SpecifiedBinaryFileNameText" csv:"17"`
	SpecifiedBineryFileURIIdentifier                         *string  `json:"SpecifiedBineryFileURIIdentifier" csv:"18"`
	SpecifiedBineryFileMIMECode                              *string  `json:"SpecifiedBineryFileMIMECode" csv:"19"`
	SpecifiedBineryFileEncodingCode                          *string  `json:"SpecifiedBineryFileEncodingCode" csv:"20"`
	SpecifiedBineryFileCode                                  *string  `json:"SpecifiedBineryFileCode" csv:"21"`
	SpecifiedBinaryFileDescriptionText                       *string  `json:"SpecifiedBinaryFileDescriptionText" csv:"22"`
	TradeSellerIdentifier                                    *string  `json:"TradeSellerIdentifier" csv:"23"`
	TradeSellerGlobalIdentifier                              *string  `json:"TradeSellerGlobalIdentifier" csv:"24"`
	TradeSellerName                                          *string  `json:"TradeSellerName" csv:"25"`
	TradeBillFromPartyRegisteredIdentifier                   *string  `json:"TradeBillFromPartyRegisteredIdentifier" csv:"26"`
	TradeContactSellerIdentifier                             *string  `json:"TradeContactSellerIdentifier" csv:"27"`
	TradeContactSellerPersonName                             *string  `json:"TradeContactSellerPersonName" csv:"28"`
	TradeContactSellerDepartmentName                         *string  `json:"TradeContactSellerDepartmentName" csv:"29"`
	SellerTelephoneNumber                                    *string  `json:"SellerTelephoneNumber" csv:"30"`
	SellerFaxNumber                                          *string  `json:"SellerFaxNumber" csv:"31"`
	SellerEmailAddress                                       *string  `json:"SellerEmailAddress" csv:"32"`
	SellerAddressPostalCode                                  *string  `json:"SellerAddressPostalCode" csv:"33"`
	SellerAddress                                            *string  `json:"SellerAddress" csv:"34"`
	TradeBuyerIdentifier                                     *string  `json:"TradeBuyerIdentifier" csv:"35"`
	TradeBuyerGlobalIdentifier                               *string  `json:"TradeBuyerGlobalIdentifier" csv:"36"`
	TradeBuyerName                                           *string  `json:"TradeBuyerName" csv:"37"`
	TradeContactBuyerIdentifier                              *string  `json:"TradeContactBuyerIdentifier" csv:"38"`
	TradeContactBuyerPersonName                              *string  `json:"TradeContactBuyerPersonName" csv:"39"`
	TradeContactBuyerDepartmentName                          *string  `json:"TradeContactBuyerDepartmentName" csv:"40"`
	BuyerTelephoneNumber                                     *string  `json:"BuyerTelephoneNumber" csv:"41"`
	BuyerFaxNumber                                           *string  `json:"BuyerFaxNumber" csv:"42"`
	BuyerEmailAddress                                        *string  `json:"BuyerEmailAddress" csv:"43"`
	BuyerAddressPostalCode                                   *string  `json:"BuyerAddressPostalCode" csv:"44"`
	BuyerAddress                                             *string  `json:"BuyerAddress" csv:"45"`
	ReferencedQuotationReplyDocumentIssureAssignedIdentifier *string  `json:"ReferencedQuotationReplyDocumentIssureAssignedIdentifier" csv:"46"`
	ReferencedQuotationReplyDocumentIssueDate                *string  `json:"ReferencedQuotationReplyDocumentIssueDate" csv:"47"`
	ReferencedQuotationReplyDocumentRevisionIdentifier       *string  `json:"ReferencedQuotationReplyDocumentRevisionIdentifier" csv:"48"`
	ReferencedQuotationReplyDocumentInformationText          *string  `json:"ReferencedQuotationReplyDocumentInformationText" csv:"49"`
	ReferencedQuotationReplyDocumentPurposeCode              *string  `json:"ReferencedQuotationReplyDocumentPurposeCode" csv:"50"`
	ReferencedQuotationReplyDocumentSubtypeCode              *string  `json:"ReferencedQuotationReplyDocumentSubtypeCode" csv:"51"`
	ReferencedSalesOrderDocumentIssureAddignedIdentifier     *string  `json:"ReferencedSalesOrderDocumentIssureAddignedIdentifier" csv:"52"`
	ReferencedSalesOrderDocumentIssueDate                    *string  `json:"ReferencedSalesOrderDocumentIssueDate" csv:"53"`
	ReferencedSalesOrderDocumentRevisionIdentifier           *string  `json:"ReferencedSalesOrderDocumentRevisionIdentifier" csv:"54"`
	ReferencedSalesOrderDocumentInformationText              *string  `json:"ReferencedSalesOrderDocumentInformationText" csv:"55"`
	ReferencedSalesOrderDocumentPurposeCode                  *string  `json:"ReferencedSalesOrderDocumentPurposeCode" csv:"56"`
	ReferencedSalesOrderDocumentSubtypeCode                  *string  `json:"ReferencedSalesOrderDocumentSubtypeCode" csv:"57"`
	RelevantTradePartyIdentifier                             *string  `json:"RelevantTradePartyIdentifier" csv:"58"`
	RelevantTradePartyGlobalIdentifier                       *string  `json:"RelevantTradePartyGlobalIdentifier" csv:"59"`
	RelevantTradePartyName                                   *string  `json:"RelevantTradePartyName" csv:"60"`
	RelevantTradePartyRoleCode                               *string  `json:"RelevantTradePartyRoleCode" csv:"61"`
	RelevantTradeContactCode                                 *string  `json:"RelevantTradeContactCode" csv:"62"`
	RelevantTradeContactPersonName                           *string  `json:"RelevantTradeContactPersonName" csv:"63"`
	ProjectIdentifier                                        *string  `json:"ProjectIdentifier" csv:"64"`
	ProjectName                                              *string  `json:"ProjectName" csv:"65"`
	InspectionEventTypeCode                                  *string  `json:"InspectionEventTypeCode" csv:"66"`
	InspectionEventDescriptionText                           *string  `json:"InspectionEventDescriptionText" csv:"67"`
	ProjectPeriodStartDate                                   *string  `json:"ProjectPeriodStartDate" csv:"68"`
	ProjectPeriodEndDate                                     *string  `json:"ProjectPeriodEndDate" csv:"69"`
	TradeShipToPartyIdentifier                               *string  `json:"TradeShipToPartyIdentifier" csv:"70"`
	TradeShipToPartyGlobalIdentifier                         *string  `json:"TradeShipToPartyGlobalIdentifier" csv:"71"`
	TradeShipToPartyName                                     *string  `json:"TradeShipToPartyName" csv:"72"`
	TradeShipToPartyContactIdentifier                        *string  `json:"TradeShipToPartyContactIdentifier" csv:"73"`
	TradeShipToPartyContactPersonName                        *string  `json:"TradeShipToPartyContactPersonName" csv:"74"`
	TradeShipToPartyContactDepartmentName                    *string  `json:"TradeShipToPartyContactDepartmentName" csv:"75"`
	TradeShipToPartyContactPersonIdentifier                  *string  `json:"TradeShipToPartyContactPersonIdentifier" csv:"76"`
	ShipToPartyTelephoneNumber                               *string  `json:"ShipToPartyTelephoneNumber" csv:"77"`
	ShipToPartyFaxNumber                                     *string  `json:"ShipToPartyFaxNumber" csv:"78"`
	ShipToPartyEmailAddress                                  *string  `json:"ShipToPartyEmailAddress" csv:"79"`
	ShipToPartyAddressPostalCode                             *string  `json:"ShipToPartyAddressPostalCode" csv:"80"`
	ShipToPartyAddress                                       *string  `json:"ShipToPartyAddress" csv:"81"`
	TradeShipFromPartyIdentifier                             *string  `json:"TradeShipFromPartyIdentifier" csv:"82"`
	TradeShipFromPartyName                                   *string  `json:"TradeShipFromPartyName" csv:"83"`
	SupplyChainOperationEventIdentifier                      *string  `json:"SupplyChainOperationEventIdentifier" csv:"84"`
	SupplyChainOperationEventOccurrenceDate                  *string  `json:"SupplyChainOperationEventOccurrenceDate" csv:"85"`
	SupplyChainEventDeliveryTypeCode                         *string  `json:"SupplyChainEventDeliveryTypeCode" csv:"86"`
	SupplyChainEventDeliveryDescription                      *string  `json:"SupplyChainEventDeliveryDescription" csv:"87"`
	InstructedTemperatureControlCode                         *string  `json:"InstructedTemperatureControlCode" csv:"88"`
	SupplyChainTradeCurrencyCode                             *string  `json:"SupplyChainTradeCurrencyCode" csv:"89"`
	TradeTaxCalculatedAmount                                 *float32 `json:"TradeTaxCalculatedAmount" csv:"90"`
	TradeTaxTypeCode                                         *string  `json:"TradeTaxTypeCode" csv:"91"`
	TradeTaxBasisAmount                                      *float32 `json:"TradeTaxBasisAmount" csv:"92"`
	TradeTaxCategoryCode                                     *string  `json:"TradeTaxCategoryCode" csv:"93"`
	TradeTaxCategoryName                                     *string  `json:"TradeTaxCategoryName" csv:"94"`
	TradeTaxRatePercent                                      *string  `json:"TradeTaxRatePercent" csv:"95"`
	TradeTaxGrandTotalAmount                                 *float32 `json:"TradeTaxGrandTotalAmount" csv:"96"`
	TradeTaxCalculationMethod                                *string  `json:"TradeTaxCalculationMethod" csv:"97"`
	TradePaymentTermsDescriptionText                         *string  `json:"TradePaymentTermsDescriptionText" csv:"98"`
	TradePaymentTermsTypeCode                                *string  `json:"TradePaymentTermsTypeCode" csv:"99"`
	TradeSettlementMonetarySummationTotalTaxAmount           *float32 `json:"TradeSettlementMonetarySummationTotalTaxAmount" csv:"100"`
	TradeOrdersSettlementMonetarySummationGrandTotalAmount   *float32 `json:"TradeOrdersSettlementMonetarySummationGrandTotalAmount" csv:"101"`
	TradeOrdersSettlementMonetarySummationNetTotalAmount     *float32 `json:"TradeOrdersSettlementMonetarySummationNetTotalAmount" csv:"102"`
	TradeOrdersMonetarySummationIncludingTaxesTotalAmount    *float32 `json:"TradeOrdersMonetarySummationIncludingTaxesTotalAmount" csv:"103"`
}

type OrdersItem struct {
	ExchangedOrdersDocumentIdentifier                                      string   `json:"ExchangedOrdersDocumentIdentifier" csv:"1"`
	OrdersDocumentItemlineIdentifier                                       string   `json:"OrdersDocumentItemlineIdentifier" csv:"104"`
	OrdersDocumentItemlineStatusCode                                       *string  `json:"OrdersDocumentItemlineStatusCode" csv:"105"`
	OrdersDocumentItemlineStatusReasonCode                                 *string  `json:"OrdersDocumentItemlineStatusReasonCode" csv:"106"`
	OrdersDocumentItemIdentifier                                           *string  `json:"OrdersDocumentItemIdentifier" csv:"107"`
	OrdersDocumentItemBuyerAssignedCategoryCode                            *string  `json:"OrdersDocumentItemBuyerAssignedCategoryCode" csv:"108"`
	NoteOrdersItemItemSubjectText                                          *string  `json:"NoteOrdersItemItemSubjectText" csv:"109"`
	NoteOrdersItemContentText                                              *string  `json:"NoteOrdersItemContentText" csv:"110"`
	NoteOrdersItemIdentifier                                               *string  `json:"NoteOrdersItemIdentifier" csv:"111"`
	ReferencedQuotationReplyDocumentItemIssuerAssignedIdentifier           *string  `json:"ReferencedQuotationReplyDocumentItemIssuerAssignedIdentifier" csv:"112"`
	ReferencedQuotationReplyDocumentItemLineIdentifier                     *string  `json:"ReferencedQuotationReplyDocumentItemLineIdentifier" csv:"113"`
	ReferencedQuotationReplyDocumentItemRevisionIdentifier                 *string  `json:"ReferencedQuotationReplyDocumentItemRevisionIdentifier" csv:"114"`
	ReferencedQuotationReplyDocumentItemInformationText                    *string  `json:"ReferencedQuotationReplyDocumentItemInformationText" csv:"115"`
	ReferencedDocumentItemIssureAssignedIdentifier                         *string  `json:"ReferencedDocumentItemIssureAssignedIdentifier" csv:"116"`
	TradePriceTypeCode                                                     *string  `json:"TradePriceTypeCode" csv:"117"`
	TradeOrdersPriceChargeAmount                                           *float32 `json:"TradeOrdersPriceChargeAmount" csv:"118"`
	TradePriceBasisQuantity                                                *float32 `json:"TradePriceBasisQuantity" csv:"119"`
	TradePriceBasisUnitCode                                                *string  `json:"TradePriceBasisUnitCode" csv:"120"`
	TradeDeliveryTermsDescriptionText                                      *string  `json:"TradeDeliveryTermsDescriptionText" csv:"121"`
	SupplyChainTradeDeliveryRequestedQuantity                              *float32 `json:"SupplyChainTradeDeliveryRequestedQuantity" csv:"122"`
	SupplyChainTradeDeliveryPerPackageUnitQuantity                         *float32 `json:"SupplyChainTradeDeliveryPerPackageUnitQuantity" csv:"123"`
	SupplyChainTradeDeliveryPackageQuantity                                *float32 `json:"SupplyChainTradeDeliveryPackageQuantity" csv:"124"`
	SupplyChainTradeDeliveryProductUnitQuantity                            *float32 `json:"SupplyChainTradeDeliveryProductUnitQuantity" csv:"125"`
	QuantityUnitCode                                                       *string  `json:"QuantityUnitCode" csv:"126"`
	ItemTradeDeliverToPartyIdentifier                                      *string  `json:"ItemTradeDeliverToPartyIdentifier" csv:"127"`
	ItemTradeDeliverToPartyGlobalIdentifier                                *string  `json:"ItemTradeDeliverToPartyGlobalIdentifier" csv:"128"`
	ItemTradeDeliverToPartyName                                            *string  `json:"ItemTradeDeliverToPartyName" csv:"129"`
	ItemTradeDeliverToPartyContactPersonName                               *string  `json:"ItemTradeDeliverToPartyContactPersonName" csv:"130"`
	ItemTradeDeliverToPartyContactDepartment                               *string  `json:"ItemTradeDeliverToPartyContactDepartment" csv:"131"`
	ItemDeliverToPartyPhoneNumber                                          *string  `json:"ItemDeliverToPartyPhoneNumber" csv:"132"`
	ItemDeliverToPartyFaxNumber                                            *string  `json:"ItemDeliverToPartyFaxNumber" csv:"133"`
	ItemDeliverToPartyEMailAddress                                         *string  `json:"ItemDeliverToPartyEMailAddress" csv:"134"`
	ItemDeliverToPartyAddressPostalCode                                    *string  `json:"ItemDeliverToPartyAddressPostalCode" csv:"135"`
	ItemDeliverToPartyAddress                                              *string  `json:"ItemDeliverToPartyAddress" csv:"136"`
	DeliverInstructionsHandlingCode                                        *string  `json:"DeliverInstructionsHandlingCode" csv:"137"`
	SupplyChainEventRequirementOccurrenceDate                              *string  `json:"SupplyChainEventRequirementOccurrenceDate" csv:"138"`
	SupplyChainEventTypeCode                                               *string  `json:"SupplyChainEventTypeCode" csv:"139"`
	SupplyChainEventRequirementOccurrenceTime                              *string  `json:"SupplyChainEventRequirementOccurrenceTime" csv:"140"`
	LogisticsLocationIdentification                                        *string  `json:"LogisticsLocationIdentification" csv:"141"`
	LogisticsLocationName                                                  *string  `json:"LogisticsLocationName" csv:"142"`
	TradeTaxTypeCode                                                       *string  `json:"TradeTaxTypeCode" csv:"143"`
	ItemTradeTaxBasisAmount                                                *float32 `json:"ItemTradeTaxBasisAmount" csv:"144"`
	ItemTradeTaxCategoryCode                                               *string  `json:"ItemTradeTaxCategoryCode" csv:"145"`
	TradeTaxCategoryName                                                   *string  `json:"TradeTaxCategoryName" csv:"146"`
	ItemTradeTaxRateApplicablePercent                                      *float32 `json:"ItemTradeTaxRateApplicablePercent" csv:"147"`
	ItemTradeTaxGrandTotalAmount                                           *float32 `json:"ItemTradeTaxGrandTotalAmount" csv:"148"`
	ItemTradeOrdersSettlementMonetarySummationNetTotalAmount               *float32 `json:"ItemTradeOrdersSettlementMonetarySummationNetTotalAmount" csv:"149"`
	ItemTradeOrdersSettlementMonetarySummationIncludingTaxesNetTotalAmount *float32 `json:"ItemTradeOrdersSettlementMonetarySummationIncludingTaxesNetTotalAmount" csv:"150"`
	TradeProductIdentifier                                                 *string  `json:"TradeProductIdentifier" csv:"151"`
	TradeProductGlobalIdentifier                                           *string  `json:"TradeProductGlobalIdentifier" csv:"152"`
	TradeProductSellerAssignedIdentifier                                   *string  `json:"TradeProductSellerAssignedIdentifier" csv:"153"`
	TradeProductBuyerAssignedIdentifier                                    *string  `json:"TradeProductBuyerAssignedIdentifier" csv:"154"`
	TradeProductManufacturerAssignedIdentifier                             *string  `json:"TradeProductManufacturerAssignedIdentifier" csv:"155"`
	TradeProductName                                                       *string  `json:"TradeProductName" csv:"156"`
	TradeProductDescription                                                *string  `json:"TradeProductDescription" csv:"157"`
	TradeProductTypeCode                                                   *string  `json:"TradeProductTypeCode" csv:"158"`
	TradeProductEndItemTypeCode                                            *string  `json:"TradeProductEndItemTypeCode" csv:"159"`
	TradeProductSizeCode                                                   *string  `json:"TradeProductSizeCode" csv:"160"`
	TradeProductSizeDescriptionText                                        *string  `json:"TradeProductSizeDescriptionText" csv:"161"`
	ProductCharacteristicIdentifier                                        *string  `json:"ProductCharacteristicIdentifier" csv:"162"`
	ProductCharacteristicTypeCode                                          *string  `json:"ProductCharacteristicTypeCode" csv:"163"`
	ProductCharacteristicDescriptionText                                   *string  `json:"ProductCharacteristicDescriptionText" csv:"164"`
	ProductCharacteristicContentTypeCode                                   *string  `json:"ProductCharacteristicContentTypeCode" csv:"165"`
	TradeProductInstanceGlobalSerialIdentifier                             *string  `json:"TradeProductInstanceGlobalSerialIdentifier" csv:"166"`
	TradeProductInstanceBatchIdentifier                                    *string  `json:"TradeProductInstanceBatchIdentifier" csv:"167"`
	ReferenceQualityInspectionDocumentInformationText                      *string  `json:"ReferenceQualityInspectionDocumentInformationText" csv:"168"`
	TradeManufacturerIdentifier                                            *string  `json:"TradeManufacturerIdentifier" csv:"169"`
	TradeManufacturerName                                                  *string  `json:"TradeManufacturerName" csv:"170"`
	ReferencedTechnicalDocumentIssuerAssignedIdentifier                    *string  `json:"ReferencedTechnicalDocumentIssuerAssignedIdentifier" csv:"171"`
	ReferencedTechnicalDocumentRevisionIdentifier                          *string  `json:"ReferencedTechnicalDocumentRevisionIdentifier" csv:"172"`
	ReferencedTechnicalDocumentName                                        *string  `json:"ReferencedTechnicalDocumentName" csv:"173"`
	ReferencedTechnicalDocumentInformationText                             *string  `json:"ReferencedTechnicalDocumentInformationText" csv:"174"`
	ReferencedTechnicalDocumentAttachment                                  *string  `json:"ReferencedTechnicalDocumentAttachment" csv:"175"`
	ReferencedSupplyDocumentIssuerAssignedIdentifier                       *string  `json:"ReferencedSupplyDocumentIssuerAssignedIdentifier" csv:"176"`
	ReferencedSupplyDocumentAttachmentBinaryObject                         *string  `json:"ReferencedSupplyDocumentAttachmentBinaryObject" csv:"177"`
	ReferencedLogisticsPackageUnitQuantity                                 *float32 `json:"ReferencedLogisticsPackageUnitQuantity" csv:"178"`
	ReferencedLogisticsPackageQuantityUnitCode                             *string  `json:"ReferencedLogisticsPackageQuantityUnitCode" csv:"179"`
	ReferencedLogisticsPackageTypeCode                                     *string  `json:"ReferencedLogisticsPackageTypeCode" csv:"180"`
	GoodsProductionIdentifier                                              *string  `json:"GoodsProductionIdentifier" csv:"181"`
	GoodsProductionManufacturingProcessDescriptionText                     *string  `json:"GoodsProductionManufacturingProcessDescriptionText" csv:"182"`
}
