package arraydiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_array_diff(t *testing.T) {
	type input[T comparable] struct {
		a []T
		b []T
	}

	type testCase[T comparable] struct {
		name     string
		args     input[T]
		expected []T
	}

	// Todo: Accept test cases for any comparable type
	tests := []testCase[int]{
		{
			name: "arrays are empty",
			args: input[int]{
				a: []int{},
				b: []int{},
			},
			expected: []int{},
		},
		{
			name: "remove one value",
			args: input[int]{
				a: []int{1, 2},
				b: []int{1},
			},
			expected: []int{2},
		},
		{
			name: "value to remove appears multiple times",
			args: input[int]{
				a: []int{1, 2, 2, 2, 3},
				b: []int{2},
			},
			expected: []int{1, 3},
		},
		{
			name: "entire list is subtracted",
			args: input[int]{
				a: []int{1, 2, 3},
				b: []int{3, 2, 1},
			},
			expected: []int{},
		},
		// {
		// 	name: "using strings",
		// 	args: input[string]{
		// 		a: []string{"1", "2", "2", "2", "3"},
		// 		b: []string{"2"},
		// 	},
		// 	expected: []string{"1", "3"},
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ArrayDiff(test.args.a, test.args.b)
			assert.Equal(t, test.expected, actual)
		})
	}
}
