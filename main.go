package main

import "go-fist/client"

func main() {
	client, _ := client.NewFistClient("localhost", "5575")
	client.Exit()
}
