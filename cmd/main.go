package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"github.com/n7olkachev/imgdiff/pkg/imgdiff"
	"log"
	"os"
	"sync"

	"github.com/alexflint/go-arg"
	. "github.com/logrusorgru/aurora"
)

func loadImages(filePathes ...string) []image.Image {
	images := make([]image.Image, len(filePathes))

	wg := sync.WaitGroup{}

	for i, path := range filePathes {
		wg.Add(1)

		go func(i int, path string) {
			file, err := os.Open(path)

			defer file.Close()

			if err != nil {
				log.Fatalf("can't open image %s %s", path, err.Error())
			}

			image, _, err := image.Decode(file)

			if err != nil {
				log.Fatalf("can't decode image %s %s", path, err.Error())
			}

			images[i] = image

			wg.Done()
		}(i, path)
	}

	wg.Wait()

	return images
}

func main() {
	var args struct {
		Threshold    float64 `arg:"-t,--threshold" help:"Color difference threshold (from 0 to 1). Less more precise." default:"0.1"`
		DiffImage    bool    `arg:"--diff-image" help:"Render image to the diff output instead of transparent background." default:"false"`
		FailOnLayout bool    `arg:"--fail-on-layout" help:"Do not compare images and produce output if images layout is different." default:"false"`
		Base         string  `arg:"positional" help:"Base image."`
		Compare      string  `arg:"positional" help:"Image to compare with."`
		Output       string  `arg:"positional" help:"Output image path."`
	}

	arg.MustParse(&args)

	images := loadImages(args.Base, args.Compare)

	image1, image2 := images[0], images[1]

	if args.FailOnLayout && !image1.Bounds().Eq(image2.Bounds()) {
		fmt.Println(Red("Failure!").Bold(), "Images have different layout.")

		os.Exit(2)
	}

	result := imgdiff.Diff(image1, image2, &imgdiff.Options{
		Threshold: args.Threshold,
		DiffImage: args.DiffImage,
	})

	if result.Equal {
		fmt.Println(Green("Success!").Bold(), "Images are equal.")
		return
	}

	enc := &png.Encoder{
		CompressionLevel: png.BestSpeed,
	}

	f, _ := os.Create(args.Output)
	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	enc.Encode(writer, result.Image)

	fmt.Println(Red("Failure!").Bold(), "Images are different.")

	fmt.Printf("Different pixels: %d\n", Red(result.DiffPixelsCount).Bold())

	os.Exit(1)
}
