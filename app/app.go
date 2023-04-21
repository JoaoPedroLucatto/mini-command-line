package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Vai gerar e retorna a aplicação de linha de comando pronta a ser executada
// exemplo: go run main.go ip --host amazon.com.br
// exemplo2: go run main.go server --host amazon.com.br
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de comando"
	app.Usage = "Buscar Ips e nomes  de servidores  na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "amazon.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Buscando Ips de endereço de Internet",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "server",
			Usage:  "Buscando servidores",
			Flags:  flags,
			Action: findNs,
		},
	}
	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Printf("Erro find Ips: %s", err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findNs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupNS(host)
	if err != nil {
		log.Printf("Erro find Ns: %s", err)
	}

	for _, ip := range ips {
		fmt.Println(ip.Host)
	}
}
