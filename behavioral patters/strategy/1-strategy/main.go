package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct {
}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct {
}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("	<li>" + item + "</li>\n")
}

type TextProcessor struct {
	builder      *strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(ls ListStrategy) *TextProcessor {
	return &TextProcessor{&strings.Builder{}, ls}
}

func (t *TextProcessor) SetOurputFormat(f OutputFormat) {
	switch f {
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}
	case Html:
		t.listStrategy = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	s := t.listStrategy
	s.Start(t.builder)
	for _, item := range items {
		s.AddListItem(t.builder, item)
	}
	s.End(t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	t := NewTextProcessor(&MarkdownListStrategy{})
	t.AppendList([]string{"foo", "bar", "bax"})
	fmt.Println(t)

	t.Reset()
	t.SetOurputFormat(Html)
	t.AppendList([]string{"foo", "bar", "bax"})
	fmt.Println(t)
}
