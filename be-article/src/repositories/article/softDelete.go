package articlerepository

import (
	"log"
	"time"
)

func (i *sArticleRepository) SoftDelete(articleId *int) error {
	currentTime := time.Now()

	_, err := i.DB.Exec("UPDATE articles SET deleted_at = ? WHERE id = ?;", currentTime, articleId)

	if err != nil {
		log.Printf("Error soft deleting article: %s", err)
		return err
	}

	return nil
}
