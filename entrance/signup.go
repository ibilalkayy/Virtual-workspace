package entrance

/* Importing libraries */
import (
	"bufio"
	"fmt"
	"os"

	"github.com/ibilalkayy/VWS/database"
)

/* Declaring variables in struct */
type PersonDetails struct {
	email, password string
}

var pd = PersonDetails{}

type AccountVariables struct {
	Name, Email, Password, Business string
}

var Combined = make(map[string]string)

/* Function to accept string/integer with space */
func WithSpace(key string) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		value := scanner.Text()
		Combined[key] = value
	}
}

/* Function for registration */
func Signup() {
	fmt.Println("--------")
	fmt.Println("Sign Up")
	fmt.Println("--------")

	fmt.Printf("Enter your name: ")
	WithSpace("name")

	fmt.Printf("Enter your email address: ")
	fmt.Scanln(&pd.email)

	fmt.Printf("Enter your password: ")
	fmt.Scanln(&pd.password)

	fmt.Printf("Enter your business name: ")
	WithSpace("business")

	/* Inserting data into database */
	data := AccountVariables{Combined["name"], pd.email, pd.password, Combined["business"]}
	database.Insertdata(data)
}
