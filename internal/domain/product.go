package domain

type Product struct {
	Name  string `json:"name" binding:"required"`
	Price string `binding:"required" json:"price"`
	Stock string `binding:"required" json:"stock"`
}
