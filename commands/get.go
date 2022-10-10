package commands

import (
	"main/helpers"
	"main/models"

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
	configFilePath := helpers.GetConfigFilePath()
	configuration := helpers.ReadConfigurationFile(configFilePath)
	helpers.Authenticate(models.Credentials{Login: configuration.Username, Password: configuration.Password})
}
