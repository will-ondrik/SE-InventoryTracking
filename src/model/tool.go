package model

import (
	"time"
)

type Tool struct {
	Id string
	Name string
	IsAvailable bool
	LastServiced time.Time
	AssignedTo User
	Condition string
	QrCode QrCode
	Job Job
}

