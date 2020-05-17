package main

import (
	"net/http"
)

func main() {
	p("My ChitChat", version(), "started at", config.Address)

	mux := http.NewServeMux()
	// @todo: http.Dir is string, but set as interface that has Open() function in fs.go. why?
	files := http.FileServer(http.Dir(config.Static))
}
