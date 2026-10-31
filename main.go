package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"

	"golang.org/x/term"
)

func main() {
	cmd := "serve"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "serve":
		s := New(false) // no cli mode
		runServer(s)
	case "add":
		s := New(true) // cli mode to skip favicon/reaper
		createUser(s, os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\nusage: slurp [serve|add]\n", cmd)
		os.Exit(2)
	}
}

func runServer(s *Site) {
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
	http.HandleFunc("DELETE /saved/{id}", s.saveHandler)
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

func createUser(s *Site, args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	
	var username string
	fs.StringVar(&username, "username", "", "username")
	fs.StringVar(&username, "u", "", "username (shorthand)")

	fs.Parse(args)
	if username == "" {
		log.Fatal("--username (or -u) is required")
	}

	fmt.Printf("Creating new user %s...\n", username)

	fmt.Fprint(os.Stderr, "Password: ")
	b1, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Fprint(os.Stderr, "Reenter password: ")
	b2, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	if !bytes.Equal(b1, b2) {
		log.Fatal("Passwords do not match. Aborting...")
	}

	if err := s.register(username, string(b1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User '%s' successfully created", username)
}