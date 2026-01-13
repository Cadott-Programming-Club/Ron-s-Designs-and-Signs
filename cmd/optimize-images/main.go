package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func main() {
	inputDir := "images/work"
	outputDir := "images/work"

	files, err := filepath.Glob(filepath.Join(inputDir, "*.png"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d PNG images to optimize\n", len(files))

	for _, file := range files {
		fmt.Printf("Processing %s...\n", filepath.Base(file))

		img, err := loadImage(file)
		if err != nil {
			log.Printf("Error loading %s: %v\n", file, err)
			continue
		}

		baseName := strings.TrimSuffix(filepath.Base(file), ".png")

		// Original dimensions: 1280x1280 (square)
		originalWidth := uint(img.Bounds().Dx())
		originalHeight := uint(img.Bounds().Dy())

		// Generate responsive sizes (WebP)
		sizes := []struct {
			width  uint
			suffix string
		}{
			{400, "-400w"},
			{800, "-800w"},
			{1280, ""},
		}

		for _, size := range sizes {
			resized := resize.Resize(size.width, 0, img, resize.Lanczos3)
			outputPath := filepath.Join(outputDir, baseName+size.suffix+".webp")
			if err := saveWebP(resized, outputPath, 80); err != nil {
				log.Printf("Error saving %s: %v\n", outputPath, err)
			} else {
				fmt.Printf("  Created %s\n", filepath.Base(outputPath))
			}
		}

		// Generate 9:16 portrait crop for mobile (720x1280)
		if originalWidth >= 720 && originalHeight >= 1280 {
			// Center crop to 9:16 ratio
			targetWidth := 720
			targetHeight := 1280

			// Calculate crop bounds (center crop)
			cropX := (int(originalWidth) - targetWidth) / 2
			cropY := (int(originalHeight) - targetHeight) / 2

			if cropX < 0 {
				cropX = 0
			}
			if cropY < 0 {
				cropY = 0
			}

			croppedImg := cropImage(img, cropX, cropY, targetWidth, targetHeight)
			mobilePath := filepath.Join(outputDir, baseName+"-mobile.webp")
			if err := saveWebP(croppedImg, mobilePath, 80); err != nil {
				log.Printf("Error saving %s: %v\n", mobilePath, err)
			} else {
				fmt.Printf("  Created %s (9:16 mobile)\n", filepath.Base(mobilePath))
			}
		}

		// Generate blur placeholder (20x20)
		blurredImg := resize.Resize(20, 0, img, resize.Bilinear)
		blurPath := filepath.Join(outputDir, baseName+"-blur.webp")
		if err := saveWebP(blurredImg, blurPath, 50); err != nil {
			log.Printf("Error saving %s: %v\n", blurPath, err)
		} else {
			fmt.Printf("  Created %s (blur placeholder)\n", filepath.Base(blurPath))
		}
	}

	fmt.Println("\n✅ Image optimization complete!")
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveWebP(img image.Image, path string, quality float32) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	options := &webp.Options{
		Lossless: false,
		Quality:  quality,
	}

	return webp.Encode(file, img, options)
}

func cropImage(img image.Image, x, y, width, height int) image.Image {
	bounds := img.Bounds()

	// Ensure crop bounds are within image
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	if x+width > bounds.Dx() {
		width = bounds.Dx() - x
	}
	if y+height > bounds.Dy() {
		height = bounds.Dy() - y
	}

	cropRect := image.Rect(x, y, x+width, y+height)
	croppedImg := image.NewRGBA(cropRect)

	for py := y; py < y+height; py++ {
		for px := x; px < x+width; px++ {
			croppedImg.Set(px, py, img.At(px, py))
		}
	}

	return croppedImg
}
