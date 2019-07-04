package main

import (
	"fmt"
	fistclient "github.com/sonirico/go-fist/client"
)

func main() {
	client, _ := fistclient.NewFistClient("localhost", "5575")
	// Obtain server version
	version, _ := client.Version()
	fmt.Println("Server version is " + version)
	// Index some data
	client.Index("todo", "wash the car")
	client.Index("todo", "walk the dog")
	client.Index("podcasts", "DSE - Daily software engineering")
	// Search for it
	documents := client.Search("the")
	fmt.Println(documents) // ["todo"]
}
