package pixelcombiner

import (
	"image"
	"image/png"
	"image/color"
	"os"
)

func calculateNumberOfPixelsFromImage(img image.Image) int {
	return img.Bounds().Max.X * img.Bounds().Max.Y
}

func generatePixelData(imageData image.Image) []color.Color {
	size := calculateNumberOfPixelsFromImage(imageData)
	pixeldata := make([]color.Color, size)
	var iter int = 0

	for y := 0; y < imageData.Bounds().Max.Y; y++ {
		for x := 0; x < imageData.Bounds().Max.X; x++ {
			color := imageData.At(x, y)
			// fmt.Printf("iter %d color %d", iter, color)
			pixeldata[iter] = color
			iter++
		}
	}
	return pixeldata
}

func Combine(imageData1 image.Image, imageData2 image.Image) image.Image {
	
	// pixelData1 := generatePixelData(imageData1)
	pixelData2 := generatePixelData(imageData2)

	emptyImage := image.NewRGBA(image.Rect(0, 0, imageData2.Bounds().Max.X , imageData2.Bounds().Max.Y))

	// for i:=0; i < len(pixelData1); i++ {
	//	fmt.Printf("color: %d\n", pixelData1[i])
	// }

	var increment = 0
	for i:=0; i < emptyImage.Bounds().Max.Y; i++ {
		for j:=0; j < emptyImage.Bounds().Max.X; j++ {
			// r1, g1, b1, a1 := pixelData1[i].RGBA()
			r2, g2, b2, a2 := pixelData2[increment].RGBA()
			
			emptyImage.Pix[increment] = uint8(r2)
			emptyImage.Pix[increment + 1] = uint8(g2)
			emptyImage.Pix[increment + 2] = uint8(b2)
			emptyImage.Pix[increment + 3] = uint8(a2)
			increment++
		}
	}

	//for i:=length; i < len(pixelData2); i++ {
	//	r2, g2, b2, a2 := pixelData2[i].RGBA()
//
//		emptyImage.Pix[i] = uint8(r2)
//		emptyImage.Pix[i + 1] = uint8(g2)
//		emptyImage.Pix[i + 2] = uint8(b2)
//		emptyImage.Pix[i + 3] = uint8(a2)
//	}

	outputFile, err := os.Create("test.png")
	png.Encode(outputFile, emptyImage)
	if err != nil {
    	// Handle error
    }

    outputFile.Close()

	return emptyImage
}

func writeImage() {
	// Create a blank image 100x200 pixels
	myImage := image.NewRGBA(image.Rect(0, 0, 100, 200))

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("test.png")
	if err != nil {
		// Handle error
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, myImage)

	outputFile.Close()
}

func LoadImage(path string) image.Image {
	existingImageFile, err := os.Open(path)
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	existingImageFile.Seek(0, 0)

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}
	return loadedImage
}