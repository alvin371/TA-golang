package rep

import (
	"capstone/backend/features/bookingOffline"
	"time"
)

type OfflineClassUser struct {
	ID        int
	ClassID   int    `json:"class_id"`
	UserID    int    `json:"user_id"`
	Session   string `json:"session"`
	Date      string `json:"date"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OfflineClassResponse
	User      UserResponse
}
type OfflineClassResponse struct {
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
		ID:         core.ID,
		Name:       core.Name,
		Capacity:   core.Capacity,
		ShortDesc:  core.ShortDesc,
		Desc:       core.Desc,
		MonthlyFee: core.MonthlyFee,
		Image:      core.Image,
	}
}
func ToBookingOfflineCore(core bookingOffline.OfflineClassUser) OfflineClassUser {
	return OfflineClassUser{
		ID:        core.ID,
		ClassID:   core.ClassID,
		UserID:    int(core.UserID),
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
