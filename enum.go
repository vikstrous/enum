package enum

import (
	"fmt"
	"strings"
)

// Member is an enum member, a specific value bound to a variable.
type Member struct {
	index int
	value string
}

func (m Member) String() string {
	if m.value == "" {
		panic("uninitialized enum value")
	}
	return m.value
}

func (m Member) Index() int {
	if m.value == "" {
		panic("uninitialized enum value")
	}
	return m.index
}

// iMember is the type constraint for Member used by Enum.
//
// We can't use Member directly in type constraints
// because the users create a new subtype from Member
// instead of using it directly.
//
// We also can't use a normal interface because new types
// don't inherit methods of their base type.
type isMemberWrapper interface {
	~struct{ Member }
}

// Enum is a collection of enum members.
//
// Use [New] to construct a new Enum from a list of members.
type Enum[M isMemberWrapper] struct {
	members []string
}

// Empty returns true if the enum doesn't have any members.
func (e Enum[M]) Empty() bool {
	return len(e.members) == 0
}

// Len returns how many members the enum has.
func (e Enum[M]) Len() int {
	return len(e.members)
}

// Contains returns true if the enum has the given member.
func (e Enum[M]) Contains(member M) bool {
	type withMember struct{ Member }
	for _, m := range e.members {
		if m == withMember(member).Member.value {
			return true
		}
	}
	return false
}

// Parse converts a raw value into a member of the enum.
//
// If none of the enum members has the given value, nil is returned.
func (e Enum[M]) Parse(value string) (M, error) {
	for i, m := range e.members {
		if m == value {
			return M{Member: Member{i, m}}, nil
		}
	}

	return M{}, fmt.Errorf("enum value not found: %s", value)
}

// Value returns the wrapped value of the given enum member.
func (e Enum[M]) stringValue(member M) string {
	type memberWrapper struct {
		Member
	}
	return memberWrapper(member).Member.value
}

// Members returns a slice of the members in the enum.
func (e Enum[M]) Members() []M {
	members := make([]M, len(e.members))
	for i, mStr := range e.members {
		members[i] = M{Member: Member{i, mStr}}
	}
	return members
}

// Values returns a slice of values of all members of the enum.
func (e Enum[M]) Values() []string {
	return e.members
}

// String implements [fmt.Stringer] interface.
//
// It returns a comma-separated list of values of the enum members.
func (e Enum[M]) String() string {
	return strings.Join(e.members, ", ")
}

// GoString implements [fmt.GoStringer] interface.
//
// When you print a member using "%#v" format,
// it will show the enum representation as a valid Go syntax.
func (e Enum[M]) GoString() string {
	values := make([]string, 0, len(e.members))
	for i, m := range e.members {
		values = append(values, fmt.Sprintf("%T{%#v}", M{Member: Member{i, m}}, m))
	}
	joined := strings.Join(values, ", ")
	return fmt.Sprintf("enum.New(%s)", joined)
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
type Builder[M isMemberWrapper] struct {
	members  []string
	finished bool
}

// NewBuilder creates a new [Builder], a constructor for an [Enum].
func NewBuilder[M isMemberWrapper]() Builder[M] {
	return Builder[M]{make([]string, 0), false}
}

// Add registers a new [Member] in the builder.
func (b *Builder[M]) Add(v string) M {
	if b.finished {
		panic("no more members can be added at run time")
	}
	i := len(b.members)
	b.members = append(b.members, v)
	return M{Member: Member{i, v}}
}

// Enum creates a new [Enum] with all members registered using [Builder.Add].
func (b *Builder[M]) Enum() Enum[M] {
	if b.finished {
		panic("build used to create multiple enums")
	}
	b.finished = true
	return Enum[M]{b.members}
}
