package model


type Category struct {
	Name string
	SubCategory SubCategory
}

type SubCategory struct {
	Name string
	Tools []Tool
}