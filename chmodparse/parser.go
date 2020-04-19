package chmodparse

import (
	"errors"
	"fmt"
	"strings"
)

type Parser interface {
	Parse(arg string) (string, error)
}

type NumParser struct {
	specialMode rune
	mode        string
}
type CharParser struct {
}

func NewNumParser(arg string) NumParser {
	p := NumParser{}
	if len(arg) == 4 {
		p.specialMode = arg[0]
		p.mode = strings.arg[1:3]
	} else {
		p.mode = arg
	}
	return p
}

var modes = map[string]string{
	"0": "---",
	"1": "--x",
	"2": "-w-",
	"3": "-wx",
	"4": "r--",
	"5": "r-x",
	"6": "rw-",
	"7": "rwx",
}

func (p NumParser) Parse(arg string) (string, error) {
	builder := strings.Builder{}

	if p.specialMode != 0 {
		builder.WriteString("Special Modes:\n")
		//TODO implement special modes
	}

	builder.WriteString(p.mode)

	return builder.String(), nil
}

func (p CharParser) Parse(arg string) (string, error) {
	a, err := reverseMap(modes, arg)
	if err != nil {
		return "", err
	}
	fmt.Printf("Given %s, we get %s", arg, a)

	return "IMPLEMENT CHARPARSER", nil
}

// Given rwx as val, will return 7 as k
func reverseMap(m map[string]string, val string) (string, error) {
	for k, v := range m {
		if val == v {
			return k, nil
		}
	}
	return "", errors.New("Invalid argument")
}
