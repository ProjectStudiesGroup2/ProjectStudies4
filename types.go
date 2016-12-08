package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"strconv"
)

func (i *Impl) PostTypeInCategory(w rest.ResponseWriter, r *rest.Request) {
	aType := Type{}

	new := Type{}
	if err := r.DecodeJsonPayload(&new); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseUint(r.PathParam("cat_id"), 10, 32)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	aType.TypeName = new.TypeName
	aType.CategoryID = uint(id)

	if err := i.DB.Save(&aType).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&aType)
}
