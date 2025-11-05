package model

import (
	"image"
	"image/color"

	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
)

// Recovery Level
// Low, Medium, and High
// Low -- Damaged code less likely to scan -- increases with level
	type QrData struct {
	Content string
	RecoveryLevel qrcode.RecoveryLevel 
	Size int
}

type QrCanvas struct {
	Img *image.RGBA
	Label string
	Face font.Face
	Col color.Color
	Width int
	BaselineY int
}

type QrCode struct {}

type QrSize int 
type LabelRatio float64

const (
	QrSmall QrSize = 256
	QrMedium QrSize = 384
	QrLarge QrSize = 512
	LabelPad LabelRatio = 0.17
)