package xml

type StandardBusinessDocument struct {
	StandardBusinessDocumentHeader *StandardBusinessDocumentHeader
	Invoice *Invoice
}

type StandardBusinessDocumentHeader struct {
	HeaderVersion string 
	Sender string 
	Receiver string 
	DocumentIdentification string 
	BusinessScope string 
}

type Invoice struct {
	CustomizationID string
	ProfileID string
	ID string
	IssueDate string
	DueDate string
	InvoiceTypeRecorder string
	DocumentCurrencyCode string
	OrderReference string
	AccountingSupplierParty string
	AccountingCustomerParty string
	Delivery string
	PaymnentMeans string
	AllowanceCharge string //could be multiple occurences
	TaxTotal string
	LegalMonetaryTotal string
}