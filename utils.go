package main

import (
	"fmt"
	"net/http"
	"strings"
	"errors"
	"html/template"
	"gowebprog/ch02/mychitchatexercise/data"
)

type Configuration struct {
	Address string
	Static  string
}

var config Configuration

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a)
}

// init in package
func init() {
	config = Configuration{
		Address: ":8080",
		// @note: NG: "/public", OK: "public"
		Static: "public",
	}
}

// version
func version() string {
	return "0.1"
}

// Convenience function to redirect to the error message page
func error_message(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), http.StatusFound)
}

// Checks if the user is logged in and has a session, if not err is not nil
func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)

}
