package productcontroller

import (
	"net/http"
	"html/template"
	"github.com/Snow-00/my-gocrudweb/pkg/models/productmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("../../views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {

}

func Add(w http.ResponseWriter, r *http.Request) {
	
}

func Edit(w http.ResponseWriter, r *http.Request) {
	
}

func Delete(w http.ResponseWriter, r *http.Request) {
	
}