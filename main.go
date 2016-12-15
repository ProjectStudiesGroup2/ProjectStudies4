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
	// i.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/categories", i.GetAllCategories),
		rest.Post("/categories", i.PostCategory),
		rest.Get("/categories/:cat_id", i.GetTypesInCategory),
		rest.Put("/categories/:cat_id", i.PutCategory),
		rest.Post("/categories/:cat_id", i.PostTypeInCategory),
		rest.Delete("/categories/:cat_id", i.DeleteCategory),

		rest.Get("/types/:type_id", i.GetProductsInType),
		rest.Put("/types/:type_id", i.PutType),
		rest.Post("/types/:type_id", i.PostProductInType),
		rest.Delete("/types/:type_id", i.DeleteType),

		rest.Get("/brands", i.GetAllBrands),
		rest.Post("/brands", i.PostBrand),
		rest.Get("/brands/:brand_id", i.GetBrand),
		rest.Put("/brands/:brand_id", i.PutBrand),
		rest.Delete("/brands/:brand_id", i.DeleteBrand),

		rest.Get("/products/:product_id", i.GetProduct),
		rest.Put("/products/:product_id", i.PutProduct),
		rest.Delete("/products/:product_id", i.DeleteProduct),

		rest.Post("/users", i.PostUser),
		rest.Get("/users/:user_id", i.GetUser),
		rest.Put("/users/:user_id", i.PutUser),
		rest.Delete("/users/:user_id", i.DeleteUser),

		rest.Get("/users/:user_id/addresses", i.GetAllAddresses),
		rest.Post("/users/:user_id/addresses", i.PostAddress),
		rest.Get("/users/addresses/:addr_id", i.GetAddress),
		rest.Put("/users/addresses/:addr_id", i.PutAddress),
		rest.Delete("/users/addresses/:addr_id", i.DeleteAddress),
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
		log.Fatalf("Got error when connect database: %v", err)
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
	if i.DB.HasTable(&IsInOrder{}) {
		i.DB.DropTable(&IsInOrder{})
	}
	if i.DB.HasTable(&Order{}) {
		i.DB.DropTable(&Order{})
	}
	if i.DB.HasTable(&IsInCart{}) {
		i.DB.DropTable(&IsInCart{})
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

	i.DB.CreateTable(&IsInCart{})
	i.DB.Model(&IsInCart{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	i.DB.Model(&IsInCart{}).AddForeignKey("product_id", "products(product_id)", "CASCADE", "CASCADE")

	i.DB.CreateTable(&Order{})
	i.DB.Model(&Order{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	i.DB.Model(&Order{}).AddForeignKey("address_id", "addresses(address_id)", "RESTRICT", "RESTRICT")

	i.DB.CreateTable(&IsInOrder{})
	i.DB.Model(&IsInOrder{}).AddForeignKey("product_id", "products(product_id)", "RESTRICT", "RESTRICT")
	i.DB.Model(&IsInOrder{}).AddForeignKey("order_id", "orders(order_id)", "CASCADE", "CASCADE")

	i.DB.CreateTable(&Rewiew{})
	i.DB.Model(&Rewiew{}).AddForeignKey("product_id", "products(product_id)", "CASCADE", "CASCADE")
	i.DB.Model(&Rewiew{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")

	i.DB.CreateTable(&ShippingAddress{})
	i.DB.Model(&ShippingAddress{}).AddForeignKey("address_id", "addresses(address_id)", "CASCADE", "CASCADE")
	i.DB.Model(&ShippingAddress{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
}
