package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/PMcca/phmod/chmodparse"
	"github.com/urfave/cli"
)

var regexNum = regexp.MustCompile("^[0-7]?[0-7]{3}$") // 3 or 4 digits
var regexChar = regexp.MustCompile("^a$")             //TODO

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

	fmt.Println(parser.Parse(arg))

	return nil
}
