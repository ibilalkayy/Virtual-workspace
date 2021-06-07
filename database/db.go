package database

/* Importing libraries */
import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Importing variables in struct */
type AccountVariables struct {
	Name, Email, Password, Business string
}

type ProjectVariables struct {
	ID, Projectname, Ownername, Startdate, Enddate, Description, Groupname, Status string
}

type TaskVariables struct {
	ID, Projectname, Taskname, Description, Startdate, Enddate, Workhours string
}

var Project ProjectVariables
var Account AccountVariables
var Task TaskVariables

/* Connection to database */
var CNX = Connection()

func Connection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

/* Inserting data in database */
func Insertdata(inputData interface{}) {
	collection := CNX.Database("VWS").Collection("dataStored")
	collection.InsertOne(context.TODO(), inputData)
}

/*------------------------- Handling account information -------------------------*/

/* Finding account data in database*/
func Findaccount(myEmail, myPassword string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	if err := collection.FindOne(context.TODO(), bson.M{"email": myEmail, "password": myPassword}).Decode(&Account); err != nil {
		log.Fatal(err)
	}
}

/* Upading account data in database */
func Updateaccount(key, value string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	filter := bson.M{"password": Account.Password}
	update := bson.M{
		"$set": bson.M{
			key: value,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

/* Removing account from database */
func Deleteaccount(key, value string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	_, err := collection.DeleteOne(context.TODO(), bson.M{key: value})
	if err != nil {
		log.Fatal(err)
	}
}

/*------------------------- Handling project information -------------------------*/
func Findproject(projectId, projectName string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	if err := collection.FindOne(context.TODO(), bson.M{"id": projectId, "projectname": projectName}).Decode(&Project); err != nil {
		log.Fatal(err)
	}
}

func Updateproject(key, value string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	filter := bson.M{"id": Project.ID}
	update := bson.M{
		"$set": bson.M{
			key: value,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

/* Verifying an account data and deleting it */
func Verifyaccount(myPassword, value string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	if err := collection.FindOne(context.TODO(), bson.M{"password": myPassword}).Decode(&Account); err != nil {
		log.Fatal(err)
	}
	Deleteaccount("id", value)
}

/*------------------------- Handling task information -------------------------*/
func Findtask(taskId, taskName string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	if err := collection.FindOne(context.TODO(), bson.M{"id": taskId, "taskname": taskName}).Decode(&Task); err != nil {
		log.Fatal(err)
	}
}

func Updatetask(key, value string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	filter := bson.M{"id": Task.ID}
	update := bson.M{
		"$set": bson.M{
			key: value,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

/* Entering name and email to database */
func NameandEmail(key, value string) {
	collection := CNX.Database("VWS").Collection("dataStored")
	filter := bson.M{"id": Task.ID}
	update := bson.M{
		"$push": bson.M{
			key: value,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

/* Uploading files to database */
func Uploadfile(filepath, filename string) {
	data, err := ioutil.ReadFile(filepath + filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	bucket, err := gridfs.NewBucket(
		CNX.Database("VWS"),
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	}
	log.Printf("Your data is succesfully uploaded. File size is %d M\n", fileSize)
}
