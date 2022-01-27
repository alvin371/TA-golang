package rep

import (
	"capstone/backend/features/bookingOnline"
	"time"
)

type OnlineClassUser struct {
	ID        int
	ClassID   int `json:"class_id"`
	UserID    int `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OnlineClassResponse
	User      UserResponse
}

type OnlineClassResponse struct {
	ID      int
	Name    string `json:"name"`
	Day     string `json:"day"`
	Date    string `json:"date"`
	Link    string `json:"link"`
	Time    string `json:"time"`
	Trainer string `json:"trainer"`
	Image   string `json:"image"`
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

func ToUserResponse(data bookingOnline.User) UserResponse {
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

func ToOfflineClassResponse(core bookingOnline.OnlineClassCore) OnlineClassResponse {
	return OnlineClassResponse{
		ID:      core.ID,
		Name:    core.Name,
		Day:     core.Day,
		Date:    core.Date,
		Link:    core.Link,
		Time:    core.Time,
		Trainer: core.Trainer,
		Image:   core.Image,
	}
}
func ToBookingOnlineCore(core bookingOnline.OnlineClassUser) OnlineClassUser {
	return OnlineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOfflineClassResponse(core.Class),
		User:      ToUserResponse(core.User),
	}
}

func TobookingOnlineCoreList(oc []bookingOnline.OnlineClassUser) []OnlineClassUser {
	conv := []OnlineClassUser{}

	for _, ocList := range oc {
		conv = append(conv, ToBookingOnlineCore(ocList))
	}
	return conv
}
