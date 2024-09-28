package user

import (
	"database/sql"
	"payment/app/entity"
	"payment/app/repositories"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type event struct {
	Logger *log.Entry
	Db     *sql.DB
}

func New(logger *log.Entry, db *sql.DB) repositories.User {
	return &event{
		Logger: logger,
		Db:     db,
	}
}

func (v *event) CreateUser(c echo.Context, data entity.ReqRegister) (id string, e error) {
	logger := c.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": data}).Info("repositories: CreateUser")

	sql := `INSERT INTO "user" (first_name, last_name, phone_number, address, pin) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var lastInsertID string
	e = v.Db.QueryRow(sql, data.FirstName, data.LastName, data.PhoneNumber, data.Address, data.Pin).Scan(&lastInsertID)
	if e != nil {
		log.Fatal("Unable to execute insert: ", e)
	}

	return lastInsertID, nil
}

func (v *event) GetUserByPhone(c echo.Context, phone string) (id string, e error) {
	logger := c.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": phone}).Info("repositories: GetUserByPhone")

	// Use a parameterized query to prevent SQL injection
	sql := `SELECT id FROM "user" WHERE phone_number = $1`

	// Execute the query with the phone number as a parameter
	row := v.Db.QueryRow(sql, phone)

	// Scan the result into the id variable
	row.Scan(&id)

	return id, nil
}

func (v *event) GetUser(c echo.Context, phone, pin string) (id, phoneNumber string, e error) {
	logger := c.Get("logger").(*logrus.Entry)
	logger.WithFields(logrus.Fields{"params": phone}).Info("repositories: GetUser")

	// Use a parameterized query to prevent SQL injection
	sql := `SELECT id, phone_number FROM "user" WHERE phone_number = $1 AND pin = $2`

	// Execute the query with the phone number as a parameter
	row := v.Db.QueryRow(sql, phone, pin)

	// Scan the result into the id variable
	row.Scan(&id, &phoneNumber)

	return id, phoneNumber, nil
}
