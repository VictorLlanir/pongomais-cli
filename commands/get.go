package commands

import (
	"main/actions"
	"time"

	"github.com/urfave/cli"
)

// Retorna as configurações do comando pg get
func BuildGetCommand() cli.Command {
	return cli.Command{
		Name:    "get",
		Usage:   "Busca as informações de registro de pontos do dia informado",
		Flags:   getCommandFlags(),
		Action:  configGetAction,
		Aliases: []string{"c"},
	}
}

func getCommandFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:     "day",
			Value:    "",
			Required: false,
		},
	}
}

func configGetAction(c *cli.Context) {
	day := c.String("day")
	if day == "" {
		day = time.Now().Format("2006-01-02")
	}

	actions.Get(day, c)
}
