package rep

import (
	news "capstone/backend/features/News"
	"time"
)

type News struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatorName string    `json:"creator_name"`
	Picture     string    `json:"picture"`
	Created_at  time.Time `json:"created_at" `
	Updated_at  time.Time `json:"updated_at" `
}

func ToCore(req news.NewsCore) News {
	return News{
		ID:          req.ID,
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		CreatorName: req.CreatorName,
		Picture:     req.Picture,
		Created_at:  req.Created_at,
		Updated_at:  req.Updated_at,
	}
}

func ToCoreSlice(core []news.NewsCore) []News {
	var newsArray []News
	for key := range core {
		newsArray = append(newsArray, ToCore(core[key]))
	}
	return newsArray
}
