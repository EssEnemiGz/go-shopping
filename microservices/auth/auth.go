package auth

import (
	"encoding/json"
	"fmt"
	"go-shopping/microservices/common"
	"io"
	"net/http"
	"os"
)

type Person struct {
	Username string `json:"user"`
	Password string `json:"passw"`
}

func Auth(w http.ResponseWriter, r *http.Request) {
	avaible := request.Method_verificator(r, []string{"POST"})
	if !avaible { // If avaible is false
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, err := os.Open("users.json")
	if err != nil {
		fmt.Println("Error opening users.json", err.Error())
	}
	defer file.Close()

	byteInfo, _ := io.ReadAll(file)
	var userInfo []Person
	json.Unmarshal(byteInfo, &userInfo)

	var userReq Person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userReq); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if userReq.Username == "" || userReq.Password == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	for _, actual := range userInfo {
		if userReq.Username == actual.Username || userReq.Password == actual.Password {
			fmt.Fprint(w, "Auth!")
			return
		}
	}

	fmt.Fprint(w, "Error loging your user")
}
