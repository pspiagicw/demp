package demp

import (
	"fmt"
	"strings"
)

const (
	VARIABLE = iota
	TEXT
	ILLEGAL
)

type Token struct {
	Value string
	Type  int
}

func (t Token) String() string {
	return fmt.Sprintf("Token(Value: '%s', Type: %d)", t.Value, t.Type)
}

func ResolveTemplate(template string, variables map[string]string) string {

	tokens := tokenize(template)

	values := []string{}

	for _, token := range tokens {
		if token.Type == VARIABLE {
			if val, ok := variables[token.Value]; ok {
				values = append(values, val)
			} else {
				values = append(values, token.Value)
			}
		} else {
			values = append(values, token.Value)
		}
	}

	return strings.Join(values, "")

}
func tokenize(template string) []Token {
	tokens := []Token{}

	readPos := 0
	curPos := -1
	eof := false
	curChar := ""

	advance := func() {
		curPos = readPos
		readPos++
		if readPos > len(template) {
			eof = true
			curChar = ""
		} else {
			curChar = string(template[curPos])
		}
	}

	getVar := func() string {
		advance()
		if curChar == "$" {
			advance()
			return "$"
		}

		if curChar == "{" {
			advance()
			start := curPos
			for curChar != "}" {
				advance()
			}
			advance()
			return template[start : curPos-1]
		}

		start := curPos
		for !eof && curChar != " " && curChar != "$" && curChar != `"` {
			advance()
		}

		return template[start:curPos]
	}

	advance()

	for !eof {
		token := Token{Type: ILLEGAL}
		if curChar == "$" {
			variable := getVar()
			token = Token{Value: variable, Type: VARIABLE}
		} else {
			token = Token{Value: string(curChar), Type: TEXT}
			advance()
		}

		tokens = append(tokens, token)
	}

	return tokens
}
