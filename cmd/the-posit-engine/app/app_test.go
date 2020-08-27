package app

import "testing"

func Test_Greetings(t *testing.T) {
	testCases := []struct {
		expected string
	}{
		{"The Posit!"},
	}
	for _, testCase := range testCases {
		if testCase.expected != Greetings {
			t.Errorf("Expected %v, Got %v", testCase.expected, Greetings)
		}
	}
}
