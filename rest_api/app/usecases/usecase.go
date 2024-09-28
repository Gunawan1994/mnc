package usecases

import (
	"payment/app/entity"

	"github.com/labstack/echo/v4"
)

type User interface {
	Register(ctx echo.Context, req entity.ReqRegister) (e error)
	Login(ctx echo.Context, req entity.ReqLogin) (e error)
}
