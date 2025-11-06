package service

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"

	errors "sandbox/straightedge/qr/SE-InventoryTracking/src/error"
	"sandbox/straightedge/qr/SE-InventoryTracking/src/model"

	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type QrService struct {
	BaseUrl string
}

func NewQrService(companyUrl string) *QrService {
	return &QrService{
		BaseUrl: companyUrl,
	}
}

// Create builds the QR data for a tool and generates the QR code.
func (s *QrService) Create(t model.Tool, d model.QrData, brand model.CompanyBrand) error {
	// Content URL - Scanning this redirects to check-in/checkout
	d.Content = fmt.Sprintf("%s/tool/%s", s.BaseUrl, t.Id)

	// TODO: Record QR code/data in db via repo layer
	return s.generateQrCode(d, brand)
}

func (s *QrService) Read()           {}
func (s *QrService) Update()         {}
func (s *QrService) Delete(t model.Tool) {}

// generateQrCode builds the QR image + canvas.
// TODO: Further test and modify company branding and tool info (custom fonts?)
func (s *QrService) generateQrCode(q model.QrData, brand model.CompanyBrand) error {
	qr, err := qrcode.New(q.Content, q.RecoveryLevel)
	if err != nil {
		log.Printf("%s: %v", errors.QrGenerationError, err)
		return fmt.Errorf(errors.QrGenerationError)
	}


	// Base QR image
	qrImg := qr.Image(q.Size)

	// Total canvas height (QR + label/branding area)
	canvasHeight := s.getCanvasSize(q.Size)

	// New canvas
	out := image.NewRGBA(image.Rect(0, 0, q.Size, canvasHeight))

	// White background
	draw.Draw(out, out.Bounds(), image.White, image.Point{}, draw.Src)

	// Draw QR at the top
	draw.Draw(out, qrImg.Bounds().Add(image.Pt(0, 0)), qrImg, image.Point{}, draw.Over)

	// Centered tool label (slightly above bottom so brand can sit at very bottom)
	labelCanvas := model.QrCanvas{
		Img:       out,
		Label:     q.Content,
		Face:      basicfont.Face7x13,
		Col:       color.Black,
		Width:     q.Size,
		BaselineY: canvasHeight - 40, // moved up a bit to leave room for brand
	}
	s.addCenterText(labelCanvas)

	// Company logo + name at bottom-left
	brandCanvas := model.QrCanvas{
		Img:   out,
		Label: brand.Name,
		Face:  basicfont.Face7x13,
		Col:   color.Black,
		Width: q.Size,
	}
	if err := s.addCompanyBranding(brand, brandCanvas); err != nil {
		return err
	}

	return nil
}

// addCompanyBranding draws the company name/logo in the bottom left corner of the Qr canvas
func (s *QrService) addCompanyBranding(b model.CompanyBrand, c model.QrCanvas) error {
	logoImg, _, err := image.Decode(bytes.NewReader(b.Logo))
	if err != nil {
		log.Printf("%s: %v", errors.ImageDecodeError, err)
		return fmt.Errorf(errors.ImageDecodeError)
	}

	// Resize logo to ~10% of canvas width (may need to decrease)
	targetLogoWidth := c.Width / 10
	if targetLogoWidth <= 0 {
		targetLogoWidth = 32
	}
	resizedLogo := resize.Resize(uint(targetLogoWidth), 0, logoImg, resize.Lanczos2)

	logoBounds := resizedLogo.Bounds()
	logoWidth := logoBounds.Dx()
	logoHeight := logoBounds.Dy()

	padding := 8
	imgBounds := c.Img.Bounds()

	// Bottom-left position for logo
	logoMin := image.Point{
		X: padding,
		Y: imgBounds.Max.Y - logoHeight - padding,
	}
	logoMax := image.Point{
		X: logoMin.X + logoWidth,
		Y: logoMin.Y + logoHeight,
	}
	dstRect := image.Rectangle{Min: logoMin, Max: logoMax}

	// Draw logo at bottom-left
	draw.Draw(c.Img, dstRect, resizedLogo, image.Point{}, draw.Over)

	// Now draw company name to the right of logo, baseline aligned near bottom of logo
	textCanvas := c
	textCanvas.BaselineY = logoMax.Y - 4

	textStartX := dstRect.Max.X + padding
	s.addBottomLeftText(textCanvas, textStartX)

	return nil
}

// addBottomLeftText draws label left-aligned starting at startX, baseline at c.BaselineY
func (s *QrService) addBottomLeftText(c model.QrCanvas, startX int) {
	d := &font.Drawer{
		Dst:  c.Img,
		Src:  image.NewUniform(c.Col),
		Face: c.Face,
	}

	// Fallback if BaselineY is zero
	baselineY := c.BaselineY
	if baselineY == 0 {
		bounds := c.Img.Bounds()
		baselineY = bounds.Max.Y - 10
	}

	d.Dot = fixed.P(startX, baselineY)
	d.DrawString(c.Label)
}

// addCenterText draws a single-line label that is centered horizontally at BaselineY
func (s *QrService) addCenterText(c model.QrCanvas) {
	d := &font.Drawer{
		Dst:  c.Img,
		Src:  image.NewUniform(c.Col),
		Face: c.Face,
	}

	// Measure the text width in pixels
	textWidth := d.MeasureString(c.Label).Round()

	x := (c.Width - textWidth) / 2
	if x < 0 {
		x = 0
	}

	d.Dot = fixed.P(x, c.BaselineY)
	d.DrawString(c.Label)
}

// getCanvasSize uses the typed QrSize + LabelPad (LabelPad is used to add the correct canvas space for branding/tool info)
func (s *QrService) getCanvasSize(size int) int {
	heightF := float64(size) * (1 + float64(model.LabelPad))
	return int(heightF + 0.5) // rounded
}
