package results

// func
type Results[S any, E any] struct {
	Failure      *E
	Success      *S
	IsSuccessful bool
}

func (r *Results[SuccessIFace, ErrorIFace]) Succeeded(success *SuccessIFace) *Results[SuccessIFace, ErrorIFace] {
	r.Success = success
	r.IsSuccessful = true
	return r
}

func (r *Results[SuccessIFace, ErrorIFace]) Failed(failed *ErrorIFace) *Results[SuccessIFace, ErrorIFace] {
	r.Failure = failed
	r.IsSuccessful = false
	return r
}

func (r *Results[SuccessIFace, ErrorIFace]) IsSuccess() bool {
	return r.IsSuccessful
}

func (r *Results[SuccessIFace, ErrorIFace]) IsFailure() bool {
	return !r.IsSuccessful
}

func (r *Results[SuccessIFace, ErrorIFace]) Outcome() (*SuccessIFace, *ErrorIFace) {
	if r.IsSuccessful {
		return r.Success, nil
	}

	return nil, r.Failure
}
