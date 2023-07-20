package main

import (
	"fmt"
	"strings"
	"unicode"
)

// capitalize may allocate a lot of memory
type FormattedText struct {
	plainText  string
	capitalize []bool
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{
		plainText:  plainText,
		capitalize: make([]bool, len(plainText)),
	}
}

func (f FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
			continue
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

// Flywight structure: avoid allocating so much memory by working with ranges
type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (tr *TextRange) Covers(position int) bool {
	return position >= tr.Start && position <= tr.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{
		plainText: plainText,
	}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{Start: start, End: end}
	b.formatting = append(b.formatting, r)
	return r
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.plainText); i++ {
		c := rune(b.plainText[i])
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				c = unicode.ToUpper(c)
			}
		}
		sb.WriteRune(c)
	}
	return sb.String()
}

func main() {
	text := "Hello world"
	ft := NewFormattedText(text)
	ft.Capitalize(0, 5)
	fmt.Println(ft.String())

	bft := NewBetterFormattedText(text)
	bft.Range(5, 10).Capitalize = true
	fmt.Println(bft.String())

}
