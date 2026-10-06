package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port, err := serverPort()
	if err != nil {
		log.Fatal(err)
	}

	server := &Server{Port: port}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	Port string
}

func (s *Server) Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		fmt.Fprintln(w, "Hello from Go server!")
	})

	log.Printf("Starting server on %s...", s.Port)

	return http.ListenAndServe(s.Port, nil)
}

func serverPort() (string, error) {
	inputPort := flag.String("p", "8080", "The server port")
	flag.Parse()

	port, err := strconv.Atoi(*inputPort)
	if err != nil {
		return "", fmt.Errorf("port must be a number: %q", *inputPort)
	}

	if port < 1 || port > 65535 {
		return "", fmt.Errorf("port must be between 1 and 65535: %d", port)
	}

	return ":" + strconv.Itoa(port), nil
}
