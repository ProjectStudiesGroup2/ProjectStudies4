package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.CreateTable(&Product{})

	products := Products{
		Store: map[string]*Product{},
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/products", products.GetAllProducts),
		rest.Post("/products", products.PostProduct),
		rest.Get("/products/:id", products.GetProduct),
		rest.Put("/products/:id", products.PutProduct),
		rest.Delete("/products/:id", products.DeleteProduct),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type Product struct {
	Id    string
	Name  string
	Price float64
	Desc  string
	Stock int
}

type Products struct {
	sync.RWMutex
	Store map[string]*Product
}

func (p *Products) GetAllProducts(w rest.ResponseWriter, r *rest.Request) {
	p.RLock()
	products := make([]Product, len(p.Store))
	i := 0
	for _, product := range p.Store {
		products[i] = *product
		i++
	}
	p.RUnlock()
	w.WriteJson(&products)
}

func (p *Products) GetProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	p.RLock()
	var product *Product
	if p.Store[id] != nil {
		product = &Product{}
		*product = *p.Store[id]
	}
	p.RUnlock()
	if product == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(product)
}

func (p *Products) PostProduct(w rest.ResponseWriter, r *rest.Request) {
	product := Product{}
	err := r.DecodeJsonPayload(&product)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.Lock()
	id := fmt.Sprintf("%d", len(p.Store)) // stupid
	product.Id = id
	p.Store[id] = &product
	p.Unlock()
	w.WriteJson(&product)
}

func (p *Products) PutProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	p.Lock()
	if p.Store[id] == nil {
		rest.NotFound(w, r)
		p.Unlock()
		return
	}
	product := Product{}
	err := r.DecodeJsonPayload(&product)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		p.Unlock()
		return
	}
	product.Id = id
	p.Store[id] = &product
	p.Unlock()
	w.WriteJson(&product)
}

func (p *Products) DeleteProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	p.Lock()
	delete(p.Store, id)
	p.Unlock()
	w.WriteHeader(http.StatusOK)
}
