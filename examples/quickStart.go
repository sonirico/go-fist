package main

import (
	"fmt"
	fistClient "go-fist/client"
)

func main() {
	client, err := fistClient.NewFistClient("localhost", "5575")
	if err != nil {
		fmt.Println("Connection Error! Is Fist up and running?")
		return
	}
	// Obtain server version
	version, _ := client.Version()
	fmt.Println("Server version is " + version)
	// Index some data
	client.Index("articles", "a an the")
	client.Index("TODO", "wash the car")
	client.Index("TODO", "walk the dog")
	client.Index("podcasts", "DSE - Daily software engineering")
	// Search for "the" keyword
	documents := client.Search("the")
	fmt.Println(documents) // ["articles", "TODO"]
	// Not needing articles?
	client.Delete("the")
	documents = client.Search("the")
	fmt.Println(documents) // []
}
