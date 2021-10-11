package server

import (
	"encoding/json"
	"net/http"
)

// Response structure for every response sent from server
type Response struct {
    Data        interface{}     `json:"data,omitempty"`     // omits if nil (use nil when errors)
    Route       string          `json:"route"`
    Method      string          `json:"method"`
    Status      int             `json:"status"`
    Message     interface{}     `json:"message,omitempty"`  // omits if nil (use nil when no-errors)
}


// Generate & Return JSON format of server.Response type
func NewResponse(w http.ResponseWriter, r Response) ([]byte, error) {
    res, err := json.Marshal(r)
    w.WriteHeader(r.Status)
    return res, err
}

 
// Response : Method Not Allowed
/* Creates a new response of server.Response{} type, with "Method Not Allowed" message, prints the Log to console, and sends the response. */
func MNAResponse(w http.ResponseWriter, r *http.Request)  {
    res, _ := NewResponse(w, Response {
        Data: nil, Route: r.URL.Path, Method: r.Method, Status: http.StatusMethodNotAllowed, Message: "["+r.Method+"] `"+r.URL.Path+"` Not Allowed",
    })
    Log.Printf("[%s] '%s' : Not Allowed\n", r.Method, r.URL.Path)
    w.Write(res)
}


// Response : Not Found
/* Creates a new response of server.Response{} type, with "Endpoint Not Found" message, prints the Log to console, and sends the response. */
func NFResponse(w http.ResponseWriter, r *http.Request)  {
    res, _ := NewResponse(w, Response {
        Data: nil, Route: r.URL.Path, Method: r.Method, Status: http.StatusNotFound, Message: "Endpoint `"+r.URL.Path+"` Not Found",
    })
    Log.Printf("[%s] '%s' : Not Found\n", r.Method, r.URL.Path)
    w.Write(res)
}


// Response : Invalid Query
/* Creates a new response of server.Response{} type, with "Invalid Query" message, prints the Log to console, and sends the response. */
func IVQResponse(w http.ResponseWriter, r *http.Request) {
    res, _ := NewResponse(w, Response {
        Data: nil, Route: r.URL.Path, Method: r.Method, Status: http.StatusNotAcceptable, Message: "Invalid query `?"+r.URL.RawQuery+"` for endpoint `"+r.URL.Path+"`",
    })
    Log.Printf("[%s] '%s' : Invalid Query\n", r.Method, r.URL.Path+"?"+r.URL.RawQuery)
    w.Write(res)
}