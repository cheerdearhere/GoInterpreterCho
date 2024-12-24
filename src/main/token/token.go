package token

type TokenType string //interface: 토큰타입

type Token struct {
	Type    TokenType //토큰 타입을 상수로 지정 가능
	Literal string    //상수에 대한 처리
}

// 상수 선언
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" //식별자: id, variable name...
	INT   = "INT"   //리터럴

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPARAN = "("
	RPARAN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "function"
	LET      = "let"
)
