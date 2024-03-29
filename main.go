package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	p("My ChitChat", version(), "started at", config.Address)

	mux := http.NewServeMux()
	// @todo: http.Dir is string, but set as interface that has Open() function in fs.go. why?
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// all route patterns matched here
	// route handler functions defined in other files

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined int route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,

		// @note: NG
		//ReadTimeout: time.Duration(5),
		//WriteTimeout: time.Duration(5),
		// OK
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
