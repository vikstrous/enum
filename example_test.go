package enum_test

import (
	"fmt"

	"github.com/orsinium-labs/enum"
)

type ExampleColor = enum.Member[ExampleValidator]

var (
	bExample      = enum.NewBuilder[ExampleValidator]()
	ExampleRed    = bExample.Add("red")
	ExampleGreen  = bExample.Add("green")
	_             = bExample.Add("blue")
	ExampleColors = bExample.Enum()
)

type ExampleValidator struct{}

func (ExampleValidator) Validate(value string) bool {
	return ExampleColors.Validate(value)
}

func ExampleBuilder() {

	fmt.Printf("Enum Members: %s\n", Colors.Members())
	fmt.Printf("Enum string: %s\n", Colors)
	fmt.Printf("Member string: %s\n", Green.String())
	parsed, err := Colors.Parse("red")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed: %s\n", parsed.String())
	// TODO: How do we do equality?
	fmt.Printf("Equality: %t\n", Red == parsed)

	// Output: Enum Members: [red green blue]
	// Enum string: red, green, blue
	// Member string: green
	// Parsed: red
	// Equality: true
}
