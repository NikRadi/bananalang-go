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

	Equals				// =
	Plus				// +
	Minus				// -
	Star				// *
	LeftRoundBracket 	// (
	RightRoundBracket	// )
)

var types = [...]string{
	EndOfFile: 		"EndOfFile",
	Equals:			"Equals",
	Error: 			"Error",
	Identifier:		"Identifier",
	LiteralNumber: 	"LiteralNumber",
	Minus:			"Minus",
	Plus:			"Plus",
	Star:			"Star",
}

func (tokenType Type) String() string {
	return types[tokenType]
}
