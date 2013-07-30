package infrastructure

import (
    "time"
    "log"
)

func TimeTrack(start time.Time, name string) {
  elapsed := time.Since(start)
  log.Printf("function %s took %s", name, elapsed)
}