package parser

import (
	"GoInterpreter/src/main/ast"
	"GoInterpreter/src/main/lexer"
	"GoInterpreter/src/main/token"
	"fmt"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	//두개를 읽어서 curToken, peekToken으로 세팅
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) Errors() []string {
	return p.errors
}

/*
기댓값과 다를 경우 작용.
*/
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		t,
		p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

/*
expectPeek
거의 모든 파서가 공유하는 단정(assertion)함수.
다음 토큰 타입을 검사해 토큰 간의 순서를 올바르게 강제할 용도
정확한 타입일때만 다음을 호출하는 역할로 자주사용
*/
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	//Identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	//Assignment
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//until meeting semicolon(;)
	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

/*
*
EOF를 만나기 전까지 반복
*/
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}
