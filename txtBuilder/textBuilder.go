package txtBuilder

import (
	"fmt"
	"strings"
)

func (b *TextBuilder) GetText() string {
	if len(b.text.params) == 0 {
		return b.text.template
	}
	values := []interface{}{}
	for _,v := range b.text.params {
		values = append(values, v)
	}

	countTmplParam := strings.Count(b.text.template, "%v")
	if countTmplParam != len(b.text.params) {
		return b.text.template
	}
	replacedText := fmt.Sprintf(b.text.template, values...)
	return replacedText
}