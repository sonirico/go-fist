package main

import (
	"fmt"
	"go-fist/client"
)

func main() {
	fistClient, _ := client.NewFistClient("localhost", "5575")
	fistClient.Index("saludos", "hola que tal")
	fistClient.Index("cugnadeces", "yo lo tal compre mas barato en el concesionario de mi barrio")
	documents := fistClient.Search("tal")
	if documents != nil {
		fmt.Println(documents)
	}
	fistClient.Exit()
}
