package controller

import (
	"net/http"

	"github.com/ia7ck/miniforum/model"
	"golang.org/x/crypto/bcrypt"
)

func signup(w http.ResponseWriter, r *http.Request) {
	err := executeTemplate(w, r, "templates/signup.go.html", nil)
	if err != nil {
		warn(err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{
		ScreenName: r.PostFormValue("screen_name"), // keyが存在しない場合は空文字列になる
		Password:   r.PostFormValue("password"),
	}
	err := user.Create()
	if err != nil {
		warn(err)
		http.Redirect(w, r, "/user/signup", 302) // 失敗したらもう一回登録してもらう
	} else {
		err = setSessionCookie(w, r, &user)
		if err != nil {
			warn(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	err := executeTemplate(w, r, "templates/signin.go.html", nil)
	if err != nil {
		warn(err)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	user, err := model.GetUserByScreenName(r.PostFormValue("screen_name"))
	if err != nil {
		warn(err)
		http.Redirect(w, r, "/user/signin", 302)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.PostFormValue("password"))); err == nil {
		err = setSessionCookie(w, r, &user)
		if err != nil {
			warn(err)
		}
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/user/signin", 302)
	}
}
