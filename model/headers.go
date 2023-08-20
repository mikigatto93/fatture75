package model

type FixtureHeaders struct {
	WidthCol        string
	HeightCol       string
	TypeCol         string
	DescriptionCol  string
	QuantityCol     string
	PriceCol        string
	OptionFamilyRow string
	OptionNameRow   string
	OptionMinCol    string
	OptionsMaxCol   string
}

type ExpenseHeaders struct {
	DescriptionCol string
	PriceCol       string
}

var OtherExpenseHeaders ExpenseHeaders = ExpenseHeaders{
	DescriptionCol: "B",
	PriceCol:       "G",
}

var FixtureHeadersMap map[FixtureGroup]FixtureHeaders = map[FixtureGroup]FixtureHeaders{
	GroupA: {
		WidthCol:        "D",
		HeightCol:       "F",
		TypeCol:         "I",
		DescriptionCol:  "J",
		QuantityCol:     "C",
		PriceCol:        "AF",
		OptionFamilyRow: "6",
		OptionNameRow:   "5",
		OptionMinCol:    "K",
		OptionsMaxCol:   "AE",
	},

	GroupB: {
		WidthCol:        "AI",
		HeightCol:       "AK",
		TypeCol:         "AN",
		DescriptionCol:  "AO",
		QuantityCol:     "AH",
		PriceCol:        "CP",
		OptionFamilyRow: "6",
		OptionNameRow:   "5",
		OptionMinCol:    "AP",
		OptionsMaxCol:   "CP",
	},

	GroupC: {
		WidthCol:        "CS",
		HeightCol:       "CU",
		TypeCol:         "CX",
		DescriptionCol:  "CY",
		QuantityCol:     "CR",
		PriceCol:        "ED",
		OptionFamilyRow: "6",
		OptionNameRow:   "5",
		OptionMinCol:    "CZ",
		OptionsMaxCol:   "ED",
	},

	GroupD: {
		WidthCol:        "EG",
		HeightCol:       "EI",
		TypeCol:         "EL",
		DescriptionCol:  "EM",
		QuantityCol:     "EF",
		PriceCol:        "EN",
		OptionFamilyRow: "6",
		OptionNameRow:   "5",
		OptionMinCol:    "",
		OptionsMaxCol:   "",
	},
}
