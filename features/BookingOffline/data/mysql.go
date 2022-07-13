package data

import (
	"capstone/backend/features/bookingOffline"
	"fmt"

	"gorm.io/gorm"
)

type mysqlBookingOfflineClassRepo struct {
	Conn *gorm.DB
}

func NewBookingOfflineRepository(conn *gorm.DB) bookingOffline.Data {
	return &mysqlBookingOfflineClassRepo{
		Conn: conn,
	}
}

func (book *mysqlBookingOfflineClassRepo) SelectAllBookingOffline(bookingOffline.OfflineClassUser) (list []bookingOffline.OfflineClassUser, err error) {
	var record []OfflineClassUser
	book.Conn.Find(&record)
	fmt.Println("this is record", record)
	fmt.Println("this is after ", ToBookingOfflineCoreList(record))
	// fmt.Println("this is finding data", err)
	if err != nil {
		return nil, err
	}
	return ToBookingOfflineCoreList(record), nil
}
func (book *mysqlBookingOfflineClassRepo) InsertMemberBookingOffline(userID int, classID int) (err error) {
	var bookingoffline = OfflineClassUser{}
	bookingoffline.UserID = userID
	bookingoffline.ClassID = classID
	// convData := ToBookingOfflineCore(bookingoffline)

	if err := book.Conn.Create(&bookingoffline).Error; err != nil {
		return err
	}
	return nil
}

func (book *mysqlBookingOfflineClassRepo) SelectBookingByID(id int) (bookingOffline.OfflineClassUser, error) {
	var bookingData OfflineClassUser

	err := book.Conn.Where("class_id = ?", id).Error
	fmt.Println(book.Conn.Where("class_id = ?", id))
	if bookingData.ClassID == 0 {
		return bookingOffline.OfflineClassUser{}, err
	}

	if err != nil {
		return bookingOffline.OfflineClassUser{}, err
	}

	return ToBookingOfflineCore(bookingData), nil
}
