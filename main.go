package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/draw"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("VR Merge Image Combiner")
		fmt.Println("Usage: go run main.go <image1> <image2> <output>")
		fmt.Println("Output does not need an extension!")
		fmt.Println("Example: go run main.go scene_left.jpg scene_right.jpg scenery")
		fmt.Println("This will output scenery_3DHF.jpg or scenery_3DVF.jpg as necessary.")
		return
	}
	img1Filename := os.Args[1]
	img2Filename := os.Args[2]
	outputFilename := os.Args[3]

	img1File, err := os.Open(img1Filename)
	if err != nil {
		fmt.Println("Error opening first image:", err)
		return
	}
	defer img1File.Close()

	img2File, err := os.Open(img2Filename)
	if err != nil {
		fmt.Println("Error opening second image:", err)
		return
	}
	defer img2File.Close()

	img1, format1, err := image.Decode(img1File)
	if err != nil {
		fmt.Println("Error decoding first image:", err)
		return
	}

	img2, format2, err := image.Decode(img2File)
	if err != nil {
		fmt.Println("Error decoding second image:", err)
		return
	}

	if format1 != format2 {
		fmt.Println("Error: images must be in the same format")
		return
	}

	fmt.Println("Loaded Image 1: (", img1.Bounds().Dx(), "x", img1.Bounds().Dy(), ")")
	fmt.Println("Loaded Image 2: (", img2.Bounds().Dx(), "x", img2.Bounds().Dy(), ")")

	// Determine the orientation to merge (side by side, or vertically stacked)
	// Most 3d image viewing software accepts "_3DVF" and "_3DHF" to automatically determine how to view the image.
	// Setting it here as part of 'which way do we merge?' and using it for the mode switch.
	mode := "3DVF"
	if img1.Bounds().Dy() > img1.Bounds().Dx() {
		mode = "3DHF"
	}

	outputFilename = appendModeToFilename(outputFilename, mode, format1)

	var finalImg *image.RGBA
	if mode == "3DHF" {
		fmt.Println("Combining Horizontally")
		width := img1.Bounds().Dx() + img2.Bounds().Dx()
		height := img1.Bounds().Dy()
		if img2.Bounds().Dy() > height {
			height = img2.Bounds().Dy()
		}
		finalImg = image.NewRGBA(image.Rect(0, 0, width, height))
		fmt.Println("Final Image Dimensions: (", width, "x", height, ")")
		draw.Draw(finalImg, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
		draw.Draw(finalImg, img2.Bounds().Add(image.Point{img1.Bounds().Dx(), 0}), img2, image.Point{0, 0}, draw.Src)
	} else if mode == "3DVF" {
		fmt.Println("Combining Vertically")
		width := img1.Bounds().Dx()
		if img2.Bounds().Dx() > width {
			width = img2.Bounds().Dx()
		}
		height := img1.Bounds().Dy() + img2.Bounds().Dy()
		finalImg = image.NewRGBA(image.Rect(0, 0, width, height))
		fmt.Println("Final Image Dimensions: (", width, "x", height, ")")
		draw.Draw(finalImg, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
		draw.Draw(finalImg, img2.Bounds().Add(image.Point{0, img1.Bounds().Dy()}), img2, image.Point{0, 0}, draw.Src)
	} else {
		fmt.Println("Error determining combination mode.")
		return
	}
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Error creating output image:", err)
		return
	}
	defer outputFile.Close()
	switch strings.ToLower(format1) {
	case "jpeg", "jpg":
		err = jpeg.Encode(outputFile, finalImg, nil)
	case "png":
		err = png.Encode(outputFile, finalImg)
	default:
		fmt.Println("Unsupported image format:", format1)
		return
	}
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}
	fmt.Println("Created", outputFilename)
}

func appendModeToFilename(filename, mode string, format1 string) string {
	//ext := filepath.Ext(filename)
	ext := format1
	name := strings.TrimSuffix(filename, ext)
	return fmt.Sprintf("%s_%s.%s", name, mode, ext)
}
