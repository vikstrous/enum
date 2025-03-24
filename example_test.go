package enum_test

import (
	"fmt"

	"github.com/orsinium-labs/enum"
)

func ExampleNew() {
	type Color struct {
		enum.Member
	}

	var (
		b      = enum.NewBuilder[Color]()
		_      = b.Add("red")
		_      = b.Add("green")
		_      = b.Add("blue")
		Colors = b.Enum()
	)

	fmt.Printf("%#v\n", Colors)
	// Output: enum.New(enum_test.Color{"red"}, enum_test.Color{"green"}, enum_test.Color{"blue"})
}

func ExampleMember_String() {
	type Color struct {
		enum.Member
	}

	var (
		b     = enum.NewBuilder[Color]()
		_     = b.Add("red")
		Green = b.Add("green")
		_     = b.Add("blue")
		_     = b.Enum()
	)

	fmt.Println(Green.String())
	// Output: green
}

func ExampleEnum_String() {
	type Color struct {
		enum.Member
	}

	var (
		b      = enum.NewBuilder[Color]()
		_      = b.Add("red")
		_      = b.Add("green")
		_      = b.Add("blue")
		Colors = b.Enum()
	)

	fmt.Println(Colors)
	// Output: red, green, blue
}

func ExampleEnum_GoString() {
	type Color struct {
		enum.Member
	}

	var (
		b      = enum.NewBuilder[Color]()
		_      = b.Add("red")
		_      = b.Add("green")
		_      = b.Add("blue")
		Colors = b.Enum()
	)

	fmt.Printf("%#v\n", Colors)
	// Output: enum.New(enum_test.Color{"red"}, enum_test.Color{"green"}, enum_test.Color{"blue"})
}

func ExampleEnum_Parse() {
	type Color struct {
		enum.Member
	}

	var (
		b      = enum.NewBuilder[Color]()
		_      = b.Add("red")
		_      = b.Add("green")
		_      = b.Add("blue")
		Colors = b.Enum()
	)

	parsed, err := Colors.Parse("red")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", parsed)
	// Output: enum_test.Color{Member:enum.Member{index:0, value:"red"}}
}

func ExampleMember_Index() {
	type Color struct {
		enum.Member
	}

	var (
		b     = enum.NewBuilder[Color]()
		_     = b.Add("red")
		Green = b.Add("green")
		_     = b.Add("blue")
		_     = b.Enum()
	)

	index := Green.Index()
	fmt.Println(index)
	// Output: 1
}

func ExampleEnum_Members() {
	type Color struct {
		enum.Member
	}

	var (
		b      = enum.NewBuilder[Color]()
		_      = b.Add("red")
		_      = b.Add("green")
		_      = b.Add("blue")
		Colors = b.Enum()
	)

	members := Colors.Members()
	fmt.Println(members)
	// Output: [red green blue]
}

func ExampleEnum_Values() {

	type Color struct {
		enum.Member
	}
	var (
		b      = enum.NewBuilder[Color]()
		_      = b.Add("red")
		_      = b.Add("green")
		_      = b.Add("blue")
		Colors = b.Enum()
	)

	values := Colors.Values()
	fmt.Println(values)
	// Output: [red green blue]
}
