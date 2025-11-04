package tracking

import (
	"sandbox/straightedge/qr/SE-InventoryTracking/src/category"
	"sandbox/straightedge/qr/SE-InventoryTracking/src/tool"
)

type TrackingService struct {

}

func NewService() *TrackingService {
	return &TrackingService{}
}

func (s *TrackingService) GetAllInventory() ([]category.Category, error) {
	var c []category.Category

	return c, nil
}

func (s *TrackingService) GetAllFromCategory(ctg category.Category) (category.Category, error) {
	var c category.Category

	return c, nil
}

func (s *TrackingService) GetAllFromSubCategory(sCtg string) ([]category.SubCategory, error) {
	var sctg []category.SubCategory

	return sctg, nil
}

func (s *TrackingService) GetTools(tl string) ([]tool.Tool, error) {
	var t []tool.Tool

	return t, nil
}

func (s *TrackingService) GetUniqueTool(tl string) (tool.Tool, error) {
	var t tool.Tool

	return t, nil
}