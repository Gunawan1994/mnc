package repositories

import (
	"payment/app/entity"

	"github.com/labstack/echo/v4"
)

type User interface {
	CreateUser(c echo.Context, data entity.ReqRegister) (id string, e error)
	GetUserByPhone(c echo.Context, phone string) (id string, e error)
	GetUser(c echo.Context, phone, pin string) (id, phoneNumber string, e error)
}
