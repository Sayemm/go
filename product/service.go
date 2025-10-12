package product

import "ecommerce/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

// Create Product - If logic is big - new file (create.go)

func (svc *service) Get(id int) (*domain.Product, error) {
	prod, err := svc.productRepo.Get(id)
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, nil
	}
	return prod, nil
}
func (svc *service) List() ([]*domain.Product, error) {
	prodList, err := svc.productRepo.List()
	if err != nil {
		return nil, err
	}
	if prodList == nil {
		return nil, nil
	}
	return prodList, nil
}
func (svc *service) Delete(id int) error {
	err := svc.productRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (svc *service) Update(product domain.Product) (*domain.Product, error) {
	prod, err := svc.productRepo.Update(product)

	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, nil
	}
	return prod, nil
}
