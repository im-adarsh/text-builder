package txtBuilder

import "testing"

func TestTextBuilder_Build(t *testing.T) {
	type fields struct {
		text *Text
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{
			name: "testing without param",
			fields:fields{text:&Text{
				template: "hello",
				params: []string{},
			}},
			want: "hello",
		},
		{
			name: "testing with single param",
			fields:fields{text:&Text{
				template: "hello %v",
				params: []string{"world"},
			}},
			want: "hello world",
		},
		{
			name: "testing with multiple param",
			fields:fields{text:&Text{
				template: "hello %v %v",
				params: []string{"go","world"},
			}},
			want: "hello go world",
		},
		{
			name: "testing with parameterized template but no param",
			fields:fields{text:&Text{
				template: "hello %v %v",
				params: []string{},
			}},
			want: "hello %v %v",
		},
		{
			name: "testing with parameterized template but more param",
			fields:fields{text:&Text{
				template: "hello %v %v",
				params: []string{"go", "world", "on", "github"},
			}},
			want: "hello %v %v",
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &TextBuilder{
				text: tt.fields.text,
			}
			if got := b.GetText(); got != tt.want {
				t.Errorf("TextBuilder.GetText() = %v, want %v", got, tt.want)
			}
		})
	}
}
