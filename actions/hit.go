package actions

import (
	"fmt"
	"log"
	"main/helpers"
	"main/models"

	"github.com/parnurzeal/gorequest"
	"github.com/urfave/cli"
)

// Registra o ponto na API do pontomais
func Hit(c *cli.Context) {
	hitUrl := helpers.URL + "/time_cards/register"
	configuration := helpers.ReadConfigurationFile()
	credentials, _ := helpers.Authenticate(models.Credentials{Login: configuration.Username, Password: configuration.Password})
	fmt.Println(credentials)
	payload := models.HitBody{
		TimeCardInfo: models.TimeCardInfo{
			Latitude:          configuration.Latitude,
			Longitude:         configuration.Longitude,
			Address:           configuration.Address,
			ReferenceId:       "",
			OriginalLatitude:  configuration.Latitude,
			OriginalLongitude: configuration.Longitude,
			OriginalAddress:   configuration.Address,
			LocationEdited:    true,
		},
		Path: "/registrar-ponto",
		Device: models.DeviceInfo{
			Browser: models.BrowserInfo{
				Name:                "Firefox",
				Version:             "86.0",
				VersionSearchString: "Firefox",
			},
		},
		AppVersion: "0.10.32",
	}

	request := gorequest.New()

	response, _, errs := request.Post(hitUrl).
		Type("json").
		Set("Host", "api.pontomais.com.br").
		Set("Accept", "application/json, text/plain, */*").
		Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MOB31K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.106 Mobile Safari/537.36").
		Set("token-type", "Bearer").
		Set("Content-Type", "application/json;charset=utf-8").
		Set("Api-Version", "2").
		Set("Origin", "https://app2.pontomais.com.br").
		Set("Referer", "https://app2.pontomais.com.br/").
		Set("access-token", credentials.Token).
		Set("token", credentials.Token).
		Set("client", credentials.ClientID).
		Set("uid", configuration.Username).
		Send(payload).
		End()

	if len(errs) > 0 {
		log.Println("[ERRO] Falhou para registrar o ponto.")
	}

	if response.StatusCode > 201 {
		fmt.Println(response.Status)
		fmt.Println("[SUCESSO] Ponto registrado com sucesso!")
	}
}
