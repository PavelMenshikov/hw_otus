package printer

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func PrintStaff(staff types.Employee) {
	// Removed the line with len(staff)
	str := staff.String() // Call the String() method
	fmt.Println(str)
}
