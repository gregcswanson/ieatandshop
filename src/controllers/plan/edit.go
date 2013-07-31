package planController

import (
  "net/http"
  "time"
  "src/infrastructure"  
  "src/infrastructure/views"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "PlanController.Edit")
	_, v := views.StaticLayout("src/controllers/plan/templates/edit")
    w.Write(v)
}