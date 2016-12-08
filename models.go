package main

type Address struct {
	AddressID uint   `gorm:"primary_key"`
	Recip     string ``
	Street    string ``
}

type Brand struct {
	BrandID   uint   `gorm:"primary_key"`
	BrandName string `gorm:"not null;unique"`
}

type Category struct {
	CategoryID   uint   `gorm:"primary_key"`
	CategoryName string `gorm:"not null;unique"`
}

type Features struct {
	FeaturesID    uint   `gorm:"primary_key"`
	FeaturesDatas string ``
}

func (Features) TableName() string {
	return "features"
}

type IsInCart struct {
	ProductID uint `gorm:"primary_key"`
	UserID    uint `gorm:"primary_key"`
}

type IsInOrder struct {
	OrderID   uint `gorm:"primary_key"`
	ProductID uint `gorm:"primary_key"`
}

type Order struct {
	UserID    uint   `gorm:"primary_key"`
	AddressID uint   `gorm:"primary_key"`
	Date      string `gorm:"primary_key;type:date"`
}

type Product struct {
	ProductID   uint    `gorm:"primary_key" json:"product_id"`
	ProductName string  `gorm:"not null;unique" json:"product_name"`
	Price       float64 ``
	Descr       string  ``
	Stock       int     ``
	BrandID     uint    ``
	TypeID      uint    ``
	FeaturesID  uint    ``
}

type Rewiew struct {
	ProductID uint   `gorm:"primary_key"`
	UserID    uint   `gorm:"primary_key"`
	Rate      int    ``
	Comment   string ``
}

type ShippingAddress struct {
	AddressID uint `gorm:"primary_key"`
	UserID    uint `gorm:"primary_key"`
}

type Type struct {
	TypeID     uint   `gorm:"primary_key"`
	TypeName   string `gorm:"not null;unique"`
	CategoryID uint   ``
}

type User struct {
	UserID        uint   `gorm:"primary_key"`
	Username      string ``
	Email         string `gorm:"not null;unique"`
	AddressID     uint   ``
	PaymentMethod string ``
	Settings      string ``
}
