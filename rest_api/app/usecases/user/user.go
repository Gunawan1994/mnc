package user

import (
	"database/sql"
	"net/http"
	"payment/app/entity"
	"payment/app/repositories"
	"payment/app/repositories/user"
	"payment/app/usecases"
	"payment/pkg/response"
	"payment/pkg/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type event struct {
	RepoUser repositories.User
	AuroraDb *sql.DB
}

func New(logger *log.Entry, auroraDb *sql.DB) usecases.User {
	return &event{
		RepoUser: user.New(logger, auroraDb),
	}
}

func (v *event) Register(ctx echo.Context, req entity.ReqRegister) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": req}).Info("repositories: Register")

	var idUser string
	idUser, e = v.RepoUser.GetUserByPhone(ctx, *req.PhoneNumber)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error create GetUserByPhone")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			e.Error(), nil, nil)
		return
	}

	if idUser != "" {
		e = response.SetResponse(ctx, http.StatusBadRequest,
			"User exist", nil, nil)
		return
	}

	var id string
	id, e = v.RepoUser.CreateUser(ctx, req)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error create CreateUser")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			e.Error(), nil, nil)
		return
	}

	e = response.SetResponse(ctx, http.StatusOK, "Success", nil, map[string]interface{}{
		"user_id":      id,
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"phone_Number": req.PhoneNumber,
		"address":      req.Address,
		"created_Date": time.Now(),
	})
	return
}

func (v *event) Login(ctx echo.Context, req entity.ReqLogin) (e error) {
	logger := ctx.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": req}).Info("repositories: Login")

	var id, phoneNumber string
	id, phoneNumber, e = v.RepoUser.GetUser(ctx, *req.PhoneNumber, *req.Pin)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error create CreateUser")
		e = response.SetResponse(ctx, http.StatusBadRequest,
			e.Error(), nil, nil)
		return
	}

	if id == "" {
		e = response.SetResponse(ctx, http.StatusNotFound,
			"User not found", nil, nil)
		return
	}

	token, refresh, err := utils.GenerateTokens(phoneNumber, id)
	if err != nil {
		log.Fatal("Error generating token:", err)
	}

	e = response.SetResponse(ctx, http.StatusOK, "Success", nil, map[string]interface{}{
		"access_token":  token,
		"refresh_token": refresh,
	})
	return
}
