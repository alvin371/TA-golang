package data

import (
	"capstone/backend/features/bookingOnline"
	"time"
)

type OnlineClassUser struct {
	ID        int
	ClassID   int
	UserID    int
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OnlineClassCore
	User      User
}

type OnlineClassCore struct {
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

func ToUserCore(data User) bookingOnline.User {
	return bookingOnline.User{
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

func ToOnlineClassCore(core OnlineClassCore) bookingOnline.OnlineClassCore {
	return bookingOnline.OnlineClassCore{
		ID:         core.ID,
		Name:       core.Name,
		Capacity:   core.Capacity,
		ShortDesc:  core.ShortDesc,
		Desc:       core.Desc,
		MonthlyFee: core.MonthlyFee,
		Image:      core.Image,
	}
}

func TobookingOnlineCore(core OnlineClassUser) bookingOnline.OnlineClassUser {
	return bookingOnline.OnlineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOnlineClassCore(core.Class),
		User:      ToUserCore(core.User),
	}
}

func toBookingOnline(oc bookingOnline.OnlineClassUser) OnlineClassUser {
	return OnlineClassUser{
		ID:        oc.ID,
		ClassID:   oc.ClassID,
		UserID:    oc.UserID,
		Email:     oc.Email,
		Phone:     oc.Phone,
		CreatedAt: oc.CreatedAt,
		UpdatedAt: oc.UpdatedAt,
	}
}

func TobookingOnlineCoreList(oc []OnlineClassUser) []bookingOnline.OnlineClassUser {
	conv := []bookingOnline.OnlineClassUser{}

	for _, ocList := range oc {
		conv = append(conv, TobookingOnlineCore(ocList))
	}
	return conv
}

func ToBookingCore(core OnlineClassUser) bookingOnline.OnlineClassUser {
	return bookingOnline.OnlineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		Email:     core.Email,
		Phone:     core.Phone,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOnlineClassCore(core.Class),
		User:      ToUserCore(core.User),
	}
}

func ToBookingCoreList(core []OnlineClassUser) []bookingOnline.OnlineClassUser {
	convCore := []bookingOnline.OnlineClassUser{}
	for _, core := range core {
		convCore = append(convCore, ToBookingCore(core))
	}
	return convCore
}
