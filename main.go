package main

/* Importing libraries */
import (
	"fmt"
	"os"

	"github.com/ibilalkayy/VWS/entrance"
	"github.com/ibilalkayy/VWS/taskChoice"
)

/* Function to call all the major functions */
func main() {
	var option string
	fmt.Println("---------------------------------")
	fmt.Println("Welcome to the Virtual Workspace")
	fmt.Println("---------------------------------")

	fmt.Println("1. Sign Up")
	fmt.Println("2. Log In")

	fmt.Printf("Enter your choice: ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		entrance.Signup()
		taskChoice.TaskOption()
	case "2":
		entrance.Login()
		taskChoice.TaskOption()
	default:
		fmt.Println("Error: Enter the correct option")
		os.Exit(0)
	}
}
