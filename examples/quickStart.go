package main

import (
	"fmt"
	fistclient "go-fist/client"
)

func main() {
	client, err := fistclient.NewFistClient("localhost", "55750")
	if err != nil {
		fmt.Println("Connection Error! Is Fist up and running?")
		return
	}
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
