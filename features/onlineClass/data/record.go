package data

import (
	"capstone/backend/features/onlineClass"
	"time"
)

type OnlineClassCore struct {
	ID         int
	Name       string `json:"name"`
	Capacity   int    `json:"capacity"`
	ShortDesc  string `json:"short_desc"`
	Desc       string `json:"desc"`
	MonthlyFee int    `json:"monthly_fee"`
	Image      string `json:"image"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func toClassRecord(oc onlineClass.OnlineClassCore) OnlineClassCore {
	return OnlineClassCore{
		ID:         oc.ID,
		Name:       oc.Name,
		Capacity:   oc.Capacity,
		ShortDesc:  oc.ShortDesc,
		Desc:       oc.Desc,
		MonthlyFee: oc.MonthlyFee,
		Image:      oc.Image,
		CreatedAt:  oc.CreatedAt,
		UpdatedAt:  oc.UpdatedAt,
	}
}

func toClassCore(oc OnlineClassCore) onlineClass.OnlineClassCore {
	return onlineClass.OnlineClassCore{
		ID:         oc.ID,
		Name:       oc.Name,
		Capacity:   oc.Capacity,
		ShortDesc:  oc.ShortDesc,
		Desc:       oc.Desc,
		MonthlyFee: oc.MonthlyFee,
		Image:      oc.Image,
		CreatedAt:  oc.CreatedAt,
		UpdatedAt:  oc.UpdatedAt,
	}
}

func toOnlineClassCoreList(oc []OnlineClassCore) []onlineClass.OnlineClassCore {
	convOC := []onlineClass.OnlineClassCore{}

	for _, ocList := range oc {
		convOC = append(convOC, toClassCore(ocList))
	}
	return convOC
}
