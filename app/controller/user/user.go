package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"lopmartyn-gin-swagger/app/helper"
	"lopmartyn-gin-swagger/app/model"
	"net/http"
	"time"
)

// GetAllUsersInfo to get all user data
// @Summary Get all user data
// @Schemes http
// @Description get user information
// @Tags users
// @Accept json
// @Produce json
// @Response 200 {object} model.Users "OK"
// @Response 400 {object} helper.HTTPError "Bad Request"
// @Response 401 {object} helper.HTTPError "Unauthorized"
// @Response 404 {object} helper.HTTPError "A Gender was not found"
// @Response 500 {object} helper.HTTPError "Internal Server Error"
// @Router /user/get/all [get]
func GetAllUsersInfo(ctx *gin.Context) {
	suppliers, err := model.GetAllUsersInfo()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetAllRegistrationInfo",
			"package":  "controller",
		}).Error(err)
		helper.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, suppliers)
}

// Registration user
// @Summary Register user
// @Schemes http
// @Description Create registers by record
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.Users true "Create supplier"
// @Response 200 {object} model.Users "OK"
// @Response 400 {object} helper.HTTPError "Bad Request"
// @Response 401 {object} helper.HTTPError "Unauthorized"
// @Response 404 {object} helper.HTTPError "A User ID with the specified ID was not found"
// @Response 500 {object} helper.HTTPError "Internal Server Error"
// @Router /user/create [post]
func Registration(ctx *gin.Context) {
	var param Parameter
	if err := ctx.ShouldBindJSON(&param); err != nil {
		//Get validator.ValidationErrors Error of type
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			log.WithFields(log.Fields{
				"function": "Registration",
				"package":  "Controller",
			}).Error(err.Error())
			helper.ResponseCustomError(ctx, http.StatusBadRequest, "Parameter validation failed")
			return
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			log.WithFields(log.Fields{
				"value": fieldErr.Value(),
				"param": fieldErr.Param(),
			}).Error("Parameter validation failed")
			helper.ResponseCustomError(ctx, http.StatusBadRequest, fieldErr.Field()+" "+fieldErr.Type().String()+" "+fieldErr.Tag()+" "+fieldErr.Param())
			return
		}
	}

	dup := model.CheckDuplicateEmail(param.Email)
	if dup {
		log.WithFields(log.Fields{
			"function": "Registration",
			"package":  "controller",
		}).Warn("Email already exits: " + param.Email)
		helper.ResponseCustomError(ctx, http.StatusBadRequest, "Email already exits: "+param.Email)
		return
	}
	var register model.Users
	register.UUID = uuid.New().String()
	register.Email = param.Email
	register.RegisteredAt = time.Now()
	register.Status = "Active"
	register.Password = param.Password

	res := model.Registration(register)
	affected, err := res.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "Registration",
			"package":  "controller",
		}).Error(err)
		helper.ResponseError(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"row affected": affected,
	})
}
