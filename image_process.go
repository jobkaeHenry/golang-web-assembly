package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"syscall/js"
)

// 흑백으로 변환하는 함수
func grayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayColor := color.GrayModel.Convert(img.At(x, y))
			grayImg.Set(x, y, grayColor)
		}
	}
	return grayImg
}

// 이미지 처리 함수
func processImage(this js.Value, p []js.Value) interface{} {
	inputBytes := make([]byte, p[0].Get("byteLength").Int())
	js.CopyBytesToGo(inputBytes, p[0])

	img, _, err := image.Decode(bytes.NewReader(inputBytes))
	if err != nil {
		return js.ValueOf(map[string]interface{}{
			"error": err.Error(),
		})
	}

	grayImg := grayscale(img)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, grayImg, nil) // JPEG로 인코딩
	if err != nil {
		return js.ValueOf(map[string]interface{}{
			"error": err.Error(),
		})
	}

	// 인코딩된 이미지 바이트 슬라이스를 JavaScript로 복사
	outputBytes := buf.Bytes()
	output := js.Global().Get("Uint8Array").New(len(outputBytes))
	js.CopyBytesToJS(output, outputBytes)

	return js.ValueOf(map[string]interface{}{
		"image": output,
	})
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("processImage", js.FuncOf(processImage))
	<-c
}
