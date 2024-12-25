package lexer

import (
	"GoInterpreter/src/main/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){};,"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPARAN, "("},
		{token.RPARAN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		t.Logf("testing[%d] \n\tType: %q\n\tLiteral: %q\n", i, tok.Type, tok.Literal)
		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q / get=%q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q / get=%q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
