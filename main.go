package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// this is a global counter that gets incremented each time a new
// record is added to the database
var recordID int
var db *sql.DB
var m sync.Mutex

func main() {
	db = initDb("mySqliteDb.sqlite3")
	defer db.Close()
	createTable(db)
	fmt.Println("Serving on 0.0.0.0:8080")

	r := mux.NewRouter()
	// serve a static page at 'http://domain.com/'
	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.HandleFunc("/shorten", handleShorten)
	r.HandleFunc("/s/{shortURL}", handleRedirect)
	http.ListenAndServe(":8080", r)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	rtrn := shortenURL(r.FormValue("url"))
	w.Write([]byte("localhost/s/" + rtrn))
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortURL"]
	rec := readItem(db, shortURL)
	fmt.Println(rec.longURL)
	http.Redirect(w, r, rec.longURL, http.StatusSeeOther)
}
