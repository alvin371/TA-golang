package presentation

import (
	// Import Echo

	user "capstone/backend/features/User"
	"capstone/backend/features/User/presentation/rep"
	"capstone/backend/features/User/presentation/req"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userBussiness user.Bussiness
}

func NewHandlerAccount(userBussiness user.Bussiness) *UserHandler {
	return &UserHandler{userBussiness: userBussiness}
}

func (usrHandler *UserHandler) CreateUserHandler(e echo.Context) error {
	newAccount := req.User{}
	if err := e.Bind(&newAccount); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	// fmt.Println("this new account data", newAccount)
	if len([]rune(newAccount.Password)) < 8 {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Password must be greather than 8 Characters",
		})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newAccount.Password), 5)
	newAccount.Password = string(hash)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	fmt.Println("new account data", newAccount)
	if err := usrHandler.userBussiness.CreateUser(newAccount.ToUserCore()); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newAccount,
	})
}

func (usrHandler *UserHandler) GetUserById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Isi ID : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := usrHandler.userBussiness.GetUserById(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    rep.ToUserCore(data),
	})
}

func (usrHandler *UserHandler) GetAllUserHandler(e echo.Context) error {
	result, err := usrHandler.userBussiness.GetAllUser(user.User{})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "All is Well",
		"data":    rep.ToUserCoreSlice(result),
	})
}

func (usrHandler *UserHandler) LoginUserHandler(e echo.Context) error {
	AccountAuth := req.UserAuth{}
	e.Bind(&AccountAuth)
	data, err := usrHandler.userBussiness.LoginUser(AccountAuth.ToUserAuth())
	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    rep.ToUserLoginResponse(data),
	})
}
func (usrHandler *UserHandler) LoginAdminHandler(e echo.Context) error {
	AccountAuth := req.UserAuth{}
	e.Bind(&AccountAuth)
	data, err := usrHandler.userBussiness.LoginAdmin(AccountAuth.ToUserAuth())
	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    rep.ToUserLoginResponse(data),
	})
}

func (usrHandler *UserHandler) UpdateAccountHandler(e echo.Context) error {
	newAccount := req.User{}
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Test : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := e.Bind(&newAccount); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	fmt.Println(newAccount)
	data, err := usrHandler.userBussiness.EditUser(id, newAccount.ToUserCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    rep.ToUserCore(data),
	})
}
