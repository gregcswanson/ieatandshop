package homeController

import (
  "net/http"
  "time"
  "src/infrastructure" 
)

func FiveZeroZero(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "controllers.fivezerozero")
  w.Write([]byte(fiveZeroZeroLayout))
}

const fiveZeroZeroLayout = `
<!doctype html>
<head>
	<title>Crez GO App Engine</title>
    <link rel="stylesheet" href="/stylesheets/style.min.css" />
</head>
<body>
  <h1>500</h1>
  <script src="/javascripts/vendor/angular.min.js" ></script>
</body>
</html>
`