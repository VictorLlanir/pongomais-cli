package main

import (
	"log"
	"main/app"
	"os"
)

func main() {
	application := app.Configure()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
