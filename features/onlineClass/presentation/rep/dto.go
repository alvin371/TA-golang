package rep

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

func ToCore(req onlineClass.OnlineClassCore) OnlineClassCore {
	return OnlineClassCore{
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

func ToCoreSlice(core []onlineClass.OnlineClassCore) []OnlineClassCore {
	var onlineClassArray []OnlineClassCore
	for key := range core {
		onlineClassArray = append(onlineClassArray, ToCore(core[key]))
	}
	return onlineClassArray
}
