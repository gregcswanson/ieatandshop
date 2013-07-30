package routes

import (
  "net/http"
  "src/controllers/home"
  "src/controllers/recipes"
  "src/controllers/users"
  "src/controllers/shopping"
  "src/controllers/plan"
  "src/controllers/echo"
)

func Configure() {
  http.HandleFunc("/", MainPage)
  http.HandleFunc("/login", homeController.Login)
  http.HandleFunc("/logout", homeController.Logout)
  http.HandleFunc("/plan/edit", planController.Edit)
  http.HandleFunc("/styleguide", homeController.Styleguide)
  http.HandleFunc("/user", usersController.Current)
  http.HandleFunc("/recipes", recipesController.Index)
  http.HandleFunc("/recipes/add", recipesController.Add)
  http.HandleFunc("/recipes/view", recipesController.View)
  http.HandleFunc("/shopping", shoppingController.Index)
  http.HandleFunc("/echo", echoController.Index)
  http.HandleFunc("/404", homeController.FourZeroFour)
  http.HandleFunc("/500", homeController.FiveZeroZero)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" || r.URL.Path != "/" {
		homeController.FourZeroFour(w, r)
  } else {
    homeController.Index(w, r)
  }
  return
}