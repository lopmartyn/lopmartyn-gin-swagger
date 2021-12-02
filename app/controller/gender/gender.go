package gender

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"lopmartyn-gin-swagger/app/helper"
	"lopmartyn-gin-swagger/app/model"
	"net/http"
)

// GetAllGenders to get all gender data
// @Summary Get all gender data
// @Schemes http
// @Description get gender information
// @Tags gender
// @Accept json
// @Produce json
// @Response 200 {object} model.Gender "OK"
// @Response 400 {object} helper.HTTPError "Bad Request"
// @Response 401 {object} helper.HTTPError "Unauthorized"
// @Response 404 {object} helper.HTTPError "A Gender was not found"
// @Response 500 {object} helper.HTTPError "Internal Server Error"
// @Router /gender/get/all [get]
func GetAllGenders(ctx *gin.Context) {
	suppliers, err := model.GetAllGenders()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetAllGenders",
			"package":  "controller",
		}).Error(err)
		helper.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, suppliers)
}
