package product

import "ecommerce/domain"

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	prod, err := svc.productRepo.Create(p)
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, nil
	}
	return prod, nil
}
