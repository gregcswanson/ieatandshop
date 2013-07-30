package shoppingController

import (
  "net/http"
  "time"
  "src/infrastructure"  
  "src/infrastructure/views"
)

func Index(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "ShoppingController.Index")
	_, v := views.StaticLayout("controllers/shopping/templates/index")
    w.Write(v)
}