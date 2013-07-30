package infrastructure

import (
	"net/http"
	"encoding/json"
	"src/vendor/gorilla/schema"
)

// TO DO: Write tests using http://golang.org/pkg/net/http/httptest/

var decoder = schema.NewDecoder() // global decoder that caches struct reflections

type JSendResponse struct {
	Status	string
	Message	string
	Data	interface{}
}

func(j *JSendResponse) SendError(err error, w http.ResponseWriter) {
	j.Status = "error"
  j.Message = err.Error()
  b, _ := json.Marshal(j)
	w.Write(b)
}

func SendError(err error, w http.ResponseWriter) {
	//j := JSendResponse {"error", err.Error(), ""}
  	//b, _ := json.Marshal(j)
  	http.Error(w, err.Error(), http.StatusInternalServerError)
	//w.Write(b)
}

/*func(j *JSendResponse) SendData(data interface{}, w http.ResponseWriter) {
	j.Status = "success"
  j.Data = data
  b, _ := json.Marshal(j)
	w.Write(b)
}*/

func SendData(data interface{}, w http.ResponseWriter) {
	j := JSendResponse {"success", "", data}
  	b, _ := json.Marshal(j)
  	w.Write(b)
}

func SendDataWithMessage(data interface{}, message string, w http.ResponseWriter) {
	j := JSendResponse {"success", message, data}
  	b, _ := json.Marshal(j)
  	w.Write(b)
}