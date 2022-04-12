package routes

import (
	"net/http"

	c "github.com/poseidon-code/go-server/controllers"
	s "github.com/poseidon-code/go-server/server"
)

// User `/user/` route handler for all http Methods
func User(w http.ResponseWriter, r *http.Request) {
    s.Headers(&w)                            // include Request Headers

    query := r.URL.Query()
    
    if r.URL.Path == "/user/" {             // handle `/user/` route
        switch r.Method {                   // handle methods on `/user/`
        case http.MethodGet:                // handle GET on `/user/`
            if query["id"]!=nil && len(query)==1 {          // use query 'id' (`/user/?id=`)
                c.GetUserId(w, r)
            } else if len(query)==0 {                       // default (`/user/`)
                c.GetRandomUser(w, r)
            } else {                                        // invalid query on GET `/user/`
                s.IVQResponse(w, r)
            }
        default:                            // handle other unhandled methods
            s.MNAResponse(w, r)
        }
    } else {                                // there is no nested paths of `/user/`
        s.NFResponse(w, r)
    }
}