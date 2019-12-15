package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/alecwest/advent-of-code-2019/advent"
)

const (
	// BLACK color
	BLACK = 0

	// WHITE color
	WHITE = 1

	// TRANSPARENT (no) color
	TRANSPARENT = 2
)

// Image is the full image
type Image struct {
	layers []Layer
	length int
	width  int
}

// Layer is one layer of an image
type Layer struct {
	pixels []int
}

func parseImage(input string, length, width int) Image {
	numPixels := length * width
	numLayers := len(input) / numPixels
	layers := make([]Layer, numLayers)

	for i := 0; i < numLayers; i++ {
		firstPixel := numPixels * i
		layer := input[firstPixel : firstPixel+numPixels]
		layers[i].pixels = make([]int, numPixels)
		for j := 0; j < numPixels; j++ {
			val, _ := strconv.Atoi(string(layer[j]))
			layers[i].pixels[j] = val
		}
	}

	return Image{layers, length, width}
}

func printImage(image Image) {
	for i := 0; i < image.length; i++ {
		for j := 0; j < image.width; j++ {
			pixel := topPixel(image, i, j)
			if pixel == 1 {
				fmt.Printf("%d ", pixel)
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}

func topPixel(image Image, row, col int) int {
	for _, layer := range image.layers {
		currPixel := layer.pixels[row*image.width+col]
		if currPixel != TRANSPARENT {
			return currPixel
		}
	}
	return image.layers[len(image.layers)].pixels[row*image.width+col]
}

func main() {
	input := advent.ReadStringInput()
	image := parseImage(input, 6, 25)
	printImage(image)
}

func pixelCount(layer Layer) map[int]int {
	count := make(map[int]int)
	for _, pixel := range layer.pixels {
		count[pixel]++
	}
	return count
}

func layerWithLeastZeros(image Image) (Layer, map[int]int) {
	var bestLayer Layer
	var bestMap map[int]int
	numZeros := math.MaxInt64
	for _, layer := range image.layers {
		m := pixelCount(layer)
		if m[0] < numZeros {
			numZeros = m[0]
			bestLayer = layer
			bestMap = m
		}
	}
	return bestLayer, bestMap
}
