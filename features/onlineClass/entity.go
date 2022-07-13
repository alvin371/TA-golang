package onlineClass

import "time"

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

type Bussiness interface {
	// crud
	GetAllClass(OnlineClassCore) (class []OnlineClassCore, err error)
	GetClassById(id int) (OnlineClassCore, error)
	CreateClass(data OnlineClassCore) (err error)
	EditClass(id int, data OnlineClassCore) (news OnlineClassCore, err error)
	DeleteClass(id int) (news OnlineClassCore, err error)
}

type Data interface {
	SelectAllClass(OnlineClassCore) (class []OnlineClassCore, err error)
	SelectClassById(id int) (OnlineClassCore, error)
	InsertClass(data OnlineClassCore) (err error)
	UpdateClass(id int, data OnlineClassCore) (news OnlineClassCore, err error)
	DestryoClass(id int) (news OnlineClassCore, err error)
}
