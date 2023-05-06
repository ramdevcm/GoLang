package main

import "fmt"

const LoginToken = "qwertyuiop" // first letter capital means the variable is global

func main() {
	fmt.Println("VAriables")

	var username string = "Ramdev"
	fmt.Println(username)

	// boolean
	var isLoogedIn bool
	fmt.Println(isLoogedIn)
	fmt.Printf("Data type is : %T \n\n", isLoogedIn)

	// implicit type
	var website = "ramdev.cm"
	fmt.Println(website)

	// no var style
	numberOfUser := 12345667
	fmt.Println(numberOfUser)

	// public
	fmt.Println(LoginToken)
	fmt.Printf("Data type is : %T \n\n", LoginToken)
}
