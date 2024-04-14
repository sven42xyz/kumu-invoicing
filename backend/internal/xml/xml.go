package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

// StandardBusinessDocument represents the root element of the XML
type StandardBusinessDocument struct {
    XMLName xml.Name `xml:"StandardBusinessDocument"`
    Header  StandardBusinessDocumentHeader
    Invoice Invoice `xml:"http://www.oasis-open.org/ubl:Invoice"`
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
    Identifier string `xml:"Identifier"`
}

// Receiver represents the Receiver element
type Receiver struct {
    Identifier string `xml:"Identifier"`
}

// DocumentIdentification represents the DocumentIdentification element
type DocumentIdentification struct {
    Standard             string `xml:"Standard"`
    TypeVersion          string `xml:"TypeVersion"`
    InstanceIdentifier   string `xml:"InstanceIdentifier"`
    Type                 string `xml:"Type"`
    CreationDateAndTime  string `xml:"CreationDateAndTime"`
}

// BusinessScope represents the BusinessScope element
type BusinessScope struct {
    Scope []Scope `xml:"Scope"`
}

// Scope represents the Scope element
type Scope struct {
    Type              string `xml:"Type"`
    InstanceIdentifier string `xml:"InstanceIdentifier"`
    Identifier        string `xml:"Identifier"`
}

// Invoice represents the Invoice element
type Invoice struct {
    CustomizationID         string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}CustomizationID"`
    ProfileID               string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ProfileID"`
    ID                      string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ID"`
    IssueDate               string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}IssueDate"`
    DueDate                 string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}DueDate"`
    InvoiceTypeCode         string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}InvoiceTypeCode"`
    DocumentCurrencyCode    string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}DocumentCurrencyCode"`
    OrderReference          OrderReference
    AccountingSupplierParty AccountingParty
    AccountingCustomerParty AccountingParty
    Delivery                Delivery
    PaymentMeans            PaymentMeans
    AllowanceCharge         []AllowanceCharge
    TaxTotal                TaxTotal
    LegalMonetaryTotal      LegalMonetaryTotal
    InvoiceLine             InvoiceLine
}

// OrderReference represents the OrderReference element
type OrderReference struct {
    ID string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ID"`
}

// AccountingParty represents the AccountingSupplierParty and AccountingCustomerParty elements
type AccountingParty struct {
    Party Party
}

// Party represents the Party element
type Party struct {
    EndpointID         string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}EndpointID,attr"`
    PartyName          PartyName
    PostalAddress      PostalAddress
    PartyTaxScheme     PartyTaxScheme
    PartyLegalEntity   PartyLegalEntity
    Contact            Contact
}

// PartyName represents the PartyName element
type PartyName struct {
    Name string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}Name"`
}

// PostalAddress represents the PostalAddress element
type PostalAddress struct {
    StreetName  string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}StreetName"`
    CityName    string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}CityName"`
    PostalZone  string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}PostalZone"`
    Country     Country
}

// Country represents the Country element
type Country struct {
    IdentificationCode string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}IdentificationCode"`
}

// PartyTaxScheme represents the PartyTaxScheme element
type PartyTaxScheme struct {
    CompanyID   string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}CompanyID"`
    TaxScheme   TaxScheme
}

// TaxScheme represents the TaxScheme element
type TaxScheme struct {
    ID string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ID"`
}

// PartyLegalEntity represents the PartyLegalEntity element
type PartyLegalEntity struct {
    RegistrationName string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}RegistrationName"`
}

// Contact represents the Contact element
type Contact struct {
    Name           string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}Name"`
    Telephone      string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}Telephone"`
    ElectronicMail string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ElectronicMail"`
}

// Delivery represents the Delivery element
type Delivery struct {
    ActualDeliveryDate string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ActualDeliveryDate"`
}

// PaymentMeans represents the PaymentMeans element
type PaymentMeans struct {
    PaymentMeansCode   string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}PaymentMeansCode"`
    PaymentID          string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}PaymentID"`
    PayeeFinancialAccount PayeeFinancialAccount
}

