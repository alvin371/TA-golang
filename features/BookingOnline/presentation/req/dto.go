package req

import (
	"capstone/backend/features/bookingOnline"
)

type OnlineClassUser struct {
	ClassID int    `json:"class_id" param:"class_id" form:"class_id"`
	UserID  int    `json:"user_id" param:"user_id" form:"user_id"`
	Email   string `json:"email" param:"email" form:"email"`
	Phone   string `json:"phone" param:"phone" form:"phone"`
}

func (app *OnlineClassUser) ToCore() bookingOnline.OnlineClassUser {
	return bookingOnline.OnlineClassUser{
		ClassID: app.ClassID,
		UserID:  app.UserID,
		Email:   app.Email,
		Phone:   app.Phone,
	}
}
