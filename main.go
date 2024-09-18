package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	return img, err
}

func main() {
	// 假设这里有四个瓦片图像
	tile1, err := loadImage("docs/png/123.png")
	if err != nil {
		log.Fatal(err)
	}
	tile2, err := loadImage("docs/png/223.png")
	if err != nil {
		log.Fatal(err)
	}
	tile3, err := loadImage("docs/png/133.png")
	if err != nil {
		log.Fatal(err)
	}
	tile4, err := loadImage("docs/png/233.png")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个新的图像，大小为四个瓦片拼接后的大小
	totalWidth := tile1.Bounds().Dx() * 2
	totalHeight := tile1.Bounds().Dy() * 2
	img := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))

	// 绘制瓦片图像
	draw.Draw(img, image.Rect(0, 0, tile1.Bounds().Dx(), tile1.Bounds().Dy()), tile1, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(tile1.Bounds().Dx(), 0, tile1.Bounds().Dx()*2, tile1.Bounds().Dy()), tile2, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(0, tile1.Bounds().Dy(), tile1.Bounds().Dx(), tile1.Bounds().Dy()*2), tile3, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(tile1.Bounds().Dx(), tile1.Bounds().Dy(), tile1.Bounds().Dx()*2, tile1.Bounds().Dy()*2), tile4, image.Point{}, draw.Src)

	// 添加标记（这里只是简单的示例，实际中需要根据具体标记位置和样式进行绘制）
	markColor := image.Black
	markSize := 5
	for i := -markSize; i < markSize; i++ {
		for j := -markSize; j < markSize; j++ {
			img.Set(tile1.Bounds().Dx()+10+i, tile1.Bounds().Dy()+10+j, markColor)
		}
	}

	// 保存拼接后的图像
	outFile, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatal(err)
	}
}
