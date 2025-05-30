package handler

import (
	"net/http"
	"html/template"
	"strconv"
	"strings"

	"thready/src/models"
    "thready/src/utils"
)

func HandleThreads(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/threads/index.html"))
        threads, err := models.GetAllThreads()
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
        tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
            "Threads": threads,
        })
    case http.MethodPost:
        r.ParseForm()
        title := strings.TrimSpace(r.FormValue("title"))

        title, errMsg := utils.ValidateThreadTitle(title)
        if errMsg != "" {
            tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/threads/new.html"))
            tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
                "Error": errMsg,
                "Title": title,
            })
            return
        }

        id, err := models.CreateThread(title)
        if err != nil {
            http.Error(w, "スレッドの作成に失敗しました", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/threads/"+strconv.Itoa(id), http.StatusSeeOther)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// 新規スレッドフォーム
func HandleThreadNew(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/threads/new.html"))
    tpl.ExecuteTemplate(w, "layout", nil)
}
