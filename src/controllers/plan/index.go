package planController

import (
  "net/http"
  "time"
  "src/infrastructure"  
  "src/infrastructure/views"
)

func Index(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "PlanController.Index")
	_, v := views.StaticLayout("controllers/plan/templates/index")
    w.Write(v)
}