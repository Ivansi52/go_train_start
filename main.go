package main

import (
	"encoding/json"
	"fmt"
	"go_train/models"
	"go_train/services"
	"net/http"
)

func main() {
	service := services.PlayerService{}

	service.AddPlayer(models.Player{
		Nickname: "Ivan",
		Level:    20,
		Gold:     1500,
		Online:   true,
	})

	service.AddPlayer(models.Player{
		Nickname: "Oleg",
		Level:    33,
		Gold:     800,
		Online:   false,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		nickname := r.URL.Query().Get("nickname")
		if nickname == "" {
			//status code bad request
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Nickname is empty")
			return
		}

		player, err := service.FindPlayer(nickname)
		if err != nil {
			//not found
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Not Found")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonErr := json.NewEncoder(w).Encode(player)
		if jsonErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "JSON Encode Error")
		}
	})

	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		player := models.Player{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&player)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Decode Error")
			return
		}
		service.AddPlayer(player)
		json.NewEncoder(w).Encode(player)
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
