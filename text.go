package rtfdoc

import (
	"fmt"
	"strings"
)

func (text text) compose() string {
	var res string

	var emphTextSlice []string
	if text.isBold {
		emphTextSlice = append(emphTextSlice, "\\b")
	}
	if text.isItalic {
		emphTextSlice = append(emphTextSlice, "\\i")
	}
	if text.isScaps {
		emphTextSlice = append(emphTextSlice, "\\scaps")
	}
	if text.isStrike {
		emphTextSlice = append(emphTextSlice, "\\strike")
	}
	if text.isSub {
		emphTextSlice = append(emphTextSlice, "\\sub")
	}
	if text.isSuper {
		emphTextSlice = append(emphTextSlice, "\\super")
	}
	if text.isUnderlining {
		emphTextSlice = append(emphTextSlice, "\\ul")
	}

	PreparedText := convertNonASCIIToUTF16(text.content)

	res += fmt.Sprintf("\n\\fs%d\\f%d \\cf%d {%s%s}\\f0", text.fontSize*2, text.fontCode, text.colorCode, strings.Join(emphTextSlice, " "), PreparedText)
	return res
}

// AddText returns new text instance
func (p *paragraph) AddText(textStr string, fontSize int, fontCode string, colorCode string) *text {

	fn := 0
	for i, f := range *p.generalSettings.fontColor {
		if f.code == fontCode {

			fn = i
		}
	}

	fc := 0
	for i, c := range *p.generalSettings.colorTable {
		if c.name == colorCode {

			fc = i + 1
		}
	}
	txt := text{
		fontSize:  fontSize,
		fontCode:  fn,
		colorCode: fc,
		content:   textStr,
		generalSettings: generalSettings{
			colorTable: p.colorTable,
			fontColor:  p.fontColor,
		},
	}
	p.content = append(p.content, &txt)
	return &txt
}

//AddNewLine adds new line into paragraph text
func (p *paragraph) AddNewLine() *paragraph {
	txt := text{
		content: "\\line",
	}
	p.content = append(p.content, &txt)
	return p
}

// SetBold function sets text to Bold
func (text *text) SetBold() *text {
	text.isBold = true
	return text
}

// SetItalic function sets text to Italic
func (text *text) SetItalic() *text {
	text.isItalic = true
	return text
}

// SetUnderlining function sets text to Underlining
func (text *text) SetUnderlining() *text {
	text.isUnderlining = true
	return text
}

// SetSuper function sets text to Super
func (text *text) SetSuper() *text {
	text.isSuper = true
	return text
}

// SetSub function sets text to Sub
func (text *text) SetSub() *text {
	text.isSub = true
	return text
}

// SetScaps function sets text to Scaps
func (text *text) SetScaps() *text {
	text.isScaps = true
	return text
}

// SetStrike function sets text to Strike
func (text *text) SetStrike() *text {
	text.isStrike = true
	return text
}

func (text *text) getEmphasis() string {
	return text.emphasis
}

// SetColor sets text color
func (text *text) SetColor(colorCode string) *text {
	for i := range *text.colorTable {
		if (*text.colorTable)[i].name == colorCode {
			// Присваиваем тексту порядковый номер шрифта
			text.colorCode = i + 1
		}
	}

	return text
}
