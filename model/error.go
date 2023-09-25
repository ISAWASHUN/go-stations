package model

import (
	"fmt"
	"time"
)

type ErrNotFound struct {
	When time.Time
	What string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("at %v, %s not found",
		e.When, e.What)
}


func NewErrNotFound(what string) *ErrNotFound {
	return &ErrNotFound{
		When: time.Now(),
		What: what,
	}
}