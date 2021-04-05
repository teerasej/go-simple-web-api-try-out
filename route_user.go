package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// app.get('/users')
// app.post('/users')

func GetUsers(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		userResponse := UserModel{
			UserId:   123448,
			Email:    "pon@nextflow.in.th",
			Password: "4567",
			Cart: []Product{
				Product{
					Name:  "iPhone",
					Price: 40000,
				},
			},
		}

		jsonEncoder := json.NewEncoder(w)
		jsonEncoder.Encode(userResponse)

	case http.MethodPost:

		bodyBytes, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Cannot read request body"))
			return
		}

		var newUser UserModel
		err = json.Unmarshal(bodyBytes, &newUser)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Cannot parse json"))
			return
		}

		fmt.Println(newUser)
		w.Write([]byte("/users POST"))
	}

}
