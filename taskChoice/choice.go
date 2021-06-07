package taskChoice

/* Importing libraries */
import (
	"fmt"
	"os"

	"github.com/ibilalkayy/VWS/profile"
	"github.com/ibilalkayy/VWS/projectWork"
)

/* Options to show task to perform */
func TaskOption() {
	var option string
	fmt.Println("Which task do you want to perform?")
	fmt.Println("1. Account Info")
	fmt.Println("2. Projects")
	fmt.Println("3. Sign Out")

	fmt.Printf("Enter your choice: ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		profile.AccountInfo()
	case "2":
		projectWork.Project()
	case "3":
		fmt.Println("Signed Out")
		os.Exit(0)
	default:
		fmt.Println("Error: Enter the correct choice")
		os.Exit(1)
	}

}
