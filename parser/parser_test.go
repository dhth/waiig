package parser

import (
	"testing"

	"monkey/ast"
	"monkey/lexer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	require.NotNil(t, program)
	require.Equal(t, len(program.Statments), 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statments[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	t.Helper()
	require.Equal(t, s.TokenLiteral(), "let")

	letStmt, ok := s.(*ast.LetStatement)
	require.True(t, ok)

	require.Equal(t, letStmt.Name.Value, name)
	require.Equal(t, letStmt.Name.TokenLiteral(), name)
}

func checkParserErrors(t *testing.T, p *Parser) {
	t.Helper()
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	// GIVEN
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l)

	// WHEN
	program := p.ParseProgram()

	// THEN
	checkParserErrors(t, p)

	require.Equal(t, 3, len(program.Statments))

	for _, stmt := range program.Statments {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		assert.True(t, ok)
		assert.Equal(t, "return", returnStmt.TokenLiteral())
	}
}
