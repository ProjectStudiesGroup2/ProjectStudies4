package main

import (
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

	category.CategoryName = new.CategoryName

	if err := i.DB.Save(&category).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&category)
}

func (i *Impl) GetTypesInCategory(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("cat_id")
	category := Category{}
	if i.DB.First(&category, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	types := []Type{}
	i.DB.Model(&category).Related(&types)
	
	w.WriteJson(&types)
}

func (i *Impl) PutCategory(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("cat_id")
	category := Category{}
	if i.DB.First(&category, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Category{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category.CategoryName = updated.CategoryName

	if err := i.DB.Save(&category).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&category)
}

func (i *Impl) DeleteCategory(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("cat_id")
	category := Category{}
	if i.DB.First(&category, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := i.DB.Delete(&category).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
