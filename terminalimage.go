package terminalimage

import (
	"image"
	"strings"

	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/gookit/color"
	_ "golang.org/x/image/webp"
)

// func init() {
// 	fmt.Println("Simple interest package initialized")
// }

// func Example() {
// 	str, err := ImageToString("img.jfif", 20, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(str)
// }

func ImgDataToArray(imageData image.Image, height int, pixelSplit bool) ([]string, error) {
	maxX := imageData.Bounds().Max.X
	maxY := imageData.Bounds().Max.Y
	YBlockSize := maxY / height
	width := maxX / YBlockSize
	XBlockSize := maxX / width
	var result []string
	for YBlockNr := 0; YBlockNr < height; YBlockNr++ {
		lineResult := ""
		if pixelSplit {
			for XBlockNr := 0; XBlockNr < width*2; XBlockNr++ {
				sumR := 0
				sumG := 0
				sumB := 0
				for XPixelNr := 0; XPixelNr < XBlockSize/2; XPixelNr++ {
					for YPixelNr := 0; YPixelNr < YBlockSize; YPixelNr++ {
						r, g, b, _ := imageData.At(XPixelNr+int(float64(XBlockNr)/float64(2)*float64(XBlockSize)), YPixelNr+(YBlockNr*(YBlockSize))).RGBA()
						sumR += int(r / 256)
						sumG += int(g / 256)
						sumB += int(b / 256)
					}
				}
				lineResult += color.Sprintf("<fg=%d,%d,%d;bg=%d,%d,%d>█</>", int(sumR/((XBlockSize/2)*(YBlockSize))), int(sumG/((XBlockSize/2)*(YBlockSize))), int(sumB/((XBlockSize/2)*(YBlockSize))), int(sumR/((XBlockSize/2)*(YBlockSize))), int(sumG/((XBlockSize/2)*(YBlockSize))), int(sumB/((XBlockSize/2)*(YBlockSize))))
			}
		} else {
			for XBlockNr := 0; XBlockNr < width; XBlockNr++ {
				sumR := 0
				sumG := 0
				sumB := 0
				for XPixelNr := 0; XPixelNr < XBlockSize; XPixelNr++ {
					for YPixelNr := 0; YPixelNr < YBlockSize; YPixelNr++ {
						r, g, b, _ := imageData.At(XPixelNr+(XBlockNr*XBlockSize), YPixelNr+(YBlockNr*YBlockSize)).RGBA()
						sumR += int(r / 256)
						sumG += int(g / 256)
						sumB += int(b / 256)
					}
				}
				lineResult += color.Sprintf("<fg=%d,%d,%d;bg=%d,%d,%d>██</>", int(sumR/(XBlockSize*YBlockSize)), int(sumG/(XBlockSize*YBlockSize)), int(sumB/(XBlockSize*YBlockSize)), int(sumR/(XBlockSize*YBlockSize)), int(sumG/(XBlockSize*YBlockSize)), int(sumB/(XBlockSize*YBlockSize)))

			}
		}
		result = append(result, lineResult)
	}
	return result, nil
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
