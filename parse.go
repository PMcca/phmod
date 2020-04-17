package main

import (
	"fmt"
	"regexp"

	"github.com/PMcca/phmod/chmodparse"
	"github.com/urfave/cli"
)

var regexNum = regexp.MustCompile("^[0124]?\\d{3}$")

func parse(ctx *cli.Context) error {
	var parser chmodparse.Parser
	parser = chmodparse.NewNumParser()
	parser.Parse()

	fmt.Println(regexNum.MatchString("1234"))

	return nil
}
