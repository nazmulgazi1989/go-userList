package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func main() {
	UserList = []User{
		{Name: "Tushar Mahmud", Email: "tushar@gmail.com", Phone: "+8544454342424", City: "Dhaka"},
		{Name: "John Doe", Email: "john@gmail.com", Phone: "+4646564343", City: "Khulna"},
	}
	handleRequests()
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	City  string `json:"city"`
}

var UserList []User

func returnAllUserList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(UserList)
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", returnAllUserList)
	log.Fatal(http.ListenAndServe(":4000", nil))

	// make a post api to add new user in golang

}
