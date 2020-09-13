package Repository

import (

	// JSon Parser
	"encoding/json"
	"fmt"

	// package used to covert string into int type
	"strconv"

	// import userModel

	UserController "go-postgres/controller"
	userModel "go-postgres/models"

	"log"

	//  access the request and response object of the api
	"net/http"

	"github.com/gorilla/mux"
)

// ================================== Repositories to handle controllers =========================================

// CreateUser create a user in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user userModel.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Error  %v", err)
	}

	// call insert user function
	insertID := UserController.InsertUser(user)

	// format a response object
	res := response{
		UserID:  insertID,
		Message: "Пользователь успешно добавлен.",
	}

	// Response
	json.NewEncoder(w).Encode(res)
}

// GetUser Get single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request
	params := mux.Vars(r)

	// call the get User function
	user, err := UserController.GetUser(params["user_id"])

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	// Response
	json.NewEncoder(w).Encode(user)
}

// Get All the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {

	// retreive all the users in db
	users, err := UserController.GetAllUsers()

	if err != nil {
		log.Fatalf("Unable to get users. %v", err)
	}

	// send all the user details
	json.NewEncoder(w).Encode(users)
}

// UpdateUser
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error.  %v", err)
	}

	var user userModel.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update user
	updatedRows := UserController.UpdateUser(int64(id), user)
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	deletedRows := UserController.DeleteUser(int64(id))
	msg := fmt.Sprintf("Removed User successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
