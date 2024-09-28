package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"payment/app/entity"
	"payment/app/usecases"
	"payment/app/usecases/user"
	"payment/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	ilog "github.com/sirupsen/logrus"
)

type HTTP struct {
	usecaseUser usecases.User
}

func Handler(logger *ilog.Entry, auroradb *sql.DB) *HTTP {
	return &HTTP{
		usecaseUser: user.New(logger, auroradb),
	}
}

func (c *HTTP) RegisterHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: RegisterHandler")

	req := entity.ReqRegister{}
	if e = ctx.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"Missing mandatory parameter", nil, nil)
		return
	}

	validate := validator.New()
	if e = validate.Struct(&req); e != nil {
		errs := e.(validator.ValidationErrors)
		for _, fieldErr := range errs {
			logger.WithField("error", e.Error()).Error(fmt.Printf("field %s: %s\n", fieldErr.Field(), fieldErr.Tag()))
			e = response.SetResponse(ctx, http.StatusBadRequest,
				fmt.Sprintf("Missing mandatory parameter %s", fieldErr.Field()), nil, nil)
			return
		}
		return
	}

	e = c.usecaseUser.Register(ctx, req)

	return
}

func (c *HTTP) LoginHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: LoginHandler")

	req := entity.ReqLogin{}
	if e = ctx.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"Missing mandatory parameter", nil, nil)
		return
	}

	e = c.usecaseUser.Login(ctx, req)

	return
}

func (c *HTTP) TopUpHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: TopUpHandler")

	req := entity.TopUp{}
	if e = ctx.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"Missing mandatory parameter", nil, nil)
		return
	}

	return
}

func (c *HTTP) PaymentHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: PaymentHandler")

	req := entity.Payment{}
	if e = ctx.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"Missing mandatory parameter", nil, nil)
		return
	}

	return
}

func (c *HTTP) TransferHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: TransferHandler")

	req := entity.Transfer{}
	if e = ctx.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"Missing mandatory parameter", nil, nil)
		return
	}

	return
}

func (c *HTTP) ReportTransactionHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: ReportTransactionHandler")

	return
}

func (c *HTTP) UpdateProfileHandler(ctx echo.Context) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.Info("handler: UpdateProfileHandler")
	return
}
