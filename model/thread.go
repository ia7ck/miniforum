package model

import "time"

// Thread はpostが複数くっつくこともある
type Thread struct {
	ID        int
	Content   string
	User      User
	CreatedAt time.Time
}

// CreateThread userがthredを作る
func (user User) CreateThread(content string) (threadID int, err error) {
	err = DB.QueryRow(`INSERT INTO threads
		(content, user_id, created_at)
		VALUES
		($1, $2, $3)
		RETURNING id
	`, content, user.ID, time.Now()).Scan(&threadID)
	return
}

// GetAllThreads はthreadをぜんぶ取り出す
func GetAllThreads() (threads []Thread, err error) {
	rows, err := DB.Query(`SELECT
		threads.id, content, user_id, created_at,
		users.screen_name, users.password
		FROM threads INNER JOIN users ON user_id = users.id
		ORDER BY created_at DESC
	`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		thread := Thread{}
		err = rows.Scan(&thread.ID, &thread.Content, &thread.User.ID, &thread.CreatedAt,
			&thread.User.ScreenName, &thread.User.Password)
		if err != nil {
			return
		}
		threads = append(threads, thread)
	}
	return
}

// GetThreadByID はidからthreadを取得する
// 存在しなかったら sql.ErrNoRows を返す
func GetThreadByID(threadID int) (thread Thread, err error) {
	err = DB.QueryRow(`SELECT
		threads.id, content, user_id, created_at,
		users.screen_name, users.password
		FROM threads INNER JOIN users ON user_id = users.id
		WHERE threads.id = $1
	`, threadID).Scan(&thread.ID, &thread.Content, &thread.User.ID, &thread.CreatedAt,
		&thread.User.ScreenName, &thread.User.Password)
	return
}
