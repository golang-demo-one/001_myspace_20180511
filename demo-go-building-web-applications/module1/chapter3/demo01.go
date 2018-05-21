package main

import (
	"database/sql"
	"fmt"
	"html/template"
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
	Title string
	//Content string
	RawContent string
	Content    template.HTML
	Date       string
	GUID       string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//pageID := vars["id"]
	pageGUID := vars["guid"]
	thisPage := Page{}
	//fmt.Println(pageID)
	fmt.Println(pageGUID)
	//err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE id=?", pageID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)
	err := database.QueryRow("SELECT page_title,page_content,page_date,page_guid FROM pages WHERE page_guid=?", pageGUID).Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date, &thisPage.GUID)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		//log.Println("Couldn't get page: " + pageID)
		log.Println("Couldn't get page: " + pageGUID)
		log.Println(err.Error)
	}
	thisPage.Content = template.HTML(thisPage.RawContent)
	/*
		html := `<html><head><title>` + thisPage.Title +
			`</title></head><body><h1>` + thisPage.Title +
			`</h1><div>` + thisPage.Content +
			`</div></body></title>`
		fmt.Fprintln(w, html)
	*/
	t, _ := template.ParseFiles("templates/blog.html")
	t.Execute(w, thisPage)

}

func RedirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", 301)
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	var Pages = []Page{}
	pages, err := database.Query("SELECT page_title,page_content,page_date,page_guid FROM pages ORDER BY ? DESC", "page_date")
	if err != nil {
		fmt.Fprintln(w, err.Error)
	} else {
		fmt.Println(*pages)
	}
	defer pages.Close()
	for pages.Next() {
		thisPage := Page{}
		pages.Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date, &thisPage.GUID)
		thisPage.Content = template.HTML(thisPage.RawContent)
		Pages = append(Pages, thisPage)
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, Pages)
}

func (p Page) TruncatedText() template.HTML {
	chars := 0
	for i, _ := range p.RawContent {
		chars++
		if chars > 150 {
			p.Content = template.HTML(p.RawContent[:i] + ` ...`)
			return p.Content
		}
	}
	p.Content = template.HTML(p.RawContent)
	return p.Content
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
	//routes.HandleFunc("/pages/{id:[0-9]+}", ServePage)
	routes.HandleFunc("/pages/{guid:[0-9a-zA-Z\\-]+}", ServePage)
	routes.HandleFunc("/", RedirIndex)
	routes.HandleFunc("/home", ServeIndex)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
}
