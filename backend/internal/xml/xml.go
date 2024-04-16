package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

// StandardBusinessDocument represents the root element of the XML
type StandardBusinessDocument struct {
	XMLName                        xml.Name `xml:"StandardBusinessDocument"`
	XMLNS                          string   `xml:"xmlns,attr"`
	StandardBusinessDocumentHeader StandardBusinessDocumentHeader
	Invoice                        Invoice `xml:"ubl:Invoice"`
}

// StandardBusinessDocumentHeader represents the StandardBusinessDocumentHeader element
type StandardBusinessDocumentHeader struct {
	HeaderVersion          string `xml:"HeaderVersion"`
	Sender                 Sender
	Receiver               Receiver
	DocumentIdentification DocumentIdentification
	BusinessScope          BusinessScope
}

// Sender represents the Sender element
type Sender struct {
	Identifier Identifier `xml:"Identifier"`
}

// Receiver represents the Receiver element
type Receiver struct {
	Identifier Identifier `xml:"Identifier"`
}

type Identifier struct {
	Identifier string `xml:",chardata"`
	Authority  string `xml:"Authority,attr"`
}

// DocumentIdentification represents the DocumentIdentification element
type DocumentIdentification struct {
	Standard            string `xml:"Standard"`
	TypeVersion         string `xml:"TypeVersion"`
	InstanceIdentifier  string `xml:"InstanceIdentifier"`
	Type                string `xml:"Type"`
	CreationDateAndTime string `xml:"CreationDateAndTime"`
}

// BusinessScope represents the BusinessScope element
type BusinessScope struct {
	Scope []Scope `xml:"Scope"`
}

// Scope represents the Scope element
type Scope struct {
	Type               string `xml:"Type"`
	InstanceIdentifier string `xml:"InstanceIdentifier"`
	Identifier         string `xml:"Identifier"`
}

// Invoice represents the Invoice element
type Invoice struct {
	XMLNS_ubl               string             `xml:"xmlns:ubl,attr"`
	XMLNS_cac               string             `xml:"xmlns:cac,attr"`
	XMLNS_cbc               string             `xml:"xmlns:cbc,attr"`
	XMLNS_xsi               string             `xml:"xmlns:xsi,attr"`
	XMLNS_qdt               string             `xml:"xmlns:qdt,attr"`
	XMLNS_udt               string             `xml:"xmlns:udt,attr"`
	XMLNS_ccts              string             `xml:"xmlns:ccts,attr"`
	XMLNS                   string             `xml:"xmlns,attr"`
	SchemaLocation          string             `xml:"xsi:schemaLocation,attr"`
	CustomizationID         CustomizationID    `xml:"cbc:CustomizationID"`
	ProfileID               string             `xml:"cbc:ProfileID"`
	ID                      string             `xml:"cbc:ID"`
	IssueDate               string             `xml:"cbc:IssueDate"`
	DueDate                 string             `xml:"cbc:DueDate"`
	InvoiceTypeCode         string             `xml:"cbc:InvoiceTypeCode"`
	DocumentCurrencyCode    string             `xml:"cbc:DocumentCurrencyCode"`
	OrderReference          CBC_ID     `xml:"cac:OrderReference"`
	AccountingSupplierParty AccountingParty    `xml:"cac:AccountingSupplierParty"`
	AccountingCustomerParty AccountingParty    `xml:"cac:AccountingCustomerParty"`
	Delivery                Delivery           `xml:"cac:Delivery"`
	PaymentMeans            PaymentMeans       `xml:"cac:PaymentMeans"`
	AllowanceCharge         []AllowanceCharge  `xml:"cac:AllowanceCharge"`
	TaxTotal                TaxTotal           `xml:"cac:TaxTotal"`
	LegalMonetaryTotal      LegalMonetaryTotal `xml:"cac:LegalMonetaryTotal"`
	InvoiceLine             InvoiceLine        `xml:"cac:InvoiceLine"`
}

type CustomizationID struct {
	CustomizationID string `xml:",chardata"`
	XMLNS_cbc       string `xml:"xmlns:cbc,attr"`
}

// OrderReference represents the OrderReference element
type CBC_ID struct {
	ID string `xml:"cbc:ID"`
}

// AccountingParty represents the AccountingSupplierParty and AccountingCustomerParty elements
type AccountingParty struct {
	Party Party `xml:"cac:Party"`
}

