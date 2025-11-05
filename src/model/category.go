package model

import "time"


type Category struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SubCategory struct {
	Id string `json:"id,omitempty"`
	CategoryId string `json:"categoryId"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}