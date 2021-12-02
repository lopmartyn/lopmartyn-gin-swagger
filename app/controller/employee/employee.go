package employee

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"lopmartyn-gin-swagger/app/helper"
	"lopmartyn-gin-swagger/app/model"
	"net/http"
	"strings"
	"time"
)

// DeleteEmployeeByID to delete employee
// @Summary Delete employee data
// @Schemes http
// @Description delete employee information
// @Tags employees
// @Accept json
// @Produce json
// @Param employeeID path string true "employeeID of employee to get the data"
// @Response 200 {object} model.Employee "OK"
// @Response 400 {object} helper.HTTPError "Bad Request"
// @Response 401 {object} helper.HTTPError "Unauthorized"
// @Response 404 {object} helper.HTTPError "A Employee was not found"
// @Response 500 {object} helper.HTTPError "Internal Server Error"
// @Router /employee/delete/{employeeID} [delete]
func DeleteEmployeeByID(ctx *gin.Context) {
	employeeID := ctx.Param("employeeID")
	_, err := model.GetEmployeeByID(employeeID)
	if err == sql.ErrNoRows {
		log.WithFields(log.Fields{
			"function": "DeleteEmployeeByID",
			"package":  "controller",
		}).Warn("Employee ID doesn't exist: " + employeeID)
		helper.ResponseCustomError(ctx, http.StatusNotFound, "Employee ID doesn't exist: "+employeeID)
		return
	}

	res := model.DeleteEmployeeByID(employeeID)
	affected, err := res.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "DeleteEmployeeByID",
			"package":  "controller",
		}).Error(err)
		helper.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"row affected": affected,
	})
}

// GetEmployeeByID to get all employee data
// @Summary Get employee data by Employee ID
// @Schemes http
// @Description get employee information
// @Tags employees
// @Accept json
// @Produce json
// @Param employeeID path string true "employeeID of employee to get the data"
// @Response 200 {object} model.Employee "OK"
// @Response 400 {object} helper.HTTPError "Bad Request"
// @Response 401 {object} helper.HTTPError "Unauthorized"
// @Response 404 {object} helper.HTTPError "A Employee was not found"
// @Response 500 {object} helper.HTTPError "Internal Server Error"
// @Router /employee/get/{employeeID} [get]
func GetEmployeeByID(ctx *gin.Context) {
	employee, err := model.GetEmployeeByID(ctx.Param("employeeID"))
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetEmployeeByID",
			"package":  "controller",
		}).Error(err)
		helper.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, employee)
}

// InsertEmployee insert employee data
// @Summary Insert employee data
// @Schemes http
// @Description insert employee information
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body model.Employee true "Create employee"
// @Response 200 {object} model.Employee "OK"
// @Response 400 {object} helper.HTTPError "Bad Request"
// @Response 401 {object} helper.HTTPError "Unauthorized"
// @Response 404 {object} helper.HTTPError "A Employee was not found"
// @Response 500 {object} helper.HTTPError "Internal Server Error"
// @Router /employee/create/ [post]
func InsertEmployee(ctx *gin.Context) {
	var param Parameter
	if err := ctx.ShouldBindJSON(&param); err != nil {
		//Get validator.ValidationErrors Error of type
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			log.WithFields(log.Fields{
				"function": "InsertEmployee",
				"package":  "Controller",
			}).Error(err.Error())
			helper.ResponseCustomError(ctx, http.StatusBadRequest, "Parameter validation failed")
			return
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			log.WithFields(log.Fields{
				"function": "InsertEmployee",
				"package":  "Controller",
				"type":     fieldErr.Type(),
				"value":    fieldErr.Value(),
				"param":    fieldErr.Param(),
			}).Error("Parameter validation failed")
			helper.ResponseCustomError(ctx, http.StatusBadRequest, fieldErr.Field()+" "+fieldErr.Type().String()+" "+fieldErr.Tag()+" "+fieldErr.Param())
			return
		}
	}

	dup := model.CheckDuplicateEmployeeID(param.EmpID)
	if dup {
		log.WithFields(log.Fields{
			"function": "InsertEmployee",
			"package":  "controller",
		}).Warn("Employee ID already exits: " + param.EmpID)
		helper.ResponseCustomError(ctx, http.StatusBadRequest, "Employee ID already exits: "+param.EmpID)
		return
	}
	var employee model.Employee
	employee.UUID = uuid.New().String()
	employee.RequestID = strings.TrimSpace(param.RequestID)
	employee.TitleID = param.TitleID
	employee.FirstName = strings.TrimSpace(param.FirstName)
	employee.LastName = strings.TrimSpace(param.LastName)
	employee.GenderID = param.GenderID
	employee.DepartmentID = param.DepartmentID
	employee.PhoneNo = strings.TrimSpace(param.PhoneNo)
	employee.CreatedAt = time.Now()
	employee.EmpID = strings.TrimSpace(param.EmpID)

	res := model.InsertEmployee(employee)
	affected, err := res.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "InsertEmployee",
			"package":  "controller",
		}).Error(err)
		helper.ResponseError(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"row affected": affected,
	})
}

// UpdateEmployee to update employee data
// @Summary Update employee data
// @Schemes http
// @Description update employee information
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body model.Employee true "Update employee"
// @Response 200 {object} model.Employee "OK"
// @Response 400 {object} helper.HTTPError "Bad Request"
// @Response 401 {object} helper.HTTPError "Unauthorized"
// @Response 404 {object} helper.HTTPError "A Employee was not found"
// @Response 500 {object} helper.HTTPError "Internal Server Error"
// @Router /employee/update/ [put]
func UpdateEmployee(ctx *gin.Context) {
	var param Parameter
	if err := ctx.ShouldBindJSON(&param); err != nil {
		//Get validator.ValidationErrors Error of type
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			log.WithFields(log.Fields{
				"function": "UpdateEmployee",
				"package":  "Controller",
			}).Error(err.Error())
			helper.ResponseCustomError(ctx, http.StatusBadRequest, "Parameter validation failed")
			return
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			log.WithFields(log.Fields{
				"function": "UpdateEmployee",
				"package":  "Controller",
				"type":     fieldErr.Type(),
				"value":    fieldErr.Value(),
				"param":    fieldErr.Param(),
			}).Error("Parameter validation failed")
			helper.ResponseCustomError(ctx, http.StatusBadRequest, fieldErr.Field()+" "+fieldErr.Type().String()+" "+fieldErr.Tag()+" "+fieldErr.Param())
			return
		}
	}

	employee, err := model.GetEmployeeByID(param.EmpID)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "UpdateEmployee",
			"package":  "controller",
		}).Error(err)
		helper.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	employee.RequestID = strings.TrimSpace(param.RequestID)
	employee.TitleID = param.TitleID
	employee.FirstName = strings.TrimSpace(param.FirstName)
	employee.LastName = strings.TrimSpace(param.LastName)
	employee.GenderID = param.GenderID
	employee.DepartmentID = param.DepartmentID
	employee.PhoneNo = strings.TrimSpace(param.PhoneNo)
	employee.UpdatedAt = time.Now()
	employee.EmpID = strings.TrimSpace(param.EmpID)

	res := model.UpdateEmployee(employee)
	affected, err := res.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "UpdateEmployee",
			"package":  "controller",
		}).Error(err)
		helper.ResponseError(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"row affected": affected,
	})
}
