package controller

import "testing"

func TestStringInSlice(t *testing.T) {

	tests := []struct {
		slice    []string
		target   string
		expected bool
	}{
		{[]string{"apple", "banana", "cherry", "dates"}, "banana", true},
		{[]string{"apple", "banana", "cherry", "dates"}, "grape", false},
		{[]string{}, "apple", false},
		{[]string{"apple", "apple", "apple"}, "apple", true},
	}

	for _, test := range tests {
		result := stringInSlice(test.target, test.slice)
		if result != test.expected {
			t.Errorf("stringInSlice(%s, %v) => Expected %v, but got %v", test.target, test.slice, test.expected, result)
		}
	}
}
