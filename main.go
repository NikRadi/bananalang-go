package main

import (
	"fmt"
	"strconv"
)

// token
type (
	Type int

	Token struct {
		Type 	Type
		Value	string
	}
)

const (
	Error Type = iota
	EndOfFile
	LiteralNumber
)

var tokens = [...]string{
	Error: 			"Error",
	EndOfFile: 		"EndOfFile",
	LiteralNumber: 	"LiteralNumber",
}

func (token Token) String() string {
	return "{" + tokens[token.Type] + ", " + token.Value + "}"
}


// lexer
type Lexer struct {
	Code 		string
	CodeIndex	int
	token		Token
}

func NewLexer(code string) *Lexer {
	lexer := &Lexer{
		Code: 		code,
		CodeIndex: 	0,
		token:		Token{},
	}

	// Initialize the first Token
	lexer.EatToken()
	return lexer
}

func (lexer *Lexer) EatToken() {
	c := lexer.peekChar()
	switch c {
	default:
		if isDigit(c) {
			lexer.token = Token{Type: LiteralNumber, Value: string(c)}
			lexer.CodeIndex += 1
		} else {
			lexer.token = Token{Type: Error}
		}
	}
}

func (lexer *Lexer) PeekToken() Token {
	return lexer.token
}

func (lexer *Lexer) peekChar() uint8 {
	return lexer.Code[lexer.CodeIndex]
}

func isDigit(c uint8) bool {
	return '0' <= c && c <= '9'
}


// ast
type Literal struct {
	Type 	Type
	Value 	string
}


// parser
type Parser struct {
	Lexer *Lexer
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{
		Lexer: lexer,
	}
}

func (parser *Parser) Parse() Literal {
	tok := parser.Lexer.PeekToken()
	switch tok.Type {
	case LiteralNumber:
		return Literal{
			Type: 	LiteralNumber,
			Value:	tok.Value,
		}
	default:
		fmt.Print("error")
		return Literal{}
	}
}


// instruction
type Instruction byte
const (
	// Push a value to the stack
	Push Instruction = iota

	// Print the top value of the stack
	Print
)


// compiler
func Compile(literal Literal) []Instruction {
	var instructions []Instruction
	if literal.Type == LiteralNumber {
		value, err := strconv.Atoi(literal.Value)
		if err != nil {
			fmt.Println("Compile error: Invalid number")
			return nil
		}

		instructions = append(instructions, Push)
		instructions = append(instructions, Instruction(value))
	}

	instructions = append(instructions, Print)
	return instructions
}


// interpreter
func Interpret(instructions []Instruction) {
	var stack [64]int
	sp := 0
	for i := 0; i < len(instructions); i += 1 {
		instruction := instructions[i]
		switch instruction {
		case Push:
			i += 1
			value := instructions[i]
			stack[sp] = int(value)
			sp += 1
		case Print:
			top := stack[sp - 1]
			fmt.Println(top)
		}
	}
}


func main() {
	const code = "2"
	lexer := NewLexer(code)
	parser := NewParser(lexer)
	tree := parser.Parse()
	instructions := Compile(tree)
	Interpret(instructions)
}
