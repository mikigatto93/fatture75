package model

import (
	"encoding/xml"
	"os"
	"strings"
)

type rawXmlDescriptionItem struct {
	Value string `xml:"Value,attr"`
	Name  string `xml:"Name,attr"`
}

type rawXmlProductData struct {
	ProductId           string
	UnitListPrice       float32
	QuantityListPrice   float32
	Quantity            int
	DescriptionExtended []rawXmlDescriptionItem `xml:"DescriptionExtended>DescriptionItem"`
}

type rawXmlDocumentData struct {
	XMLName xml.Name            `xml:"Order"`
	Rows    []rawXmlProductData `xml:"Rows>Row"`
}

type XmlProductData struct {
	ProductId string
	Height    string
	Width     string
	Quantity  int
	Price     float32
}

type XmlCollector struct {
	filePath    string
	ProductData []XmlProductData
}

func NewXmlCollector(filePath string) *XmlCollector {
	return &XmlCollector{
		filePath:    filePath,
		ProductData: make([]XmlProductData, 0),
	}
}

func (c *XmlCollector) LoadData() error {
	content, err := os.ReadFile(c.filePath)
	if err != nil {
		return err
	}

	xmlData := rawXmlDocumentData{}

	err = xml.Unmarshal(content, &xmlData)

	if err != nil {
		return err
	}

	c.ProductData = c.parseData(xmlData)
	return nil
}

func (c *XmlCollector) parseData(data rawXmlDocumentData) []XmlProductData {

	prodList := make([]XmlProductData, 0)

	for _, p := range data.Rows {

		var w, h string

		for _, val := range p.DescriptionExtended {
			if val.Name == "Altezza" {
				h = strings.Replace(val.Value, "mm", "", 1)
			} else if val.Name == "Larghezza" {
				w = strings.Replace(val.Value, "mm", "", 1)
			}
		}

		prodData := XmlProductData{
			ProductId: p.ProductId,
			Height:    h,
			Width:     w,
			Price:     p.QuantityListPrice, // total price
			Quantity:  p.Quantity,
		}

		prodList = append(prodList, prodData)
	}

	return prodList
}
