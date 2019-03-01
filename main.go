package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// this is a global counter that gets incremented each time a new
// record is added to the database
var recordID int
var db *sql.DB
var m sync.Mutex
var templates = template.Must(template.ParseGlob("public/templates/*"))

//Page is used for templating html pages
type Page struct {
	Title  string
	Result string
}

func main() {
	db = initDb("mySqliteDb.sqlite3")
	defer db.Close()
	createTable(db)

	router := mux.NewRouter()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/about", aboutPage).Methods("GET")
	router.HandleFunc("/shorten", handleShorten).Methods("POST")
	router.HandleFunc("/s/{shortURL}", handleRedirect).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":7777", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	display(w, "main", Page{Title: "Home", Result: "test"})
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	display(w, "about", Page{Title: "About", Result: "test"})
}

func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	longURL := r.Form.Get("url")
	rtrn, err := shortenURL(longURL)
	if err != nil {
		w.Write([]byte("ruh roh"))
	}
	display(w, "result", Page{Title: "Result", Result: rtrn})
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortURL"]
	rec := readItem(db, shortURL)
	http.Redirect(w, r, rec.longURL, http.StatusSeeOther)
}
