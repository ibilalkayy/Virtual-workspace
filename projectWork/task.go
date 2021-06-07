package projectWork

/* Importing libraries */
import (
	"fmt"
	"os"
	"strconv"

	"github.com/ibilalkayy/VWS/database"
	"github.com/ibilalkayy/VWS/entrance"
)

/* Declaring variables in struct */
type TaskVariables struct {
	ID, Projectname, Taskname, Description, Startdate, Enddate, Workhours string
}

/* Function to add and modify tasks */
func Task() {
	var choice string
	var teamMembers int
	fmt.Printf("Do you want to add a task? [y/n]: ")
	fmt.Scanln(&choice)

	switch choice {
	case "y":
		fmt.Printf("Enter the task id: ")
		fmt.Scanln(&pd.id)

		fmt.Printf("Enter the project name: ")
		entrance.WithSpace("projectName")

		fmt.Printf("Enter the task name: ")
		entrance.WithSpace("taskName")

		fmt.Println("Enter the task description: ")
		entrance.WithSpace("description")

		fmt.Printf("Enter the starting date: ")
		fmt.Scanln(&pd.startDate)

		fmt.Printf("Enter the ending date: ")
		fmt.Scanln(&pd.endDate)

		fmt.Printf("Enter the working hours: ")
		fmt.Scanln(&pd.workHours)

		/* Insert data into database */
		data := TaskVariables{pd.id, entrance.Combined["projectName"], entrance.Combined["taskName"], entrance.Combined["description"], pd.startDate, pd.endDate, pd.workHours}
		database.Insertdata(data)
	case "n":
		fmt.Println("Okay")
	default:
		fmt.Println("Error: Enter the correct choice")
		os.Exit(0)
	}

	fmt.Printf("Do you want to edit a task? [y/n]: ")
	fmt.Scanln(&choice)

	switch choice {
	case "y":
		fmt.Println("-------------------")
		fmt.Println("Access your account")
		fmt.Println("-------------------")

		fmt.Printf("Enter your task id: ")
		fmt.Scanln(&pd.id)

		fmt.Printf("Enter your task name: ")
		entrance.WithSpace("taskName")

		/* Find info and show it */
		database.Findtask(pd.id, entrance.Combined["taskName"])
		fmt.Printf("Here is your task name: %+v\n", database.Task.Taskname)
		fmt.Println("Here is your task description")
		fmt.Printf("%+v\n", database.Task.Description)

		fmt.Println("Editable information")
		fmt.Println("1. Task name")
		fmt.Println("2. Task description")
		fmt.Println("3. Ending date")
		fmt.Println("4. Working hours")
		fmt.Println("5. Task status")
		fmt.Println("6. Attach a file")
		fmt.Println("7. Add a team")

		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&choice)

		/* Update the task info */
		switch choice {
		case "1":
			fmt.Printf("Enter the task name: ")
			entrance.WithSpace("taskName")
			database.Updatetask("taskname", entrance.Combined["taskName"])
			fmt.Println("Your task name is successfully updated")
		case "2":
			fmt.Println("Enter the description: ")
			entrance.WithSpace("description")
			database.Updatetask("description", entrance.Combined["description"])
			fmt.Println("Your description is successfully updated")
		case "3":
			fmt.Printf("Enter the ending date: ")
			fmt.Scanln(&pd.endDate)
			database.Updatetask("enddate", pd.endDate)
			fmt.Println("Ending date is successfully updated")
		case "4":
			fmt.Printf("Enter the working hours: ")
			fmt.Scanln(&pd.workHours)
			database.Updatetask("workhours", pd.workHours)
			fmt.Println("Working hours are successfully updated")
		case "5":
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
				database.Updatetask("status", "active")
				fmt.Println("Task status is active")
			case "2":
				database.Updatetask("status", "in progress")
				fmt.Println("Task status is in progress")
			case "3":
				database.Updatetask("status", "delayed")
				fmt.Println("Task status is delayed")
			case "4":
				database.Updatetask("status", "in testing")
				fmt.Println("Task status is in testing")
			case "5":
				database.Updatetask("status", "approved")
				fmt.Println("Task status is approved")
			case "6":
				database.Updatetask("status", "cancelled")
				fmt.Println("Task status is cancelled")
			case "7":
				database.Updatetask("status", "completed")
				fmt.Println("Task status is completed")
			}
		case "6":
			fmt.Printf("Enter the filename to attach: ")
			entrance.WithSpace("fileName")
			filepath := "/home/bilal/Videos/"
			filename := entrance.Combined["fileName"]
			database.Uploadfile(filepath, filename)
			fmt.Println("File is successfully attached")
		case "7":
			fmt.Printf("Enter the team name: ")
			entrance.WithSpace("teamName")
			database.Updatetask("teamname", entrance.Combined["teamName"])

			fmt.Printf("Enter the team lead: ")
			entrance.WithSpace("teamLead")
			database.Updatetask("teamlead", entrance.Combined["teamLead"])

			fmt.Printf("Enter the number of members: ")
			fmt.Scanln(&teamMembers)
			database.Updatetask("teammembers", strconv.Itoa(teamMembers))

			fmt.Println("Enter their names and email addresses")
			for i := 0; i < teamMembers; i++ {
				fmt.Printf("Member name: ")
				entrance.WithSpace("memberName")
				database.NameandEmail("membername", entrance.Combined["memberName"])
			}
			for i := 0; i < teamMembers; i++ {
				fmt.Printf("Member email: ")
				entrance.WithSpace("memberEmail")
				database.NameandEmail("memberemail", entrance.Combined["memberEmail"])
			}
		}
	case "n":
		fmt.Println("Okay")
	default:
		fmt.Println("Error: Enter the correct choice")
		os.Exit(1)
	}
}
