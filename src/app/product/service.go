package product

import "sans-api-product/model"

type Service interface {
	Create(productRequest ProductRequest) (model.Product, error)
	FindByID(id int) (model.Product, error)
	FindLIKE(keyword string) ([]model.Product, error)
	FindAll() ([]model.Product, error)
	Update(id int, productRequest ProductRequest) (model.Product, error)
	Delete(id int) (model.Product, error)
}

// create private model 'service'
type service struct {
	repository Repository
}

// get model 'service' and use get access for all method in this model
func NewService(r Repository) *service {
	return &service{repository: r}
}

// Get Product
func (s *service) FindAll() ([]model.Product, error) {
	result, err := s.repository.FindAll()
	return result, err
}

// Craete Product
func (s *service) Create(productRequest ProductRequest) (model.Product, error) {
	product := model.Product{
		Name:  productRequest.Name,
		Qty:   productRequest.Qty,
		Price: productRequest.Price,
	}
	result, err := s.repository.Create(product)

	return result, err
}

// Get Product
func (s *service) FindByID(id int) (model.Product, error) {
	result, err := s.repository.FindByID(id)
	return result, err
}

// Get Product by name
func (s *service) FindLIKE(keyword string) ([]model.Product, error) {
	result, err := s.repository.FindLIKE(keyword)
	return result, err
}

// Update Product
func (s *service) Update(id int, productRequest ProductRequest) (model.Product, error) {
	product, _ := s.FindByID(id)

	product.Name = productRequest.Name
	product.Price = productRequest.Price
	product.Qty = productRequest.Qty

	result, err := s.repository.Update(product)

	return result, err
}

// Delete Product
func (s *service) Delete(id int) (model.Product, error) {
	product, _ := s.repository.FindByID(id)

	deletedProduct, err := s.repository.Delete(product)

	return deletedProduct, err
}
