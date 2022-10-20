package models

type Product struct {
	ID     int           `json:"id" gorm:"primary_key:auto_increment"`
	Title  string        `json:"name" form:"name" gorm:"type: varchar(255)"`
	Price  int           `json:"price" form:"price" gorm:"type: int"`
	Image  string        `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int           `json:"userId"`
	User   UsersResponse `json:"user"`
}

type ProductResponse struct {
	ID     int           `json:"id"`
	Title  string        `json:"Title"`
	Price  int           `json:"price"`
	Image  string        `json:"image"`
	UserID int           `json:"-"`
	User   UsersResponse `json:"user"`
}
type ProductUserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	UserID int    `json:"-"`
	User   int    `json:"user"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
