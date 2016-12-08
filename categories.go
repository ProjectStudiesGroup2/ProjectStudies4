package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func (i *Impl) GetAllCategories(w rest.ResponseWriter, r *rest.Request) {
	categories := []Category{}
	i.DB.Find(&categories)
	w.WriteJson(&categories)
}

func (i *Impl) PostCategory(w rest.ResponseWriter, r *rest.Request) {
	category := Category{}

	new := Category{}
	if err := r.DecodeJsonPayload(&new); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(new)
	category.CategoryName = new.CategoryName

	if err := i.DB.Save(&category).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&category)
}
