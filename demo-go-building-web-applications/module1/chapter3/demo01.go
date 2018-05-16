package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	DBHost  = "db4free.net"
	DBPort  = ":3306"
	DBUser  = "testuser0912"
	DBPass  = "ycliu912"
	DBDbase = "cms001"
	PORT    = ":8080"
)

var database *sql.DB

type Page struct {
	Title   string
	Content string
	Date    string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	thisPage := Page{}
	fmt.Println(pageID)
	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE id=?", pageID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page: " + pageID)
		log.Println(err.Error)
	}
	html := `<html><head><title>` + thisPage.Title +
		`</title></head><body><h1>` + thisPage.Title +
		`</h1><div>` + thisPage.Content +
		`</div></body></title>`
	fmt.Fprintln(w, html)
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect!")
		log.Println(err.Error)
	} else {
		log.Println("Connect successfully!")
	}
	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/pages/{id:[0-9]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
}
