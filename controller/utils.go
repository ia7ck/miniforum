package controller

import (
	"html/template"
	"log"
	"miniforum/model"
	"net/http"
	"strings"
	"time"
)

func warn(args ...interface{}) {
	log.Print(args)
}

func executeTemplate(w http.ResponseWriter, r *http.Request, file string, data interface{}) (err error) {
	// ファイルたちを解析する 失敗したらpanic
	// define "layout" した部分がExecuteTemplateで実行される
	t := template.Must(template.New("layout").Funcs(funcMap).ParseFiles("templates/layout.go.html", file))
	var signedIn bool
	if _, err := loadSessionCookie(r); err != nil {
		signedIn = false
	} else {
		signedIn = true
	}
	err = t.ExecuteTemplate(w, "layout", struct {
		Data     interface{}
		SignedIn bool
	}{
		Data:     data,
		SignedIn: signedIn,
	})
	return
}

// requestに含まれるcookieを読んで、（もしあるなら）sessionIDを返す
func loadSessionCookie(r *http.Request) (sessionID string, err error) {
	cookie, err := r.Cookie("miniforum_cookie")
	if err != nil {
		return
	}
	sessionID = cookie.Value
	return
}

// 認証用のcookieを設定する
func setSessionCookie(w http.ResponseWriter, r *http.Request, user *model.User) (err error) {
	session, err := user.CreateSession()
	if err != nil {
		return
	}
	cookie := http.Cookie{ // session cookie (ブラウザを閉じたら消える)
		Name:     "miniforum_cookie",
		Value:    session.SessionID,
		HttpOnly: true, // javascriptから見れないようにする
		Path:     "/",  // これ要る
	}
	http.SetCookie(w, &cookie)
	return
}

var funcMap = template.FuncMap{
	"newLineToBreak": func(str string) template.HTML {
		// あらかじめescapeする！
		return template.HTML(strings.Replace(template.HTMLEscapeString(str), "\n", "<br>", -1))
	},
	"dateFormat": func(tm time.Time) string {
		return tm.Format("2006-01-02 03:04")
	},
}
