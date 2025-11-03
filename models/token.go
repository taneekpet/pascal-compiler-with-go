package models

type TokenType string

const (
	Reserved TokenType = "reserved"
	Program  TokenType = "program"
	Variable TokenType = "variable"
)

type Token struct {
	Lexemes     []rune
	IdReference *string
	Type        TokenType
	ExpandedTo  []Token
}
