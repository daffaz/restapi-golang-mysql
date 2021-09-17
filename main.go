package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-mysql/models"
	"golang-mysql/movie"
	"golang-mysql/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		context, cancel := context.WithCancel(context.Background())
		defer cancel()

		movies, err := movie.GetAll(context)

		if err != nil {
			log.Println(err)
		}

		utils.ResponseJSON(w, movies, http.StatusOK)
		return
	}
	http.Error(w, "Error...", http.StatusNotFound)
}

func PostMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-type") != "application/json" {
			http.Error(w, "Only accepting application/json request", http.StatusBadRequest)
			return
		}
		context, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mov models.Movie
		if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
		}

		if err := movie.InsertMovie(context, mov); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
		}

		response := map[string]string{
			"message": "Succes",
		}

		utils.ResponseJSON(w, response, http.StatusCreated)
		return
	}
	http.Error(w, "Error...", http.StatusNotFound)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if r.Method == "PUT" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Only accepting application/json request", http.StatusBadRequest)
			return
		}

		context, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mov models.Movie
		if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		idMovie := params.ByName("id")

		if err := movie.UpdateMovieById(context, mov, idMovie); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, response, http.StatusCreated)
		return
	}
	http.Error(w, "Error...", http.StatusNotFound)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method == "DELETE" {
		context, cancel := context.WithCancel(context.Background())
		defer cancel()

		idMovie := params.ByName("id")
		if err := movie.DeleteMovie(context, idMovie); err != nil {
			response := map[string]string{
				"message": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, response, http.StatusBadRequest)
			return
		}
		response := map[string]string{
			"message": "Succesfully delete data",
		}
		utils.ResponseJSON(w, response, http.StatusOK)
		return
	}
	http.Error(w, "Error...", http.StatusNotFound)
}

func main() {
	// Inisialisasi
	router := httprouter.New()

	router.GET("/movie", GetMovie)
	router.POST("/movie", PostMovie)
	router.PUT("/movie/:id", UpdateMovie)
	router.DELETE("/movie/:id", DeleteMovie)

	log.Println("Running in port 10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}
