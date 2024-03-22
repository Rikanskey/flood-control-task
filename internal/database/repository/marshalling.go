package repository

import "time"

type UserQuery struct {
	UserId int64
	Time   time.Time `json:"time"`
	Tokens int       `json:"tokens"`
}
