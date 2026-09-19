package main

import (
	"encoding/json"
	"fmt"
	"go_train/models"
	"go_train/services"
	"net/http"
	"time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func timer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Since(start)

		fmt.Println("TIMER:", r.Method, r.URL.Path, "took", elapsed)
	})
}

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

	getPlayers := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "List of all players")
	}

	getPlayer := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintln(w, "Get player with id", id)
	}

	addPlayer := func(w http.ResponseWriter, r *http.Request) {
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
	}

	deletePlayer := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintln(w, "Deleted player with id: ", id)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", helloWorld)
	mux.HandleFunc("GET /players", getPlayers)
	mux.HandleFunc("GET /players/{id}", getPlayer)
	mux.HandleFunc("POST /players", addPlayer)
	mux.HandleFunc("DELETE /players/{id}", deletePlayer)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", timer(logger(mux)))

}
