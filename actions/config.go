package actions

import (
	"bufio"
	"main/helpers"
	"strings"

	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func Configure(c *cli.Context) {
	login := c.String("login")
	password := c.String("password")

	address, latitude, longitude := getLocationData()
	configFolderPath := helpers.GetConfigFolderPath()
	configFilePath := configFolderPath + "\\config.txt"

	if _, err := os.Stat(configFolderPath); os.IsNotExist(err) {
		fmt.Println("Criando diretório...")
		helpers.CreateDirectory(configFolderPath)
	}

	lines := []string{
		login, password, address, latitude, longitude,
	}
	helpers.CreateFile(configFilePath, lines)
}

func getLocationData() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Por favor, informe as seguintes informações:")
	fmt.Print(" - Endereço: ")

	address, err_address := reader.ReadString('\n')
	if err_address != nil {
		log.Fatal(err_address)
	}

	fmt.Print(" - Latitude: ")
	latitude, err_latitude := reader.ReadString('\n')
	if err_latitude != nil {
		log.Fatal(err_latitude)
	}

	fmt.Print(" - Longitude: ")
	longitude, err_longitude := reader.ReadString('\n')
	if err_longitude != nil {
		log.Fatal(err_longitude)
	}

	address = strings.Replace(address, "\n", "", -1)
	latitude = strings.Replace(latitude, "\n", "", -1)
	longitude = strings.Replace(longitude, "\n", "", -1)
	return address, latitude, longitude
}
