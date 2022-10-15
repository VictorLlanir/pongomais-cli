package commands

import (
	"main/actions"

	"github.com/urfave/cli"
)

// Retorna as configurações do comando pg hit
func BuildHitCommand() cli.Command {
	return cli.Command{
		Name:    "hit",
		Usage:   "Bate o ponto no horário atual",
		Action:  configHitAction,
		Aliases: []string{"h"},
	}
}

// Executa a ação do comando pg hit
func configHitAction(c *cli.Context) {
	actions.Hit(c)
}
