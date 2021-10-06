package controllers

import (
	"math/rand"
	"net/http"
	"strconv"

	m "github.com/poseidon-code/go-server/models"
	s "github.com/poseidon-code/go-server/server"
)

// method: GET `/user/`
func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(3)

	res, _ := s.NewResponse(w, s.Response{
		Data: u[random],Route: r.URL.Path, Method: r.Method, Status: http.StatusOK, Message: nil,
	})
	w.Write(res)
}


// method: GET `/user/?id=`
func GetUserId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var f m.User
	
	for _, us := range u {
		if strconv.Itoa(us.ID) == id {
			f = us
		}
	}

	if f == (m.User{}) {
		res, _ := s.NewResponse(w, s.Response {
			Data: nil, Route: r.URL.Path, Method: r.Method, Status: http.StatusNotFound, Message: "User with {ID: "+id+"} Not Found",
		})
		w.Write(res)
	} else {
		res, _ := s.NewResponse(w, s.Response {
			Data: f, Route: r.URL.Path, Method: r.Method, Status: http.StatusOK, Message: nil,
		})
		w.Write(res)
	}
}
