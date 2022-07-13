package bookingOffline

import (
	"time"
)

type OfflineClassUser struct {
	ID        int
	ClassID   int `gorm:"primaryKey"`
	UserID    int `gorm:"primaryKey"`
	Session   string
	Date      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OfflineClassCore
	User      User
}

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
type User struct {
	ID           uint
	Username     string `json:"username" gorm:"unique;not null"`
	Role         string `json:"role" gorm:"default:user"`     //user
	Email        string `json:"email" gorm:"unique;not null"` //unique,
	Password     string `json:"password" gorm:"not null"`
	Token        string `jason:"token" form:"token"`
	Avatar       string `json:"avatar" form:"avatar"`
	Goals        string `json:"goals" form:"goals"`
	MemberStatus string `json:"member_status" gorm:"default:reguler"`
	Created_at   time.Time
	Updated_at   time.Time
}

type Bussiness interface {
	GetListBookingOffline(OfflineClassUser) (list []OfflineClassUser, err error)
	GetBookingByID(id int) (OfflineClassUser, error)
	MemberBookingOffline(userID int, classID int) (err error)
}

type Data interface {
	SelectAllBookingOffline(OfflineClassUser) (list []OfflineClassUser, err error)
	SelectBookingByID(id int) (OfflineClassUser, error)
	InsertMemberBookingOffline(userID int, classID int) (err error)
}
