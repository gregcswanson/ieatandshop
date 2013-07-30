package routes

import (
  "net/http"
  "src/controllers/users"
  "src/controllers/echo"
)

func ConfigureAPI() {
  http.HandleFunc("/api/user", UserApiRoutes)
  http.HandleFunc("/api/echo/post", echoController.Post)
  http.HandleFunc("/api/echo/get", echoController.Get)
}

func UserApiRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"  {
		usersController.Current(w, r)
	} else if r.Method == "POST" {
		usersController.Current(w, r)
  	}
}