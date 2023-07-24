package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Element interface {
	Value() int
}

type Integer struct {
	value int
}

func (i *Integer) Value() int {
	return i.value
}

func NewInteger(value int) *Integer {
	return &Integer{value}
}

type Operation int

const (
	Addition = iota
	Substraction
)

type BinaryOperation struct {
	Type        Operation
	Left, Right Element
}

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Substraction:
		return b.Left.Value() - b.Right.Value()
	default:
		panic("unsupported operation")
	}
}

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	Lpar
	Rpar
)

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return fmt.Sprintf("`%s", t.Text)
}

func Lex(input string) []Token {
	var result []Token
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, Token{Plus, "+"})
		case '-':
			result = append(result, Token{Minus, "-"})
		case '(':
			result = append(result, Token{Lpar, "("})
		case ')':
			result = append(result, Token{Rpar, ")"})
		default:
			sb := strings.Builder{}
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteRune(rune(input[j]))
					i++
				} else {
					result = append(result, Token{Int, sb.String()})
					i--
					break
				}
			}
		}
	}
	return result
}

func Parse(tokens []Token) Element {
	result := BinaryOperation{}
	haveLhs := false
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Type {
		case Int:
			n, _ := strconv.Atoi(token.Text)
			integer := Integer{n}
			if !haveLhs {
				result.Left = &integer
				haveLhs = true
			} else {
				result.Right = &integer
			}
		case Plus:
			result.Type = Addition
		case Minus:
			result.Type = Substraction
		case Lpar:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rpar {
					break
				}
			}
			var subexp []Token
			for k := i + 1; k < j; k++ {
				subexp = append(subexp, tokens[k])
			}
			element := Parse(subexp)
			if !haveLhs {
				result.Left = element
				haveLhs = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return &result
}

func main() {
	input := "(13+4)-(3+1)"
	tokens := Lex(input)
	fmt.Println(tokens)
	res := Parse(tokens)
	fmt.Printf("%s = %d \n", input, res.Value())
}
