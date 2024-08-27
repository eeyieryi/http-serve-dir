package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port string
	dir  string
)

func init() {
	const (
		defaultPort string = "8000"
		usagePort   string = "the port of the http server"
	)
	flag.StringVar(&port, "port", defaultPort, usagePort)
	flag.StringVar(&port, "p", defaultPort, usagePort)

	const (
		defaultDir string = "." // current working directory
		usageDir   string = "the directory to serve over http"
	)
	flag.StringVar(&dir, "dir", defaultDir, usageDir)
	flag.StringVar(&dir, "d", defaultDir, usageDir)
}

func main() {
	handler := http.FileServer(http.Dir(dir))
	srv := http.Server{
		Addr:    "localhost:" + port,
		Handler: handler,
	}
	log.Printf("Serving dir %s on %s ...", dir, srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
