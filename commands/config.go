package commands

import (
	"main/actions"

	"github.com/urfave/cli"
)

// Retorna as configurações do comando pg config
func BuildConfigCommand() cli.Command {
	return cli.Command{
		Name:    "config",
		Usage:   "Configura as credenciais de autenticação do pontomais e outras informações necessárias para o registro dos pontos",
		Flags:   getConfigCommandFlags(),
		Action:  configCommandAction,
		Aliases: []string{"c"},
	}
}

// Retorna as flags do comando de autenticação
func getConfigCommandFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "login",
			Value: "",
		},
		cli.StringFlag{
			Name:  "password",
			Value: "",
		},
	}
}

// Executa a ação de configuração
func configCommandAction(c *cli.Context) {
	actions.Configure(c)
}
