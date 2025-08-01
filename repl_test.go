package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "eat veg every day",
			expected: []string{"eat", "veg", "every", "day"},
		},
		{
			input: "  I like to  go to	the store  *",
			expected: []string{"i", "like", "to", "go", "to", "the", "store", "*"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Actual lenth %d != Expected length of %d", len(actual), len(c.expected))
			t.Fail()
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Word %s != expected word %s", word, expectedWord)
				t.Fail()
			}
		}
	}
	}
}


