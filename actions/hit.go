package actions

import (
	"main/helpers"
	"main/models"

	"github.com/parnurzeal/gorequest"
	"github.com/urfave/cli"
)

func Hit(c *cli.Context) {
	hitUrl := helpers.URL + "/time_cards/register"
	configuration := helpers.ReadConfigurationFile()
	credentials, _ := helpers.Authenticate(models.Credentials{Login: configuration.Username, Password: configuration.Password})

	request := gorequest.New()

	_, body, errs := request.Post(hitUrl).
		Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MOB31K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.106 Mobile Safari/537.36").
		Set("Content-Type", "application/json").
		Set("X-Requested-With", "br.com.pontomais.pontomais").
		Set("access-token", credentials.Token).
		Set("client", credentials.ClientID).
		Set("uid", configuration.Username).
		Send().
		End()
}
