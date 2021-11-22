package models

type User struct {
	UserID   int64  `db:"user_id" json:"user_id"`
	UserName string `db:"username" json:"user_name"`
	Password string `db:"password" json:"password"`
}
