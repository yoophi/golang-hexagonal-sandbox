package entity

import "time"

type Todo struct {
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
}
