package main

type Address struct {
	AddressID uint   `gorm:"primary_key" json:"address_id"`
	Recip     string `json:"recip"`
	Street    string `json:"street"`
}

type Brand struct {
	BrandID   uint   `gorm:"primary_key" json:"brand_id"`
	BrandName string `gorm:"not null;unique" json:"brand_name"`
}

type Category struct {
	CategoryID   uint   `gorm:"primary_key" json:"cat_id"`
	CategoryName string `gorm:"not null;unique" json:"cat_name"`
	Types        []Type `json:"-"`
}

type Features struct {
	FeaturesID    uint   `gorm:"primary_key"`
	FeaturesDatas string ``
}

func (Features) TableName() string {
	return "features"
}

type IsInCart struct {
	ProductID  uint `gorm:"primary_key"`
	UserID     uint `gorm:"primary_key"`
	CartAmount int  ``
}

type IsInOrder struct {
	OrderID     uint `gorm:"primary_key"`
	ProductID   uint `gorm:"primary_key"`
	OrderAmount int  ``
}

type Order struct {
	OrderID   uint   `gorm:"primary_key"`
	UserID    uint   `gorm:"unique_index:idx_order"`
	AddressID uint   `gorm:"unique_index:idx_order"`
	Date      string `gorm:"unique_index:idx_order;type:date"`
}

type Product struct {
	ProductID     uint    `gorm:"primary_key" json:"product_id"`
	ProductName   string  `gorm:"not null;unique" json:"product_name"`
	Price         float64 `json:"price"`
	Description   string  `json:"desc"`
	Stock         int     `json:"stock"`
	BrandID       uint    `json:"-"`
	BrandName     string  `gorm:"-" json:"brand"`
	TypeID        uint    `json:"-"`
	TypeName      string  `gorm:"-" json:"type"`
	FeaturesID    uint    `json:"-"`
	FeaturesDatas string  `gorm:"-" json:"features"`
}

type Rewiew struct {
	ProductID uint   `gorm:"primary_key"`
	UserID    uint   `gorm:"primary_key"`
	Rate      int    ``
	Comment   string ``
}

type ShippingAddress struct {
	AddressID uint   `gorm:"primary_key" json:"address_id"`
	Recip     string `gorm:"-" json:"recip"`
	Street    string `gorm:"-" json:"street"`
	UserID    uint   `gorm:"primary_key" json:"-"`
}

type Type struct {
	TypeID     uint   `gorm:"primary_key" json:"type_id"`
	TypeName   string `gorm:"not null;unique" json:"type_name"`
	CategoryID uint   `json:"cat_id"`
}

type User struct {
	UserID         uint   `gorm:"primary_key" json:"user_id"`
	Username       string `json:"username"`
	Email          string `gorm:"not null;unique" json:"email"`
	AddressID      uint   `json:"-"`
	RealName       string `gorm:"-" json:"real_name"`
	BillingAddress string `gorm:"-" json:"billing_address"`
	PaymentMethod  string `json:"payment_method"`
	Settings       string `json:"settings"`
}
