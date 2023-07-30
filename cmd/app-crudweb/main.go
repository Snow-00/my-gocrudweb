package main

import (
  "github.com/Snow-00/my-gocrudweb/pkg/config"
  "net/http"
  "log"
)

func main() {
  config.ConnectDB()
  
  log.Println("Server running on port: 8000")
  http.ListenAndServe(":8000", nil)
}