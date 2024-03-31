package lexer

import (
	"bananalang/token"
	"fmt"
)

type Lexer struct {
	Code 		string
	CodeIndex	int
	token		token.Token
}

func NewLexer(code string) *Lexer {
	lexer := &Lexer{
		Code: 		code,
		CodeIndex: 	0,
		token:		token.Token{},
	}

	// Initialize the first Token
	lexer.EatToken()
	return lexer
}

func (lexer *Lexer) EatToken() {
	for lexer.CodeIndex < len(lexer.Code) {
		c := lexer.peekChar()
		if c == '\n' || c == ' ' {
			lexer.eatChar()
		} else {
			break
		}
	}

	if lexer.CodeIndex == len(lexer.Code) {
		lexer.token = token.Token{Type: token.EndOfFile}
		return
	}

	c := lexer.peekChar()
	switch c {
	case '+':
		lexer.token = token.Token{Type: token.Plus}
		lexer.eatChar()
	case '-':
		lexer.token = token.Token{Type: token.Minus}
		lexer.eatChar()
	case '*':
		lexer.token = token.Token{Type: token.Star}
		lexer.eatChar()
	default:
		if isDigit(c) {
			lexer.token = token.Token{Type: token.LiteralNumber, Value: string(c)}
			lexer.eatChar()
		} else {
			lexer.token = token.Token{Type: token.Error}
			lexer.eatChar()
			fmt.Println("Error token")
		}
	}
}

func (lexer *Lexer) PeekToken() token.Token {
	return lexer.token
}

func (lexer *Lexer) eatChar() {
	lexer.CodeIndex += 1
}

func (lexer *Lexer) peekChar() uint8 {
	return lexer.Code[lexer.CodeIndex]
}

func isDigit(c uint8) bool {
	return '0' <= c && c <= '9'
}