// Party represents the Party element
type Party struct {
	EndpointID       EndpointID       `xml:"cbc:EndpointID"`
	PartyName        PartyName        `xml:"cac:PartyName"`
	PostalAddress    PostalAddress    `xml:"cac:PostalAddress"`
	PartyTaxScheme   PartyTaxScheme   `xml:"cac:PartyTaxScheme"`
	PartyLegalEntity PartyLegalEntity `xml:"cac:PartyLegalEntity"`
	Contact          Contact          `xml:"cac:Contact"`
}

type EndpointID struct {
	EndpointID string `xml:",chardata"`
	SchemeID   string `xml:"schemeID,attr"`
}

// PartyName represents the PartyName element
type PartyName struct {
	Name string `xml:"cbc:Name"`
}

// PostalAddress represents the PostalAddress element
type PostalAddress struct {
	StreetName string  `xml:"cbc:StreetName"`
	CityName   string  `xml:"cbc:CityName"`
	PostalZone string  `xml:"cbc:PostalZone"`
	Country    Country `xml:"cac:Country"`
}

// Country represents the Country element
type Country struct {
	IdentificationCode string `xml:"cbc:IdentificationCode"`
}

// PartyTaxScheme represents the PartyTaxScheme element
type PartyTaxScheme struct {
	CompanyID string    `xml:"cbc:CompanyID"`
	TaxScheme CBC_ID `xml:"cac:TaxScheme"`
}

// PartyLegalEntity represents the PartyLegalEntity element
type PartyLegalEntity struct {
	RegistrationName string `xml:"cbc:RegistrationName"`
}

// Contact represents the Contact element
type Contact struct {
	Name           string `xml:"cbc:Name"`
	Telephone      string `xml:"cbc:Telephone"`
	ElectronicMail string `xml:"cbc:ElectronicMail"`
}

// Delivery represents the Delivery element
type Delivery struct {
	ActualDeliveryDate string `xml:"cbc:ActualDeliveryDate"`
}

// PaymentMeans represents the PaymentMeans element
type PaymentMeans struct {
	PaymentMeansCode      string                `xml:"cbc:PaymentMeansCode"`
	PaymentID             string                `xml:"cbc:PaymentID"`
	PayeeFinancialAccount CBC_ID `xml:"cac:PayeeFinancialAccount"`
}

// AllowanceCharge represents the AllowanceCharge element
type AllowanceCharge struct {
	ChargeIndicator       string      `xml:"cbc:ChargeIndicator"`
	AllowanceChargeReason string      `xml:"cbc:AllowanceChargeReason"`
	Amount                Amount      `xml:"cbc:Amount"`
	TaxCategory           TaxCategory `xml:"cac:TaxCategory"`
}

type Amount struct {
	Amount   string `xml:",chardata"`
	Currency string `xml:"currencyID,attr"`
}

// TaxCategory represents the TaxCategory element
type TaxCategory struct {
	ID        string    `xml:"cbc:ID"`
	Percent   string    `xml:"cbc:Percent"`
	TaxScheme CBC_ID    `xml:"cac:TaxScheme"`
}

// TaxTotal represents the TaxTotal element
type TaxTotal struct {
	TaxAmount   Amount      `xml:"cbc:TaxAmount"`
	TaxSubtotal TaxSubtotal `xml:"cac:TaxSubtotal"`
}

// TaxSubtotal represents the TaxSubtotal element
type TaxSubtotal struct {
	TaxableAmount Amount      `xml:"cbc:TaxableAmount"`
	TaxAmount     Amount      `xml:"cbc:TaxAmount"`
	TaxCategory   TaxCategory `xml:"cac:TaxCategory"`
}

// LegalMonetaryTotal represents the LegalMonetaryTotal element
type LegalMonetaryTotal struct {
	LineExtensionAmount Amount `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount  Amount `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount  Amount `xml:"cbc:TaxInclusiveAmount"`
	ChargeTotalAmount   Amount `xml:"cbc:ChargeTotalAmount"`
	PayableAmount       Amount `xml:"cbc:PayableAmount"`
}

// InvoiceLine represents the InvoiceLine element
type InvoiceLine struct {
	ID                  string `xml:"cbc:ID"`
	InvoicedQuantity    InvoicedQuantity `xml:"cbc:InvoicedQuantity"`
	LineExtensionAmount Amount `xml:"cbc:LineExtensionAmount"`
	Item                Item `xml:"cac:Item"`
	Price               Price `xml:"cac:Price"`
}

