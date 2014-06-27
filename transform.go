package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"runtime"
	"io"
	"image/png"
	"bytes"
)

const (
	RATIO = 0.18672199170124
)

func makeImages(rect *image.Rectangle, file io.Reader) ([]*bytes.Buffer, error) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ar := aspectRatio(rect)

	if ar < RATIO {
		return nil, fmt.Errorf("makeImages: aspect ratio is too small")
	} else if ar > RATIO {
		cropRatio(rect)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	img = imaging.Clone(img)

	top := rect.Min.Y
	square := squareHeight(rect)
	space := spaceHeight(rect)

	buffers := make([]*bytes.Buffer, 5, 5)

	for i := 0; i < 5; i++ {
		new_rect := image.Rect(rect.Min.X, top, rect.Max.X, top+square)
		new_img := imaging.Crop(img, new_rect)

		err = png.Encode(buffers[i], new_img)
		if err != nil {
			panic(err)
		}

		top = top + square + space
	}

	return buffers, nil
}

func squareHeight(rect *image.Rectangle) int {
	return int(float64(rect.Dy()) * RATIO)
}

func spaceHeight(rect *image.Rectangle) int {
	return int(float64(rect.Dy()) * 0.01659751037344)
}

func aspectRatio(rect *image.Rectangle) float64 {
	return float64(rect.Dx()) / float64(rect.Dy())
}

func cropRatio(rect *image.Rectangle) {
	if aspectRatio(rect) <= RATIO {
		return
	}

	ideal := int(float64(rect.Dy()) * RATIO)

	trim := (rect.Dx() - ideal) / 2

	// Make corrections
	rect.Min.X += trim
	rect.Max.X -= trim
}
