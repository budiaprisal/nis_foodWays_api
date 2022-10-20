package productdto

type ProductRequest struct {
	Title  string `json:"title" form:"desc" validate:"required"`
	Price  int    `json:"price" form:"price" validate:"required"`
	Image  string `json:"image" form:"image" validate:"required"`
	UserID int    `json:"user" form:"user_id" validate:"required"`
}
