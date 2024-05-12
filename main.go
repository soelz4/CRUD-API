package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies = make([]Movie, 0)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	// Save given ID in /movies/{id} in params
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// Using append function to combine two slices
			// first slice is the slice of all the elements before the given index
			// second slice is the slice of all the elements after the given index
			// append function appends the second slice to the end of the first slice
			// returning a slice, so we store it in the form of a slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	// Save given ID in /movies/{id} in params
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Set json Content-Type
	w.Header().Set("Content-Type", "apllication/json")
	// params
	params := mux.Vars(r)
	// Loop Over the Movies with Range
	// Delete Movie with ID
	// Add a New Movie - the Movie that we Send in the Body of POSTMAN
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}

func main() {
	// HTTP Request Multiplexer
	r := mux.NewRouter()

	// Initialize movie Slice
	movies = append(
		movies,
		Movie{
			ID:       "1",
			Isbn:     "45621",
			Title:    "Movie One",
			Director: &Director{FirstName: "David", LastName: "Fincher"},
		},
		Movie{
			ID:       "2",
			Isbn:     "87356",
			Title:    "Movie Two",
			Director: &Director{FirstName: "Martin", LastName: "Scorsese"},
		},
	)

	// GET All Movies
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// GET Movie by ID
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// CREATE Movie
	r.HandleFunc("/movies/", createMovie).Methods("POST")
	// UPDATE Movie by ID
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// DELETE Movie by ID
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at PORT 8000")
	// Create Server - err = Error or Null
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
