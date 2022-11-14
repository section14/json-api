package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var data []byte

func main() {
	//get port as argument
	port := fmt.Sprintf(":%s", os.Args[1])

	//get JSON file as second arguent
	jsonFilename := os.Args[2]

	//validate and parse supplied JSON file
	jsonData, err := validateJson(jsonFilename)
	if err != nil {
		log.Fatal(err)
	}
	data = jsonData

	//create new router for API
	r := mux.NewRouter()

	//headers
	headersOk := handlers.AllowedHeaders([]string{
		"X-Requested-With",
		"Content-Type",
		"User-Token"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//cors handling - limit the url's that can hit this here
	//currently set to all
	originsOk := handlers.AllowedOrigins([]string{"*"})

	//supply router to handlers
	routes(r)

	//set up listener
	l, err := net.Listen("tcp4", port)
	if err != nil {
		log.Fatal(err)
	}

	//config server
	server := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(r),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	//launch server
	log.Fatal(server.Serve(l))
}

func routes(r *mux.Router) {
	r.HandleFunc("/", TestRoute).Methods("GET")
	r.HandleFunc("/data", ServeJson).Methods("GET")
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	log.Println("API is running")
	w.Write([]byte("API is up and running"))
}

func ServeJson(w http.ResponseWriter, r *http.Request) {
	w.Write(data)
}

func validateJson(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	ok := json.Valid(file)
	if !ok {
		return []byte{}, errors.New("invalid JSON format")
	}

	return file, nil
}
