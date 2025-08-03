package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Endpoint struct {
    Url string
    Label  string
}

type HomePageData struct {
    PageTitle string
    Endpoints     []Endpoint
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := HomePageData{
		PageTitle: "GitHub User Activity",
		Endpoints: []Endpoint{
			{Url: "/events/{your_username}", Label: "Display all event activities of a user"},
			{Url: "/repos/{your_username}", Label: "Display all repositories of a user"},
		},
	}
    tmpl.Execute(w, data)
}

func getAllUserActivties(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	resp, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error fetching data: %s", resp.Status), resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// return the body in json format
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)	//  or use this -> w.Write([]byte(body))
}

func getAllUserRepos(w http.ResponseWriter, r * http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	resp, err := http.Get("https://api.github.com/users/" + username + "/repos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// return the body in json format
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
}

func main() {
	r := mux.NewRouter()

    r.HandleFunc("/", homePage).Methods("GET").Schemes("https")
    r.HandleFunc("/events/{username}", getAllUserActivties).Methods("GET").Schemes("https")
    r.HandleFunc("/repos/{username}", getAllUserRepos).Methods("GET").Schemes("https")

	fmt.Println("Server is running...")
    http.ListenAndServe(":3030", r)
}
