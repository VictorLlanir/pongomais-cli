package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"main/models"

	"github.com/parnurzeal/gorequest"
)

func Authenticate(c models.Credentials) (error, models.Credentials) {
	url := "http://api.pontomaisweb.com.br/api/auth/sign_in"
	requestData := fmt.Sprintf(`{
		"email": "%s",
		"password": "%s"}`,
		c.Login, c.Password)
	request := gorequest.New()

	response, body, errs := request.Post(url).
		Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MOB31K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.106 Mobile Safari/537.36").
		Set("Content-Type", "application/json").
		Set("X-Requested-With", "br.com.pontomais.pontomais").
		Send(requestData).
		End()

	var err error
	if len(errs) > 0 {
		log.Println("[ERROR] Falhou autenticando na api do pontomais.")
		return err, c
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)

	if response.StatusCode > 201 || err != nil {
		msg := fmt.Sprintf("HttpStatus: %d. Email: %s. error: %s", response.StatusCode, c.Login, err)
		err = errors.New(msg)
		return err, c
	}

	c.Token = result["token"].(string)
	c.ClientID = result["client_id"].(string)

	fmt.Println(c.Token)
	fmt.Println(c.ClientID)
	return err, c
}
