package models

import (
	"fmt"
	"log"
	"time"

	"github.com/Murodkadirkhanoff/uiux-design-service/db"
)

type Design struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	CategoryID    *int64    `json:"category_id,omitempty"`
	Tags          *string   `json:"tags,omitempty"`
	Visibility    string    `json:"visibility"`
	Status        string    `json:"status"`
	ViewsCount    int64     `json:"views_count"`
	LikesCount    int64     `json:"likes_count"`
	CommentsCount int64     `json:"comments_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func GetAllProjects() ([]Design, error) {
	rows, err := db.DB.Query("SELECT * FROM designs")
	if err != nil {
		log.Fatal("Could not fetch designs:", err.Error())
	}
	defer rows.Close()

	var designs []Design

	for rows.Next() {
		var design Design
		err := rows.Scan(
			&design.ID,
			&design.UserID,
			&design.Title,
			&design.Slug,
			&design.Description,
			&design.CategoryID,
			&design.Tags,
			&design.Visibility,
			&design.Status,
			&design.ViewsCount,
			&design.LikesCount,
			&design.CommentsCount,
			&design.CreatedAt,
			&design.UpdatedAt,
		)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}

		designs = append(designs, design)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}

	// Выводим результат
	fmt.Printf("Fetched %d designs\n", len(designs))
	return designs, nil
}
