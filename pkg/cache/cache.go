package cache

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/valkey-io/valkey-go"
	"github.com/will-x86/r-place/pkg/helper"
	"github.com/will-x86/r-place/pkg/models"
	"github.com/will-x86/r-place/pkg/vk"
)

var can *PixelCanvas
var cachedPng []byte

type PixelCanvas struct {
	img *image.RGBA
}

func (pc *PixelCanvas) updatePngCache() {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, pc.img); err != nil {
		panic(fmt.Sprintf("failed to encode png: %v", err))
	}
	cachedPng = buf.Bytes()
}

func GetCachedPng() []byte {
	return cachedPng
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func SetColor(x, y int, hex string) error {
	c, err := ParseHexColor(hex)
	if err != nil {
		return err
	}

	can.SetColor(x, y, c)
	return nil
}
func (pc *PixelCanvas) SetColor(x, y int, color color.RGBA) {
	pc.img.Set(x, y, color)
	pc.updatePngCache()
}
func (pc *PixelCanvas) SetPixel(x, y int, r, g, b, a uint8) {
	c := color.RGBA{r, g, b, a}
	pc.img.Set(x, y, c)
	pc.updatePngCache()
}
func InitialCache() error {
	max_x, max_y := helper.GetMaxXY()
	can = NewPixelCanvas(max_x, max_y)
	for x := range max_x {
		for y := range max_y {
			hex, err := vk.GetSingleCanvasValue(context.Background(), models.Canvas{
				X: x,
				Y: y,
			})
			if err != nil {
				if !valkey.IsValkeyNil(err) {
					return err
				}
				hex = "#FFFFFF"
			}
			col, err := ParseHexColor(hex)
			if err != nil {
				return err
			}
			can.img.Set(x, y, col)
		}
	}
	can.updatePngCache()
	return nil
}

func NewPixelCanvas(width, height int) *PixelCanvas {
	rect := image.Rect(0, 0, width, height)

	img := image.NewRGBA(rect)

	return &PixelCanvas{
		img: img,
	}
}
