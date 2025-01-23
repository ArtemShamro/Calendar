package domain

import "time"

type Event struct {
	Id 				int			`db:"id"`
	Owner			string		`db:"owner"`
	CreatedAt		time.Time	`db:"created_at"`
	Name 			string		`db:"name"`
	Description 	string		`db:"description"`
	WhenOccurs		time.Time	`db:"when_occurs"`
	WhenRemind		time.Time	`db:"when_remind"`
}