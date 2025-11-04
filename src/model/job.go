package model

type Job struct {
	JobberNumber int
	Address Address
}

type Address struct {
	StreetNumber int
	StreetName string
	City string
	Province string
	Country string
}