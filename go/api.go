package main

import (
	"net/http"
	"time"
"encoding/json"
	m "github.com/gorilla/mux"
)

type ResponseWriter = http.ResponseWriter
type Request = http.Request

func main() {
	r := m.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	println("listening...")
	server.ListenAndServe()
}

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"descrption"`
	CreatedOn   time.Time `json:"createdon"`
}

var (
	noteStore     = make(map[int]Note)
	id        = 0 
)

func PostNoteHandler(w ResponseWriter, r *Request){
 var note Note
 if err := json.NewDecoder(r.Body).Decode(&note); err!=nil{
	 panic(err)
 }
 note.CreatedOn = time.Now()
 id++
 noteStore[id] = note
 w.Header().Set("Content-Type","application/json")
 w.WriteHeader(http.StatusCreated)
 err := json.NewEncoder(w).Encode(note)
 if err!=nil{
	 panic(err)
 }
}

func GetNoteHandler(w ResponseWriter,r *Request){}
func PutNoteHandler(w ResponseWriter, r *Request){}
func DeleteNoteHandler(ResponseWriter,*Request){}