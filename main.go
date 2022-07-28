package main

import (
	"fmt"

	"github.com/jamesparry2/GoLangROP/pkg/results"
)

/*
* Example Implementation
 */

func main() {
	success, err := NewResult(&SuccessModel[int]{}, &ErrorModel[string]{}).
		Map(GetData).
		Map(GetApi).
		Map(GetNeverHere).
		Outcome()

	fmt.Print(success)
	fmt.Print(err)
}

func GetData(s *SuccessModel[int], f *ErrorModel[string], r *results.Results[SuccessModel[int], ErrorModel[string]]) *results.Results[SuccessModel[int], ErrorModel[string]] {
	s.SetValue(1)

	return r.Succeeded(s)
}

func GetApi(s *SuccessModel[int], f *ErrorModel[string], r *results.Results[SuccessModel[int], ErrorModel[string]]) *results.Results[SuccessModel[int], ErrorModel[string]] {
	f.SetValue("invalid memory error")

	return r.Failed(f)
}

func GetNeverHere(s *SuccessModel[int], f *ErrorModel[string], r *results.Results[SuccessModel[int], ErrorModel[string]]) *results.Results[SuccessModel[int], ErrorModel[string]] {
	s.SetValue(2)

	return r.Succeeded(s)
}

func NewResult(s *SuccessModel[int], e *ErrorModel[string]) *results.Results[SuccessModel[int], ErrorModel[string]] {
	return &results.Results[SuccessModel[int], ErrorModel[string]]{
		Failure:      e,
		Success:      s,
		IsSuccessful: true,
	}
}

type SuccessModel[T int] struct {
	Value T
}

func (v *SuccessModel[int]) SetValue(value int) {
	v.Value = value
}

func (v *SuccessModel[int]) GetValue() int {
	return 0
}

type ErrorModel[T string] struct {
	Value T
}

func (v *ErrorModel[string]) SetValue(value string) {
	v.Value = value
}

func (v *ErrorModel[string]) GetValue() string {
	return v.Value
}
