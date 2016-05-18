package main

import (
	"flag"
	"fmt"
	"net/http"
)

var serverDirectory = flag.String("dir", ".", "Specify a filesystem path for the server root")
var serverPort = flag.String("port", ":8080", "Specify an address to serve at")

func main() {
	root := http.Dir(serverDirectory)
	http.Handle("/", http.FileServer(root))

	if err := http.ListenAndServe(serverPort, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
