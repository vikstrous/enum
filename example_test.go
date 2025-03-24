package enum_test

import (
	"fmt"

	"github.com/orsinium-labs/enum"
)

func ExampleBuilder() {
	type Color struct {
		enum.Member
	}

	var (
		b      = enum.NewBuilder[Color]()
		_      = b.Add("red")
		Green  = b.Add("green")
		_      = b.Add("blue")
		Colors = b.Enum()
	)

	fmt.Printf("Enum Members: %s\n", Colors.Members())
	fmt.Printf("Enum string: %s\n", Colors)
	fmt.Printf("Member string: %s\n", Green.String())
	parsed, err := Colors.Parse("red")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed: %s\n", parsed.String())
	// Output: Enum Members: [red green blue]
	// Enum string: red, green, blue
	// Member string: green
	// Parsed: red
}
