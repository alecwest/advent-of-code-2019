package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/alecwest/advent-of-code-2019/advent"
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

func main() {
	input := advent.ReadStringInput()
	image := parseImage(input, 6, 25)
	leastZeros, numDigits := layerWithLeastZeros(image)
	fmt.Printf("%+v\n", leastZeros)
	fmt.Printf("%d\n", numDigits[1]*numDigits[2])
}
