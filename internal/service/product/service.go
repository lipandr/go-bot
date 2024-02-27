package product

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) GetProduct(idx int) (*Product, error) {
	if idx > len(s.List())-1 {
		return nil, errors.New("product is not found")
	}
	return &allProducts[idx], nil
}

// 1 2 3 4 5
// 0 1 2 3 4
// 6
