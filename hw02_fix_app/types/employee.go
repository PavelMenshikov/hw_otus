package types

import "fmt"

type Employee struct {
	UserID       int    `json:"userid"`
	Age          int    `json:"age"`
	FirstName    string `json:"name"`
	DepartmentID int    `json:"departmentid"`
}

func (e Employee) String() string {
	return fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ", e.UserID, e.Age, e.FirstName, e.DepartmentID)
}
