package userrepository

import (
	"log"
)

func (i *sUserRepository) UpdatePassword(userID int64, password string) error {
	_, err := i.DB.Exec("UPDATE users SET password = ?, reset_token = NULL, reset_token_expires = NULL WHERE id = ?",
		password,
		userID,
	)

	if err != nil {
		log.Printf("Error updating user: %s", err)
		return err
	}

	return nil
}
