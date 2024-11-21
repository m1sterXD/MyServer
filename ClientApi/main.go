package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const UserAPIBaseURL = "http://server:8080"

func CreateUser(user User) (*User, error) {
	url := fmt.Sprintf("%s/create", UserAPIBaseURL)

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	var CreatedUser User

	err = json.NewDecoder(resp.Body).Decode(&CreatedUser)
	return &CreatedUser, err
}
func getById(id int) (*User, error) {
	url := fmt.Sprintf("%s/get?id=%d", UserAPIBaseURL, id)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("User not found")
	}

	var foundUser User

	err = json.NewDecoder(resp.Body).Decode(&foundUser)
	if err != nil {
		return nil, err
	}
	return &foundUser, nil
}
func main() {

	user := User{
		Name: "Senya",
	}
	for {
		CraetedUser, err := CreateUser(user)

		if err != nil {
			log.Fatal(err, "error creating user\n")
			return
		}

		fetched, err := getById(CraetedUser.ID)
		if err != nil {
			log.Fatal(err, "error fetching user\n")
			return
		}

		fmt.Printf("fetched user: %v\n", fetched)
		time.Sleep(3*time.Second)
	}

}
