package req

import "capstone/backend/features/bookingOffline"

type OfflineClassUser struct {
	ClassID int    `json:"class_id"`
	UserID  int    `json:"user_id"`
	Session string `json:"session"`
	Date    string `json:"date"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func (app *OfflineClassUser) ToCore() bookingOffline.OfflineClassUser {
	return bookingOffline.OfflineClassUser{
		ClassID: app.ClassID,
		UserID:  app.UserID,
		Session: app.Session,
		Date:    app.Date,
		Email:   app.Email,
		Phone:   app.Phone,
	}
}
