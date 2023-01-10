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
	YBlockSize := int(float64(maxY) / float64(height))
	width := int(float64(maxX) / (float64(maxY) / float64(height)))
	XBlockSize := int(float64(maxX) / (float64(maxX) / (float64(maxY) / float64(height))))
	var result []string
	Dividor := 1
	if pixelSplit {
		Dividor = 2
	}
	for YBlockNr := 0; YBlockNr < height; YBlockNr++ {
		result = append(result, blockLine(imageData, Dividor, width, XBlockSize, YBlockSize, YBlockNr))
	}
	return result, nil
}

func blockLine(imageData image.Image, Dividor int, width int, XBlockSize int, YBlockSize int, YBlockNr int) string {
	lineResult := ""
	PixelSize := 2
	if Dividor == 2 {
		PixelSize = 1
	}
	for XBlockNr := 0; XBlockNr < width*Dividor; XBlockNr++ {
		sumR := 0
		sumG := 0
		sumB := 0
		for XPixelNr := 0; XPixelNr < XBlockSize/Dividor; XPixelNr++ {
			for YPixelNr := 0; YPixelNr < YBlockSize; YPixelNr++ {
				r, g, b, _ := imageData.At(XPixelNr+int(float64(XBlockNr)/float64(Dividor)*float64(XBlockSize)), int(float64(YPixelNr)+(float64(YBlockNr)*(float64(YBlockSize))))).RGBA()
				sumR += int(r / 256)
				sumG += int(g / 256)
				sumB += int(b / 256)
			}
		}
		lineResult += color.Sprintf("<fg=%d,%d,%d;bg=%d,%d,%d>%s</>", int(float64(sumR)/((float64(XBlockSize)/float64(Dividor))*(float64(YBlockSize)))), int(float64(sumG)/((float64(XBlockSize)/float64(Dividor))*(float64(YBlockSize)))), int(float64(sumB)/((float64(XBlockSize)/float64(Dividor))*(float64(YBlockSize)))), int(float64(sumR)/((float64(XBlockSize)/float64(Dividor))*(float64(YBlockSize)))), int(float64(sumG)/((float64(XBlockSize)/float64(Dividor))*(float64(YBlockSize)))), int(float64(sumB)/((float64(XBlockSize)/float64(Dividor))*(float64(YBlockSize)))), strings.Repeat("█", PixelSize))
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
