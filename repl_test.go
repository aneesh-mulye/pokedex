package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanInput(t *testing.T) {
	a := assert.New(t)
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		a.Equal(actual, c.expected)
	}
}
