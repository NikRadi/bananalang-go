package token

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
	Identifier
	LiteralNumber

	Plus				// +
	Minus				// -
	Star				// *
	LeftRoundBracket 	// (
	RightRoundBracket	// )

	Equals				// =
	TwoEquals			// ==
	NotEquals			// !=
	LessThan			// <
	LessThanEquals		// <=
	GreaterThan			// >
	GreaterThanEquals	// >=
)

var types = [...]string{
	EndOfFile: 		"EndOfFile",
	Equals:			"Equals",
	Error: 			"Error",
	Identifier:		"Identifier",
	LiteralNumber: 	"LiteralNumber",
	Minus:			"Minus",
	NotEquals:		"NotEquals",
	Plus:			"Plus",
	Star:			"Star",
	TwoEquals:		"TwoEquals",
}

func (tokenType Type) String() string {
	return types[tokenType]
}
