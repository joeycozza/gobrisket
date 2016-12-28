package main

import (
	"net/http"

	"github.com/bahlo/goat"
)

//User totally rocks
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	Initialize()
	r := goat.New()
	r.Get("/users", "users_url", usersHandler)
	r.Run(":3000")
}

func usersHandler(w http.ResponseWriter, r *http.Request, p goat.Params) {
	u := User{
		Name: "Jimmy",
		Age:  29,
	}
	goat.WriteJSON(w, u)
}
