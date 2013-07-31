package recipesController

import (
  "net/http"
  "time"
  "src/infrastructure"  
  "src/infrastructure/views"
)

func Add(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "RecipesController.Add")
	_, v := views.StaticLayout("src/controllers/recipes/templates/add")
    w.Write(v)
}