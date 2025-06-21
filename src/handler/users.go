package handler

import (
	"net/http"
	"html/template"
	"strconv"
	"strings"

	"thready/src/models"
	"thready/src/utils"
)

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/sign_up.html"))
			tpl.ExecuteTemplate(w, "layout", nil)
		case http.MethodPost:
			r.ParseForm()
			username := strings.TrimSpace(r.FormValue("username"))
			password := strings.TrimSpace(r.FormValue("password"))
			username, userErrMsg := utils.ValidateUsername(username)
			if userErrMsg != "" {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/sign_up.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   userErrMsg,
					"Username": username,
					"Password": password,
				})
				return
			}
			password, passwordErrMsg := utils.ValidatePassword(password)
			if passwordErrMsg != "" {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/sign_up.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   passwordErrMsg,
					"Username": username,
					"Password": password,
				})
				return
			}
			user_id, err := models.CreateUser(username, password)
			if err != nil {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/sign_up.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   "ユーザーの作成に失敗しました",
					"Username": username,
					"Password": password,
				})
				return
			}
			cookie := &http.Cookie{
				Name:  "user_id",
				Value: strconv.Itoa(user_id),
				Path:  "/",
				HttpOnly: true,
				MaxAge:  86400,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/mypage", http.StatusSeeOther)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleMyPage(w http.ResponseWriter, r *http.Request) {
	// ユーザー情報を取得
	user_id, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(user_id.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	user, err := models.GetCurrentUser(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/mypage.html"))
	tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
		"User": user,
	})
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/sign_in.html"))
			tpl.ExecuteTemplate(w, "layout", nil)
		case http.MethodPost:
			r.ParseForm()
			username := strings.TrimSpace(r.FormValue("username"))
			password := strings.TrimSpace(r.FormValue("password"))
			username, userErrMsg := utils.ValidateUsername(username)
			if userErrMsg != "" {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/sign_in.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   userErrMsg,
					"Username": username,
					"Password": password,
				})
				return
			}
			password, passwordErrMsg := utils.ValidatePassword(password)
			if passwordErrMsg != "" {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/sign_in.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   passwordErrMsg,
					"Username": username,
					"Password": password,
				})
				return
			}
			hashPassword, hashErr := utils.HashPassword(password)
			if hashErr != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			user, err := models.FindUserByLogin(username, hashPassword)
			if err != nil {
				tpl := template.Must(template.ParseFiles("templates/layout.html", "templates/users/sign_in.html"))
				tpl.ExecuteTemplate(w, "layout", map[string]interface{}{
					"Error":   "ユーザー名またはパスワードが間違っています",
					"Username": username,
					"Password": password,
				})
				return
			}
			cookie := &http.Cookie{
				Name:  "user_id",
				Value: strconv.Itoa(user.ID),
				Path:  "/",
				HttpOnly: true,
				MaxAge:  86400,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/mypage", http.StatusSeeOther)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		cookie := &http.Cookie{
				Name:   "user_id",
				Value:  "",
				Path:   "/",
				MaxAge: -1,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
