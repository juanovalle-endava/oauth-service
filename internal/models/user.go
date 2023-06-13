package models

type User struct {
	Id       string `db:"id"`
	UserId   string `db:"user_id" json:"userId"`
	Password string `db:"password" json:"password"`
}
