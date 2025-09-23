package canvas

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/valkey-io/valkey-go"
	"github.com/will-x86/r-place/pkg/cache"
	"github.com/will-x86/r-place/pkg/helper"
	"github.com/will-x86/r-place/pkg/models"
	"github.com/will-x86/r-place/pkg/vk"
)

// Delete request for deleting square
type DeleteRequest struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

func DeleteCanvasSection(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("invalid json body: %w", err), http.StatusBadRequest)
		return
	}

	// ensure x1 < x2 and y1 < y2
	if req.X1 > req.X2 {
		req.X1, req.X2 = req.X2, req.X1
	}
	if req.Y1 > req.Y2 {
		req.Y1, req.Y2 = req.Y2, req.Y1
	}

	maxX, maxY := helper.GetMaxXY()
	if req.X1 < 0 || req.Y1 < 0 || req.X2 >= maxX || req.Y2 >= maxY {
		helper.ReturnJsonError(w, fmt.Errorf("coordinates are out of bounds"), http.StatusBadRequest)
		return
	}

	// set each pixel to white
	for x := req.X1; x <= req.X2; x++ {
		for y := req.Y1; y <= req.Y2; y++ {
			pixel := models.Canvas{X: x, Y: y, Hex: "#FFFFFF"}
			// Set in Valkey
			if err := vk.SetCanvasValue(r.Context(), pixel); err != nil {
				// best-effort deletion
				log.Printf("Error deleting pixel at (%d, %d): %v", x, y, err)
			}
			cache.SetColor(pixel.X, pixel.Y, pixel.Hex)
		}
	}

	log.Printf("Admin cleared section from (%d, %d) to (%d, %d)", req.X1, req.Y1, req.X2, req.Y2)
	w.WriteHeader(http.StatusNoContent) // 204 No Content DELETE
}

// Get a value on the canvas
func GetCanvas(w http.ResponseWriter, r *http.Request) {
	png := cache.GetCachedPng()
	w.Header().Add("Content-Type", "image/png")
	w.Write(png)
}

// Set a value on the canvas query params
func SetCanvasQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	xStr := query.Get("x")
	if xStr == "" {
		helper.ReturnJsonError(w, fmt.Errorf("missing required parameter 'x'"), http.StatusBadRequest)
		return
	}
	x, err := strconv.Atoi(xStr)
	if err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("invalid 'x' parameter: must be an integer"), http.StatusBadRequest)
		return
	}

	yStr := query.Get("y")
	if yStr == "" {
		helper.ReturnJsonError(w, fmt.Errorf("missing required parameter 'y'"), http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("invalid 'y' parameter: must be an integer"), http.StatusBadRequest)
		return
	}

	hex := query.Get("hex")

	c := models.Canvas{
		X:   x,
		Y:   y,
		Hex: hex,
	}

	if err := vk.SetCanvasValue(r.Context(), c); err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("Error setting database %w", err), http.StatusInternalServerError)
		return
	}

	cache.SetColor(c.X, c.Y, c.Hex)

	response, _ := json.Marshal(map[string]any{"ok": "ok"})
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	log.Printf("Successfully wrote to co-ords X:%d , Y:%d with hex %s\n", c.X, c.Y, c.Hex)
}

// Set a value on the canvas
func SetCanvas(w http.ResponseWriter, r *http.Request) {
	var c models.Canvas
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("Error with json input, see error: %w", err), http.StatusBadRequest)
		return
	}
	if err := vk.SetCanvasValue(r.Context(), c); err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("Error setting database %w", err), http.StatusInternalServerError)
		return
	}
	cache.SetColor(c.X, c.Y, c.Hex)

	// Ignoring error, I'll cut my socks off if this errors
	response, _ := json.Marshal(map[string]any{"ok": "ok"})
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	log.Printf("Successfully wrote to co-ords X:%d , Y:%d with hex %s\n", c.X, c.Y, c.Hex)

}

