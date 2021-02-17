package imgdiff

import (
	"image"
	"image/color"
	"github.com/n7olkachev/imgdiff/pkg/yiq"
	"runtime"
	"sync"
	"sync/atomic"
)

// Options struct.
type Options struct {
	Threshold float64
	DiffImage bool
}

// Result struct.
type Result struct {
	Equal           bool
	Image           image.Image
	DiffPixelsCount uint64
}

// Diff between two images.
func Diff(image1 image.Image, image2 image.Image, options *Options) *Result {
	diffPixelsCount := uint64(0)

	maxDelta := yiq.MaxDelta * options.Threshold * options.Threshold

	diff := image.NewNRGBA(image1.Bounds())

	wg := sync.WaitGroup{}

	cpus := runtime.NumCPU()

	for i := 0; i < cpus; i++ {
		wg.Add(1)

		go func(i int) {
			diffPixelsCounter := 0

			for y := i; y <= image1.Bounds().Max.Y; y += cpus {
				for x := 0; x <= image1.Bounds().Max.X; x++ {
					pixel1, pixel2 := image1.At(x, y), image2.At(x, y)

					if pixel1 != pixel2 {
						delta := yiq.Delta(pixel1, pixel2)

						if delta > maxDelta {
							diff.SetNRGBA(x, y, color.NRGBA{R: 255, G: 0, B: 0, A: 255})

							diffPixelsCounter++
						}
					} else if options.DiffImage {
						diff.Set(x, y, pixel1)
					}
				}
			}

			if diffPixelsCounter > 0 {
				atomic.AddUint64(&diffPixelsCount, uint64(diffPixelsCounter))
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	return &Result{
		Equal:           diffPixelsCount == 0,
		DiffPixelsCount: diffPixelsCount,
		Image:           diff,
	}
}
