package main

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"
)

// Structs
type token struct {
	kind  string
	value string
	line  int16
}

// Tokens types
var tokenTypes = map[string][]string{
	"keyword":   []string{"if", "then", "while", "print"}, // Keyword values
	"vartype":   []string{"int", "str", "bool"},           // Variable types
	"boolean":   []string{"true", "false"},                // Boolean values
	"assign":    []string{"="},                            // Assignment
	"operators": []string{"+", "-"},                       // Possible operators
	"compare":   []string{"==", "!="},                     // Comparison operators
	"quote":     []string{"\""},                           // Start/end string characters
	"space":     []string{" ", "\t"},                      // Space characters
	"newline":   []string{"\n", "\r\n"},                   // New line characters
}

// Check if a string is inside of a slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

// Get token if available
func tokenCheck(word string, line int16) token {
	var tokenKind token
	for kind, list := range tokenTypes {
		if stringInSlice(word, list) {
			tokenKind = token{kind, word, line}
		}
	}
	return tokenKind
}

// Parse string for tokens
func tokenParse(input string) (tokens []token, err error) {
	var blankToken token
	var word string
	var inString bool
	var innerString token
	var line int16

	// Loop through input string
	for i, w := 0, 0; i < len(input); i += w {
		val, width := utf8.DecodeRuneInString(input[i:])
		w = width

		word += string(val)
		currToken := tokenCheck(word, line)

		if inString {
			if currToken = tokenCheck(string(val), line); currToken.kind == "quote" {
				innerString = token{"innerString", word[0 : len(word)-1], line}
				tokens = append(tokens, innerString)
			}
		} else if currToken.kind == "newline" {
			line++
		}

		if currToken != blankToken {
			if currToken.kind == "quote" {
				inString = !inString
			}

			tokens = append(tokens, currToken)
			word = ""
		}
	}

	// Lexing errors
	if word != "" && inString {
		err = errors.New(fmt.Sprintf("Missing end quote for quote on line %x", line))
	} else if word != "" {
		err = errors.New(fmt.Sprintf("Invalid syntax found after %#v."+
			" String remaining => %v\n", tokens[len(tokens)-1].value, word))
	}

	return tokens, err
}

func main() {
	// Parse string for tokens
	tokens, err := tokenParse("if true then print \"fun\" \n print \"t\"")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[ Lexer ] Tokens found => %#v\n", tokens)
}
