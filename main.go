package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func cliInit() *cli.App {
	app := cli.NewApp()

	app.Name = "Phmod"
	app.Description = "Translator for chmod strings, both digit and char representations."
	app.Usage = "phmod [OPTION] <args>"
	app.Version = "1.0.0"
	cli.AppHelpTemplate = `{{.Name}}{{if .Description}} - {{.Description}}{{end}}

{{if .Usage}}Usage: {{.Usage}}{{end}}

{{if .VisibleFlags}}Options:
{{range $index, $option := .VisibleFlags}}{{if $index}}
{{end}}{{$option}}{{end}}{{end}}`

	// app.Flags = []cli.Flag{
	// 	cli.BoolFlag{
	// 		Name:  "verbose, v",
	// 		Usage: "",
	// 	},
	// }
	return app
}

func main() {
	app := cliInit()
	if err := app.Run(os.Args); err != nil {
		log.Fatal("ERROR - ", err)
	}
}
