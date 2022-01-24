package product

// Create model for Product Request and using this for validation
type ProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Qty   int    `json:"qty" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
