package tool

import "sandbox/straightedge/qr/SE-InventoryTracking/src/model"


type ToolService struct {
	// DB
}

func NewService() *ToolService {
	return &ToolService{}
}

func (s *ToolService) Create(t model.Tool) (bool, error) {
	return true, nil
}

func (s *ToolService) Read() ([]model.Tool, error){
	var tool []model.Tool

	return tool, nil
}

func (s *ToolService) Update(t model.Tool) (bool, error) {
	return true, nil
}

func (s *ToolService) Delete() (bool, error) {
	return true, nil
}

