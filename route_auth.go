package main

import (
	"net/http"
	"gowebprog/ch02/chitchat/data"
)

// GET /login
// Show the login page
func login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

// GET /logout
// Logs the user out
func logout(w http.ResponseWriter, r * http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		warning(err, "FAiled to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}