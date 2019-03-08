package model

import (
	"time"

	"github.com/google/uuid"
)

// Session はログイン状態を判断するのに使う
type Session struct {
	ID        int
	SessionID string // UUID
	User      User
	CreatedAt time.Time
}

// CreateSession はsessionをつくる
func (user *User) CreateSession() (session Session, err error) {
	err = DB.QueryRow(`INSERT INTO sessions
		(session_id, user_id, created_at)
		VALUES
		($1, $2, $3)
		RETURNING id, session_id, user_id, created_at
	`, uuid.Must(uuid.NewRandom()), user.ID, time.Now()).Scan(&session.ID, &session.SessionID, &session.User.ID, &session.CreatedAt)
	return
}
