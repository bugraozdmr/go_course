package main

import (
	"fmt"
	"structs_l/user"
)



func main() {
	firstName := user.GetUserData("Please enter your first name: ")
	lastName := user.GetUserData("Please enter your last name: ")
	birthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	admin := user.NewAdmin("admin@superusr.online", "strongpassword")

	user.OutputUserDetails(appUser)

	appUser.ClearUsername()

	// direkt struct'dan erişmek
	appUser.OutputUserDetailsAttached()

	fmt.Println("Sys Admin: ", admin)
}

