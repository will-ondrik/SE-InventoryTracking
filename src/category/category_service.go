package category

import "sandbox/straightedge/qr/SE-InventoryTracking/src/model"

type CategoryService struct {}

func NewService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) Create() (bool, error) {
	return true, nil
}

func (s *CategoryService) Read() ([]model.Category, error) {
	var c []model.Category

	return c, nil
}

func (s *CategoryService) Update() (bool, error) {
	return true, nil
}

func (s *CategoryService) Delete() (bool, error) {
	return true, nil
}