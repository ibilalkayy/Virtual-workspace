package profile

/* Importing libraries */
import (
	"fmt"
	"os"

	"github.com/ibilalkayy/VWS/database"
	"github.com/ibilalkayy/VWS/entrance"
)

/* Declaring variables in struct */
type PersonDetails struct {
	email, password, gender, language, phone string
}

var pd = PersonDetails{}

type AccountVariables struct {
	Name, Email, Password, Business string
}

var Account AccountVariables

/* Showing account info and updating them */
func AccountInfo() {
	var option string
	fmt.Printf("Here is your name: %+v\n", database.Account.Name)
	fmt.Printf("Here is your email address: %+v\n", database.Account.Email)
	fmt.Printf("Here is your password: %+v\n", database.Account.Password)
	fmt.Printf("Here is your business name: %+v\n", database.Account.Business)

	fmt.Println("Editable information")
	fmt.Println("1. Name")
	fmt.Println("2. Email address")
	fmt.Println("3. Password")
	fmt.Println("4. Gender")
	fmt.Println("5. State")
	fmt.Println("6. Country")
	fmt.Println("7. Language")
	fmt.Println("8. Phone number")
	fmt.Println("9. Close your account")

	fmt.Printf("Enter your choice: ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		fmt.Printf("Enter your name: ")
		entrance.WithSpace("name")
		database.Updateaccount("name", entrance.Combined["name"])
		fmt.Println("Your name is succesfully updated")
	case "2":
		fmt.Printf("Enter your email address: ")
		fmt.Scanln(&pd.email)
		database.Updateaccount("email", pd.email)
		fmt.Println("Your email address is succesfully updated")
	case "3":
		fmt.Printf("Enter your password: ")
		fmt.Scanln(&pd.password)
		database.Updateaccount("password", pd.password)
		fmt.Println("Your password is succesfully updated")
	case "4":
		fmt.Printf("Enter your gender: ")
		fmt.Scanln(&pd.gender)
		database.Updateaccount("gender", pd.gender)
		fmt.Println("Your gender is succesfully updated")
	case "5":
		fmt.Printf("Enter your state: ")
		entrance.WithSpace("state")
		database.Updateaccount("state", entrance.Combined["state"])
		fmt.Println("Your state is succesfully updated")
	case "6":
		fmt.Printf("Enter your country: ")
		entrance.WithSpace("country")
		database.Updateaccount("country", entrance.Combined["country"])
		fmt.Println("Your country is succesfully updated")
	case "7":
		fmt.Printf("Enter your language: ")
		fmt.Scanln(&pd.language)
		database.Updateaccount("language", pd.language)
		fmt.Println("Your language is succesfully updated")
	case "8":
		fmt.Printf("Enter your phone number: ")
		fmt.Scanln(&pd.phone)
		database.Updateaccount("phone", pd.phone)
		fmt.Println("Your phone number is succesfully updated")
	case "9":
		fmt.Printf("Confirm your password: ")
		fmt.Scanln(&pd.password)
		if pd.password == database.Account.Password {
			database.Deleteaccount("password", database.Account.Password)
			fmt.Println("Your account is closed")
			os.Exit(0)
		} else {
			fmt.Println("Error: Wrong Password")
			os.Exit(1)
		}
	default:
		fmt.Println("Error: Enter the correct choice")
		os.Exit(2)
	}
}
