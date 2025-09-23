package main

import "testing"

type normalizeTestCase struct {
	input string
	want  string
}

func TestNormalize(t *testing.T) {
	testCases := []normalizeTestCase{
		{"+7 (904) 675 60-79", "79046756079"},
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"(123)456-7892", "1234567892"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			normalized := normalize(tc.input)
			if normalized != tc.want {
				t.Errorf("got: |%v|; want: |%v|", normalized, tc.want)
			}
		})
	}
}
