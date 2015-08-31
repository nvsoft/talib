package main

import (
	"github.com/nvsoft/talib"
	"fmt"
	"math"
)

func main() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
	// => [0 1]
}
