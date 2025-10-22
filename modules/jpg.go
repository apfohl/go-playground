package modules

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"

	"github.com/spf13/cobra"
)

func NewJpgCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "jpg",
		Short: "Generate a base64-encoded JPEG image",
		Run:   runJpg,
	}
}

func runJpg(_ *cobra.Command, _ []string) {
	width, height := 100, 100
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	white := color.RGBA{255, 255, 255, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, white)
		}
	}

	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		fmt.Printf("Error encoding image: %v\n", err)
		return
	}

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println(base64Str)
}
