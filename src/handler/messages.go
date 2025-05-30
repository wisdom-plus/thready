package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"thready/src/models"
    "thready/src/utils"
)

func HandleThreadShowOrPost(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/threads/")
    parts := strings.Split(idStr, "/")

    if len(parts) == 0 {
        http.NotFound(w, r)
        return
    }

    id, err := strconv.Atoi(parts[0])
    if err != nil {
        http.NotFound(w, r)
        return
    }

    thread, err := models.FindThreadByID(id)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    messages, err := models.FindMessageByThreadID(id)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    switch r.Method {
    case http.MethodGet:
        tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/threads/show.html"))
        tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
            "Thread":   thread,
            "Messages": messages,
        })
    case http.MethodPost:
        r.ParseForm()
        content := strings.TrimSpace(r.FormValue("content"))

        content, errMsg := utils.ValidateMessageContent(content)
        if errMsg != "" {
            tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/threads/show.html"))
            tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
                "Thread":   thread,
                "Messages": messages,
                "Error": errMsg,
            })
            return
        }

        _, err := models.CreateMessage(id, content)
        if err != nil {
            http.Error(w, "メッセージの作成に失敗しました", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/threads/"+strconv.Itoa(thread.ID), http.StatusSeeOther)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
	}
