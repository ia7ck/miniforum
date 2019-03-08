package model

// User はthreadを作ったりpostをしたりできる
type User struct {
	ID         int
	ScreenName string
	Password   string // hash化したもの！
}

// Create はuserを作る
func (user *User) Create() (err error) {
	err = DB.QueryRow(`INSERT INTO users
		(screen_name, password)
		VALUES
		($1, $2)
		RETURNING id
	`, user.ScreenName, hash(user.Password)).Scan(&user.ID)
	return
}

// GetUserByScreenName はscreen_nameからuserを得る
func GetUserByScreenName(screenName string) (user User, err error) {
	err = DB.QueryRow(`SELECT
		id, screen_name, password
		FROM users
	`).Scan(&user.ID, &user.ScreenName, &user.Password)
	return
}

// GetUserBySessionID はsession_idからuserを得る
func GetUserBySessionID(sessionID string) (user User, err error) {
	err = DB.QueryRow(`SELECT
		users.id, screen_name, password
		FROM users INNER JOIN sessions ON users.id = sessions.user_id
		WHERE sessions.session_id = $1
`, sessionID).Scan(&user.ID, &user.ScreenName, &user.Password)
	return
}
