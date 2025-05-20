package models

type User struct {
	ID                string      `db:"id"`
	Username          string      `db:"username"`
	Password          string      `db:"password"`
	ConfirmationToken string      `db:"confirmation_token"`
	Enabled           bool        `db:"enabled"`
	Roles             string      `db:"roles"`
	Profile           UserProfile `db:"profile"`
}
