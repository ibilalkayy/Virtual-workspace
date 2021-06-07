package projectWork

/* Importing libraries */
import (
	"fmt"
	"os"

	"github.com/ibilalkayy/VWS/database"
	"github.com/ibilalkayy/VWS/entrance"
)

/* Declaring variables in struct */
type ProjectDetails struct {
	id, startDate, endDate, groupChoice, status, workHours string
}
type ProjectVariables struct {
	ID, Projectname, Ownername, Startdate, Enddate, Description, Groupname, Status string
}

var pd = ProjectDetails{}

/* Function to create, update, and delete project */
func Project() {
	var option, confirm string
	fmt.Println("Enter an option")
	fmt.Println("1. Create a new project")
	fmt.Println("2. Work on existing project")
	fmt.Println("3. Delete a project")

	fmt.Printf("Enter your choice: ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		fmt.Println("---------------------")
		fmt.Println("Create a new project")
		fmt.Println("---------------------")

		fmt.Printf("Enter the Project id: ")
		fmt.Scanln(&pd.id)

		fmt.Printf("Enter the Project name: ")
		entrance.WithSpace("projectName")

		fmt.Printf("Enter the owner's name: ")
		entrance.WithSpace("ownerName")

		fmt.Printf("Enter the starting date: ")
		fmt.Scanln(&pd.startDate)

		fmt.Printf("Enter the ending date: ")
		fmt.Scanln(&pd.endDate)

		fmt.Printf("Enter the description: ")
		entrance.WithSpace("description")

		fmt.Printf("Do you want the project to be in a group? [y/n]: ")
		fmt.Scanln(&pd.groupChoice)

		switch pd.groupChoice {
		case "y":
			fmt.Printf("Enter the group name to create: ")
			entrance.WithSpace("groupName")
		case "n":
			fmt.Println("Okay")
		default:
			fmt.Println("Error: Enter the correct choice")
			os.Exit(0)
		}

		fmt.Printf("Choose the project status [private/public]: ")
		fmt.Scanln(&pd.status)

		switch pd.status {
		case "private":
			fmt.Println("The private status is selected")
		case "public":
			fmt.Println("The public status is selected")
		default:
			fmt.Println("Error: Enter the correct choice")
			os.Exit(1)
		}
		/* Inserting data in database */
		data := ProjectVariables{pd.id, entrance.Combined["projectName"], entrance.Combined["ownerName"], pd.startDate, pd.endDate, entrance.Combined["description"], entrance.Combined["groupName"], pd.status}
		database.Insertdata(data)
	case "2":
		fmt.Println("--------------------")
		fmt.Println("Access your account")
		fmt.Println("--------------------")

		fmt.Printf("Enter your Project id: ")
		fmt.Scanln(&pd.id)

		fmt.Printf("Enter your Project name: ")
		entrance.WithSpace("projectName")

		/* Find project info and show it */
		database.Findproject(pd.id, entrance.Combined["projectName"])
		fmt.Printf("Here is your Project name: %+v\n", database.Project.Projectname)
		fmt.Printf("Here is your Project status: %+v\n", database.Project.Status)

		fmt.Println("Editable information")
		fmt.Println("1. Project name")
		fmt.Println("2. Owner name")
		fmt.Println("3. Ending date")
		fmt.Println("4. Project description")
		fmt.Println("5. Group name")
		fmt.Println("6. Project status")
		fmt.Println("7. Manage tasks")

		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&option)

		/* Updating project details */
		switch option {
		case "1":
			fmt.Printf("Enter the project name: ")
			entrance.WithSpace("projectName")
			database.Updateproject("projectname", entrance.Combined["projectName"])
			fmt.Println("Project name is succesfully updated")
		case "2":
			fmt.Printf("Enter the owner's name: ")
			entrance.WithSpace("ownerName")
			database.Updateproject("ownername", entrance.Combined["ownerName"])
			fmt.Println("Owner's name is succesfully updated")
		case "3":
			fmt.Printf("Enter the ending date: ")
			fmt.Scanln(&pd.endDate)
			database.Updateproject("enddate", pd.endDate)
			fmt.Println("Ending date is succesfully updated")
		case "4":
			fmt.Println("Enter the project description: ")
			entrance.WithSpace("description")
			database.Updateproject("description", entrance.Combined["description"])
			fmt.Println("Project description is succesfully updated")
		case "5":
			fmt.Printf("Enter the group name: ")
			entrance.WithSpace("groupName")
			database.Updateproject("groupname", entrance.Combined["groupName"])
			fmt.Println("Group name is succesfully updated")
		case "6":
			fmt.Println("1. Active")
			fmt.Println("2. In Progress")
			fmt.Println("3. Delayed")
			fmt.Println("4. In Testing")
			fmt.Println("5. Approved")
			fmt.Println("6. Cancelled")
			fmt.Println("7. Completed")

			fmt.Printf("Select the project status: ")
			fmt.Scanln(&pd.status)

			switch pd.status {
			case "1":
				database.Updateproject("status", "active")
				fmt.Println("Project status is active")
			case "2":
				database.Updateproject("status", "in progress")
				fmt.Println("Project status is in progress")
			case "3":
				database.Updateproject("status", "delayed")
				fmt.Println("Project status is delayed")
			case "4":
				database.Updateproject("status", "in testing")
				fmt.Println("Project status is in testing")
			case "5":
				database.Updateproject("status", "approved")
				fmt.Println("Project status is approved")
			case "6":
				database.Updateproject("status", "cancelled")
				fmt.Println("Project status is cancelled")
			case "7":
				database.Updateproject("status", "completed")
				fmt.Println("Project status is completed")
			}
		case "7":
			Task()
		default:
			fmt.Println("Error: Enter the correct choice")
		}
	/* Deleting the data from database */
	case "3":
		var id string
		fmt.Printf("Enter the password to confirm: ")
		fmt.Scanln(&confirm)
		fmt.Printf("Enter the project id: ")
		fmt.Scanln(&id)
		if confirm == database.Account.Password && pd.id == database.Project.ID {
			database.Verifyaccount(confirm, id)
			fmt.Println("Your project is deleted")
		} else {
			fmt.Println("Error: Wrong Id or Password")
			os.Exit(3)
		}
	default:
		fmt.Println("Error: Enter the correct choice")
		os.Exit(5)
	}
}
