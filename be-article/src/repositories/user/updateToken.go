package userrepository

import (
	"log"
	"time"
)

func (i *sUserRepository) UpdateResetToken(userID int64, token string, expiry time.Time) error {
	_, err := i.DB.Exec("UPDATE users SET reset_token = ?, reset_token_expires = ? WHERE id = ?",
		token,
		expiry,
		userID,
	)

	if err != nil {
		log.Printf("Error updating user: %s", err)
		return err
	}

	return nil
}
