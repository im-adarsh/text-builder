package txtBuilder

import (
	"fmt"
	"strings"
)

func (b *TextBuilder) GetText() string {
	countTmplParam := strings.Count(b.text.template, "%v")
	countParam := len(b.text.params)

	if countTmplParam != countParam  || countParam == 0{
		return b.text.template
	}

	values := []interface{}{}
	for _,v := range b.text.params {
		values = append(values, v)
	}

	return fmt.Sprintf(b.text.template, values...)
}