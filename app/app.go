package app

import (
	"main/commands"

	"github.com/urfave/cli"
)

// Configura uma nova aplicação de linha de comando
func Configure() *cli.App {
	app := cli.NewApp()
	app.Name = "Pongomais CLI"
	app.Usage = "Criado para facilitar a comunicação com a API do pontomais"

	app.Commands = []cli.Command{
		commands.BuildGetCommand(),
		commands.BuildHitCommand(),
		commands.BuildConfigCommand(),
	}

	return app
}
