package rep

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

func ToCore(req offlineClass.OfflineClassCore) OfflineClassCore {
	return OfflineClassCore{
		ID:         req.ID,
		Name:       req.Name,
		Capacity:   req.Capacity,
		ShortDesc:  req.ShortDesc,
		Desc:       req.Desc,
		MonthlyFee: req.MonthlyFee,
		Image:      req.Image,
		CreatedAt:  req.CreatedAt,
		UpdatedAt:  req.UpdatedAt,
	}
}

func ToCoreSlice(core []offlineClass.OfflineClassCore) []OfflineClassCore {
	var onlineClassArray []OfflineClassCore
	for key := range core {
		onlineClassArray = append(onlineClassArray, ToCore(core[key]))
	}
	return onlineClassArray
}
