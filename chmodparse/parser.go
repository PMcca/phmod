package chmodparse

type Parser interface {
	Parse(s string) string
}

type NumParser struct {
}
type CharParser struct {
}

func (p NumParser) Parse(s string) string {
	return "IMPLEMENT NUMPARSER"
}

func (p CharParser) Parse(s string) string {
	return "IMPLEMENT CHARPARSER"
}
