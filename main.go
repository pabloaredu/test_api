package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Models
type app struct {
	Router *mux.Router
}

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Testing API",
		Description: "Creating a small api to continue practicing",
	},
}

// Main
func main() {
	app := app{}
	app.initializeApp()
}

// Utils
func (app *app) initializeApp() {
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router.HandleFunc("/", homelink)
	app.Router.HandleFunc("/", createEvent).Methods("POST")
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

// Handlers
func homelink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
	fmt.Fprintf(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO handle error
	}

	json.Unmarshal(body, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)

}
