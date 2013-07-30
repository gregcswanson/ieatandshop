package homeController

import (
  "net/http"
  "time"
  "src/infrastructure" 
  "src/infrastructure/views"
)

func FourZeroFour(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "controllers.fourzerofour")
  _, v := views.Static("404")
  w.Write(v)
}