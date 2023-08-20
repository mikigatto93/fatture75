package model

import fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"

/*
type DocumentBuilder struct {
	collector       *DataCollector

}

func NewDocumentBuilder(collector *DataCollector) *DocumentBuilder {
	return &DocumentBuilder{
		collector: collector,
	}
}
*/

type Document struct {
	Client         *fattureincloud.Entity
	collector      *DataCollector
	IssuedDocument fattureincloud.IssuedDocument
}

func NewDocument(docType fattureincloud.IssuedDocumentType, collector *DataCollector) *Document {

	doc := Document{
		collector:      collector,
		IssuedDocument: *fattureincloud.NewIssuedDocument(),
	}

	client := *fattureincloud.NewEntity().
		SetName(collector.Costumer.Name).
		SetAddressStreet(collector.Costumer.Address).
		SetAddressCity(collector.Costumer.Municipality).
		SetAddressProvince(collector.Costumer.Discrict).
		SetCountry("Italia")

	doc.IssuedDocument.
		SetEntity(client).
		SetType(docType).
		SetCurrency(*fattureincloud.NewCurrency().SetId("EUR")).
		SetLanguage(*fattureincloud.NewLanguage().SetCode("it").SetName("italiano"))

	return &doc
}

func (d *Document) FillItems() {
	itemsList := []fattureincloud.IssuedDocumentItemsListItem{}

	//fills fixture
	for _, fixture := range d.collector.Products {
		newItem := *fattureincloud.NewIssuedDocumentItemsListItem().
			SetName(fixture.Type).
			SetDescription(fixture.GetExtensiveDescription()).
			SetNetPrice(fixture.Price / float32(fixture.Quantity)).
			SetDiscount(0).
			SetQty(float32(fixture.Quantity)).
			SetVat(*fattureincloud.NewVatType().SetId(3)) // 10%
		itemsList = append(itemsList, newItem)
	}

	//fills expenses
	for _, expense := range d.collector.OtherExpenses {
		newItem := *fattureincloud.NewIssuedDocumentItemsListItem().
			SetName(expense.Type).
			SetDescription(expense.Description).
			SetNetPrice(expense.Price).
			SetDiscount(0).
			SetQty(1).
			SetVat(*fattureincloud.NewVatType().SetId(3)) // 10%

		itemsList = append(itemsList, newItem)
	}

	d.IssuedDocument.SetItemsList(itemsList)
}
