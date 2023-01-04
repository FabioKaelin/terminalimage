# Terminal Image

With this package can you display a image in your terminal.

The height if the generated image is customizable.

The supported image-types are png (.png), jpeg (.jpg, .jpeg, .jfif), singleframe gif (.gif) and webP (.webp).

## Usage

To import the package use `go get github.com/fabiokaelin/terminalimage`.

```go
package main

import (
	"fmt"

	"github.com/fabiokaelin/terminalimage"
)

func main() {
	imageString, err := terminalimage.ImageToString("img.png", 20, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(imageString)
}

```

## Image Example

To show the convertion, here is the original image:

![Icon of the superman soundtrack on spotify](./docs/images/img.png)

The following are the converte ones:

64 pixel with pixelsplit on:

![64 Pixel with pixelsplit on](./docs/images/64-true.png)

64 pixel with pixelsplit off:

![64 Pixel with pixelsplit off](./docs/images/64-false.png)

20 pixel with pixelsplit on:

![20 Pixel with pixelsplit on](./docs/images/20-true.png)

20 pixel with pixelsplit off:

![20 Pixel with pixelsplit off](./docs/images/20-false.png)
