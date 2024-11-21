package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Users = make(map[int]User)
var CreateCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "http_Create_request_total",
	Help: "Total number of Create HTTP request",
})

var requestCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "http_request_total",
	Help: "Total number of HTTP request",
})

func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(CreateCount)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/create", CreateUser)
	http.HandleFunc("/get", ReadUser)
	http.HandleFunc("/update", UpdateUser)
	http.HandleFunc("/delete", DeleteUser)

	fmt.Println("Server started at :8080")
	Users[1] = User{1, "Senya"}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	CreateCount.Inc()
	log.Println("Received request: ", req.URL.Path)
	log.Println("CREATING USER")
	var user User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		log.Fatal(err)
		return
	}

	for {
		user.ID = rand.Intn(1000)
		if _, exist := Users[user.ID]; !exist {
			Users[user.ID] = user
			break
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	log.Printf("USER [%s] HAS BEEN CREATED\n", user.Name)
	// CreateCount.Dec()
}
func ReadUser(w http.ResponseWriter, req *http.Request) {
	requestCount.Inc()
	log.Println("Received request: ", req.URL.Path)
	log.Println("READING USER")

	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
		return
	}
	user, exist := Users[id]
	if !exist {
		log.Fatal("no user")
		return
	}
	json.NewEncoder(w).Encode(user)
	log.Printf("USER [%s] HAS BEEN READED\n", user.Name)
	requestCount.Dec()
}
func UpdateUser(w http.ResponseWriter, req *http.Request) {
	log.Println("Received request: ", req.URL.Path)
	log.Println("UPDATING USER")

	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
		return
	}
	_, exist := Users[id]
	if !exist {
		log.Fatal("no user")
		return
	}

	var user User

	if err = json.NewDecoder(req.Body).Decode(&user); err != nil {
		log.Fatal(err)
		return
	}
	user.ID = id
	Users[id] = user
	json.NewEncoder(w).Encode(user)
	log.Printf("USER [%s] HAS BEEN UPDATED\n", user.Name)

}
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	log.Println("Received request: ", req.URL.Path)
	log.Println("DELETING USER")

	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	delete(Users, id)
	w.WriteHeader(http.StatusNoContent)
	log.Printf("USER [%s] HAS BEEN DELETED\n", Users[id].Name)

}
