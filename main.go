package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	i := Impl{}
	i.InitDB()
	i.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/products", i.GetAllProducts),
		rest.Post("/products", i.PostProduct),
		rest.Get("/products/:id", i.GetProduct),
		rest.Put("/products/:id", i.PutProduct),
		rest.Delete("/products/:id", i.DeleteProduct),
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
	Id        int64     `json:"id"`
	Name      string    `sql:"size:256" json:"name"`
}

type Impl struct {
	DB *gorm.DB
}

func (i *Impl) InitDB() {
	var err error
	i.DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(true)
}

func (i *Impl) InitSchema() {
	i.DB.AutoMigrate(&Product{})
}

func (i *Impl) GetAllProducts(w rest.ResponseWriter, r *rest.Request) {
    products := []Product{}
    i.DB.Find(&products)
    w.WriteJson(&products)
}

func (i *Impl) GetProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	product := Product{}
	if i.DB.First(&product, id).Error == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&product)
}

func (i *Impl) PostProduct(w rest.ResponseWriter, r *rest.Request) {
	product := Product{}
	if err := r.DecodeJsonPayload(&product); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := i.DB.Save(&product).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.WriteJson(&product)
}

func (i *Impl) PutProduct(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")
	product := Product{}
	if i.DB.First(&product, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Product{}
	if err := r.DecodeJsonPayload(&product); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("    \\/")
	fmt.Println(updated.Name)
	fmt.Println("    /\\")
	product.Name = updated.Name

	if err := i.DB.Save(&product).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&product)
}

func (i *Impl) DeleteProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
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
