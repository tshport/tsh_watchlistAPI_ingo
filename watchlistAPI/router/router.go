package router

import (
	"watchlistAPI/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.AddMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.UpdateIdAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovieId).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
