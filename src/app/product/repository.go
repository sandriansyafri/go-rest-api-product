package product

import (
	"sans-api-product/model"

	"gorm.io/gorm"
)

// interface for product
type Repository interface {
	FindAll() ([]model.Product, error)
	FindByID(id int) (model.Product, error)
	FindLIKE(keyword string) ([]model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(product model.Product) (model.Product, error)
}

//create private  model for repository
type repository struct {
	db *gorm.DB
}

// get visibility for private  method (repository) and use all method in model for service
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Create product
func (r *repository) Create(product model.Product) (model.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

//Update Product
func (r *repository) Update(product model.Product) (model.Product, error) {
	err := r.db.Save(product).Error

	return product, err
}

//Get All Product
func (r *repository) FindAll() ([]model.Product, error) {
	products := []model.Product{}
	err := r.db.Find(&products).Error

	return products, err
}

// Get Product by ID
func (r *repository) FindByID(id int) (model.Product, error) {
	product := model.Product{}
	err := r.db.First(&product, id).Error

	return product, err
}

// Get Product  By name and use %name% for get data products
func (r *repository) FindLIKE(keyword string) ([]model.Product, error) {
	products := []model.Product{}

	err := r.db.Where("name LIKE ?", "%"+keyword+"%").Find(&products).Error

	return products, err

}

func (r *repository) Delete(product model.Product) (model.Product, error) {
	err := r.db.Delete(product).Error

	return product, err
}
