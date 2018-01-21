package main

import (
	"log"
	"image"
	_ "image/jpeg"
	"os"
) 

func main() {
	decode(os.Args[1])
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