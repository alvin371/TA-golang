package presentation

import (
	"capstone/backend/features/bookingOffline"
	presentation_response "capstone/backend/features/bookingOffline/presentation/rep"
	presentation_request "capstone/backend/features/bookingOffline/presentation/req"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingOfflineHandler struct {
	bookingOfflineBussiness bookingOffline.Bussiness
}

func NewBookingOfflineHandler(bookingOfflineBussiness bookingOffline.Bussiness) BookingOfflineHandler {
	return BookingOfflineHandler{bookingOfflineBussiness}
}

func (boh *BookingOfflineHandler) GetListBookingOffline(e echo.Context) error {
	data, err := boh.bookingOfflineBussiness.GetListBookingOffline(bookingOffline.OfflineClassUser{})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get List Booking Offline",
		"data":    presentation_response.ToBookingOfflineCoreList(data),
	})
}

func (boh *BookingOfflineHandler) InsertMemberBookingOffline(e echo.Context) error {
	newBookingOffline := presentation_request.OfflineClassUser{}

	if err := e.Bind(&newBookingOffline); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	fmt.Println("testing", newBookingOffline)

	if err := boh.bookingOfflineBussiness.MemberBookingOffline(newBookingOffline.ToCore().ClassID, newBookingOffline.UserID); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newBookingOffline,
	})
}

func (boh *BookingOfflineHandler) SelectBookingByClassID(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Isi ID : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := boh.bookingOfflineBussiness.GetBookingByID(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToBookingOfflineCore(data),
	})
}
