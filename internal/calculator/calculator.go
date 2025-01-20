package calculator

import (
	"fmt"
	"math"
)

const num = 14

func calcEven(x float64) float64 {
	p1 := math.Log((1 - num) / math.Sin(x+num))
	p2 := math.Abs(math.Cos(math.Log(x)) / num)
	return math.Max(p1, p2)
}

func Run(x float64) (r float64, err error) {
	r = calcEven(x)

	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, fmt.Errorf("ERROR")
	}

	return
}
