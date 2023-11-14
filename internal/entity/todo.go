package entity

import "time"

type Todo struct {
	Title     string
	Completed bool
	Priority  int
	CretedAt  time.Time
}
