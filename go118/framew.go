package main

import (
	"encoding/json"
	//"io"
	"net/http"
	"time"
	//"net/http/httptest"
	//"os"
	//"strings"
	"log"
	"io"
)

// type server struct {
//     router *someRouter
// }

func main() {
	// req := httptest.NewRequest("POST", "/", strings.NewReader(`{"email":"name@example.com", "name":"John Doe"}`))
	// w := httptest.NewRecorder()

	// handler := createHandler(UpdateUser)
	// handler(w, req)

	// io.Copy(os.Stdout, w.Result().Body)

	mux := http.NewServeMux()
	update := createHandler(UpdateUser)
	hello := createHandler(Hello)

	srv := &http.Server{
		Addr: ":8000",
	    ReadTimeout:  5 * time.Second,
	    WriteTimeout: 10 * time.Second,
	    IdleTimeout:  120 * time.Second,
	    Handler:      mux,
	}
	mux.HandleFunc("/update", update)
	mux.HandleFunc("/hello", hello)
	log.Fatal(srv.ListenAndServe())

	// http.HandleFunc("/hello", createHandler(Hello))
	// log.Fatal(http.ListenAndServe(":8000", nil))
}

// -------------------------------------------
// The "framework"
// -------------------------------------------

type handlerFunc[Req any, Resp any] func(Req) (Resp, error)

func createHandler[Req any, Resp any](f handlerFunc[Req, Resp]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req

		// empty Body will return EOF e.g. for GET request.
		err := json.NewDecoder(r.Body).Decode(&req) 
		if err != nil && err != io.EOF {
			http.Error(w, "error unmarshaling request body", http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		log.Println(req)

		resp, err := f(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode(resp)
	}
}

// -------------------------------------------
// The RPC handler
// -------------------------------------------

type UpdateUserRequest struct {
	Email string  `json:"email"`
	Name  *string `json:"name"`
}

type UpdateUserResponse struct {
	Status string `json:"status"`
	User UpdateUserRequest `json:"user"`
}

func UpdateUser(req UpdateUserRequest) (*UpdateUserResponse, error) {
	// do something with the request

	return &UpdateUserResponse{"ok", req}, nil
}

func Hello(req UpdateUserRequest) (*UpdateUserResponse, error) {
	// do something with the request

	return &UpdateUserResponse{Status: "Hello"}, nil
}
