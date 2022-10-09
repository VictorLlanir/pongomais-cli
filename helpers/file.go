package helpers

import (
	"bufio"
	"log"
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

// Cria o diretório informado
func CreateDirectory(directoryPath string) {
	if err := os.Mkdir(directoryPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// Criar o arquivo no caminho informado, gravando os dados
func CreateFile(filePath string, lines []string) {
	f, err := os.Create(filePath)
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
