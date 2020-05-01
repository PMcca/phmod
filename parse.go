package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/PMcca/phmod/chmodparse"
	"github.com/urfave/cli"
)

var regexNum = regexp.MustCompile("^[1-7]?[0-7]{3}$") // 3 or 4 digits
var regexChar = regexp.MustCompile("^(?:r|-)(?:w|-)(?:x|-).{6}$")

var parser chmodparse.Parser

func parse(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		return errors.New("Invalid number of arguments")
	}

	arg := ctx.Args().First()
	switch {
	case regexNum.MatchString(arg):
		parser = chmodparse.NewNumParser(arg)
	case regexChar.MatchString(arg):
		parser = chmodparse.CharParser{}
	default:
		return errors.New("Invalid argument")
	}

	var parsed string

	if ctx.Bool("l") {
		parsed = parser.ParseVerbose(arg)
	} else {
		parsed = parser.Parse(arg)
	}

	fmt.Println(parsed)

	return nil
}
