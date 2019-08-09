package main

import (
	"encoding/json"
	users2 "github.com/jakubgoralewski/workshop/service/users"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		userJson, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		user := &users2.User{}
		json.Unmarshal(userJson, user)

		users2.RepositoryInstance().Add(user)
		users2.RepositoryInstance()
	}

	usersJson, _ := json.Marshal(users2.RepositoryInstance())
	// fmt.Printf("%+v", *users.RepositoryInstance().Users[0])

	w.Write(usersJson)
}