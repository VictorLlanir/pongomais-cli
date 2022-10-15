package helpers

import (
	"bufio"
	"log"
	"main/models"
	"os"
	"runtime"
)

// Retorna o caminho para a pasta %APPDATA% no Windows e o caminho para a variável HOME no Linux
func GetConfigFolderPath() string {
	if runtime.GOOS == "windows" {
		home, _ := os.UserCacheDir()
		return home + "\\pongomais"
	}
	return os.Getenv("HOME")
}

func getConfigFilePath() string {
	return GetConfigFolderPath() + "\\config.txt"
}

// Cria o diretório informado
func CreateDirectory(directoryPath string) {
	if err := os.Mkdir(directoryPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// Criar o arquivo no caminho informado, gravando os dados
func CreateFile(filePath string, lines []string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	buffer := bufio.NewWriter(f)

	for _, line := range lines {
		if _, err := buffer.WriteString(line + "\n"); err != nil {
			log.Fatal(err)
		}
	}

	if err := buffer.Flush(); err != nil {
		log.Fatal(err)
	}
}

func ReadConfigurationFile() models.Configuration {
	file, err := os.Open(getConfigFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	configuration := models.Configuration{
		Username:  lines[0],
		Password:  lines[1],
		Address:   lines[2],
		Longitude: lines[3],
		Latitude:  lines[4],
	}

	return configuration
}
