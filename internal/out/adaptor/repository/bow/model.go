package bow

import (
	"time"

	"github.com/zippoxer/bow"
)

type TodoModel struct {
	Id          bow.Id
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
}
