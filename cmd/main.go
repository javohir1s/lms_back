package main

import (
	"fmt"
	"lms_back/config"
	"lms_back/controller"
	"lms_back/storage/postgres"
	"net/http"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	con := controller.NewController(store)

	http.HandleFunc("/admin", con.Admin)
	http.HandleFunc("/branch", con.Branch)
	http.HandleFunc("/group", con.Group)
	http.HandleFunc("/payment", con.Payment)
	http.HandleFunc("/schedule", con.Schedule)
	http.HandleFunc("/student", con.Student)
	http.HandleFunc("/teacher", con.Teacher)


	fmt.Println("programm is running on localhost:8080...")
	http.ListenAndServe(":8080", nil)

}
