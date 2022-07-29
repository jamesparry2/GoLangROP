package main

import (
	"fmt"

	"github.com/jamesparry2/GoLangROP/internal"
)

func main() {
	apiHandler := internal.NewHandler()

	if success, err := apiHandler.
		Results.
		Map(apiHandler.GetCarDataHandler).
		// Map(apiHandler.GetFailure).
		Map(apiHandler.GetCarDataHandler).
		Outcome(); err != nil {
		fmt.Println(err.GetValue())
	} else {
		fmt.Println(success.GetValue())
	}
}
