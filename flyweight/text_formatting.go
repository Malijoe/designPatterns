package flyweight

import (
    "fmt"
    "strings"
    "unicode"
)

// inefficient implementation; a huge boolean array for each letter in the text field

type FormattedText struct {
    plainText string
    capitalize []bool
}

func (f *FormattedText) String() string {
    sb := strings.Builder{}
    for i := 0; i < len(f.plainText); i ++ {
        c := f.plainText[i]
        if f.capitalize[i] {
            sb.WriteRune(unicode.ToUpper(rune(c)))
        } else {
            sb.WriteRune(rune(c))
        }
    }
    return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
    for i := start; i <= end; i++ {
        f.capitalize[i] = true
    }
}

func NewFormattedText(plainText string) *FormattedText {
    return &FormattedText{ plainText: plainText, capitalize: make([]bool, len(plainText))}
}

type TextRange struct {
    Start, End int
    Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
    return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
    plainText string
    formatting []*TextRange
}

func NewBetterFormattedText (plainText string) *BetterFormattedText {
    return &BetterFormattedText{plainText: plainText}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
    r := &TextRange{Start: start, End: end}
    b.formatting = append(b.formatting, r)
    return r
}

func (b *BetterFormattedText) String() string {
    sb := strings.Builder{}

    for i := 0; i < len(b.plainText); i++ {
        c := b.plainText[i]
        for _, r := range b.formatting {
            if r.Covers(i) && r.Capitalize {
                c = uint8(unicode.ToUpper(rune(c)))
            }
        }
        sb.WriteRune(rune(c))
    }
    return sb.String()
}

func TextFormat() {
    text := "This is a brave new world"
    ft := NewFormattedText(text)
    ft.Capitalize(10, 15)
    fmt.Println(ft.String())

    bft := NewBetterFormattedText(text)
    bft.Range(16, 19).Capitalize = true
    fmt.Println(bft.String())
}