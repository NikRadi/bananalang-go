package parser

import (
	"bananalang/ast"
	"bananalang/lexer"
	"bananalang/token"
	"fmt"
	"os"
)

type (
	parseInfixOperatorFunction 	func(ast.Expression) ast.Expression
	parsePrefixOperatorFunction func() ast.Expression
)

type Parser struct {
	lexer 						*lexer.Lexer
	infixOperators				map[token.Type]parseInfixOperatorFunction
	infixOperatorPrecedences 	map[token.Type]int
	prefixOperators 			map[token.Type]parsePrefixOperatorFunction
}

func NewParser(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: 						lex,
		infixOperators: 			make(map[token.Type]parseInfixOperatorFunction),
		infixOperatorPrecedences:	make(map[token.Type]int),
		prefixOperators: 			make(map[token.Type]parsePrefixOperatorFunction),
	}

	parser.infixOperators[token.Equals] = parser.parseBinaryOperation
	parser.infixOperators[token.Plus] 	= parser.parseBinaryOperation
	parser.infixOperators[token.Minus] 	= parser.parseBinaryOperation
	parser.infixOperators[token.Star] 	= parser.parseBinaryOperation

	parser.infixOperatorPrecedences[token.Equals] 	=  9 // 10 - 1 because it is right associative
	parser.infixOperatorPrecedences[token.Plus] 	= 20
	parser.infixOperatorPrecedences[token.Minus] 	= 20
	parser.infixOperatorPrecedences[token.Star] 	= 30

	parser.prefixOperators[token.LiteralNumber] = parser.parseNumber
	parser.prefixOperators[token.Identifier] 	= parser.parseVariable

	return parser
}

func (parser *Parser) Parse() ast.Expression {
	return parser.parseExpression(0)
}

func (parser *Parser) parseBinaryOperation(leftExpression ast.Expression) ast.Expression {
	tok := parser.lexer.PeekToken()
	parser.lexer.EatToken()

	precedence := parser.infixOperatorPrecedences[tok.Type]
	rightExpression := parser.parseExpression(precedence)

	return ast.BinaryOperator{
		Operator: 			tok.Type,
		LeftExpression: 	leftExpression,
		RightExpression:	rightExpression,
	}
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	tok := parser.lexer.PeekToken()
	parsePrefixOperator, ok := parser.prefixOperators[tok.Type]
	if !ok {
		fmt.Println("Expected expression", tok)
		os.Exit(1)
	}

	leftExpression := parsePrefixOperator()

	tok = parser.lexer.PeekToken()
	parseInfixOperator := parser.infixOperators[tok.Type]
	for precedence < parser.infixOperatorPrecedences[tok.Type] {
		leftExpression = parseInfixOperator(leftExpression)

		tok = parser.lexer.PeekToken()
		parseInfixOperator = parser.infixOperators[tok.Type]
	}

	return leftExpression
}

func (parser *Parser) parseNumber() ast.Expression {
	tok := parser.lexer.PeekToken()
	parser.lexer.EatToken()
	return ast.Literal{Type: token.LiteralNumber, Value: tok.Value}
}

func (parser *Parser) parseVariable() ast.Expression {
	tok := parser.lexer.PeekToken()
	parser.lexer.EatToken()
	return ast.Literal{Type: token.Identifier, Value: tok.Value}
}
