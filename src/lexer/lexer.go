package lexer

type Token struct {
	Typ TokenType
	Buf []rune
}



func Tokenize( str []rune) ([]*Token, error) {
	tokenText:=new([]rune)

	toks:= make([]*Token, 0)
	state:=INIT
	var p *TokenType= &state
	temptoken:=&Token{}
	var least rune
	for _,ch:= range str {
		least=ch
		switch state {
		case INIT:
			initToken(ch,temptoken,tokenText,&toks,p)
			break
		case Identifier:
			if (isAlpha(ch) || isDigit(ch)) {
				*tokenText=append(*tokenText, ch);       //保持标识符状态
			} else {
				initToken(ch,temptoken,tokenText,&toks,p)     //退出标识符状态，并保存Token
			}
			break;

		case IntLiteral:
			if (isDigit(ch)) {
				*tokenText=append(*tokenText, ch);    //继续保持在数字字面量状态
			} else {
				initToken(ch,temptoken,tokenText,&toks,p)      //退出当前状态，并保存Token
			}
			break;
		case GT:
			if (ch == '=') {
				state = GE;  //转换成GE
				temptoken.Typ =GE
				*tokenText=append(*tokenText, ch);
			} else {
				initToken(ch,temptoken,tokenText,&toks,p)  //退出GT状态，并保存Token
			}
			break;
		case GE:
		case Assign:
			initToken(ch,temptoken,tokenText,&toks,p)      //退出当前状态，并保存Token
			break;
		case Plus:
		case Minus:
		case Star:
		case Slash:
		case SemiColon:
			initToken(ch,temptoken,tokenText,&toks,p)      //退出当前状态，并保存Token
			break;
		case LeftParen:
		case RightParen:
			initToken(ch,temptoken,tokenText,&toks,p)      //退出当前状态，并保存Token
			break;
		}

	}
	// 把最后一个token送进去
	if (len(*tokenText) > 0) {
		initToken(least,temptoken,tokenText,&toks,p)
	}
	return  toks,nil
}

func initToken(ch rune, token *Token, tokentest *[]rune, tokens *[]*Token, tokenType *TokenType) {
	if len(*tokentest) > 0 {
		if typ, ok := keywords[string(*tokentest)]; ok {
			tok := &Token{
				Typ: typ,
				Buf: *tokentest,
			}
			*tokens =append(*tokens, tok)
		}else {
			tok := &Token{
				Typ: *tokenType,
				Buf: *tokentest,
			}
			*tokens =append(*tokens, tok)
		}
		*tokentest = make([]rune, 0)
		*token =Token{}

	}
	if isAlpha(ch) {
		*tokenType=Identifier
		*tokentest = append(*tokentest, ch)
	} else if isDigit(ch) {
		*tokenType=IntLiteral
		*tokentest = append(*tokentest, ch)
	} else if ch == '>' {
		*tokenType=GT
		*tokentest = append(*tokentest, ch)
	} else if ch == '=' {
		*tokenType=Assign
		*tokentest = append(*tokentest, ch)
	} else if ch == '+' {
		*tokenType=Plus
		*tokentest = append(*tokentest, ch)
	} else if ch == '-' {
		*tokenType=Minus
		*tokentest = append(*tokentest, ch)
	} else if ch == '*' {
		*tokenType=Star
		*tokentest = append(*tokentest, ch)
	} else if (ch == '(') {
		*tokenType=LeftParen
		*tokentest = append(*tokentest, ch)
	} else if (ch == ')') {
		*tokenType=RightParen
		*tokentest = append(*tokentest, ch)
	}else if ch == '/' {
		*tokenType=Slash
		*tokentest = append(*tokentest, ch)

	}else if ch == ';' {
		*tokenType=SemiColon
		*tokentest = append(*tokentest, ch)

	} else {
		*tokenType=INIT
	}
}


func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func (t TokenType) Str() string {
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