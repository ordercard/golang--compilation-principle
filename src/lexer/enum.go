package lexer
type TokenType int
const (
	EOF TokenType = iota
	INT
	DOUBLE
	STRING
	BOOL
	NIL
	NewLine

	Identifier

	GT
	GE
	EQ
	Assign
	Plus
	PlusPlus
	PlusEQ
	Minus
	MinusMinus
	MinusEQ
	Star
	StarEQ
	Slash
	SlashEQ
	SemiColon
	LeftParen
	RightParen

	Assignment   //=
	IntLiteral  //整型字面量 数字 16
	UintLiteral //16位整形字面量  19
	DoubleLiteral
	INIT

)

func (t TokenType) Strings() string {
	switch t {
	case Assignment:
		return "Assignment"
	case SemiColon:
		return "SemiColon"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"

	case INT:
		return "INT"
	case DOUBLE:
		return "DOUBLE"
	case STRING:
		return "STRING"
	case BOOL:
		return "BOOL"
	case NIL:
		return "NIL"
	case Identifier:
		return "Identifier"
	case GT:
		return "GT"
	case GE:
		return "GE"
	case EQ:
		return "EQ"
	case Assign:
		return "Assign"
	case Plus:
		return "Plus"
	case PlusPlus:
		return "PlusPlus"
	case PlusEQ:
		return "PlusEQ"
	case Minus:
		return "Minus"
	case MinusMinus:
		return "MinusMinus"
	case MinusEQ:
		return "MinusEQ"
	case Star:
		return "Star"
	case StarEQ:
		return "StarEQ"
	case Slash:
		return "Slash"
	case SlashEQ:
		return "SlashEQ"
	case IntLiteral:
		return "IntLiteral"
	case UintLiteral:
		return "UintLiteral"
	case DoubleLiteral:
		return "DoubleLiteral"
	default:
		return "Unknown TokenType"
	}
}
var keywords = map[string]TokenType{
	"int":    INT,
	"double": DOUBLE,
	"string": STRING,
	"bool":   BOOL,
	"nil":    NIL,
}