package database

import (
	"GoStart/models"
	"fmt"
)

func GetAllUsers() ([]models.User, error) {
	query := `
	SELECT
		u.id AS id,
		u.username,
		u.password,
		u.confirmation_token,
		u.enabled,
		u.roles,

		up.id AS profile_id,
		up.second_id,
		up.nickname,
		up.created_at,
		up.last_activity_at,
		up.user_id
	FROM users u
	LEFT JOIN user_profile up ON u.id = up.user_id
	`

	rows, err := DB.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var profile models.UserProfile

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.ConfirmationToken,
			&user.Enabled,
			&user.Roles,

			&profile.ID,
			&profile.SecondID,
			&profile.Nickname,
			&profile.CreatedAt,
			&profile.LastActivityAt,
			&profile.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		user.Profile = profile
		users = append(users, user)
	}

	return users, nil
}
