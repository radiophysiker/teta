
package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "aaaabccddddde", expected: "a4b1c2d5e1"},
		{input: "abccd", expected: "a1b1c2d1"},
		{input: "abcd", expected: "abcd"},
		{input: "baa", expected: "a2b1"},
		{input: "cccccaabbb", expected: "a2b3c5"},
		{input: "aaabbbcccccaaaaa", expected: "a8b3c5"},
		{input: "", expected: ""},
		{input: "zzzzcccUUUuu", expected: "U3c3u2z4"},
		{input: "ЯЯЯБББддд", expected: "Б3Я3д3"},
		{input: "ЯЯЯБББдддaa", expected: "a2Б3Я3д3"},
		{input: "", expected: ""},
		{input: "m11", expected: "12m1"},
		{input: "zzzzcccUUUuuЯЯЯБББддд", expected: "U3c3u2z4Б3Я3д3"},
		{input: "aaabbbccccc", expected: "a3b3c5"},
		{input: "acb", expected: "acb"},
		{input: "a", expected: "a"},
		
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := Pack(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
