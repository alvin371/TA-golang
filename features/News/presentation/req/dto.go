package req

import (
	news "capstone/backend/features/News"
	"time"
)

type News struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" param:"title"`
	Description string    `json:"description" param:"description"`
	Content     string    `json:"content" param:"content"`
	CreatorName string    `json:"creator_name" param:"creator_name"`
	Picture     string    `json:"picture" param:"picture"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromCore(core News) news.NewsCore {
	return news.NewsCore{
		ID:          core.ID,
		Title:       core.Title,
		Description: core.Description,
		Content:     core.Content,
		CreatorName: core.CreatorName,
		Picture:     core.Picture,
		Created_at:  core.CreatedAt,
		Updated_at:  core.CreatedAt,
	}
}
func (core *News) ToNewsCore() news.NewsCore {
	return news.NewsCore{
		ID:          core.ID,
		Title:       core.Title,
		Content:     core.Content,
		Description: core.Description,
		CreatorName: core.CreatorName,
		Picture:     core.Picture,
		Created_at:  core.CreatedAt,
		Updated_at:  core.CreatedAt,
	}
}
