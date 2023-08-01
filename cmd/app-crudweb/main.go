package main

import (
  "github.com/Snow-00/my-gocrudweb/pkg/config"
  "net/http"
  "log"
  "github.com/Snow-00/my-gocrudweb/pkg/controllers/homecontroller"
)

func main() {
  config.ConnectDB()

  // 1. homepage
  http.HandleFunc("/", homecontroller.Welcome)
  
  log.Println("Server running on port: 8000")
  http.ListenAndServe(":8000", nil)
}