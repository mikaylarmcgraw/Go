package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Chore struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Age   int    `json:"age"`
}

func allChores(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	chores := Chores{
		Chore{Title: "Make bed", Desc: "", Age: 5},
		Chore{Title: "Pick up clothes", Desc: "Pick up all clothes on floor and put in basket", Age: 2},
		Chore{Title: "Get dressed", Desc: "", Age: 3},
		Chore{Title: "Bursh hair", Desc: "", Age: 4},
	}
	fmt.Println(("Endpoint Hit: All Articles Endpoint"))
	json.NewEncoder(w).Encode(chores)
}

type Chores []Chore

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/chores", allChores)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	handleRequest()
}
