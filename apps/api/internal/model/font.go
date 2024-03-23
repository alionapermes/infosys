package model

type FontFormat string

const (
	FontFormatPlain  FontFormat = "plain"
	FontFormatItalic FontFormat = "italic"
	FontFormatBold   FontFormat = "bold"
)

type Font struct {
	Family string     `json:"family"`
	Name   string     `json:"name"`
	Format FontFormat `json:"format"`
}
