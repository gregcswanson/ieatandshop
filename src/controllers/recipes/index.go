package recipesController

import (
  "net/http"
  "time"
  "src/infrastructure"  
  "src/infrastructure/views"
)

func Index(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "RecipesController.Index")
	_, v := views.StaticLayout("controllers/recipes/templates/index")
    w.Write(v)
}