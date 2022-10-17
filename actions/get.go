package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"main/helpers"
	"main/models"

	"github.com/parnurzeal/gorequest"
	"github.com/urfave/cli"
)

// Busca as informações de registro de pontos do dia informado no formato yyyy-MM-dd
func Get(day string, c *cli.Context) {
	workdayUrl := helpers.URL + "/time_card_control/current/work_days/" + day

	configuration := helpers.ReadConfigurationFile()
	credentials, _ := helpers.Authenticate(models.Credentials{Login: configuration.Username, Password: configuration.Password})

	request := gorequest.New()

	_, body, errs := request.Get(workdayUrl).
		Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MOB31K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.106 Mobile Safari/537.36").
		Set("Content-Type", "application/json").
		Set("X-Requested-With", "br.com.pontomais.pontomais").
		Set("access-token", credentials.Token).
		Set("client", credentials.ClientID).
		Set("uid", configuration.Username).
		End()

	if len(errs) > 0 {
		log.Println("[ERRO] Falhou autenticando na api do pontomais.")
	}

	var times models.Times
	json.Unmarshal([]byte(body), &times)

	if len(times.Workday.TimeCards) == 0 {
		fmt.Println("Não existem registros de ponto para " + day)
		return
	}

	fmt.Println("Tempos registrados em " + day)
	fmt.Println("+---------------+---------------+---------------+")
	fmt.Println("| Registro      | Recibo        | Horário       |")
	fmt.Println("+---------------+---------------+---------------+")
	for i, time := range times.Workday.TimeCards {
		entryType := formatEntryType(i)
		line := fmt.Sprintf("| %-13s | %-13s | %-13s |", entryType, time.Receipt, time.Time)
		fmt.Println(line)
	}
	fmt.Println("+---------------+---------------+---------------+")
}

// Retornar o tipo do registro: Entrada ou Saída
func formatEntryType(index int) string {
	if index%2 == 0 {
		return "Entrada"
	} else {
		return "Saída"
	}
}
