package enum

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Member is an enum member, a specific value bound to a variable.
type Member[VA Validator] struct {
	value string
}

type Validator interface {
	Validate(string) bool
}

func (m Member[VA]) String() string {
	if m.value == "" {
		panic("uninitialized enum value")
	}
	return m.value
}

func (m Member[VA]) MarshalJSON() ([]byte, error) {
	if m.value == "" {
		return nil, errors.New("uninitialized enum value")
	}
	return json.Marshal(m.value)
}

func (m *Member[VA]) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}
	var validator VA
	valid := validator.Validate(value)
	if !valid {
		return fmt.Errorf("invalid enum value %s", value)
	}
	m.value = value
	return nil
}

// Enum is a collection of enum members.
//
// Use [New] to construct a new Enum from a list of members.
type Enum[VA Validator] struct {
	memberStrings []string
	members       []Member[VA]
}

// Parse converts a raw value into a member of the enum.
//
// If none of the enum members has the given value, nil is returned.
func (e Enum[VA]) Parse(value string) (Member[VA], error) {
	for i, m := range e.memberStrings {
		if m == value {
			return e.members[i], nil
		}
	}

	return Member[VA]{}, fmt.Errorf("enum value not found: %s", value)
}

func (e Enum[VA]) Validate(value string) bool {
	for _, m := range e.memberStrings {
		if m == value {
			return true
		}
	}

	return false
}

// Members returns a slice of the members in the enum.
func (e Enum[VA]) Members() []Member[VA] {
	return e.members
}

// String implements [fmt.Stringer] interface.
//
// It returns a comma-separated list of values of the enum members.
func (e Enum[VA]) String() string {
	return strings.Join(e.memberStrings, ", ")
}

// GoString implements [fmt.GoStringer] interface.
//
// When you print a member using "%#v" format,
// it will show the enum representation as a valid Go syntax.
func (e Enum[VA]) GoString() string {
	values := make([]string, 0, len(e.memberStrings))
	for i, m := range e.memberStrings {
		values = append(values, fmt.Sprintf("%T{%#v}", e.members[i], m))
	}
	joined := strings.Join(values, ", ")
	return fmt.Sprintf("enum{%s}", joined)
}

// Builder is a constructor for an [Enum].
//
// Use [Builder.Add] to add new members to the future enum
// and then call [Builder.Enum] to create a new [Enum] with all added members.
//
// Builder is useful for when you have lots of enum members, and new ones
// are added over time, as the project grows. In such scenario, it's easy to forget
// to add in the [Enum] a newly created [Member].
// The builder is designed to prevent that.
type Builder[VA Validator] struct {
	enum     Enum[VA]
	finished bool
}

// NewBuilder creates a new [Builder], a constructor for an [Enum].
func NewBuilder[VA Validator]() Builder[VA] {
	return Builder[VA]{}
}

// Add registers a new [Member] in the builder.
func (b *Builder[VA]) Add(v string) Member[VA] {
	if b.finished {
		panic("no more members can be added at run time")
	}
	b.enum.memberStrings = append(b.enum.memberStrings, v)
	m := Member[VA]{v}
	b.enum.members = append(b.enum.members, m)
	return m
}

// Enum creates a new [Enum] with all members registered using [Builder.Add].
func (b *Builder[VA]) Enum() Enum[VA] {
	if b.finished {
		panic("build used to create multiple enums")
	}
	b.finished = true
	return b.enum
}
