package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"strconv"
)

func (i *Impl) GetProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("product_id")
	product := Product{}
	if i.DB.First(&product, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	if product.FillProduct(i) != nil {
		rest.NotFound(w, r)
	}

	w.WriteJson(&product)
}

func (i *Impl) PostProductInType(w rest.ResponseWriter, r *rest.Request) {
	product := Product{}

	new := Product{}
	err := r.DecodeJsonPayload(&new)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product.ProductName = new.ProductName
	product.Price = new.Price
	product.Description = new.Description
	product.Stock = new.Stock

	brandId, err := strconv.ParseUint(new.BrandName, 10, 32)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product.BrandID = uint(brandId)

	typeId, err := strconv.ParseUint(r.PathParam("type_id"), 10, 32)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product.TypeID = uint(typeId)

	newFeat := Features{}
	newFeat.FeaturesDatas = new.FeaturesDatas
	if err := i.DB.Save(&newFeat).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product.FeaturesID = newFeat.FeaturesID

	if err := i.DB.Save(&product).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if product.FillProduct(i) != nil {
		rest.NotFound(w, r)
	}
	w.WriteJson(&product)
}

func (i *Impl) PutProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("product_id")
	product := Product{}
	if i.DB.First(&product, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	update := Product{}
	err := r.DecodeJsonPayload(&update)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product.ProductName = update.ProductName
	product.Price = update.Price
	product.Description = update.Description
	product.Stock = update.Stock

	brandId, err := strconv.ParseUint(update.BrandName, 10, 32)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product.BrandID = uint(brandId)

	typeId, err := strconv.ParseUint(update.TypeName, 10, 32)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product.TypeID = uint(typeId)

	updateFeat := Features{}
	if i.DB.First(&updateFeat, product.FeaturesID).Error != nil {
		rest.NotFound(w, r)
		return
	}
	updateFeat.FeaturesDatas = update.FeaturesDatas
	if err := i.DB.Save(&updateFeat).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := i.DB.Save(&product).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if product.FillProduct(i) != nil {
		rest.NotFound(w, r)
	}
	w.WriteJson(&product)
}

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

func (p *Product) FillProduct(i *Impl) error {
	brand := Brand{}
	if err := i.DB.First(&brand, p.BrandID).Error; err != nil {
		return err
	}
	p.BrandName = brand.BrandName

	features := Features{}
	if err := i.DB.First(&features, p.FeaturesID).Error; err != nil {
		return err
	}
	p.FeaturesDatas = features.FeaturesDatas

	aType := Type{}
	if err := i.DB.First(&aType, p.TypeID).Error; err != nil {
		return err
	}
	p.TypeName = aType.TypeName

	return nil
}
