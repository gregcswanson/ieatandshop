package recipesController

import (
  "net/http"
  "time"
  "src/infrastructure"  
  "src/infrastructure/views"
)

func View(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "RecipesController.View")
	_, v := views.StaticLayout("controllers/recipes/templates/view")
    w.Write(v)
}