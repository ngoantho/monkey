package lexer

type Lexer struct {
	input        string
	position     int  // current position
	readPosition int  // current reading position
	ch           byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
