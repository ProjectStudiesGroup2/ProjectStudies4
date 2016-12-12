package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func (i *Impl) GetProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("product_id")
	product := Product{}
	if i.DB.First(&product, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	brand := Brand{}
	if i.DB.First(&brand, product.BrandID).Error != nil {
		rest.NotFound(w, r)
		return
	}
	product.BrandName = brand.BrandName

	features := Features{}
	if i.DB.First(&features, product.FeaturesID).Error != nil {
		rest.NotFound(w, r)
		return
	}
	product.FeaturesDatas = features.FeaturesDatas

	aType := Type{}
	if i.DB.First(&aType, product.TypeID).Error != nil {
		rest.NotFound(w, r)
		return
	}
	product.TypeName = aType.TypeName

	w.WriteJson(&product)
}

// func (i *Impl) PostProduct(w rest.ResponseWriter, r *rest.Request) {
// 	product := Product{}
//
// 	new := Product{}
// 	if err := r.DecodeJsonPayload(&new); err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	product.ProductName = new.ProductName
// 	product.Price = new.Price
// 	product.Descr = new.Descr
// 	product.Stock = new.Stock
//
// 	if err := i.DB.Save(&product).Error; err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteJson(&product)
// }
//
// func (i *Impl) PutProduct(w rest.ResponseWriter, r *rest.Request) {
// 	id := r.PathParam("product_id")
// 	product := Product{}
// 	if i.DB.First(&product, id).Error != nil {
// 		rest.NotFound(w, r)
// 		return
// 	}
//
// 	updated := Product{}
// 	if err := r.DecodeJsonPayload(&updated); err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	product.ProductName = updated.ProductName
// 	product.Price = updated.Price
// 	product.Descr = updated.Descr
// 	product.Stock = updated.Stock
//
// 	if err := i.DB.Save(&product).Error; err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteJson(&product)
// }

func (i *Impl) DeleteProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("product_id")
	product := Product{}
	if i.DB.First(&product, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := i.DB.Delete(&product).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
