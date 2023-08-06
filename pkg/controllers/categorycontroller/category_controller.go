package cateogrycontroller

import (
	"net/http"
	"github.com/Snow-00/my-gocrudweb/pkg/models/categorymodel"
	"github.com/Snow-00/my-gocrudweb/pkg/entities"
	"html/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any {
		"categories": categories,
	}

	temp, err := template.ParseFiles("../../views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	// mengarah ke tampilan tambah category
	if r.Method == "GET" {
		temp, err := template.ParseFiles("../../views/category/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}
	
	// action utk simpan data
	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("../../views/category/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	
}

func Delete(w http.ResponseWriter, r *http.Request) {
	
}