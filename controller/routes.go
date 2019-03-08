package controller

import (
	"miniforum/model"
	"net/http"
)

// Routes はルーティングを設定する
func Routes(mux *http.ServeMux) {
	mux.HandleFunc("/", index)

	mux.HandleFunc("/user/signup", signup)
	mux.HandleFunc("/user/create", createUser)
	mux.HandleFunc("/user/signin", signin)
	mux.HandleFunc("/user/auth", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/show", showThread)

	mux.HandleFunc("/thread/post/create", createPost)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // マッチしないやつはぜんぶルートに来るので
		http.NotFound(w, r)
		return
	}
	threads, err := model.GetAllThreads()
	if err != nil {
		warn(err)
	}
	err = executeTemplate(w, r, "templates/index.go.html", threads)
	if err != nil {
		warn(err)
	}
}
