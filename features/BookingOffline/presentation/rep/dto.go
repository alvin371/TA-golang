package rep

import (
	"capstone/backend/features/bookingOffline"
	"time"
)

type OfflineClassUser struct {
	ID        int
	ClassID   int `json:"class_id"`
	UserID    int `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OfflineClassResponse
	User      UserResponse
}

type OfflineClassResponse struct {
	ID       int
	Name     string `json:"name"`
	Day      string `json:"day"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Time     string `json:"time"`
	Trainer  string `json:"trainer"`
	Image    string `json:"image"`
}
type UserResponse struct {
	ID           uint
	Username     string `json:"username"`
	Role         string `json:"role"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	Goals        string `json:"goals"`
	MemberStatus string `json:"member_status"`
}

func ToUserResponse(data bookingOffline.User) UserResponse {
	return UserResponse{
		ID:           data.ID,
		Username:     data.Username,
		Role:         data.Role,
		Email:        data.Email,
		Avatar:       data.Avatar,
		Goals:        data.Goals,
		MemberStatus: data.MemberStatus,
	}
}

func ToOfflineClassResponse(core bookingOffline.OfflineClassCore) OfflineClassResponse {
	return OfflineClassResponse{
		ID:       core.ID,
		Name:     core.Name,
		Day:      core.Day,
		Date:     core.Date,
		Location: core.Location,
		Time:     core.Time,
		Trainer:  core.Trainer,
		Image:    core.Image,
	}
}
func ToBookingOfflineCore(core bookingOffline.OfflineClassUser) OfflineClassUser {
	return OfflineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOfflineClassResponse(core.Class),
		User:      ToUserResponse(core.User),
	}
}

func ToBookingOfflineCoreList(oc []bookingOffline.OfflineClassUser) []OfflineClassUser {
	conv := []OfflineClassUser{}

	for _, ocList := range oc {
		conv = append(conv, ToBookingOfflineCore(ocList))
	}
	return conv
}
