package controller

import (
	"net/http"
	"strconv"

	"github.com/ia7ck/miniforum/model"
)

func createPost(w http.ResponseWriter, r *http.Request) {
	sessionID, err := loadSessionCookie(r)
	if err != nil {
		warn(err)
		http.Redirect(w, r, "/signin", 302)
		return
	}
	user, err := model.GetUserBySessionID(sessionID)
	if err != nil {
		warn(err)
	}
	threadID, _ := strconv.Atoi(r.PostFormValue("thread_id"))
	thread, err := model.GetThreadByID(threadID)
	if err != nil {
		warn(err)
	}
	_, err = user.CreatePost(&thread, r.PostFormValue("body"))
	if err != nil {
		warn(err)
	}
	http.Redirect(w, r, "/thread/show?id="+strconv.Itoa(threadID), 302)
}
