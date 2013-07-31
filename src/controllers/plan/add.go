package planController

import (
  "net/http"
  "time"
  "src/infrastructure"  
  "src/infrastructure/views"
)

func Add(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "PlanController.Add")
	_, v := views.StaticLayout("src/controllers/plan/templates/add")
    w.Write(v)
}