type InvoicedQuantity struct {
    InvoicedQuantity string `xml:",chardata"`
    UnitCode         string `xml:"currencyID,attr"`
}

// Item represents the Item element
type Item struct {
	Name                      string `xml:"cbc:Name"`
	SellersItemIdentification CBC_ID `xml:"cac:SellersItemIdentification"`
	ClassifiedTaxCategory     ClassifiedTaxCategory `xml:"cac:ClassifiedTaxCategory"`
}

// ClassifiedTaxCategory represents the ClassifiedTaxCategory element
type ClassifiedTaxCategory struct {
	ID        string `xml:"cbc:ID"`
	Percent   string `xml:"cbc:Percent"`
	TaxScheme CBC_ID `xml:"cac:TaxScheme"`
}

// Price represents the Price element
type Price struct {
	PriceAmount  Amount `xml:"cbc:PriceAmount"`
	BaseQuantity string `xml:"cbc:BaseQuantity"`
}

func NewXMLFile() error {
	// Create a new instance of the XML struct with sample data
	xmlData := StandardBusinessDocument{
		XMLNS: "http://www.unece.org/cefact/namespaces/StandardBusinessDocumentHeader",
		StandardBusinessDocumentHeader: StandardBusinessDocumentHeader{
			HeaderVersion: "1.0",
			Sender: Sender{
				Identifier: Identifier{
					"9938:LU21086120",
					"iso6523-actorid-upis",
				},
			},
			Receiver: Receiver{
				Identifier: Identifier{
					"9938:LU28959578",
					"iso6523-actorid-upis",
				},
			},
			DocumentIdentification: DocumentIdentification{
				Standard:            "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2",
				TypeVersion:         "2.1",
				InstanceIdentifier:  "6c4f9576-26fd-4674-b787-ef0b4b06c150",
				Type:                "Invoice",
				CreationDateAndTime: "2024-04-14T10:35:51.173Z",
			},
			BusinessScope: BusinessScope{
				Scope: []Scope{
					{
						Type:               "DOCUMENTID",
						InstanceIdentifier: "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2::Invoice##urn:cen.eu:en16931:2017#compliant#urn:fdc:peppol.eu:2017:poacc:billing:3.0::2.1",
						Identifier:         "busdox-docid-qns",
					},
					{
						Type:               "PROCESSID",
						InstanceIdentifier: "urn:fdc:peppol.eu:2017:poacc:billing:01:1.0",
						Identifier:         "cenbii-procid-ubl",
					},
				},
			},
		},
		Invoice: Invoice{
            XMLNS_ubl: "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2",
            XMLNS_cac: "urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2",
            XMLNS_cbc: "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2",
            XMLNS_xsi: "http://www.w3.org/2001/XMLSchema-instance",
            XMLNS_qdt: "urn:oasis:names:specification:ubl:schema:xsd:QualifiedDataTypes-2",
            XMLNS_udt: "urn:oasis:names:specification:ubl:schema:xsd:UnqualifiedDataTypes-2",
            XMLNS_ccts: "urn:un:unece:uncefact:documentation:2",
            XMLNS: "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2",
            SchemaLocation: "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2 http://docs.oasis-open.org/ubl/os-UBL-2.1/xsd/maindoc/UBL-Invoice-2.1.xsd",
			CustomizationID: CustomizationID{
                CustomizationID: "urn:cen.eu:en16931:2017#compliant#urn:fdc:peppol.eu:2017:poacc:billing:3.0",
                XMLNS_cbc: "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2",
            },
			ProfileID:            "urn:fdc:peppol.eu:2017:poacc:billing:01:1.0",
			ID:                   "1026127",
			IssueDate:            "2024-04-14",
			DueDate:              "2024-05-14",
			InvoiceTypeCode:      "380",
			DocumentCurrencyCode: "EUR",
		},
	}

	// Create a new XML file
	xmlFile, err := os.Create("invoice.xml")
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	// Marshal the XML data into bytes
	xmlBytes, err := xml.MarshalIndent(xmlData, "", "    ")
	if err != nil {
		return err
	}

	// Write the XML data to the file
	_, err = xmlFile.Write(xmlBytes)
	if err != nil {
		return err
	}

	fmt.Println("XML file created successfully")

	return nil
}
