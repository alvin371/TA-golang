package data

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

func toClassRecord(oc offlineClass.OfflineClassCore) OfflineClassCore {
	return OfflineClassCore{
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

func toClassCore(oc OfflineClassCore) offlineClass.OfflineClassCore {
	return offlineClass.OfflineClassCore{
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

func toOfflineClassCoreList(oc []OfflineClassCore) []offlineClass.OfflineClassCore {
	convOC := []offlineClass.OfflineClassCore{}

	for _, ocList := range oc {
		convOC = append(convOC, toClassCore(ocList))
	}
	return convOC
}
