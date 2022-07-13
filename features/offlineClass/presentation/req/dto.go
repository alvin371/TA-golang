package req

import (
	"capstone/backend/features/offlineClass"
	"time"
)

type OfflineClassCore struct {
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

func FromCore(core OfflineClassCore) offlineClass.OfflineClassCore {
	return offlineClass.OfflineClassCore{
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

func (core *OfflineClassCore) ToClassCore() offlineClass.OfflineClassCore {
	return offlineClass.OfflineClassCore{
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
