package main

import (
  "github.com/Snow-00/my-gocrudweb/pkg/config"
  "net/http"
  "log"
  "github.com/Snow-00/my-gocrudweb/pkg/controllers/homecontroller"
  "github.com/Snow-00/my-gocrudweb/pkg/controllers/categorycontroller"
  "github.com/Snow-00/my-gocrudweb/pkg/controllers/productcontroller"
)

func main() {
  config.ConnectDB()

  // 1. homepage
  http.HandleFunc("/", homecontroller.Welcome)
  
  // 2. Categories
  http.HandleFunc("/categories", cateogrycontroller.Index)
  http.HandleFunc("/categories/add", cateogrycontroller.Add)
  http.HandleFunc("/categories/edit", cateogrycontroller.Edit)
  http.HandleFunc("/categories/delete", cateogrycontroller.Delete)
  
  // 3. Products
  http.HandleFunc("/products", productcontroller.Index)
  http.HandleFunc("/products/add", productcontroller.Add)
  http.HandleFunc("/products/detail", productcontroller.Detail)
  http.HandleFunc("/products/edit", productcontroller.Edit)
  http.HandleFunc("/products/delete", productcontroller.Delete)

  log.Println("Server running on port: 8000")
  http.ListenAndServe(":8000", nil)
}