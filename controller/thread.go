package controller

import (
	"miniforum/model"
	"net/http"
	"strconv"
)

func newThread(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err := executeTemplate(w, r, "templates/new.go.html", nil)
	if err != nil {
		warn(err)
	}
}

func createThread(w http.ResponseWriter, r *http.Request) {
	sessionID, err := loadSessionCookie(r)
	if err != nil {
		warn(err)
		http.Redirect(w, r, "/signin", 302) // ログインしていないユーザ
		return
	}
	user, err := model.GetUserBySessionID(sessionID)
	if err != nil {
		warn(err)
	}
	threadID, err := user.CreateThread(r.PostFormValue("content"))
	if err != nil {
		warn(err)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/thread/show?id="+strconv.Itoa(threadID), 302)
	}
}

func showThread(w http.ResponseWriter, r *http.Request) {
	threadID, _ := strconv.Atoi(r.FormValue("id")) // idが数字以外だとAtoiしたら0になる
	thread, err := model.GetThreadByID(threadID)   // threadはnilにならないみたい
	if err != nil {
		warn(err)
		err = executeTemplate(w, r, "templates/show.go.html", nil) // 明示的にnilを渡す
	} else {
		err = executeTemplate(w, r, "templates/show.go.html", &thread)
	}
	if err != nil {
		warn(err)
	}
}
