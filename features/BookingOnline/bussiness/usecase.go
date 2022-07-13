package bussiness

import (
	"capstone/backend/features/bookingOnline"
	"fmt"
)

type bookingOnlineUseCase struct {
	bData bookingOnline.Data
}

func NewBussinessBookingOnlineClass(bOCData bookingOnline.Data) bookingOnline.Bussiness {
	return &bookingOnlineUseCase{
		bData: bOCData,
	}
}

func (booku *bookingOnlineUseCase) GetListBookingOnline(data bookingOnline.OnlineClassUser) (list []bookingOnline.OnlineClassUser, err error) {
	offlineClass, err := booku.bData.SelectAllBookingOnline(data)
	fmt.Println(offlineClass, "ini datanya online class", err, "kalau ada error", data, "APA SIH ANJING ISI DATANYA")
	if err != nil {
		return nil, err
	}
	return offlineClass, nil
}

func (booku *bookingOnlineUseCase) MemberBookingOnline(data bookingOnline.OnlineClassUser) (err error) {
	if err := booku.bData.InsertMemberBookingOnline(data); err != nil {
		return err
	}
	return nil
}
