package parser

import (
	"bananalang/ast"
	"bananalang/lexer"
	"bananalang/token"
	"fmt"
	"os"
)

type (
	parseInfixOperatorFunction func(ast.Expression) ast.Expression
	parsePrefixOperatorFunction func() ast.Expression
)

type Parser struct {
	Lexer 						*lexer.Lexer
	infixOperators				map[token.Type]parseInfixOperatorFunction
	infixOperatorPrecedences 	map[token.Type]int
	prefixOperators 			map[token.Type]parsePrefixOperatorFunction
}

func NewParser(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		Lexer: 						lex,
		infixOperators: 			make(map[token.Type]parseInfixOperatorFunction),
		infixOperatorPrecedences:	make(map[token.Type]int),
		prefixOperators: 			make(map[token.Type]parsePrefixOperatorFunction),
	}

	parser.infixOperators[token.Plus] 	= parser.parseBinaryOperation
	parser.infixOperators[token.Minus] 	= parser.parseBinaryOperation
	parser.infixOperators[token.Star] 	= parser.parseBinaryOperation

	parser.infixOperatorPrecedences[token.Plus] 	= 10
	parser.infixOperatorPrecedences[token.Minus] 	= 10
	parser.infixOperatorPrecedences[token.Star] 	= 20

	parser.prefixOperators[token.LiteralNumber] = parser.parseNumber

	return parser
}

func (parser *Parser) Parse() ast.Expression {
	return parser.parseExpression(0)
}

func (parser *Parser) parseBinaryOperation(leftExpression ast.Expression) ast.Expression {
	tok := parser.Lexer.PeekToken()
	parser.Lexer.EatToken()

	precedence := parser.infixOperatorPrecedences[tok.Type]
	rightExpression := parser.parseExpression(precedence)

	return ast.BinaryOperator{
		Operator: 			tok.Type,
		LeftExpression: 	leftExpression,
		RightExpression:	rightExpression,
	}
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	tok := parser.Lexer.PeekToken()
	parsePrefixOperator, ok := parser.prefixOperators[tok.Type]
	if !ok {
		fmt.Println("Expected expression", tok)
		os.Exit(1)
	}

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

func (parser *Parser) parseNumber() ast.Expression {
	tok := parser.Lexer.PeekToken()
	parser.Lexer.EatToken()
	return ast.Literal{Type: token.LiteralNumber, Value: tok.Value}
}
