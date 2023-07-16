package user

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Username  string
	Password  string
	ProfileId int
	CreatedAt time.Time
	UpdatedAt time.Time
}