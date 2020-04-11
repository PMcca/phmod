package chmodparse

type Parser interface {
	Parse() string
}

type NumParser struct {
}
type CharParser struct {
}

func (p NumParser) Parse() string {
	return "123"
}

func NewNumParser() *NumParser {
	return &NumParser{}
}
