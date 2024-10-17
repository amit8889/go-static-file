package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.Info("==============Main function started=========")
	fileServe := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServe)
	http.HandleFunc("POST /submit", handleSubmit)
	done := make(chan os.Signal, 1)
	go func() {
		err := http.ListenAndServe("0.0.0.0:8001", nil)
		if err != nil {
			slog.Error("Error starting server:", err)
		}
	}()
	<-done

}
func handleSubmit(w http.ResponseWriter, r *http.Request) {
	slog.Info("=========form submitter=============")
	w.Header().Set("Content-type", "application/json")
	err := r.ParseForm()
	if err != nil {
		slog.Error("Error decoding JSON:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Error decoding JSON",
		})
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	slog.Info("Name:", name)
	slog.Info("Email:", email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Form submitted successfully",
		"name":    name,
		"email":   email,
	})

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Welcome to the server",
	})
}
