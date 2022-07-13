package rep

import (
	"capstone/backend/features/bookingOnline"
	"time"
)

type OnlineClassUser struct {
	ID        int
	ClassID   int       `json:"class_id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Class     OnlineClassResponse
	User      UserResponse
}

type OnlineClassResponse struct {
	ID         int
	Name       string    `json:"name"`
	Capacity   int       `json:"capacity"`
	ShortDesc  string    `json:"short_desc"`
	Desc       string    `json:"desc"`
	MonthlyFee int       `json:"monthly_fee"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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

func ToOnlineClassResponse(core bookingOnline.OnlineClassCore) OnlineClassResponse {
	return OnlineClassResponse{
		ID:         core.ID,
		Name:       core.Name,
		Capacity:   core.Capacity,
		ShortDesc:  core.ShortDesc,
		Desc:       core.Desc,
		MonthlyFee: core.MonthlyFee,
		Image:      core.Image,
		CreatedAt:  core.CreatedAt,
		UpdatedAt:  core.UpdatedAt,
	}
}
func ToBookingOnlineCore(core bookingOnline.OnlineClassUser) OnlineClassUser {
	return OnlineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		Email:     core.Email,
		Phone:     core.Phone,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOnlineClassResponse(core.Class),
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
