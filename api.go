package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	store         Storage
}

type ApiError struct {
	Error string `json:"error"`
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

	//Mux's Router is used to define routes
	router := mux.NewRouter()

	//func signature to be passed as 2nd arg has to be of type HTTP handler  i.e  func(http.ResponseWriter, *http.Request) exactly, without the error response
	//(or HandlerFunc as defined in HTTP Package - type HandlerFunc func(ResponseWriter, *Request))
	//so we convert our internal api func to the handler type
	router.HandleFunc("/account", apiFuncToHTTPHandler(s.handleAccount))
	router.HandleFunc("/account/{id}", apiFuncToHTTPHandler(s.handleAccountById))
	router.HandleFunc("/transfer", apiFuncToHTTPHandler(s.handleTransfer))

	log.Println("JSON api running on port", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}

//Handler for Collection Operations
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	//handle HTTP methods
	switch r.Method {

	case "GET":
		return s.handleGetAccount(w, r)

	case "POST":
		return s.handleCreateAccount(w, r)

	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

//Handler for Item Operations
func (s *APIServer) handleAccountById(w http.ResponseWriter, r *http.Request) error {

	//handle HTTP methods
	switch r.Method {

	case "GET":
		return s.handleGetAccountById(w, r)

	case "DELETE":
		return s.handleDeleteAccountById(w, r)

	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {

	//fetch vars from uri/body
	vars := mux.Vars(r)
	idStr := vars["id"] //Id is string

	//Convert to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	account, err := s.store.GetAccountById(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, 200, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	//this checks if the input body matches the struct type

	// createAccountRequest := &CreateAccountRequest{}
	createAccountRequest := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}
	defer r.Body.Close()

	//Create new account using data
	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccountById(w http.ResponseWriter, r *http.Request) error {
	//fetch vars from uri/body
	vars := mux.Vars(r)
	idStr := vars["id"] //Id is string

	//Convert to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	err = s.store.DeleteAccount(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Account with id %d deleted", id)})
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferRequest := new(TransferRequest)
	if err := json.NewDecoder(r.Body).Decode(transferRequest); err != nil {
		return err
	}
	defer r.Body.Close()

	return WriteJSON(w, http.StatusOK, transferRequest)

}
