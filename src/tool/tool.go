package tool

import (
	"sandbox/straightedge/qr/SE-InventoryTracking/src/job"
	qrcode "sandbox/straightedge/qr/SE-InventoryTracking/src/qr_code"
	"sandbox/straightedge/qr/SE-InventoryTracking/src/user"
	"time"
)


type Tool struct {
	Id string
	Name string
	IsAvailable bool
	LastServiced time.Time
	AssignedTo user.User
	Condition string
	QrCode qrcode.QrCode
	Job job.Job
}

