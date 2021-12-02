package employee

type Parameter struct {
	UUID         string `json:"-"`
	RequestID    string `json:"requestID" binding:"required,min=8,max=8"`
	TitleID      int    `json:"titleID" binding:"required,numeric"`
	FirstName    string `json:"firstName" binding:"required,min=3,max=250"`
	LastName     string `json:"lastName" binding:"required,min=3,max=250"`
	GenderID     int    `json:"genderID" binding:"required,numeric"`
	DepartmentID int    `json:"departmentID" binding:"required,numeric"`
	PhoneNo      string `json:"phoneNo"`
	EmpID        string `json:"empID" binding:"required,min=8,max=8"`
}