// PayeeFinancialAccount represents the PayeeFinancialAccount element
type PayeeFinancialAccount struct {
    ID string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ID"`
}

// AllowanceCharge represents the AllowanceCharge element
type AllowanceCharge struct {
    ChargeIndicator      string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ChargeIndicator"`
    AllowanceChargeReason string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}AllowanceChargeReason"`
    Amount               string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}Amount"`
    TaxCategory          TaxCategory
}

// TaxCategory represents the TaxCategory element
type TaxCategory struct {
    ID      string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ID"`
    Percent string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}Percent"`
    TaxScheme TaxScheme
}

// TaxTotal represents the TaxTotal element
type TaxTotal struct {
    TaxAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}TaxAmount"`
    TaxSubtotal TaxSubtotal
}

// TaxSubtotal represents the TaxSubtotal element
type TaxSubtotal struct {
    TaxableAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}TaxableAmount"`
    TaxAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}TaxAmount"`
    TaxCategory TaxCategory
}

// LegalMonetaryTotal represents the LegalMonetaryTotal element
type LegalMonetaryTotal struct {
    LineExtensionAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}LineExtensionAmount"`
    TaxExclusiveAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}TaxExclusiveAmount"`
    TaxInclusiveAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}TaxInclusiveAmount"`
    ChargeTotalAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ChargeTotalAmount"`
    PayableAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}PayableAmount"`
}

// InvoiceLine represents the InvoiceLine element
type InvoiceLine struct {
    ID string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ID"`
    InvoicedQuantity string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}InvoicedQuantity"`
    LineExtensionAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}LineExtensionAmount"`
    Item Item
    Price Price
}

// Item represents the Item element
type Item struct {
    Name string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}Name"`
    SellersItemIdentification SellersItemIdentification
    ClassifiedTaxCategory ClassifiedTaxCategory
}

// SellersItemIdentification represents the SellersItemIdentification element
type SellersItemIdentification struct {
    ID string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ID"`
}

// ClassifiedTaxCategory represents the ClassifiedTaxCategory element
type ClassifiedTaxCategory struct {
    ID string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}ID"`
    Percent string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}Percent"`
    TaxScheme TaxScheme
}

// Price represents the Price element
type Price struct {
    PriceAmount string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}PriceAmount"`
    BaseQuantity string `xml:"{urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2}BaseQuantity"`
}

func NewXMLFile() error {
	// Create a new instance of the XML struct with sample data
	xmlData := StandardBusinessDocument{
		Header: StandardBusinessDocumentHeader{
			HeaderVersion: "1.0",
			Sender: Sender{
				Identifier: "9938:LU21086120",
			},
			Receiver: Receiver{
				Identifier: "9938:LU28959578",
			},
			DocumentIdentification: DocumentIdentification{
				Standard:             "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2",
				TypeVersion:          "2.1",
				InstanceIdentifier:   "6c4f9576-26fd-4674-b787-ef0b4b06c150",
				Type:                 "Invoice",
				CreationDateAndTime:  "2024-04-14T10:35:51.173Z",
			},
			BusinessScope: BusinessScope{
				Scope: []Scope{
					{
						Type:              "DOCUMENTID",
						InstanceIdentifier: "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2::Invoice##urn:cen.eu:en16931:2017#compliant#urn:fdc:peppol.eu:2017:poacc:billing:3.0::2.1",
						Identifier:        "busdox-docid-qns",
					},
					{
						Type:              "PROCESSID",
						InstanceIdentifier: "urn:fdc:peppol.eu:2017:poacc:billing:01:1.0",
						Identifier:        "cenbii-procid-ubl",
					},
				},
			},
		},
		Invoice: Invoice{
			CustomizationID:      "urn:cen.eu:en16931:2017#compliant#urn:fdc:peppol.eu:2017:poacc:billing:3.0",
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