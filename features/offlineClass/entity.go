package offlineClass

import "time"

type OfflineClassCore struct {
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

type Bussiness interface {
	// crud
	GetAllClass(OfflineClassCore) (class []OfflineClassCore, err error)
	GetClassById(id int) (OfflineClassCore, error)
	CreateClass(data OfflineClassCore) (err error)
	EditClass(id int, data OfflineClassCore) (news OfflineClassCore, err error)
	DeleteClass(id int) (news OfflineClassCore, err error)
}

type Data interface {
	SelectAllClass(OfflineClassCore) (class []OfflineClassCore, err error)
	SelectClassById(id int) (OfflineClassCore, error)
	InsertClass(data OfflineClassCore) (err error)
	UpdateClass(id int, data OfflineClassCore) (news OfflineClassCore, err error)
	DestroyClass(id int) (news OfflineClassCore, err error)
}
