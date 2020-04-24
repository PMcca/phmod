package chmodparse

import (
	"errors"
	"fmt"
	"strings"
)

type Parser interface {
	Parse(arg string) (string, error)
	ParseVerbose(arg string) (string, error)
}

type numParser struct {
	specialMode rune
	mode        string
}
type CharParser struct {
}

func NewNumParser(arg string) numParser {
	p := numParser{}
	if len(arg) == 4 {
		p.specialMode = rune(arg[0])
		p.mode = string(arg[1:])
	} else {
		p.mode = arg
	}
	return p
}

var modes = map[rune]string{
	'0': "---",
	'1': "--x",
	'2': "-w-",
	'3': "-wx",
	'4': "r--",
	'5': "r-x",
	'6': "rw-",
	'7': "rwx",
}

// NumParser
func (p numParser) Parse(arg string) (string, error) {
	builder := strings.Builder{}
	// TODO special modes
	for _, r := range p.mode {
		builder.WriteString(modes[rune(r)])
	}
	return builder.String(), nil
}

func (p numParser) ParseVerbose(arg string) (string, error) {
	builder := strings.Builder{}
	if p.specialMode != 0 {
		builder.WriteString("Special Modes:\n")
		//TODO implement special modes
		arg = arg[1:]
	}

	// User modes
	builder.WriteString("\nUser Modes:\n")
	m := modes[rune(arg[0])]
	builder.WriteString(m + "\n")
	longMode(m, &builder)

	// Group modes
	builder.WriteString("\nGroup Modes:\n")
	m = modes[rune(arg[1])]
	builder.WriteString(m + "\n")
	longMode(m, &builder)

	//Other modes
	builder.WriteString("\nOther Modes:\n")
	m = modes[rune(arg[2])]
	builder.WriteString(m + "\n")
	longMode(m, &builder)

	return builder.String(), nil
}

// CharParser
func (p CharParser) Parse(arg string) (string, error) {
	a, err := reverseMap(modes, arg)
	if err != nil {
		return "", err
	}
	fmt.Printf("Given %s, we get %s", arg, a)

	return "IMPLEMENT CHARPARSER", nil
}

func (p CharParser) ParseVerbose(arg string) (string, error) {
	//TODO
	return "", nil
}

// Given rwx as val, will return 7 as k
func reverseMap(m map[rune]string, val string) (rune, error) {
	for k, v := range m {
		if val == v {
			return k, nil
		}
	}
	return 0, errors.New("Invalid argument")
}

// Given rw-, add Read, Write to builder
func longMode(m string, b *strings.Builder) {
	for _, s := range m {
		switch s {
		case 'r':
			b.WriteString("Read\n")
		case 'w':
			b.WriteString("Write\n")
		case 'x':
			b.WriteString("Execute\n")
		}
	}
}
