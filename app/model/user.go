package model

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"lopmartyn-gin-swagger/config/database"
	"time"
)

type Users struct {
	bun.BaseModel `bun:"users,alias:u" json:"-" swaggerignore:"true"`
	UUID          string `bun:",pk" json:"-"`
	Email         string
	RegisteredAt  time.Time `json:"-"`
	Status        string    `json:"-"`
	Password      string
}

// CheckDuplicateEmail get user information
func CheckDuplicateEmail(email string) bool {
	user := new(Users)
	err := database.PGConnection.NewSelect().
		Model(user).
		Where("u.email = ?", email).
		Scan(ctx)
	if err != sql.ErrNoRows {
		return true
	}
	return false
}

// GetAllUsersInfo get user information
func GetAllUsersInfo() ([]Users, error) {
	users := new([]Users)
	err := database.PGConnection.NewSelect().
		Model(users).
		Order("u.registered_at ASC").
		Scan(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetAllRegistrationInfo",
			"package":  "Model",
		}).Error(err)
		return *users, err
	}
	return *users, nil
}

// Registration : get all gender
func Registration(user Users) sql.Result {
	tx, err := database.PGConnection.BeginTx(ctx, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "Registration",
			"package":  "Model",
			"action":   "Begin the transaction",
		}).Fatal(err)
	}

	res, err := database.PGConnection.NewInsert().
		Model(&user).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"function": "Registration",
			"package":  "Model",
			"action":   "Insert record",
		}).Fatal(err)
		return res
	}

	err = tx.Commit()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "Registration",
			"package":  "Model",
			"action":   "Commit the transaction",
		}).Fatal(err)
	}
	return res
}
