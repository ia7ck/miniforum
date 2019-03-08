package model

import "time"

// Post はthreadにくっつく
type Post struct {
	ID        int
	Body      string
	User      User
	Thread    Thread
	CreatedAt time.Time
}

// CreatePost はuserがthreadにpostをつくる（内容はbody）
func (user *User) CreatePost(thread *Thread, body string) (post Post, err error) {
	err = DB.QueryRow(`INSERT INTO posts
		(body, user_id, thread_id, created_at)
		VALUES
		($1, $2, $3, $4)
		RETURNING id
	`, body, user.ID, thread.ID, time.Now()).Scan(&post.ID)
	return
}

// GetPostsByThread はthreadに投稿されたpostsを得る
func (thread *Thread) GetPostsByThread() (posts []Post, err error) {
	rows, err := DB.Query(`SELECT
		posts.id, body, user_id, thread_id, created_at,
		users.screen_name, users.password
		FROM posts INNER JOIN users ON posts.user_id = users.id
		WHERE thread_id = $1
	`, thread.ID)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{Thread: *thread} // ポインタのほうがいいかも
		err = rows.Scan(&post.ID, &post.Body, &post.User.ID, &post.Thread.ID, &post.CreatedAt,
			&post.User.ScreenName, &post.User.Password)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}
