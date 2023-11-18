package main

import (
	//	"bufio"
	"fmt"
	//	"os"
	"net/http"
	"regexp"
	"strings"
)

type TokenType string

const (
	Number       TokenType = "Number"
	RealNumber   TokenType = "RealNumber"
	Identifier   TokenType = "Identifier"
	ReservedWord TokenType = "ReservedWord"
	ArithmeticOp TokenType = "ArithmeticOp"
	ComparisonOp TokenType = "ComparisonOp"
	LogicalOp    TokenType = "LogicalOp"
	AssignmentOp TokenType = "AssignmentOp"
	IncrementOp  TokenType = "IncrementOp"
	Parenthesis  TokenType = "Parenthesis"
	Braces       TokenType = "Braces"
	Terminal     TokenType = "Terminal"
	Separator    TokenType = "Separator"
	Hexadecimal  TokenType = "Hexadecimal"
	String       TokenType = "String"
	CommentLine  TokenType = "CommentLine"
	CommentBlock TokenType = "CommentBlock"
)

type Token struct {
	Type  TokenType
	Value string
}

func Lex(input string) []Token {
	tokens := []Token{}

	words := strings.Fields(input)

	for _, word := range words {
		switch {
		case word == "+", word == "-", word == "*", word == "/":
			tokens = append(tokens, Token{ArithmeticOp, word})
		case word == "==", word == "!=", word == ">", word == "<", word == ">=", word == "<=":
			tokens = append(tokens, Token{ComparisonOp, word})
		case word == "YY", word == "OO", word == "!":
			tokens = append(tokens, Token{LogicalOp, word})
		case word == "=", word == "+=", word == "-=", word == "*=", word == "/=":
			tokens = append(tokens, Token{AssignmentOp, word})
		case word == "++", word == "--":
			tokens = append(tokens, Token{IncrementOp, word})
		case word == "(", word == ")":
			tokens = append(tokens, Token{Parenthesis, word})
		case word == "{", word == "}":
			tokens = append(tokens, Token{Braces, word})
		case word == ";":
			tokens = append(tokens, Token{Terminal, word})
		case word == ",":
			tokens = append(tokens, Token{Separator, word})
		default:
			if isNaturalNumber(word) {
				tokens = append(tokens, Token{Number, word})
			} else if isRealNumber(word) {
				tokens = append(tokens, Token{RealNumber, word})
			} else if isReservedWord(word) {
				tokens = append(tokens, Token{ReservedWord, word})
			} else if isIdentifier(word) {
				tokens = append(tokens, Token{Identifier, word})
			} else if isHexadecimal(word) {
				tokens = append(tokens, Token{Hexadecimal, word})
			} else if isString(word) {
				tokens = append(tokens, Token{String, word})
			} else if isCommentLine(word) {
				tokens = append(tokens, Token{CommentLine, word})
			} else if isCommentBlock(word) {
				tokens = append(tokens, Token{CommentBlock, word})
			}
		}
	}

	return tokens
}

func isNaturalNumber(word string) bool {
	match, _ := regexp.MatchString(`^\d+$`, word)
	return match
}

func isRealNumber(word string) bool {
	match, _ := regexp.MatchString(`^\d+\.\d+$`, word)
	return match
}

func isIdentifier(word string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z_]\w{0,9}$`, word)
	return match
}

func isReservedWord(word string) bool {
	reservedWords := []string{"if", "else", "while", "for", "function", "return", "String", "int"} // Ejemplo de palabras reservadas
	for _, reserved := range reservedWords {
		if word == reserved {
			return true
		}
	}
	return false
}

func isHexadecimal(word string) bool {
	match, _ := regexp.MatchString(`^0[xX][0-9A-Fa-f]+$`, word)
	return match
}

func isString(word string) bool {
	match, _ := regexp.MatchString(`^".*"$`, word)
	return match
}

func isCommentLine(word string) bool {
	match, _ := regexp.MatchString(`^¿.*$`, word)
	return match
}


func isCommentBlock(word string) bool {
	match, _ := regexp.MatchString(`^#\*.*\*#$`, word)
	return match
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			text := r.FormValue("text")
			tokens := Lex(text)

			for _, token := range tokens {
				fmt.Fprintf(w, "Type: %s, Value: %s\n", token.Type, token.Value)
			}
			fmt.Fprintf(w, "El texto ingresado es: %s", text)
		} else {
			fmt.Fprint(w, "<html><body><form method='post'><label>Ingresa un texto:</label><br><textarea name='text'></textarea><br><br><input type='submit' value='Enviar'></form></body></html>")
		}
	})

	fmt.Println("Servidor iniciado en http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
