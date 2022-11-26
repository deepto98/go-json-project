package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	store         Storage
}

type ApiError struct {
	Error string
}

//Func to return JSON
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

/*
this is the func signature of the API functions we define internally - func(http.ResponseWriter, *http.Request) error())
wrapped into the type  apiFunc so it can be easily passed to apiFuncToHTTPHandler
to convert into http.HandlerFunc required by router.HandleFunc
*/
type apiFunc func(http.ResponseWriter, *http.Request) error

func apiFuncToHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func newAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *APIServer) Run() {

	//Mux's ROuter is used to define routes
	router := mux.NewRouter()

	//func signature to be passed as 2nd arg has to be of type HTTP handler  i.e  func(http.ResponseWriter, *http.Request) exactly, without the error response
	//(or HandlerFunc as defined in HTTP Package - type HandlerFunc func(ResponseWriter, *Request))
	//so we convert our internal api func to the handler type
	router.HandleFunc("/account", apiFuncToHTTPHandler(s.handleAccount))
	router.HandleFunc("/account/{id}", apiFuncToHTTPHandler(s.handleGetAccount))

	log.Println("JSON api running on port", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	//handle HTTP methods
	switch r.Method {

	case "GET":
		return s.handleGetAccount(w, r)

	case "POST":
		return s.handleCreateAccount(w, r)

	case "DELETE":
		return s.handleDeleteAccount(w, r)

	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	//fetch vars from uri/body
	// vars := mux.Vars(r)
	// id := vars["id"]

	account := NewAccount("Deepto", "Gopher")
	return WriteJSON(w, 200, account)
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil

}
