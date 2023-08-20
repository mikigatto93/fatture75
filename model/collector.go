package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func cell(row int, col string) string {
	return fmt.Sprintf("%s%d", col, row)
}

func getCellInt(file *excelize.File, sheet string, cellCoords string) (int, error) {
	val, err := file.GetCellValue(sheet, cellCoords)
	if err != nil {
		return 0, err
	}

	//remove the , that gets added at the thousands ex: 1,000
	formattedVal := strings.ReplaceAll(val, ",", "")

	intValue, err := strconv.Atoi(formattedVal)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func getCellFloat(file *excelize.File, sheet string, cellCoords string) (float32, error) {
	val, err := file.GetCellValue(sheet, cellCoords)
	if err != nil {
		return 0, err
	}
	//fmt.Println(val)
	floatValue, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return 0, err
	}

	return float32(floatValue), nil
}

type DataCollector struct {
	filePath      string
	file          *excelize.File
	Costumer      CostumerData
	Products      []Fixture
	OtherExpenses []OtherExpenses
}

func NewDataCollector(filePath string) *DataCollector {
	col := DataCollector{
		filePath:      filePath,
		Products:      make([]Fixture, 0),
		OtherExpenses: make([]OtherExpenses, 0),
	}
	return &col
}

func (c *DataCollector) LoadData() error {
	file, err := excelize.OpenFile(c.filePath)
	if err != nil {
		return err
	}

	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	c.file = file

	c.loadCostumer()
	c.loadProducts()
	c.loadOtherExpenses()
	return nil

}

func (c *DataCollector) loadCostumer() {
	name, err1 := c.file.GetCellValue(InizioSheet, ConstumerNameCell)
	address, err2 := c.file.GetCellValue(InizioSheet, ConstumerAddressCell)
	municipality, err3 := c.file.GetCellValue(InizioSheet, ConstumerMunicipalityCell)

	errs := [3]error{err1, err2, err3}
	for i := 0; i < 3; i++ {
		if errs[i] != nil {
			fmt.Println(fmt.Errorf("error in loading constumer data %d: %v", i, errs[i]))
			return
		}
	}

	c.Costumer = NewCostumer(name, address, municipality)

}

func (c *DataCollector) loadProducts() {

	for i := MinFixtureRow; i <= MaxFixtureRow; i++ {

		fixtureGroup := c.getRowFixtureGroup(i)

		if fixtureGroup != "" {
			headers := FixtureHeadersMap[fixtureGroup]

			prod, err := c.buildFixture(i, headers)

			if err != nil {
				fmt.Println(fmt.Errorf("error in loading the product line %d: %v", i, err))
			} else {
				c.Products = append(c.Products, prod)
			}
		}

	}

}

func (c *DataCollector) loadOtherExpenses() {
	c.loadComplementaryWorks("Opere complementari")
	c.loadOptionalServices("Servizi eventuali")
}

func (c *DataCollector) loadComplementaryWorks(expenseType string) {
	for i := MinComplementaryWorksRow; i <= MaxComplementaryWorksRow; i++ {
		p, err1 := getCellFloat(c.file,
			Check1Sheet, cell(i, OtherExpenseHeaders.PriceCol))
		if err1 != nil {
			fmt.Println(fmt.Errorf("error in loading the expense line %d: %v", i, err1))
		} else {
			if p > 0 {
				desc, err2 := c.file.GetCellValue(
					Check1Sheet, cell(i, OtherExpenseHeaders.DescriptionCol))

				if err2 != nil {
					fmt.Println(fmt.Errorf("error in loading the expense line %d: %v", i, err2))
				} else {

					ex := OtherExpenses{
						Price: p,

						Description: ApplyRules(desc, i, OtherExpenseHeaders.DescriptionCol, Check1Sheet),

						Type: expenseType,
					}

					c.OtherExpenses = append(c.OtherExpenses, ex)
				}
			}
		}
	}
}

func (c *DataCollector) loadOptionalServices(expenseType string) {
	for i := MinOptionalServicesRow; i <= MaxOptionalServicesRow; i++ {
		p, err1 := getCellFloat(c.file,
			Check1Sheet, cell(i, OtherExpenseHeaders.PriceCol))
		if err1 != nil {
			fmt.Println(fmt.Errorf("error in loading the expense line %d: %v", i, err1))
		} else {
			if p > 0 {
				desc, err2 := c.file.GetCellValue(
					Check1Sheet, cell(i, OtherExpenseHeaders.DescriptionCol))

				if err2 != nil {
					fmt.Println(fmt.Errorf("error in loading the expense line %d: %v", i, err2))
				} else {
					ex := OtherExpenses{
						Price: p,

						Description: ApplyRules(desc, i, OtherExpenseHeaders.DescriptionCol, Check1Sheet),

						Type: expenseType,
					}

					c.OtherExpenses = append(c.OtherExpenses, ex)
				}
			}
		}
	}
}

func (c *DataCollector) getRowFixtureGroup(rowIndex int) FixtureGroup {

	for group, headers := range FixtureHeadersMap {

		value, err := c.file.GetCellValue(
			SerramentiSheet, cell(rowIndex, headers.QuantityCol))

		if err == nil && value != "" {
			return group
		}
	}
	return ""
}

func (c *DataCollector) buildFixture(rowIndex int, headers FixtureHeaders) (Fixture, error) {

	h, err1 := getCellInt(c.file,
		SerramentiSheet, cell(rowIndex, headers.HeightCol))

	w, err2 := getCellInt(c.file,
		SerramentiSheet, cell(rowIndex, headers.WidthCol))

	q, err3 := getCellInt(c.file,
		SerramentiSheet, cell(rowIndex, headers.QuantityCol))

	d, err4 := c.file.GetCellValue(
		SerramentiSheet, cell(rowIndex, headers.DescriptionCol))

	t, err5 := c.file.GetCellValue(
		SerramentiSheet, cell(rowIndex, headers.TypeCol))

	p, err6 := getCellFloat(c.file,
		SerramentiSheet, cell(rowIndex, headers.PriceCol))

	// TODO: ADD OPTIONS SUPPORT

	errs := [6]error{err1, err2, err3, err4, err5, err6}
	for i := 0; i < 6; i++ {
		if errs[i] != nil {
			return Fixture{}, errs[i]
		}
	}

	prod := Fixture{
		Height:   h,
		Width:    w,
		Quantity: q,
		Price:    p,

		Description: ApplyRules(d, rowIndex, headers.DescriptionCol, SerramentiSheet),

		Type: ApplyRules(t, rowIndex, headers.TypeCol, SerramentiSheet),

		Options: nil,
	}

	return prod, nil

}
