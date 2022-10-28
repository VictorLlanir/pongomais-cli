package actions

import (
	"crypto/tls"
	"fmt"
	"log"
	"main/helpers"
	"main/models"

	"github.com/google/uuid"
	"github.com/parnurzeal/gorequest"
	"github.com/urfave/cli"
	"golang.org/x/exp/slices"
)

// Registra o ponto na API do pontomais
func Hit(c *cli.Context) {
	hitUrl := helpers.URL + "/time_cards/register"
	configuration := helpers.ReadConfigurationFile()
	credentials, _ := helpers.Authenticate(models.Credentials{Login: configuration.Username, Password: configuration.Password})

	requestData := fmt.Sprintf(`{
		"_path": "/meu_ponto/registro_de_ponto",
		"time_card": {
			"accuracy": 600,
			"accuracy_method": true,
			"address": "%s",
			"latitude": "%s",
			"longitude": "%s",
			"location_edited": false,
			"original_address": "%s",
			"original_latitude": "%s",
			"original_longitude": "%s",
			"reference_id": null
		}}`, configuration.Address, configuration.Latitude, configuration.Longitude, configuration.Address, configuration.Latitude, configuration.Longitude)

	request := gorequest.New()

	response, _, errs := request.Post(hitUrl).
		Set("Authority", "api.pontomais.com.br").
		Set("Accept", "application/json. text/plain, */*").
		Set("Api-Version", "2").
		Set("Origin", "https://app.pontomaisweb.com.br").
		Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MOB31K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.106 Mobile Safari/537.36").
		Set("Content-Type", "application/json;charset=UTF-8").
		Set("X-Requested-With", "br.com.pontomais.pontomais").
		Set("token-type", "Bearer").
		Set("uid", credentials.Login).
		Set("access-token", credentials.Token).
		Set("client", credentials.ClientID).
		Set("uuid", uuid.New().String()).
		TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Send(requestData).
		End()

	errorStatusCodes := []int{
		400, 403, 404, 401,
	}
	if len(errs) > 0 || slices.Contains(errorStatusCodes, response.StatusCode) {
		log.Println("[ERRO] Falhou para registrar o ponto.")
	}

	if response.StatusCode == 201 || response.StatusCode == 202 {
		fmt.Println("[SUCESSO] Ponto registrado com sucesso!")
	}
}
