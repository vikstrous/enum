package enum_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/enum"
)

type Color = enum.Member[ColorValidator]

type ColorValidator struct{}

func (ColorValidator) Validate(value string) bool {
	return Colors.Validate(value)
}

var (
	bColors = enum.NewBuilder[ColorValidator]()
	Red     = bColors.Add("red")
	Green   = bColors.Add("green")
	Blue    = bColors.Add("blue")
	Colors  = bColors.Enum()
)

func TestMember_String(t *testing.T) {
	is := is.New(t)
	is.Equal(Red.String(), "red")
	is.Equal(Green.String(), "green")
	is.Equal(Blue.String(), "blue")
}

func TestMember_GoString(t *testing.T) {
	is := is.New(t)
	is.Equal(fmt.Sprintf("%#v", Colors), `enum{enum.Member[github.com/orsinium-labs/enum_test.ColorValidator]{"red"}, enum.Member[github.com/orsinium-labs/enum_test.ColorValidator]{"green"}, enum.Member[github.com/orsinium-labs/enum_test.ColorValidator]{"blue"}}`)
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

func TestEnum_JSON(t *testing.T) {
	is := is.New(t)
	_, err := json.Marshal(Color{})
	if err == nil {
		is.Fail()
	}
	redBytes, err := json.Marshal(Red)
	is.NoErr(err)
	is.Equal([]byte(`"red"`), redBytes)
	var newColor Color
	err = json.Unmarshal(redBytes, &newColor)
	is.NoErr(err)
	is.Equal(Red, newColor)
	err = json.Unmarshal([]byte(`"rainbow"`), &newColor)
	if err == nil {
		is.Fail()
	}
	err = json.Unmarshal([]byte(`{}`), &newColor)
	if err == nil {
		is.Fail()
	}
}

func TestEnum_String_Panic(t *testing.T) {
	is := is.New(t)
	defer func() {
		r := recover()
		is.Equal(r, "uninitialized enum value")
	}()
	_ = Color{}.String()
}

type Country = enum.Member[CountryValidator]

var (
	bC        = enum.NewBuilder[CountryValidator]()
	NL        = bC.Add("Netherlands")
	FR        = bC.Add("France")
	BE        = bC.Add("Belgium")
	Countries = bC.Enum()
)

type CountryValidator struct{}

func (CountryValidator) Validate(value string) bool {
	return Countries.Validate(value)
}
func TestBuilder(t *testing.T) {
	is := is.New(t)
	is.Equal(Countries.Members(), []Country{NL, FR, BE})
	func() {
		defer func() {
			r := recover()
			is.Equal(r, "no more members can be added at run time")
		}()
		_ = bColors.Add("Other")
	}()
	func() {
		defer func() {
			r := recover()
			is.Equal(r, "build used to create multiple enums")
		}()
		_ = bColors.Enum()
	}()
}
