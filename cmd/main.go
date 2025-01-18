package cmd

import "github.com/bobopylabepolhk/sbpstu-algo-1/internal/output"

func main() {
	params, err := output.ReadParams()
	if err != nil {
		panic(err)
	}
	output.Print(params)
}
