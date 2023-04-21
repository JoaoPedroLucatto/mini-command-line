package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ips publicos do servidores !")

	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
