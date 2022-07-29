package results

func NewResult[S any, F any](success *S, failure *F) *Results[S, F] {
	return &Results[S, F]{
		Failure:      failure,
		Success:      success,
		IsSuccessful: true,
	}
}

type Results[S any, E any] struct {
	Failure      *E
	Success      *S
	IsSuccessful bool
}

func (r *Results[S, E]) Succeeded(success *S) *Results[S, E] {
	r.Success = success
	r.IsSuccessful = true
	return r
}

func (r *Results[S, E]) Failed(failed *E) *Results[S, E] {
	r.Failure = failed
	r.IsSuccessful = false
	return r
}

func (r *Results[S, E]) IsSuccess() bool {
	return r.IsSuccessful
}

func (r *Results[S, E]) IsFailure() bool {
	return !r.IsSuccessful
}

func (r *Results[S, E]) Outcome() (*S, *E) {
	if r.IsSuccessful {
		return r.Success, nil
	}

	return nil, r.Failure
}
