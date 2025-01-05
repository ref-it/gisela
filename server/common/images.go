package common

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/anthonynsimon/bild/transform"
)

func ResizeImage(image string, dimensions []int) string {
	coI := strings.Index(image, ",")
	rawImage := image[coI+1:]
	jpgBytes, _ := base64.StdEncoding.DecodeString(string(rawImage))

	img, err := png.Decode(bytes.NewReader(jpgBytes))
	if err != nil {
		panic(err)
	}
	result := transform.Resize(img, dimensions[0], dimensions[1], transform.Linear)

	var imageBytesBuffer bytes.Buffer
	jpeg.Encode(&imageBytesBuffer, result, nil)
	data := imageBytesBuffer.Bytes()

	return base64.StdEncoding.EncodeToString(data)
}
