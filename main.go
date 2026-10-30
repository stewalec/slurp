package main

import (
	"log"
	"net/http"
)

func main() {
	s := New()

	http.HandleFunc("GET /{$}", s.indexHandler)
	http.HandleFunc("GET /{username}", s.userHandler)
	http.HandleFunc("GET /static/{file}", s.staticHandler)
	http.HandleFunc("GET /stir", s.stirHandler)
	http.HandleFunc("POST /stir", s.stirHandler)
	http.HandleFunc("GET /feeds", s.settingsHandler)
	http.HandleFunc("POST /feeds/submit", s.settingsSubmitHandler)
	http.HandleFunc("GET /login", s.loginHandler)
	http.HandleFunc("POST /login", s.loginHandler)
	http.HandleFunc("GET /logout", s.logoutHandler)
	http.HandleFunc("POST /logout", s.logoutHandler)
	http.HandleFunc("POST /register", s.registerHandler)
	http.HandleFunc("GET /saved", s.userSavedHandler)
	http.HandleFunc("POST /save/{url}", s.saveHandler)
	http.HandleFunc("POST /archive/{id}", s.archiveHandler)
	http.HandleFunc("GET /read/{url}", s.readHandler)
	http.HandleFunc("POST /read/{url}", s.readHandler)
	http.HandleFunc("GET /feeds/{url}", s.feedDetailsHandler)

	// backwards compatibility redirects
	http.HandleFunc("GET /settings", s.settingsRedirectHandler)
	http.HandleFunc("POST /settings/submit", s.settingsSubmitRedirectHandler)

	log.Println("main: listening on http://localhost:5544")
	log.Fatal(http.ListenAndServe(":5544", nil))
}
