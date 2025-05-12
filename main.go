package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"fmt"
	"os"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pg"
)

var db *sqlx.DB

type Thread struct {
	ID int
	Title string
}

type Message struct {
	Content string
}

var mockThreads = []Thread{
	{ID: 1, Title: "好きな映画について"},
	{ID: 2, Title: "最近読んだ本"},
	{ID: 3, Title: "Go言語の面白さ"},
}

func initDB() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sqlx.Connect("postgres", dsn)
	return err
}

func main() {
	initDB()
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/threads", handleThreads)
	http.HandleFunc("/threads/new", handleThreadNew)
	port := "8080"	
  log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func handleThreads(w http.ResponseWriter , r*http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleThreadsIndex(w, r)
	case http.MethodPost:
		handleThreadCreate(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleThreadsIndex(w http.ResponseWriter , r*http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/threads/index.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Threads": mockThreads,
	})
}

func handleThreadShow(w http.ResponseWriter , r*http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/threads/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	
	var thread *Thread
	for _, t := range mockThreads {
		if t.ID == id {
			thread = &t
			break
		}
	}

	if thread == nil {
		http.NotFound(w, r)
		return
	}

	messages := []Message{
		{Content: "こんにちは!"},
		{Content: "その映画、私は好きです!"},
	}
	tmpl := template.Must(template.ParseFiles("templates/threads/show.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Threads": thread,
		"Messages": messages,
	})
}

func handleThreadNew(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/threads/new.html"))
	tmpl.Execute(w, nil)
}

func handleThreadCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "フォームの読み取りに失敗しました。", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(r.FormValue("title"))
	if title == "" {
		tmpl := template.Must(template.ParseFiles("templates/threads/new.html"))
		tmpl.Execute(w, map[string]interface{}{
			"Error": "タイトルが入力がされていません。",
			"Title": title,
		})
	}

	if len(title) > 255 {
		tmpl := template.Must(template.ParseFiles("templates/threads/new.html"))
		tmpl.Execute(w, map[string]interface{}{
			"Error": "タイトルは255文字以内で入力してください",
			"Title": title,
		})
	}

	newID := 1
	if len(mockThreads) > 0 {
		newID = mockThreads[len(mockThreads)-1].ID + 1
	}
	newThread := Thread{ID: newID, Title: title}
	mockThreads = append(mockThreads, newThread)
	
	http.Redirect(w, r, "/threads/"+strconv.Itoa(newID), http.StatusSeeOther)

}
