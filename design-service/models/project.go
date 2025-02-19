package models

import "time"

type Project struct {
	ID          int64
	ClientID    string
	Title       string
	Description string
	Budget      float64
	Status      int8
	Deadline    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
