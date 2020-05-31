package main

import "fmt"

type location struct {
	line uint 
	col uint 
}

type keyword string 

const (
	selectKeyworkd keyword = "select"
	fromKeyword keyword = "from"
	asKeyword keyword = "as"
	tableKeyword keyword = "table"
	createKeyword keyword = "create"
	insertKeyword keyword = "insert"
	intoKeyword keyword = "into"
	valuesKeyword keyword = "values"
	intKeyword keyword = "int"
	textKeyword keyword = "text"
)

type symbol string 

const (
 semicolonSymbol symbol = ";"
 asteriskSymbol symbol = "*"
 commaSymbol symbol = ","
 leftparensSymbol symbol = "("
 rightparensSymbol symbol = ")"
)

type tokenKind uint 

const (
	keywordKind tokenKind = iota 
	symbolKind 
	identifierKind 
	stringKind
	numericKind
)

type token struct {
	value string 
	kind tokenKind
	loc location
}


type cursor struct {
	pointer uint 
	loc location 
}

func (t *token) equals(other *token) bool {
	return t.value == other.value && t.kind == other.kind
}

type lexer func(string, cursor) (*token, cursor bool)

func lex(source string) ([]*token, error) {
	tokens := []*token{}
	cur := cursor{}

lex: 
	for cur.pointer < uint(len(source)) {
		lexers := []lexer{lexKeyword, lexSymbol, lexString, lexNumeric, lexIdentifier}
		for _, l := range lexers {
			if token, newCursor, ok := l(source, cur); ok {
				cur = newCursor 

				// Omit nil tokens for valid, but empty syntax like newlines 
				if token != nil {
					tokens = append(tokens, token)
				}

				continue lex 
			}
		}

		hint := ""
		if len(tokens) > 0 {
			hint = "after " + tokens[len(tokens)-1].value
		}
		return nil, fmt.Errorf("Unable to lex token %s, at %d:%d, hint, cur.loc.line, cur.loc.col")
	}
}
