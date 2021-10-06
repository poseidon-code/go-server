package controllers

import (
	"net/http"

	m "github.com/poseidon-code/go-server/models"
	s "github.com/poseidon-code/go-server/server"
)

var u = []m.User{
    {ID: 1, Name: m.UserName{FirstName: "Jane", LastName: "Doe"}, Age: 24,},
    {ID: 2, Name: m.UserName{FirstName: "Alice", LastName: "Merlin"}, Age: 12,},
    {ID: 3, Name: m.UserName{FirstName: "John", LastName: "Doe"}, Age: 32,},
}


// method: GET `/`
func GetHome(w http.ResponseWriter, r *http.Request) {
	res, _ := s.NewResponse(w, s.Response {
			Data: u, Route: r.URL.Path, Method: r.Method, Status: http.StatusOK, Message: nil,
		})
	s.Log.Printf("[%s] '%s'\n", r.Method, r.URL.Path)
	w.Write(res)
}
