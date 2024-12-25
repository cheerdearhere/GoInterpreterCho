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

func Test_NextToken_function(t *testing.T) {
	input := `let five = 5;
			let ten = 10;
			let add = function(x,y){
				return x + y;
			};
			let result = add(five,ten);
		`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LPARAN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPARAN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPARAN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPARAN, ")"},
		{token.SEMICOLON, ";"},
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
