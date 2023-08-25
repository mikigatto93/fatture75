package model

/*type FixtureLimits struct {
	MinRow string
	MaxRow string
}

var FixtureLimitsMap map[FixtureGroup]FixtureLimits = map[FixtureGroup]FixtureLimits{
	GroupA: FixtureLimits{
		MinRow: "7",
		MaxRow: "56",
	},
	GroupB: FixtureLimits{
		MinRow: "7",
		MaxRow: "56",
	},
	GroupC: FixtureLimits{
		MinRow: "7",
		MaxRow: "56",
	},
	GroupD: FixtureLimits{
		MinRow: "7",
		MaxRow: "56",
	},
}*/

const (
	MinFixtureRow int = 7
	MaxFixtureRow int = 56

	MinComplementaryWorksRow int = 22
	MaxComplementaryWorksRow int = 35

	MinOptionalServicesRow int = 39
	MaxOptionalServicesRow int = 43

	ProfessionalExpensesRow int = 35
)
