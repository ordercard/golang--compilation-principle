package main

import (
	"errors"
	"fmt"
	"go-PlayScript/src/lexer"
)

func main() {
	str:=`
int a= 10;
`
	tokens,err:=lexer.Tokenize([]rune(str))
	if err !=nil {
		panic(errors.New("xxxxx"))
	}
	for _,value:=range tokens {
		fmt.Println(value.Typ.Str(),"----",string(value.Buf))
	}

}