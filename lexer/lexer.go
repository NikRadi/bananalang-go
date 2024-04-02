package lexer

import (
	"bananalang/token"
	"fmt"
	"os"
)

type Lexer struct {
	code 		string
	codeIndex	int
	token		token.Token
}

func NewLexer(code string) *Lexer {
	lexer := &Lexer{
		code: 		code,
		codeIndex: 	0,
		token:		token.Token{},
	}

	// Initialize the first Token
	lexer.EatToken()
	return lexer
}

func (lexer *Lexer) EatToken() {
	for lexer.codeIndex < len(lexer.code) {
		c := lexer.peekChar()
		if c == '\n' || c == ' ' {
			lexer.eatChar()
		} else {
			break
		}
	}

	if lexer.codeIndex == len(lexer.code) {
		lexer.token = token.Token{Type: token.EndOfFile}
		return
	}

	c := lexer.peekChar()
	switch c {
	case ';':
		lexer.eatChar()
		lexer.token = token.Token{Type: token.Semicolon}
	case '=':
		lexer.eatChar()
		if lexer.peekChar() == '=' {
			lexer.eatChar()
			lexer.token = token.Token{Type: token.TwoEquals}
		} else {
			lexer.token = token.Token{Type: token.Equals}
		}
	case '!':
		lexer.eatChar()
		if lexer.peekChar() == '=' {
			lexer.eatChar()
			lexer.token = token.Token{Type: token.NotEquals}
		} else {
			fmt.Println("Parsing error: unknown !")
			os.Exit(1)
		}
	case '<':
		lexer.eatChar()
		if lexer.peekChar() == '=' {
			lexer.eatChar()
			lexer.token = token.Token{Type: token.LessThanEquals}
		} else {
			lexer.token = token.Token{Type: token.LessThan}
		}
	case '>':
		lexer.eatChar()
		if lexer.peekChar() == '=' {
			lexer.eatChar()
			lexer.token = token.Token{Type: token.GreaterThanEquals}
		} else {
			lexer.token = token.Token{Type: token.GreaterThan}
		}
	case '+':
		lexer.token = token.Token{Type: token.Plus}
		lexer.eatChar()
	case '-':
		lexer.token = token.Token{Type: token.Minus}
		lexer.eatChar()
	case '*':
		lexer.token = token.Token{Type: token.Star}
		lexer.eatChar()
	case '(':
		lexer.token = token.Token{Type: token.LeftRoundBracket}
		lexer.eatChar()
	case ')':
		lexer.token = token.Token{Type: token.RightRoundBracket}
		lexer.eatChar()
	default:
		if isDigit(c) {
			lexer.token = token.Token{Type: token.LiteralNumber, Value: string(c)}
			lexer.eatChar()
		} else if isAlphabetic(c) || c == '_' {
			lexer.token = token.Token{Type: token.Identifier, Value: string(c)}
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
	lexer.codeIndex += 1
}

func (lexer *Lexer) peekChar() byte {
	return lexer.code[lexer.codeIndex]
}

func isAlphabetic(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
