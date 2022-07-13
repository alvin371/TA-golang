package req

import (
	"capstone/backend/features/onlineClass"
	"time"
)

type OnlineClassCore struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Capacity   int       `json:"capacity"`
	ShortDesc  string    `json:"short_desc"`
	Desc       string    `json:"desc"`
	MonthlyFee int       `json:"monthly_fee"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromCore(core OnlineClassCore) onlineClass.OnlineClassCore {
	return onlineClass.OnlineClassCore{
		ID:         core.ID,
		Name:       core.Name,
		Capacity:   core.Capacity,
		ShortDesc:  core.ShortDesc,
		Desc:       core.Desc,
		MonthlyFee: core.MonthlyFee,
		Image:      core.Image,
		CreatedAt:  core.CreatedAt,
		UpdatedAt:  core.UpdatedAt,
	}
}

func (core *OnlineClassCore) ToClassCore() onlineClass.OnlineClassCore {
	return onlineClass.OnlineClassCore{
		ID:         core.ID,
		Name:       core.Name,
		Capacity:   core.Capacity,
		ShortDesc:  core.ShortDesc,
		Desc:       core.Desc,
		MonthlyFee: core.MonthlyFee,
		Image:      core.Image,
		CreatedAt:  core.CreatedAt,
		UpdatedAt:  core.UpdatedAt,
	}
}
