package main

import (
	"fmt"
	"linha-de-comando/code"
	"log"
	"os"
)

func main() {
	fmt.Println("Ips publicos do servidores !")

	application := code.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
