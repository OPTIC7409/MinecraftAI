package capture

import (
	"github.com/kbinani/screenshot"
	"image"
)

func CaptureScreen(displayID int) (image.Image, error) {
	img, err := screenshot.CaptureDisplay(displayID)
	if err != nil {
		return nil, err
	}
	return img, nil
}
