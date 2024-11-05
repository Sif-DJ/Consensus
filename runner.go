package main

import (
	client "Consensus/client"
)

func main() {

	go client.StartNode("5050", "5051", true)
	go client.StartNode("5051", "5052", false)
	go client.StartNode("5052", "5050", false)

	for {

	}

}
