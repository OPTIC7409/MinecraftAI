package analyze

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// HealthInfo holds the health-related information extracted from the image.
type HealthInfo struct {
	EmptyHeartPixels int     `json:"empty_heart_pixels"`
	HeartPixels      int     `json:"heart_pixels"`
	TotalHearts      float64 `json:"total_hearts"`
	HealthPoints     int     `json:"health_points"`
}

// ExtractHealth analyzes the health display in the given image and returns the health points.
func ExtractHealth(img image.Image) int {
	// Define the updated region of interest (ROI) after moving it up by 64 pixels
	healthRegion := image.Rect(1058, 1515-64, 1299, 1364-64)
	healthImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(healthRegion)

	// Draw a rectangle around the ROI on the original image for visualization
	imgWithRect := image.NewRGBA(img.Bounds())
	draw.Draw(imgWithRect, img.Bounds(), img, image.Point{}, draw.Over)

	// Define the rectangle color and thickness
	rectColor := color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red color
	thickness := 2

	// Draw the rectangle around the ROI
	for y := healthRegion.Min.Y; y < healthRegion.Max.Y; y++ {
		for x := healthRegion.Min.X; x < healthRegion.Max.X; x++ {
			if x < healthRegion.Min.X+thickness || x >= healthRegion.Max.X-thickness || y < healthRegion.Min.Y+thickness || y >= healthRegion.Max.Y-thickness {
				imgWithRect.Set(x, y, rectColor)
			}
		}
	}

	// Save the image with the rectangle for visualization
	file, err := os.Create("debug_image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, imgWithRect)
	if err != nil {
		panic(err)
	}

	// Process the extracted image to count colors
	bounds := healthImg.Bounds()
	colorCount := make(map[color.Color]int)
	totalPixels := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := healthImg.At(x, y)
			colorCount[c]++
			totalPixels++
		}
	}

	// Define known heart colors (based on default Minecraft textures)
	heartColor := color.RGBA{R: 255, G: 19, B: 19, A: 255}     // Example color for a heart
	emptyHeartColor := color.RGBA{R: 40, G: 40, B: 40, A: 255} // Example color for an empty heart

	// Count heart pixels
	heartPixels := colorCount[heartColor]
	emptyHeartPixels := colorCount[emptyHeartColor]

	// Print the number of empty heart pixels
	fmt.Println("Empty Heart Pixels:", emptyHeartPixels)
	fmt.Println("Heart Pixels:", heartPixels)

	// Calculate the total number of hearts based on pixel counts
	// Known that 104 pixels represent one heart
	pixelsPerHeart := 104
	totalHearts := float64(heartPixels) / float64(pixelsPerHeart)

	// Convert the number of hearts to health points
	healthPoints := int(totalHearts * 2) // Each heart represents 2 health points

	// Create a HealthInfo struct to hold the results
	healthInfo := HealthInfo{
		EmptyHeartPixels: emptyHeartPixels,
		HeartPixels:      heartPixels,
		TotalHearts:      totalHearts,
		HealthPoints:     healthPoints,
	}

	// Convert HealthInfo to JSON
	jsonData, err := json.MarshalIndent(healthInfo, "", "  ")
	if err != nil {
		panic(err)
	}

	// Print the JSON output
	fmt.Println(string(jsonData))

	return healthPoints
}
