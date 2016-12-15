package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"strconv"
)

func (i *Impl) GetAllAddresses(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("user_id")
	addresses := []ShippingAddress{}
	i.DB.Raw("SELECT * FROM shipping_addresses WHERE user_id = ?", id).Scan(&addresses)
	for idx := range addresses {
		addresses[idx].FillShippingAddress(i)
	}
	w.WriteJson(&addresses)
}

func (i *Impl) PostAddress(w rest.ResponseWriter, r *rest.Request) {
	address := Address{}

	new := Address{}
	if err := r.DecodeJsonPayload(&new); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	address.Recip = new.Recip
	address.Street = new.Street

	if err := i.DB.Save(&address).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shipping := ShippingAddress{}
	id, err := strconv.ParseUint(r.PathParam("user_id"), 10, 32)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shipping.AddressID = address.AddressID
	shipping.UserID = uint(id)

	if err := i.DB.Save(&shipping).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&address)
}

func (i *Impl) GetAddress(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("addr_id")
	address := Address{}
	if i.DB.First(&address, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&address)
}

func (i *Impl) PutAddress(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("addr_id")
	address := Address{}
	if i.DB.First(&address, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Address{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	address.Recip = updated.Recip
	address.Street = updated.Street

	if err := i.DB.Save(&address).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&address)
}

func (i *Impl) DeleteAddress(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("addr_id")
	address := Address{}
	if i.DB.First(&address, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := i.DB.Delete(&address).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sa *ShippingAddress) FillShippingAddress(i *Impl) error {
	address := Address{}
	if err := i.DB.First(&address, sa.AddressID).Error; err != nil {
		return err
	}
	sa.Recip = address.Recip
	sa.Street = address.Street

	return nil
}
