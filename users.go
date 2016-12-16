package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func (i *Impl) PostUser(w rest.ResponseWriter, r *rest.Request) {
	user := User{}

	new := User{}
	if err := r.DecodeJsonPayload(&new); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Username = new.Username
	user.Email = new.Email
	user.PaymentMethod = new.PaymentMethod
	user.Settings = new.Settings

	newAddress := Address{}
	newAddress.Recip = new.RealName
	newAddress.Street = new.BillingAddress
	if err := i.DB.Save(&newAddress).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.AddressID = newAddress.AddressID

	if err := i.DB.Save(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.FillUser(i) != nil {
		rest.NotFound(w, r)
	}

	shipping := ShippingAddress{}
	shipping.AddressID = newAddress.AddressID
	shipping.UserID = user.UserID
	if err := i.DB.Save(&shipping).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&user)
}

func (i *Impl) GetUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("user_id")
	user := User{}
	if i.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	if user.FillUser(i) != nil {
		rest.NotFound(w, r)
	}

	w.WriteJson(&user)
}

func (i *Impl) PutUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("user_id")
	user := User{}
	if i.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	update := User{}
	if err := r.DecodeJsonPayload(&update); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Username = update.Username
	user.Email = update.Email
	user.PaymentMethod = update.PaymentMethod
	user.Settings = update.Settings

	updateAddress := Address{}
	if i.DB.First(&updateAddress, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	updateAddress.Recip = update.RealName
	updateAddress.Street = update.BillingAddress
	if err := i.DB.Save(&updateAddress).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := i.DB.Save(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.FillUser(i) != nil {
		rest.NotFound(w, r)
	}
	w.WriteJson(&user)
}

func (i *Impl) DeleteUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("user_id")
	user := User{}
	if i.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := i.DB.Delete(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *User) FillUser(i *Impl) error {
	address := Address{}
	if err := i.DB.First(&address, u.AddressID).Error; err != nil {
		return err
	}
	u.RealName = address.Recip
	u.BillingAddress = address.Street

	return nil
}
