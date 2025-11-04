package category

import "sandbox/straightedge/qr/SE-InventoryTracking/src/tool"

type Category struct {
	Name string
	SubCategory SubCategory
}

type SubCategory struct {
	Name string
	Tools []tool.Tool
}