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
