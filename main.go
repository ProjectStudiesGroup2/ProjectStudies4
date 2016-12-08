package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
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
	if i.DB.HasTable(&ShippingAddress{}) {
		i.DB.DropTable(&ShippingAddress{})
	}
	if i.DB.HasTable(&Rewiew{}) {
		i.DB.DropTable(&Rewiew{})
	}
	if i.DB.HasTable(&User{}) {
		i.DB.DropTable(&User{})
	}
	if i.DB.HasTable(&Product{}) {
		i.DB.DropTable(&Product{})
	}
	if i.DB.HasTable(&Type{}) {
		i.DB.DropTable(&Type{})
	}
	if i.DB.HasTable(&Features{}) {
		i.DB.DropTable(&Features{})
	}
	if i.DB.HasTable(&Category{}) {
		i.DB.DropTable(&Category{})
	}
	if i.DB.HasTable(&Brand{}) {
		i.DB.DropTable(&Brand{})
	}
	if i.DB.HasTable(&Address{}) {
		i.DB.DropTable(&Address{})
	}

	// IsInCart IsInOrder Order

	i.DB.CreateTable(&Address{})

	i.DB.CreateTable(&Brand{})

	i.DB.CreateTable(&Category{})

	i.DB.CreateTable(&Features{})

	i.DB.CreateTable(&Type{})
	i.DB.Model(&Type{}).AddForeignKey("category_id", "categories(category_id)", "CASCADE", "CASCADE")

	i.DB.CreateTable(&Product{})
	i.DB.Model(&Product{}).AddForeignKey("brand_id", "brands(brand_id)", "RESTRICT", "RESTRICT")
	i.DB.Model(&Product{}).AddForeignKey("type_id", "types(type_id)", "RESTRICT", "RESTRICT")
	i.DB.Model(&Product{}).AddForeignKey("features_id", "features(features_id)", "RESTRICT", "RESTRICT")

	i.DB.CreateTable(&User{})
	i.DB.Model(&User{}).AddForeignKey("address_id", "addresses(address_id)", "RESTRICT", "RESTRICT")

	i.DB.CreateTable(&Rewiew{})
	i.DB.Model(&Rewiew{}).AddForeignKey("product_id", "products(product_id)", "CASCADE", "CASCADE")
	i.DB.Model(&Rewiew{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")

	i.DB.CreateTable(&ShippingAddress{})
	i.DB.Model(&Rewiew{}).AddForeignKey("address_id", "addresses(address_id)", "CASCADE", "CASCADE")
	i.DB.Model(&Rewiew{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
}
