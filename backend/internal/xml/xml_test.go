package xml

import (
    "os"
    "testing"
)

func TestNewXMLFile(t *testing.T) {
    err := NewXMLFile()
    if err != nil {
        t.Errorf("NewXMLFile returned an error: %v", err)
    }

    // Check if the XML file was created
    _, err = os.Stat("invoice.xml")
    if os.IsNotExist(err) {
        t.Error("Expected invoice.xml file to be created, but it does not exist")
    } else if err != nil {
        t.Errorf("Error checking for invoice.xml file: %v", err)
    }
}
