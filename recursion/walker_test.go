package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name     string
	Address  Address
	Contacts []Contact
}

type Address struct {
	Street string
	Number int
}

type Contact struct {
	Friend Person
}

func TestWalk(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected []string
	}{
		{
			"single string field",
			struct {
				Name string
			}{"Bob"},
			[]string{"Bob"},
		},
		{
			"two string fields",
			struct {
				Firstname string
				Lastname  string
			}{"Bob", "Fischer"},
			[]string{"Bob", "Fischer"},
		},
		{
			"string and int field",
			struct {
				Name string
				Age  int
			}{"Bob", 12},
			[]string{"Bob"},
		},
		{
			"nested fields",
			Person{"Bob", Address{"Foobar", 14}, []Contact{}},
			[]string{"Bob", "Foobar"},
		},
		{
			"pointer to struct",
			&Person{"Bob", Address{"Foobar", 14}, []Contact{}},
			[]string{"Bob", "Foobar"},
		},
		{
			"slices",
			[]Address{{"Hogwarts", 11}, {"Jurassic World", 0}},
			[]string{"Hogwarts", "Jurassic World"},
		},
		{
			"nested slices",
			Person{"Bob", Address{"Foobar", 14}, []Contact{{Friend: Person{"Bobby", Address{"Crazy", 1337}, []Contact{}}}}},
			[]string{"Bob", "Foobar", "Bobby", "Crazy"},
		},
		{
			"arrays",
			[]string{"boring", "array"},
			[]string{"boring", "array"},
		},
		{
			"maps",
			map[string]string{
				"Foo": "Bar",
				"Bob": "Fischer",
			},
			[]string{"Bar", "Fischer"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertStringFields(t, test.input, test.expected...)
		})
	}
}

func TestWalkWithSpecialCases(t *testing.T) {
	t.Run("with channel", func(t *testing.T) {
		channel := make(chan Address)
		go func() {
			channel <- Address{"Munich", 1337}
			channel <- Address{"New York", 1}
			close(channel)
		}()
		assertStringFields(t, channel, "Munich", "New York")
	})
	t.Run("with function", func(t *testing.T) {
		fn := func() (Address, Address) {
			return Address{"Munich", 1337}, Address{"New York", 1}
		}
		assertStringFields(t, fn, "Munich", "New York")
	})
}

func assertStringFields(t *testing.T, object any, expected ...string) {
	t.Helper()
	var fields []string
	Walk(object, func(field string) {
		fields = append(fields, field)
	})
	assert.EqualValues(t, fields, expected)
}
