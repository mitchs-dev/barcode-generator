/*
barcode-generator is a simple CLI barcode generator that generates QR codes and saves them to a specified directory.
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var (
	barcodeWidth        = 96                   // Pixels
	barcodeHeight       = 96                   // Pixels
	barcodeFileTemplate = "barcode-$INPUT.png" // The template for the barcode file name (Ex: <outputDir>/barcode-<input>.png)
)

func main() {

	// Check if the input and output directory are provided
	// Both input and output directory are required
	if len(os.Args) < 3 {
		fmt.Println("Usage: barcode <input> <outputDir>")
		fmt.Println("Output: <outputDir>/barcode-<input>.png")
		os.Exit(1)
	}
	inputArg := strings.TrimSpace(os.Args[1])
	outputDir := strings.TrimSpace(os.Args[2])

	// Split the input by comma
	inputToList := strings.Split(inputArg, ",")

	// Loop through the input and generate the barcode
	for _, input := range inputToList {

		// Exit if the input is empty
		if input == "" {
			break
		}

		// Replace the $INPUT with the actual input
		imageFileName := strings.ReplaceAll(barcodeFileTemplate, "$INPUT", input)

		// Generate the barcode
		barCode, _ := qr.Encode(input, qr.M, qr.Auto)

		// Allows you to specify the width and height of the barcode using the vars set above
		barCode, _ = barcode.Scale(barCode, barcodeWidth, barcodeHeight)

		// Create a new image and draw the barcode on it
		img := image.NewRGBA(image.Rect(0, 0, barCode.Bounds().Dx(), barCode.Bounds().Dy()+20))
		draw.Draw(img, img.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)
		draw.Draw(img, barCode.Bounds(), barCode, barCode.Bounds().Min, draw.Src)

		// Add the label to the barcode (This can be removed if you don't want to add a label)
		addLabel(img, 10, barCode.Bounds().Dy()+15, input)

		// Create the file for the barcode
		os.MkdirAll(outputDir, os.ModePerm)
		file, _ := os.Create(outputDir + "/" + imageFileName)
		defer file.Close()

		// Save the barcode to the file
		png.Encode(file, img)
		fmt.Println("Barcode generated for: " + input + " (" + outputDir + "/" + imageFileName + ")")
	}

}

// addLabel adds a label to the barcode
func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{0, 0, 0, 255} // black color
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
