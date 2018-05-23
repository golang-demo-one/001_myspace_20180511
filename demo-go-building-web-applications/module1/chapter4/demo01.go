package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strconv"
	//"strconv"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	//DBHost = "db4free.net"
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "testuser0912"
	DBPass  = "ycliu912"
	DBDbase = "cms001"
	PORT    = ":8080"
)

var database *sql.DB

type Page struct {
	Id    int
	Title string
	//Content string
	RawContent string
	Content    template.HTML
	Date       string
	Comments   []Comment
	//Session    Session
	GUID string
}

type Comment struct {
	Id          int
	Guid        string
	Name        string
	Email       string
	CommentText string
	Date        string
}

type JSONResponse struct {
	Fields map[string]string
}

func New() *JSONResponse {
	return &JSONResponse{
		Fields: map[string]string{},
	}
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//pageID := vars["id"]
	pageGUID := vars["guid"]

	//pageName := vars["name"]
	//pageEmail := vars["email"]
	//pageCommentText := vars["comments"]

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

	comments, err := database.Query("SELECT id, comment_guid, comment_name, comment_email, comment_text, comment_date FROM comments WHERE comment_guid=? ORDER BY comment_date DESC", pageGUID)
	if err != nil {
		log.Println(err.Error)
	}
	for comments.Next() {
		var comment Comment
		comments.Scan(&comment.Id, &comment.Guid, &comment.Name, &comment.Email, &comment.CommentText, &comment.Date)
		thisPage.Comments = append(thisPage.Comments, comment)
	}

	/*
		commentsEdit, err := database.Exec("UPDATE comments SET comment_name=?, comment_email=?, comment_text=? WHERE comment_guid=?", pageName, pageEmail, pageCommentText, pageGUID)
		if err != nil {
			log.Println(err.Error)
		}
		for commentsEdit.Next() {
			var comment_edit comment
			commentsEdit.
		}
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

func APIPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}
	fmt.Println(pageGUID)
	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?", pageGUID).Scan(
		&thisPage.Title,
		&thisPage.RawContent,
		&thisPage.Date)
	thisPage.Content = template.HTML(thisPage.RawContent)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}
	APIOutput, err := json.Marshal(thisPage)
	fmt.Println(APIOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, thisPage)
}

func APICommentPost(w http.ResponseWriter, r *http.Request) {
	var commentAdded string
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}

	//fmt.Fprintln(r)
	fmt.Println(r.Context())
	fmt.Println(r.Form)
	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")
	guid := r.FormValue("guid")
	fmt.Println(name, email, comments, guid)

	res, err := database.Exec("INSERT INTO comments SET comment_name=?, comment_email=?, comment_text=?, comment_guid=?", name, email, comments, guid)
	if err != nil {
		log.Println(err.Error)
	}

	id, err := res.LastInsertId()
	if err != nil {
		commentAdded = "false"
	} else {
		commentAdded = "true"
	}

	resp := New()
	fmt.Println(id)
	fmt.Println(strconv.FormatInt(id, 10))
	id_str := strconv.FormatInt(id, 10)
	(*resp).Fields["id"] = id_str
	(*resp).Fields["added"] = commentAdded

	jsonResp, _ := json.Marshal(*resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)
	fmt.Println("==================================================")
	fmt.Println(resp)
	fmt.Println(jsonResp)
}

func APICommentPut(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	} else {
		fmt.Println("r.Form: ", r.Form)
	}
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")
	fmt.Println("vars: ", vars)

	res, err := database.Exec("UPDATE comments SET comment_name=?, comment_email=?, comment_text=? WHERE comment_id=?", name, email, comments, id)
	fmt.Println(res)
	if err != nil {
		log.Println(err.Error)
	}

	resp := New()
	jsonResp, _ := json.Marshal(*resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)

	fmt.Println(*resp, jsonResp)
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

	//routes.HandleFunc("/api/pages", APIPage).Methods("GET").Schemes("https")
	//routes.HandleFunc("/api/pages`/{guid:[0-9a-zA\\-]+}", APIPage).Methods("GET").Schemes("https")

	routes.HandleFunc("/api/pages", APIPage).Methods("GET")
	routes.HandleFunc("/api/pages/{guid:[0-9a-zA\\-]+}", APIPage).Methods("GET")
	routes.HandleFunc("/api/comments", APICommentPost).Methods("POST")
	//routes.HandleFunc("/api/comments/{guid:[0-9a-zA\\-]+}", APICommentPut).Methods("PUT")
	routes.HandleFunc("/api/comments/{id:[0-9a-zA\\-]+}", APICommentPut).Methods("PUT")

	//routes.HandleFunc("/pages/{id:[0-9]+}", ServePage)
	routes.HandleFunc("/pages/{guid:[0-9a-zA-Z\\-]+}", ServePage)
	routes.HandleFunc("/", RedirIndex)
	routes.HandleFunc("/home", ServeIndex)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
}
