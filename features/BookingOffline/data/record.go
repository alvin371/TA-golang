package data

import (
	"capstone/backend/features/bookingOffline"
	"time"
)

type OfflineClassUser struct {
	ID        int
	ClassID   int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OfflineClassCore
	User      User
}

type OfflineClassCore struct {
	ID         int
	Name       string
	Capacity   int
	ShortDesc  string
	Desc       string
	MonthlyFee int
	Image      string
}
type User struct {
	ID           uint
	Username     string
	Role         string
	Email        string
	Password     string
	Token        string
	Avatar       string
	Goals        string
	MemberStatus string
}

func ToUserCore(data User) bookingOffline.User {
	return bookingOffline.User{
		ID:           data.ID,
		Username:     data.Username,
		Role:         data.Role,
		Email:        data.Email,
		Password:     data.Password,
		Token:        data.Token,
		Avatar:       data.Avatar,
		Goals:        data.Goals,
		MemberStatus: data.MemberStatus,
	}
}

func ToOfflineClassCore(core OfflineClassCore) bookingOffline.OfflineClassCore {
	return bookingOffline.OfflineClassCore{
		ID:         core.ID,
		Name:       core.Name,
		Capacity:   core.Capacity,
		ShortDesc:  core.ShortDesc,
		Desc:       core.Desc,
		MonthlyFee: core.MonthlyFee,
		Image:      core.Image,
	}
}

func ToBookingOfflineCore(core OfflineClassUser) bookingOffline.OfflineClassUser {
	return bookingOffline.OfflineClassUser{
		ID:        core.ID,
		ClassID:   core.ClassID,
		UserID:    int(core.UserID),
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOfflineClassCore(core.Class),
		User:      ToUserCore(core.User),
	}
}

func ToBookingOfflineCoreList(oc []OfflineClassUser) []bookingOffline.OfflineClassUser {
	conv := []bookingOffline.OfflineClassUser{}

	for _, ocList := range oc {
		conv = append(conv, ToBookingOfflineCore(ocList))
	}
	return conv
}
