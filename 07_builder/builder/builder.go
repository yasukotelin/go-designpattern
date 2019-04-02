package builder

import (
	"fmt"
	"os"
	"strings"
)

type Builder interface {
	MakeTitle(title string)
	MakeString(s string)
	MakeItems(items []string)
	Close()
}

type TextBuilder struct {
	sb strings.Builder
}

func (tb *TextBuilder) MakeTitle(title string) {
	tb.sb.WriteString("==========================\n")
	tb.sb.WriteString(fmt.Sprintf("『%s』\n", title))
	tb.sb.WriteString("\n")
}

func (tb *TextBuilder) MakeString(s string) {
	tb.sb.WriteString(fmt.Sprintf("■%s\n", s))
	tb.sb.WriteString("\n")
}

func (tb *TextBuilder) MakeItems(items []string) {
	for i := 0; i < len(items); i++ {
		tb.sb.WriteString(fmt.Sprintf(" .%s\n", items[i]))
	}
	tb.sb.WriteString("\n")
}

func (tb *TextBuilder) Close() {
	tb.sb.WriteString("==========================\n")
}

func (tb *TextBuilder) GetResult() string {
	return tb.sb.String()
}

type HTMLBuilder struct {
	filename string
	file     *os.File
}

func NewHTMLBuilder(filename string) (*HTMLBuilder, error) {
	f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &HTMLBuilder{
		filename: filename,
		file:     f,
	}, nil
}

func (hb *HTMLBuilder) MakeTitle(title string) {
	hb.file.WriteString(fmt.Sprintf("<html><head><title>%s</title></head><body>", title))
	hb.file.WriteString(fmt.Sprintf("<h1>%s</h1>", title))
}

func (hb *HTMLBuilder) MakeString(s string) {
	hb.file.WriteString(fmt.Sprintf("<p>%s</p>", s))
}

func (hb *HTMLBuilder) MakeItems(items []string) {
	hb.file.WriteString("<ul>")
	for i := 0; i < len(items); i++ {
		hb.file.WriteString(fmt.Sprintf("<li>%s</li>", items[i]))
	}
	hb.file.WriteString("</ul>")
}

func (hb *HTMLBuilder) Close() {
	hb.file.WriteString("</body></html>")
	hb.file.Close()
}

func (hb *HTMLBuilder) GetResult() string {
	return hb.filename
}
