package main

import (
	"fmt"

	"github.com/krady21/raytracing-go/internal/utils"
	"github.com/schollz/progressbar/v3"
)

const (
	imageWidth = 256
	imageHeight = 256
)
type Vec3 utils.Vec3

func main() {
	fmt.Println("P3") 
	fmt.Println(imageWidth, imageHeight)
	fmt.Println(255)
	v := Vec3 {0, 0, 0}
	fmt.Print(v)

	bar := progressbar.Default(100)
	for j := imageHeight - 1; j >= 0; j-- {
		bar.Add(1)
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := 0.25

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Println(ir, ig, ib)
		}
	}
}
