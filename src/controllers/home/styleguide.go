package homeController

import (
  "net/http"
  "time"
  "src/infrastructure"
  "src/infrastructure/views"
)

func Styleguide(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "controllers.styleguide")
    _, v := views.StaticLayout("src/controllers/home/templates/styleguide")
  w.Write(v)
}