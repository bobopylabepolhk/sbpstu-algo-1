package output

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/bobopylabepolhk/sbpstu-algo-1/internal/calculator"
	"github.com/bobopylabepolhk/sbpstu-algo-1/pkg/strpad"
)

type Params struct {
	start float64
	end   float64
	delta float64
}

func ReadParams() (params Params, err error) {
	flag.Float64Var(&params.start, "start", 0, "start value of range calculation")
	flag.Float64Var(&params.end, "end", 0, "end value of range calculation")
	flag.Float64Var(&params.delta, "deltaX", 0, "increments X")
	flag.Parse()

	if (params.delta > params.end && params.end > 0) || (params.delta < params.end && params.end < 0) {
		err = fmt.Errorf("invalid deltaX. deltaX: %v end: %v", params.delta, params.end)
		return
	}

	if params.delta == 0 {
		err = fmt.Errorf("deltaX can't be zero")
		return
	}

	if params.delta > 0 && params.start >= params.end {
		err = fmt.Errorf("invalid positive range. start %v end %v", params.start, params.end)
		return
	}

	if params.delta < 0 && params.start <= params.end {
		err = fmt.Errorf("invalid negative range. end %v start %v", params.end, params.start)
		return
	}

	return params, nil
}

func printRow(iter int, x float64, pad int) string {
	r, err := calculator.Run(x)
	paddedIter := strpad.Left(fmt.Sprintf("%v", iter), pad-4)
	paddedX := strpad.Right(strpad.TrimFloat(fmt.Sprintf("%.5f", x)), pad)

	if err != nil {
		return fmt.Sprintf("%s | %s | ERROR", paddedIter, paddedX)
	}

	paddedR := strpad.Right(strpad.TrimFloat(fmt.Sprintf("%.5f", r)), pad)

	return fmt.Sprintf("%s | %s | %s", paddedIter, paddedX, paddedR)
}

func shouldRun(params Params, x float64) bool {
	if params.delta > 0 {
		return x <= params.end
	}

	return x >= params.end
}

func Print(params Params) {
	maxFlag := math.Max(params.start, params.end)
	maxColLen := len(fmt.Sprintf("%v", maxFlag)) + 4
	tableHead := fmt.Sprintf(
		"%s | %s | %s ",
		strpad.Left("#", maxColLen-4),
		strpad.Right("x", maxColLen),
		"result",
	)
	fmt.Println(tableHead)
	fmt.Println(strings.Repeat("—", len(tableHead)))

	i := 1

	for x := params.start; shouldRun(params, x); x += params.delta {
		fmt.Println(printRow(i, x, maxColLen))
		i += 1
	}
}
