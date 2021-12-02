package model

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"lopmartyn-gin-swagger/config/database"
	"time"
)

type Employee struct {
	bun.BaseModel `bun:"employees,alias:e" json:"-" swaggerignore:"true"`
	UUID          string    `bun:",pk" json:"-"`
	RequestID     string    `json:"requestID" example:"REQ-xxxx"`
	TitleID       int       `json:"titleID" example:1`
	FirstName     string    `json:"firstName" example:"Firstname"`
	LastName      string    `json:"lastName" example:"Lastname"`
	GenderID      int       `json:"genderID" example:1`
	DepartmentID  int       `json:"departmentID" example:1`
	PhoneNo       string    `json:"phoneNo" example:"02-000-0000"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	EmpID         string    `json:"empID" example:"EMP-xxxx"`
	Genders       *Gender   `bun:"rel:belongs-to,join:gender_id=id" json:"-"`
}

// CheckDuplicateEmployeeID : check duplicate employee
func CheckDuplicateEmployeeID(empID string) bool {
	employee := new(Employee)
	err := database.PGConnection.NewSelect().
		Model(employee).
		Where("e.emp_id = ?", empID).
		Scan(ctx)
	if err != sql.ErrNoRows {
		return true
	}
	return false
}

// DeleteEmployeeByID : delete data by record
func DeleteEmployeeByID(empID string) sql.Result {
	tx, err := database.PGConnection.BeginTx(ctx, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "DeleteEmployeeByID",
			"package":  "model",
			"action":   "Begin the transaction",
		}).Fatal(err)
	}
	employee := new(Employee)
	res, err := database.PGConnection.NewDelete().
		Model(employee).
		Where("e.uuid = ?", employee).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"function": "DeleteEmployeeByID",
			"package":  "model",
			"action":   "Delete Record",
		}).Fatal(err)
		return res
	}
	err = tx.Commit()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "DeleteEmployeeByID",
			"package":  "model",
			"action":   "Commit the transaction",
		}).Fatal(err)
	}
	return res
}

// GetEmployeeByID to get employee
func GetEmployeeByID(empID string) (Employee, error) {
	emp := new(Employee)
	err := database.PGConnection.NewSelect().
		Model(emp).
		Relation("Genders").
		Where("e.emp_id = ?", empID).
		Limit(1).
		Scan(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetEmployeeByID",
			"package":  "model",
		}).Error(err)
		return *emp, err
	}
	return *emp, nil
}

// InsertEmployee to insert employee
func InsertEmployee(employee Employee) sql.Result {
	tx, err := database.PGConnection.BeginTx(ctx, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "InsertEmployee",
			"package":  "Model",
			"action":   "Begin the transaction",
		}).Fatal(err)
	}

	res, err := database.PGConnection.NewInsert().
		Model(&employee).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"function": "InsertEmployee",
			"package":  "Model",
			"action":   "Insert record",
		}).Fatal(err)
		return res
	}

	err = tx.Commit()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "InsertEmployee",
			"package":  "Model",
			"action":   "Commit the transaction",
		}).Fatal(err)
	}
	return res
}

// UpdateEmployee : update the employee per record
func UpdateEmployee(employee Employee) sql.Result {
	tx, err := database.PGConnection.BeginTx(ctx, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "UpdateByRecord",
			"package":  "model",
			"action":   "Begin the transaction",
		}).Fatal(err)
	}

	res, err := database.PGConnection.NewUpdate().
		Model(&employee).
		Where("e.uuid = ?", employee.UUID).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"function": "UpdateByRecord",
			"package":  "model",
			"action":   "Update the record",
		}).Fatal(err)
		return res
	}

	err = tx.Commit()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "UpdateByRecord",
			"package":  "model",
			"action":   "Commit the transaction",
		}).Fatal(err)
	}
	return res
}
