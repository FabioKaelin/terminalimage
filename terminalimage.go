package terminalimage

import (
	_ "embed"

	"bytes"
	"fmt"
	"image"
	"strings"

	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/gookit/color"
	_ "golang.org/x/image/webp"
)

//go:embed docs/images/img.png
var img []byte

func Example() {
	imageData, _, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		panic(err)
	}
	resultArray, err := ImgDataToArray(imageData, 61, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Join(resultArray[:], "\n"))
}

func ImgDataToArray(imageData image.Image, height int, pixelSplit bool) ([]string, error) {
	maxX := imageData.Bounds().Max.X
	maxY := imageData.Bounds().Max.Y
	YBlockSize := maxY / height
	width := maxX / YBlockSize
	XBlockSize := maxX / width
	var result []string
	PixelSize := 1
	if pixelSplit {
		PixelSize = 2
	}
	for YBlockNr := 0; YBlockNr < height; YBlockNr++ {
		result = append(result, blockLine(imageData, PixelSize, width, XBlockSize, YBlockSize, YBlockNr))
	}
	return result, nil
}

func blockLine(imageData image.Image, PixelSize int, width int, XBlockSize int, YBlockSize int, YBlockNr int) string {
	lineResult := ""
	for XBlockNr := 0; XBlockNr < width*PixelSize; XBlockNr++ {
		sumR := 0
		sumG := 0
		sumB := 0
		for XPixelNr := 0; XPixelNr < XBlockSize/PixelSize; XPixelNr++ {
			for YPixelNr := 0; YPixelNr < YBlockSize; YPixelNr++ {
				r, g, b, _ := imageData.At(XPixelNr+int(float64(XBlockNr)/float64(PixelSize)*float64(XBlockSize)), int(float64(YPixelNr)+(float64(YBlockNr)*(float64(YBlockSize))))).RGBA()
				sumR += int(r / 256)
				sumG += int(g / 256)
				sumB += int(b / 256)
			}
		}
		lineResult += color.Sprintf("<fg=%d,%d,%d;bg=%d,%d,%d>█</>", int(float64(sumR)/((float64(XBlockSize)/float64(PixelSize))*(float64(YBlockSize)))), int(float64(sumG)/((float64(XBlockSize)/float64(PixelSize))*(float64(YBlockSize)))), int(float64(sumB)/((float64(XBlockSize)/float64(PixelSize))*(float64(YBlockSize)))), int(float64(sumR)/((float64(XBlockSize)/float64(PixelSize))*(float64(YBlockSize)))), int(float64(sumG)/((float64(XBlockSize)/float64(PixelSize))*(float64(YBlockSize)))), int(float64(sumB)/((float64(XBlockSize)/float64(PixelSize))*(float64(YBlockSize)))))
	}
	return lineResult
}

func ImageToString(path string, height int, pixelSplit bool) (string, error) {

	existingImageFile, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer existingImageFile.Close()

	existingImageFile.Seek(0, 0)

	imageData, _, err := image.Decode(existingImageFile)
	if err != nil {
		return "", err
	}
	resultArray, err := ImgDataToArray(imageData, height, pixelSplit)
	if err != nil {
		return "", err
	}
	return strings.Join(resultArray[:], "\n"), nil
}
