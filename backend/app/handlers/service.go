package handlers

import (
	"net/http"
	"tech-db-forum/pkg/database"
	"tech-db-forum/pkg/network"
)

func Clear(w http.ResponseWriter, r *http.Request) {
	database.Clear()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Очистка базы успешно завершена"))
}

func Status(w http.ResponseWriter, r *http.Request) {
	status := database.GetStatus()
	network.WriteResponse(w, http.StatusOK, status)
}
