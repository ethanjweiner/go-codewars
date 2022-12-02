package createphonenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePhoneNumber(t *testing.T) {
	type input struct {
		numbers [10]uint
	}

	type testCase struct {
		name     string
		args     input
		expected string
	}

	tests := []testCase{
		{
			name: "generic case",
			args: input{
				numbers: [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			expected: "(123) 456-7890",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := CreatePhoneNumber(test.args.numbers)
			assert.Equal(t, actual, test.expected)
		})
	}
}
