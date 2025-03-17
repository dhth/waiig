package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"monkey/token"
)

func TestString(t *testing.T) {
	// GIVEN
	program := Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	// WHEN
	expected := "let myVar = anotherVar;"
	got := program.String()

	// THEN
	assert.Equal(t, expected, got)
}
