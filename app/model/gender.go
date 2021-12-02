package model

import (
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"lopmartyn-gin-swagger/config/database"
)

type Gender struct {
	bun.BaseModel `bun:"genders,alias:g" json:"-" swaggerignore:"true"`
	ID            int    `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

// GetAllGenders : get all gender
func GetAllGenders() ([]Gender, error) {
	genders := new([]Gender)
	err := database.PGConnection.NewSelect().
		Model(genders).
		Order("g.id ASC").
		Scan(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetAllGenders",
			"package":  "model",
		}).Error(err)
		return *genders, err
	}
	return *genders, nil
}
