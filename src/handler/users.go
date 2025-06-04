package handler

import (
	"net/http"
	"html/template"
	"strconv"
	"strings"

	"thready/src/models"
	"thready/src/utils"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
		case http.MethodPost:
			r.ParseForm()
			username := strings.TrimSpace(r.FormValue("username"))
			password := strings.TrimSpace(r.FormValue("password"))
			username, errMsg := utils.ValidateUsername(username)
			if errMsg != "" {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/new.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   errMsg,
					"Username": username,
					"Password": password,
				})
				return
			}
			password, errMsg := utils.ValidatePassword(password)
			if errMsg != "" {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/new.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   errMsg,
					"Username": username,
					"Password": password,
				})
				return
			}
			id, err := models.CreateUser(username, password)
			if err != nil {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/new.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   "ユーザーの作成に失敗しました",
					"Username": username,
					"Password": password,
				})
				return
			}
			http.Redirect(w, r, "/mypage", http.StatusSeeOther)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlerUserNew(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/new.html"))
		tpl.ExecuteTemplate(w, "layout", nil)
}

func HandleMyPage(w http.ResponseWriter, r *http.Request) {
	// ユーザー情報を取得
	user, err := models.GetCurrentUser(r)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/mypage.html"))
	tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
		"User": user,
	})
}
