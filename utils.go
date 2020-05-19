package main

import (
	"errors"
	"fmt"
	"gowebprog/ch02/mychitchatexercise/data"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Address string
	Static  string
}

var config Configuration
var logger *log.Logger

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
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING")
	logger.Println(args...)
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

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.gohtml", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.gohtml", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)

}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR")
	logger.Println(args...)
}
