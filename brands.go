package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func (i *Impl) GetAllBrands(w rest.ResponseWriter, r *rest.Request) {
	brands := []Brand{}
	i.DB.Find(&brands)
	w.WriteJson(&brands)
}

func (i *Impl) PostBrand(w rest.ResponseWriter, r *rest.Request) {
	brand := Brand{}

	new := Brand{}
	if err := r.DecodeJsonPayload(&new); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	brand.BrandName = new.BrandName

	if err := i.DB.Save(&brand).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&brand)
}

func (i *Impl) GetBrand(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("brand_id")
	brand := Brand{}
	if i.DB.First(&brand, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&brand)
}

func (i *Impl) PutBrand(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("brand_id")
	brand := Brand{}
	if i.DB.First(&brand, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Brand{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	brand.BrandName = updated.BrandName

	if err := i.DB.Save(&brand).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&brand)
}

func (i *Impl) DeleteBrand(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("brand_id")
	brand := Brand{}
	if i.DB.First(&brand, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := i.DB.Delete(&brand).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
