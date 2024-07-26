package printer

import (
	"fmt"

	"github.com/PavelMenshikov/hw_otus/tree/hw02_fix_app/printer"
)

//func PrintStaff(staff []types.Employee) {
//	var str string
//	for i := 0; i < len(staff); i++ {
//	str += fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
//	staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
//	}

//		fmt.Println(str)
//	}
func PrintStaff(staff []types.Employee) {
	if len(staff) == 0 {
		fmt.Println("Список сотрудников пуст.")
		return
	}
	var str string
	for _, employee := range staff {
		str += fmt.Sprintf(
			"User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			employee.UserID, employee.Age, employee.Name, employee.DepartmentID)
	}
	fmt.Println(str)
}
