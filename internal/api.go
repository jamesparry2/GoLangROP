package internal

import "github.com/jamesparry2/GoLangROP/pkg/results"

/*
* Example Implementation
 */

type Handler struct {
	Results *results.Results[SuccessModel, ErrorModel]
}

func NewHandler() *Handler {
	return &Handler{
		Results: results.NewResult(NewSuccessModel(), NewErrorModel("")),
	}
}

func (h *Handler) GetCarDataHandler(s *SuccessModel) *results.Results[SuccessModel, ErrorModel] {
	s.SetValue(s.GetValue() + 1)
	return h.Results.Succeeded(s)
}

func (h *Handler) GetFailure(s *SuccessModel) *results.Results[SuccessModel, ErrorModel] {
	return h.Results.Failed(NewErrorModel("Failed: Errored"))
}
