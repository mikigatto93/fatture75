package model

import "fmt"

type FixtureGroup string

const (
	GroupA FixtureGroup = "A"
	GroupB FixtureGroup = "B"
	GroupC FixtureGroup = "C"
	GroupD FixtureGroup = "D"
)

type option struct {
	Family string
	Value  string
	Option string
}

type Fixture struct {
	Height      int
	Width       int
	Quantity    int
	Description string
	Type        string
	Price       float32
	Options     []option
}

func (f Fixture) GetExtensiveDescription() string {
	var desc string
	if f.Description == "" {
		desc = f.Description + ", "
	} else {
		desc = ""
	}

	return fmt.Sprintf("%s dimensioni: %dx%d", desc, f.Width, f.Height)
}
