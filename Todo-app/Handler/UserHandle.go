package Handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo/Database"
	"todo/Models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating the user")

	var user Models.Users

	if r.Body == nil {
		w.Write([]byte("Insert the information for registartion"))
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error hai babau bhaiay")
	}

	fmt.Println(user)

	queryString := `INSERT INTO regisuser (username, email, password)  VALUES ($1,$2,$3)`

	fmt.Println(queryString)

	res, err := Database.DBConnection.Exec(queryString, user.UserName, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error in db execution")
		return
	}

	json.NewEncoder(w).Encode("User Created Sucessfully")

	fmt.Println(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Deleting the user")

}
