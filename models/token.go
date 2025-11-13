package models

type TokenType string

const (
	Reserved TokenType = "reserved"
	Program  TokenType = "program"
	Variable TokenType = "variable"
	Block    TokenType = "block"

	Declaration         TokenType = "declaration"
	CompoundStatement   TokenType = "compoundStatement"
	VariableDeclaration TokenType = "variableDeclaration"
)

type Token struct {
	Lexemes     []rune
	IdReference *string
	Type        TokenType
	ExpandedTo  []Token
}
