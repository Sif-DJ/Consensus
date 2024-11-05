package main

func main() {

	go startServer("5050", "5051")
	go client.startServer("5051", "5052")
	go client.startServer("5052", "5050")

	for {

	}

}
