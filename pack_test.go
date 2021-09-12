
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
		{input: "aaaabccddddde", expected: "a4bc2d5e"},
		{input: "abccd", expected: "abc2d"},
		{input: "abcd", expected: "abcd"},
		{input: "cccccaabbb", expected: "a2b3c5"},
		{input: "aaabbbcccccaaaaa", expected: "a8b3c5"},
		{input: "", expected: ""},
		{input: "zzzzcccUUUuu", expected: "U3c3u2z4"},
		{input: "ЯЯЯБББддд", expected: "Б3Я3д3"},
		{input: "ЯЯЯБББдддaa", expected: "a2Б3Я3д3"},
		{input: "", expected: ""},
		{input: "11", expected: "12"},
		{input: "aaabbbccccc", expected: "a3b3c5"},
		
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := Pack(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
