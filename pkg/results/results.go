package results

func NewResult(s *Success, e *Error) *Results[Success, Error] {
	return &Results[Success, Error]{
		Success:      s,
		Failure:      e,
		isSuccessful: true,
	}
}

type Results[S Success, E Error] struct {
	Failure      *E
	Success      *S
	isSuccessful bool
}

func (r *Results[S, E]) Succeeded(success *S) *Results[S, E] {
	r.Success = success
	r.isSuccessful = true
	return r
}

func (r *Results[S, E]) Failed(failed *E) *Results[S, E] {
	r.Failure = failed
	r.isSuccessful = false
	return r
}

func (r *Results[T, E]) IsSuccess() bool {
	return r.isSuccessful
}

func (r *Results[T, E]) IsFailure() bool {
	return !r.isSuccessful
}

func (r *Results[T, E]) Outcome() interface{} {
	if r.isSuccessful {
		return r.Success
	}

	return r.Failure
}
