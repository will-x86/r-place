package cache

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/valkey-io/valkey-go"

	"github.com/will-x86/r-place/pkg/helper"
	"github.com/will-x86/r-place/pkg/models"
	"github.com/will-x86/r-place/pkg/vk"
)

var (
	can           *PixelCanvas
	cachedPng     []byte
	pngCacheMutex sync.RWMutex
	pngUpdateCh   chan struct{}
	dirtyFlag     bool
	dirtyMutex    sync.RWMutex
)

type PixelCanvas struct {
	img   *image.RGBA
	mutex sync.Mutex
}

func (pc *PixelCanvas) updatePngCache() {
	startTime := time.Now()
	buf := new(bytes.Buffer)

	// Use best compression for smaller file size
	encoder := &png.Encoder{CompressionLevel: png.BestCompression}
	if err := encoder.Encode(buf, pc.img); err != nil {
		panic(fmt.Sprintf("failed to encode png: %v", err))
	}

	pngCacheMutex.Lock()
	cachedPng = buf.Bytes()
	pngCacheMutex.Unlock()
	log.Println("Time to re-encode png:", time.Since(startTime).String())
}

func markDirty() {
	dirtyMutex.Lock()
	if !dirtyFlag {
		dirtyFlag = true
		select {
		case pngUpdateCh <- struct{}{}:
		default:
		}
	}
	dirtyMutex.Unlock()
}

func clearDirty() {
	dirtyMutex.Lock()
	dirtyFlag = false
	dirtyMutex.Unlock()
}

func GetCachedPng() []byte {
	pngCacheMutex.RLock()
	defer pngCacheMutex.RUnlock()
	// Return direct reference since it's read-only after creation
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
	pc.mutex.Lock()
	pc.img.Set(x, y, color)
	pc.mutex.Unlock()
	markDirty()
}

func (pc *PixelCanvas) SetPixel(x, y int, r, g, b, a uint8) {
	pc.mutex.Lock()
	c := color.RGBA{r, g, b, a}
	pc.img.Set(x, y, c)
	pc.mutex.Unlock()
	markDirty()
}

func InitialCache() error {
	max_x, max_y := helper.GetMaxXY()
	can = NewPixelCanvas(max_x, max_y)

	type job struct {
		x, y int
	}

	numJobs := max_x * max_y
	jobs := make(chan job, numJobs)
	errs := make(chan error, 1)
	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()
	numWorkers = min(numJobs, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				select {
				case <-errs:
					return
				default:
				}

				hex, err := vk.GetSingleCanvasValue(context.Background(), models.Canvas{
					X: j.x,
					Y: j.y,
				})
				if err != nil {
					if !valkey.IsValkeyNil(err) {
						select {
						case errs <- err:
						default:
						}
						return
					}
					hex = "#FFFFFF"
				}

				col, err := ParseHexColor(hex)
				if err != nil {
					log.Printf("Invalid hex color '%s' at position (%d, %d), using white: %v", hex, j.x, j.y, err)
					col = color.RGBA{255, 255, 255, 255}
				}

				can.mutex.Lock()
				can.img.Set(j.x, j.y, col)
				can.mutex.Unlock()
			}
		}()
	}

	for x := range max_x {
		for y := range max_y {
			jobs <- job{x: x, y: y}
		}
	}
	close(jobs)

	wg.Wait()

	select {
	case err := <-errs:
		return err
	default:
	}

	can.mutex.Lock()
	defer can.mutex.Unlock()
	can.updatePngCache()
	return nil
}

func StartPngUpdateWorker() {
	startPngUpdateWorker()
}

func startPngUpdateWorker() {
	pngUpdateCh = make(chan struct{}, 1)
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				dirtyMutex.RLock()
				if dirtyFlag {
					dirtyMutex.RUnlock()
					can.mutex.Lock()
					can.updatePngCache()
					can.mutex.Unlock()
					clearDirty()
				} else {
					dirtyMutex.RUnlock()
				}
			case <-pngUpdateCh:
			}
		}
	}()
}

func NewPixelCanvas(width, height int) *PixelCanvas {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	return &PixelCanvas{
		img: img,
	}
}
