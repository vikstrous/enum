package enum_test

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/enum"
)

type Color struct {
	enum.Member
}

var colorBuilder = enum.NewBuilder[Color]()

func (Color) Enum() enum.Enum[Color] {
	return Colors
}

var (
	Red    = colorBuilder.Add("red")
	Green  = colorBuilder.Add("green")
	Blue   = colorBuilder.Add("blue")
	Colors = colorBuilder.Enum()
)

func TestMember_String(t *testing.T) {
	is := is.New(t)
	is.Equal(Red.String(), "red")
	is.Equal(Green.String(), "green")
	is.Equal(Blue.String(), "blue")
}

func TestMember_Print(t *testing.T) {
	is := is.New(t)
	is.Equal(fmt.Sprint(Red), "red")
}

func TestEnum_Parse(t *testing.T) {
	is := is.New(t)
	parsed, err := Colors.Parse("red")
	is.NoErr(err)
	is.Equal(parsed, Red)
	parsed, err = Colors.Parse("purple")
	if err == nil {
		is.Fail()
	}
	is.Equal(parsed, Color{})
}

func TestEnum_Members(t *testing.T) {
	is := is.New(t)
	exp := []Color{Red, Green, Blue}
	is.Equal(Colors.Members(), exp)
}

func TestEnum_Values(t *testing.T) {
	is := is.New(t)
	exp := []string{"red", "green", "blue"}
	is.Equal(Colors.Values(), exp)
}

func TestEnum_Index(t *testing.T) {
	is := is.New(t)
	is.Equal(Red.Index(), 0)
	is.Equal(Green.Index(), 1)
	is.Equal(Blue.Index(), 2)
}

func TestEnum_Index_Panic(t *testing.T) {
	is := is.New(t)
	defer func() {
		r := recover()
		is.Equal(r, "uninitialized enum value")
	}()
	Color{}.Index()
}

func TestEnum_String_Panic(t *testing.T) {
	is := is.New(t)
	defer func() {
		r := recover()
		is.Equal(r, "uninitialized enum value")
	}()
	Color{}.String()
}

func TestBuilder(t *testing.T) {
	is := is.New(t)
	type Country struct {
		enum.Member
	}
	var (
		b         = enum.NewBuilder[Country]()
		NL        = b.Add("Netherlands")
		FR        = b.Add("France")
		BE        = b.Add("Belgium")
		Countries = b.Enum()
	)
	is.Equal(Countries.Members(), []Country{NL, FR, BE})
	func() {
		defer func() {
			r := recover()
			is.Equal(r, "no more members can be added at run time")
		}()
		_ = b.Add("Other")
	}()
	func() {
		defer func() {
			r := recover()
			is.Equal(r, "build used to create multiple enums")
		}()
		_ = b.Enum()
	}()
}
