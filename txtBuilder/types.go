package txtBuilder

type Text struct {
	template string
	params []string
}

// Text txtBuilder pattern code
type TextBuilder struct {
	text *Text
}

func NewTextBuilder() *TextBuilder {
	text := &Text{}
	b := &TextBuilder{text: text}
	return b
}

func (b *TextBuilder) Template(template string) *TextBuilder {
	b.text.template = template
	return b
}

func (b *TextBuilder) Params(params []string) *TextBuilder {
	b.text.params = params
	return b
}


func (b *TextBuilder) Build() (*Text, error) {
	return b.text, nil
}

