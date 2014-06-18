package main

import (
	"image"
	"testing"
)

func TestAspectRatio(t *testing.T) {
	rect := image.Rect(0, 0, 1, 4)
	ar := aspectRatio(&rect)
	if ar != 0.25 {
		t.Errorf("Ratio of %#v is %f, not %f", rect, ar, 0.25)
	}
}

func TestHeights(t *testing.T) {
	rect := &image.Rectangle{image.Pt(0, 0), image.Pt(0, 10)}

	square := squareHeight(rect)
	space := spaceHeight(rect)

	if square != 1 {
		t.Errorf("Square height of %#v is %d, not %d", rect, square, 1)
	}

	if space != 0 {
		t.Errorf("Space height of %#v is %d, not %d", rect, space, 0)
	}
}

func TestCrop(t *testing.T) {
	rect := image.Rectangle{image.Pt(0, 0), image.Pt(5, 10)}

	cropRatio(&rect)

	ideal_rect := image.Rect(2, 0, 3, 10)

	if rect != ideal_rect {
		t.Errorf("Cropped rect %#v is not equal to %#v", rect, ideal_rect)
	}
}

func TestCropLarge(t *testing.T) {
	rect := image.Rect(1264, 2304, 8736, 6789)
	ideal_rect := image.Rect(4581, 2304, 5419, 6789)

	cropRatio(&rect)

	if rect != ideal_rect {
		t.Errorf("Cropped rect %#v is not equal to %#v", rect, ideal_rect)
	}
}
