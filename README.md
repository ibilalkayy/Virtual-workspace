# Virtual-workspace
This repository contains the code of virtual workspace that will manage your projects and tasks virtually using a single platform. It is a project management software that will the tasks easily. You can just give a project or task id and their names to get access to the account and make changes in it.

The software is totally written in Go language. This repository contains six directories and in these six directories, many files are present that perform different functions. Using alphabetical order. 
- The first directory is **database** that will handle the database stuff and I used **MongoDB** database in this code. 
- The second directory is **entrance** that will handle the signup and login functionalities. 
- The third directory is **profile** that will show the account information to the user. 
- The fourth directory is **projectWork** that will handle all the projects and tasks. 
- The fifth directory is **taskChoice** that will give the user a choice to select any function.
- The sixth directory is **vendor** that contains all files related to **MongoDB** and that directory is created using ```go mod vendor```

---

### Table of Contents

The headers will be used to reference the location of destination.

- [The First Directory](#the-first-directory)
- [The Second Directory](#the-second-directory)
- [The Third Directory](#the-third-directory)
- [The Fourth Directory](#the-fourth-directory)
- [The Fifth Directory](#the-fifth-directory)
- [Author Info](#author-info)

---

# The First Directory
In the first directory, I created one file that is **db.go**. It contains twelve functions that will handle different tasks.
- The first function is ```Connection()```. It will make the connection to database that is **MongoDB**. This function is equal to the ```CNX``` variable. The database ```url``` is also given.
- The second function is ```Insertdata()```. It will enter the data to database. It uses ```InsertOne``` to insert the data.
- The third function is ```Findaccount()```. It will find the account data in database. It uses ```FindOne``` to find the data based on email address and password.
- The fourth function is ```Updateaccount()```. It will update the account details in database. It filtered the data based on password and allow the user to change that data. 
- The fifth functions is ```Deleteaccount()```. It will delete the account in database. It uses ```DeleteOne``` to remove the data from database. 
- The sixth function is ```Findproject()```. It will find the project data in database but this finding will be based on project id and project name.
- The seventh function is ```Updateproject()```. It will update the project details in database but this will be based on project id.
- The eighth function is ```Verifyaccount()```. It will find the account details and after verification, it will delete that account.
- The ninth function is ```Findtask()```. It will find the task data in database but this finding will be based on task id and task name.
- The tenth function is ```Updatetask()```. It will update the task details in database based on task id.
- The eleventh function is ```NameandEmail()```. It will just update the name and email address of a task.
- The twelveth function is ```Uploadfile()```. It will transfer a file(video, audio, etc) to database. Two parameters are used. First one is used as path and another one is used as the exact file name.

---

# The Second Directory
In the second directory, I created two files. One is **login.go** that contains one function and another is **signup.go** that also contains one function.
- The first file function is ```Login()``` that will take your email address and password and find that details in database to allow you to login.

Now come to the second file **signup.go** 

- The second file function is ```Signup()``` that will take your name, eamil address, password, and business name to make an account and insert that data in database.

---

# The Third Directory
In the third directory, I created one file that is **info.go**. It contains one function.
- The function is ```AccountInfo()``` that will show the information that a user already gave when signing up and show more options to update and at the end, it will give a choice also to close an account.

---

# The Fourth Directory
In the fourth directory, I created two file. One is **project.go** that contains one function and another is **task.go** that also contains one function.
- The first file function is ```Project()``` that will give three choices to the user, that are create, update and delete a project. 
- For project creation, it imports ```Insertdata()``` function to insert data in database.
- For project updating, it imports  ```Findproject()``` function to find the data in database and ```Updateproject()``` function to update the data in database.
- For project deletion, it imports ```Verifyaccount()``` function to verify an account and delete it.
- The second file function is ```Task()``` that will give two choices to the user, that are add and edit a task.
- Here the ```Insertdata()``` function is also imported and instead of ```Findproject()``` and ```Updateproject()```, I imported ```Findtask()``` and ```Updatetask()```. 
- Two more functions are also imported in **task.go** that are ```Uploadfile()``` that will uplaod files to database and ```NameandEmail()``` that will update name and email address.

---

# The Fifth Directory
In the fifth directory, I created one file that is ```TaskOption()```. This will provide a list of tasks and give the user a choice to select any one of them.

---

# The Sixth Directory
The sixth directory contains a lot of subdirectories and files. You can make this directory and other files out of the directory using the following commands.
- ```go mod init```
- ```go mod tidy```
- ```go mod vendor```

---

## Author Info

- YouTube - [ibilalkayy](https://www.youtube.com/channel/UCBLTfRg0Rgm4FtXkvql7DRQ)
- LinkedIn - [@ibilalkayy](https://www.linkedin.com/in/ibilalkayy/)

