package models

type User struct {
	ID       string `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}