// Get a single pixel
func GetSingle(w http.ResponseWriter, r *http.Request) {
	// If they want to use params:
	var c models.Canvas
	log.Println("url params:", r.URL.Query().Get("x"), r.URL.Query().Get("y"))
	if r.URL.Query().Get("x") != "" && r.URL.Query().Get("y") != "" {
		x, err := strconv.Atoi(r.URL.Query().Get("x"))
		if err != nil {
			helper.ReturnJsonError(w, fmt.Errorf("Invalid query params %v", err), http.StatusBadRequest)
			return
		}
		y, err := strconv.Atoi(r.URL.Query().Get("y"))
		if err != nil {
			helper.ReturnJsonError(w, fmt.Errorf("Invalid query params %v", err), http.StatusBadRequest)
			return
		}

		c.Y = y
		c.X = x
	} else {
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			helper.ReturnJsonError(w, fmt.Errorf("Invalid json body %v", err), http.StatusBadRequest)
			return
		}
	}
	if c.Hex != "" {
		helper.ReturnJsonError(w, fmt.Errorf("hex input cannot be provided on get request, hex provided: %s", c.Hex), http.StatusBadRequest)
		return
	}
	maxX, maxY := helper.GetMaxXY()
	if c.X > maxX || c.Y > maxY || c.Y <= 0 || c.X <= 0 {
		helper.ReturnJsonError(w, fmt.Errorf("Invalid X-Y co-ords. Max X= %d Max y= %d", maxX, maxY), http.StatusBadRequest)
		return
	}
	hex, err := vk.GetSingleCanvasValue(r.Context(), c)
	if err != nil {
		// If there's no value set, we set it to white
		if valkey.IsValkeyNil(err) {
			hex = "#FFFFFF"
		} else {
			helper.ReturnJsonError(w, fmt.Errorf("error getting hex from DB %w", err), http.StatusInternalServerError)
			return
		}
	}
	response, _ := json.Marshal(map[string]any{"hex": hex})
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	log.Printf("Successfully got co-ords X:%d , Y:%d with hex %s\n", c.X, c.Y, hex)

}

type Set10x10ImageRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func Set10x10Image(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("error parsing multipart form: %w", err), http.StatusBadRequest)
		return
	}

	xStr := req.FormValue("x")
	yStr := req.FormValue("y")

	if xStr == "" || yStr == "" {
		helper.ReturnJsonError(w, fmt.Errorf("x and y coordinates are required"), http.StatusBadRequest)
		return
	}

	x, err := strconv.Atoi(xStr)
	if err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("invalid x coordinate: %w", err), http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("invalid y coordinate: %w", err), http.StatusBadRequest)
		return
	}

	maxX, maxY := helper.GetMaxXY()
	if x < 0 || y < 0 || x+10 > maxX || y+10 > maxY {
		helper.ReturnJsonError(w, fmt.Errorf("10x10 image would extend beyond canvas bounds. Canvas size: %dx%d", maxX, maxY), http.StatusBadRequest)
		return
	}

	file, _, err := req.FormFile("image")
	if err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("error getting image file: %w", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		helper.ReturnJsonError(w, fmt.Errorf("error decoding PNG image: %w", err), http.StatusBadRequest)
		return
	}

	bounds := img.Bounds()
	if bounds.Dx() != 10 || bounds.Dy() != 10 {
		helper.ReturnJsonError(w, fmt.Errorf("image must be exactly 10x10 pixels, got %dx%d", bounds.Dx(), bounds.Dy()), http.StatusBadRequest)
		return
	}

	rgba := image.NewRGBA(bounds)
	for py := bounds.Min.Y; py < bounds.Max.Y; py++ {
		for px := bounds.Min.X; px < bounds.Max.X; px++ {
			rgba.Set(px, py, img.At(px, py))
		}
	}

	pixelsSet := 0
	for py := 0; py < 10; py++ {
		for px := 0; px < 10; px++ {
			canvasX := x + px
			canvasY := y + py

			r, g, b, a := rgba.At(px, py).RGBA()

			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8)

			hexColor := fmt.Sprintf("#%02x%02x%02x", r8, g8, b8)

			if a8 < 128 {
				continue
			}

			pixel := models.Canvas{X: canvasX, Y: canvasY, Hex: hexColor}

			if err := vk.SetCanvasValue(req.Context(), pixel); err != nil {
				log.Printf("Error setting pixel at (%d, %d): %v", canvasX, canvasY, err)
				continue
			}

			cache.SetColor(canvasX, canvasY, hexColor)
			pixelsSet++
		}
	}

	log.Printf("Successfully set %d pixels from 10x10 image at (%d, %d)", pixelsSet, x, y)

	response, _ := json.Marshal(map[string]any{
		"success":    true,
		"pixels_set": pixelsSet,
		"message":    fmt.Sprintf("Set %d pixels from 10x10 image", pixelsSet),
	})
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
