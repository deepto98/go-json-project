package main

func main() {
	server := newAPIServer(":8000")
	server.Run()
}
