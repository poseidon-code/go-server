package routes

import (
	"net/http"

	c "github.com/poseidon-code/go-server/controllers"
	s "github.com/poseidon-code/go-server/server"
)

// Home `/` route handler for all http Methods
func Home(w http.ResponseWriter, r *http.Request) {
    s.Headers(w)                    // include Request Headers

    if r.URL.Path == "/" {
        switch r.Method {           // handle methods on `/`
        case http.MethodOptions:
        case http.MethodGet:        // handle GET on `/`
            c.GetHome(w, r)
        default:                    // handle other unhandled methods
            s.MNAResponse(w, r)
        }
    } else {                        // there is no nested paths of `/`
        s.NFResponse(w, r)
    }
}