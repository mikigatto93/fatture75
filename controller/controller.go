package controller

import (
	"fatture75/model"
	"fmt"

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)

var (
	conf Configuration
	api  *FattureInCloudApi
)

func SetupController() {
	conf = NewConfiguration("conf.json")
	api = NewFattureInCloudApi(conf.AccessToken, conf.CompanyId)
}

func FillCalcSheet(xmlFilePath string, newFileName string) error {

	xmlCol := model.NewXmlCollector(xmlFilePath)

	err := xmlCol.LoadData()
	if err != nil {
		fmt.Println(err)
		return err
	}

	excelWriter, err := model.NewExcelWriter(conf.ModelFilePath, "conversion_map.json")

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = excelWriter.InsertProducts(xmlCol, newFileName)
	if err != nil {
		fmt.Println(err)
	}
	return err // even if it's nil
}

func GenerateNewQuote(calcFilePath string, apply75Discount bool) error {
	err := model.LoadRules("special_rules.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	collector := model.NewDataCollector(calcFilePath)
	err = collector.LoadData()
	if err != nil {
		fmt.Println(err)
		return err
	}

	doc := model.NewDocument(fattureincloud.IssuedDocumentTypes.QUOTE, collector)
	doc.FillItems()

	if apply75Discount {
		doc.ApplyDiscount(75)
	}

	_, err = api.CreateDocument(doc)

	if err != nil {
		//TODO: IMPEMENT LOGGING
		return fmt.Errorf("Error after sending the HTTP request, more info on the log file")
	} else {
		return nil
	}
}
