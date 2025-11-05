package model

import (
	"time"
)

type Tool struct {
	Id string	`json:"id,omitempty"`
	Name string	`json:"name"`
	SubCategoryId SubCategory `json:"subCategoryId"`
	IsAvailable bool `json:"isAvailable"`
	LastServiced time.Time `json:"lastServiced"`
	AssignedTo User `json:"assignedTo"`
	Condition string `json:"condition"`
	QrCode QrCode `json:"qrCode,omitempty"`
	Job Job `json:"job"`
}

