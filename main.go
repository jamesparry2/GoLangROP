package main

import (
	"fmt"

	"github.com/jamesparry2/GoLangROP/pkg/results"
)

func main() {
	output := results.NewResult(nil, nil)

	result := output.
		Map(GetData).
		Map(GetNeverHere).
		Map(GetApi)

	fmt.Print(result.Outcome())
}

func GetData(s *results.Success, r *results.Results[results.Success, results.Error]) *results.Results[results.Success, results.Error] {
	return r.Succeeded(results.NewSuccess(2))
}

func GetApi(s *results.Success, r *results.Results[results.Success, results.Error]) *results.Results[results.Success, results.Error] {
	return r.Failed(results.NewError("Error"))
}

func GetNeverHere(s *results.Success, r *results.Results[results.Success, results.Error]) *results.Results[results.Success, results.Error] {
	return r.Succeeded(results.NewSuccess("Test"))
}
