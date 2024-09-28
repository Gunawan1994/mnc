package handlers

import (
	"database/sql"
	"net/http"

	"payment/app/middlewares"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

type Routes struct {
	AuroraDb *sql.DB
}

func NewRoutes(auroradb *sql.DB) *Routes {
	return &Routes{
		AuroraDb: auroradb,
	}
}

func (route *Routes) RegisterServices(c *echo.Echo) {
	logger := log.WithFields(log.Fields{
		"job":    "RegisterServices",
		"msg_id": xid.New().String(),
	})
	logger.Debug("Running")

	handler := Handler(logger, route.AuroraDb)
	routes := c.Group("/api")
	route.setMiddleware(routes)
	routes.POST("/register", handler.RegisterHandler)
	routes.POST("/login", handler.LoginHandler)
	routes.POST("/topup", handler.TopUpHandler)
	routes.POST("/payment", handler.PaymentHandler)
	routes.POST("/transfer", handler.PaymentHandler)
	routes.GET("/report", handler.ReportTransactionHandler)
	routes.POST("/update", handler.UpdateProfileHandler)
}

func (route *Routes) setMiddleware(rGroup *echo.Group) {
	rGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRealIP},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost},
	}))

	m := middlewares.New("api")
	rGroup.Use(m.AddLoggerToContext, m.DumpRequest)

}
