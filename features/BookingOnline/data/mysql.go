package data

import (
	"capstone/backend/features/bookingOnline"
	"fmt"

	"gorm.io/gorm"
)

type mysqlbookingOnlineClassRepo struct {
	Conn *gorm.DB
}

func NewBookingOnlineRepository(conn *gorm.DB) bookingOnline.Data {
	return &mysqlbookingOnlineClassRepo{
		Conn: conn,
	}
}

func (book *mysqlbookingOnlineClassRepo) SelectAllBookingOnline(bookingOnline.OnlineClassUser) (list []bookingOnline.OnlineClassUser, err error) {
	var record []OnlineClassUser
	err = book.Conn.Joins("JOIN online_class_cores ON online_class_cores.id = online_class_users.class_id JOIN users ON users.id = online_class_users.user_id").Find(&record).Error

	if err != nil {
		return nil, err
	}
	fmt.Println(record, "harusnya masuk dari data", err, "ini kalau ada error data")
	fmt.Println(ToBookingCoreList(record), "INI DATA BOOKING CORE LIST RECORD")
	return ToBookingCoreList(record), nil
}
func (book *mysqlbookingOnlineClassRepo) InsertMemberBookingOnline(data bookingOnline.OnlineClassUser) (err error) {
	// convData := TobookingOnlineCore(bookingOnline)
	convData := toBookingOnline(data)

	if err := book.Conn.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
