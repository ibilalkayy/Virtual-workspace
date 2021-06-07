package entrance

/* Importing libraries */
import (
	"fmt"

	"github.com/ibilalkayy/VWS/database"
)

/* Function to login */
func Login() {
	fmt.Println("--------")
	fmt.Println("Log In")
	fmt.Println("--------")

	fmt.Printf("Enter your email address: ")
	fmt.Scanln(&pd.email)

	fmt.Printf("Enter your password: ")
	fmt.Scanln(&pd.password)

	/* Finding login info in database */
	database.Findaccount(pd.email, pd.password)
}
