package models

import (
	"time"
)

type UserProfile struct {
	ID             string    `db:"profile_id"`
	SecondID       string    `db:"second_id"`
	Nickname       string    `db:"nickname"`
	CreatedAt      time.Time `db:"created_at"`
	LastActivityAt time.Time `db:"last_activity_at"`
	UserID         string    `db:"user_id"`
}
