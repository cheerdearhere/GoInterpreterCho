package parser

import (
	"GoInterpreter/src/main/ast"
	"GoInterpreter/src/main/lexer"
	"GoInterpreter/src/main/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	//두개를 읽어서 curToken, peekToken으로 세팅
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
