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
	Plus
	Minus
	Star
)

var types = [...]string{
	Error: 			"Error",
	EndOfFile: 		"EndOfFile",
	LiteralNumber: 	"LiteralNumber",
	Plus:			"Plus",
	Minus:			"Minus",
	Star:			"Star",
}

func (tokenType Type) String() string {
	return types[tokenType]
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
	if lexer.CodeIndex == len(lexer.Code) {
		lexer.token = Token{Type: EndOfFile}
		return
	}

	c := lexer.peekChar()
	switch c {
	case '+':
		lexer.token = Token{Type: Plus}
		lexer.eatChar()
	case '-':
		lexer.token = Token{Type: Minus}
		lexer.eatChar()
	case '*':
		lexer.token = Token{Type: Star}
		lexer.eatChar()
	default:
		if isDigit(c) {
			lexer.token = Token{Type: LiteralNumber, Value: string(c)}
			lexer.eatChar()
		} else {
			lexer.token = Token{Type: Error}
			lexer.eatChar()
		}
	}
}

func (lexer *Lexer) PeekToken() Token {
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


// ast
type (
	Expression interface {}

	Literal struct {
		Type 	Type
		Value 	string
	}

	BinaryOperator struct {
		Operator 		Type
		LeftExpression	Expression
		RightExpression Expression
	}
)

// parser
type (
	parseInfixOperatorFunction func(Expression) Expression
	parsePrefixOperatorFunction func() Expression
)

type Parser struct {
	Lexer 						*Lexer
	infixOperators				map[Type]parseInfixOperatorFunction
	infixOperatorPrecedences 	map[Type]int
	prefixOperators 			map[Type]parsePrefixOperatorFunction
}

func NewParser(lexer *Lexer) *Parser {
	parser := &Parser{
		Lexer: 						lexer,
		infixOperators: 			make(map[Type]parseInfixOperatorFunction),
		infixOperatorPrecedences:	make(map[Type]int),
		prefixOperators: 			make(map[Type]parsePrefixOperatorFunction),
	}

	parser.infixOperators[Plus] 	= parser.parseBinaryOperation
	parser.infixOperators[Minus] 	= parser.parseBinaryOperation
	parser.infixOperators[Star] 	= parser.parseBinaryOperation

	parser.infixOperatorPrecedences[Plus] 	= 10
	parser.infixOperatorPrecedences[Minus] 	= 10
	parser.infixOperatorPrecedences[Star] 	= 20

	parser.prefixOperators[LiteralNumber] = parser.parseNumber

	return parser
}

func (parser *Parser) Parse() Expression {
	return parser.parseExpression(0)
}

func (parser *Parser) parseBinaryOperation(leftExpression Expression) Expression {
	tok := parser.Lexer.PeekToken()
	parser.Lexer.EatToken()

	precedence := parser.infixOperatorPrecedences[tok.Type]
	rightExpression := parser.parseExpression(precedence)

	return BinaryOperator{
		Operator: 			tok.Type,
		LeftExpression: 	leftExpression,
		RightExpression:	rightExpression,
	}
}

func (parser *Parser) parseExpression(precedence int) Expression {
	tok := parser.Lexer.PeekToken()
	parsePrefixOperator := parser.prefixOperators[tok.Type]
	leftExpression := parsePrefixOperator()

	tok = parser.Lexer.PeekToken()
	parseInfixOperator := parser.infixOperators[tok.Type]
	for precedence < parser.infixOperatorPrecedences[tok.Type] {
		leftExpression = parseInfixOperator(leftExpression)

		tok = parser.Lexer.PeekToken()
		parseInfixOperator = parser.infixOperators[tok.Type]
	}

	return leftExpression
}

func (parser *Parser) parseNumber() Expression {
	tok := parser.Lexer.PeekToken()
	parser.Lexer.EatToken()
	return Literal{Type: LiteralNumber, Value: tok.Value}
}


// instruction
type Instruction byte
const (
	// Push a value to the stack
	Push Instruction = iota

	// Print the top value of the stack
	Print

	// Pop the top two values, add them, and push the result
	Add

	// Pop the top two values, subtract the 2nd from the 1st, and push the result
	Sub

	// Pop the top two values, multiply them, and push the result
	Mul
)

func printInstructions(instructions []Instruction) {
	for i := 0; i < len(instructions); i += 1 {
		instruction := instructions[i]
		switch instruction {
		case Push:
			i += 1
			value := instructions[i]
			fmt.Println("Push", value)
		case Print:
			fmt.Println("Print")
		case Add:
			fmt.Println("Add")
		case Sub:
			fmt.Println("Sub")
		case Mul:
			fmt.Println("Mul")
		}
	}
}

// compiler
func Compile(expression Expression) []Instruction {
	var instructions []Instruction
	switch expr := expression.(type) {
	case Literal:
		value, err := strconv.Atoi(expr.Value)
		if err != nil {
			fmt.Println("Compile error: Invalid number")
			return nil
		}

		instructions = append(instructions, Push)
		instructions = append(instructions, Instruction(value))
	case BinaryOperator:
		instructions = append(instructions, Compile(expr.LeftExpression)...)
		instructions = append(instructions, Compile(expr.RightExpression)...)
		switch expr.Operator {
		case Plus:
			instructions = append(instructions, Add)
		case Minus:
			instructions = append(instructions, Sub)
		case Star:
			instructions = append(instructions, Mul)
		default:
			fmt.Println("Compile error: Unknown binary operator")
		}
	default:
		fmt.Println("Compile error: Unknown expression type")
	}

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
		case Add:
			value1 := stack[sp - 2]
			value2 := stack[sp - 1]
			stack[sp - 2] = value1 + value2
			sp -= 1
		case Sub:
			value1 := stack[sp - 2]
			value2 := stack[sp - 1]
			stack[sp - 2] = value1 - value2
			sp -= 1
		case Mul:
			value1 := stack[sp - 2]
			value2 := stack[sp - 1]
			stack[sp - 2] = value1 * value2
			sp -= 1
		}
	}
}


func main() {
	const code = "1+2*3-1"
	lexer := NewLexer(code)
	parser := NewParser(lexer)
	tree := parser.Parse()
	instructions := Compile(tree)
	instructions = append(instructions, Print)
	printInstructions(instructions)
	Interpret(instructions)
}
