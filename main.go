package main

import (
	"log"
	"image"
	"image/jpeg"
	"image/color"
	"os"
) 

func main() {
	img1, err := decode(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	img2, err := decode(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	if img1.Bounds() != img2.Bounds() {
		log.Fatal()
	}

	out := blend(img1, img2)

	writer, err := os.Create(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	jpeg.Encode(writer, out, &jpeg.Options{100})
}

func blend(img1, img2 image.Image) image.Image {
	width := img1.Bounds().Max.X
	height := img1.Bounds().Max.Y

	out := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color1 := img1.At(x, y)
			color2 := img2.At(x, y)

			if luminance(color1) >= luminance(color2) {
				out.Set(x, y, color1)
			} else {
				out.Set(x, y, color2)
			}
		}
	}

	return out
}

func luminance(color color.Color) uint32 {
	r, g, b, _ := color.RGBA()

	return 298 * r + 586 * g + 114 * b
}

func decode(file string) (image.Image, error) {
	reader, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	return img, nil
}