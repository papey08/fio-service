package valid

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

type nameTest struct {
	description string
	name        string
	expected    bool
}

func TestName(t *testing.T) {
	nameTests := []nameTest{
		{
			description: "valid name",
			name:        "Dmitriy",
			expected:    true,
		},
		{
			description: "invalid name with non allowed symbols",
			name:        "X Æ A-12",
			expected:    false,
		},
	}

	for _, test := range nameTests {
		t.Run(test.description, func(t *testing.T) {
			assert.Equal(t, Name(test.name), test.expected)
		})
	}
}
