package results

type ResultTeeArgs[S any] func(s *S)

// Map functions take in a result object from the previous chain and will evaluate the track to see if we take the exit path
func (r *Results[S, E]) Map(inbound func(s *S) *Results[S, E]) *Results[S, E] {
	if r.IsSuccess() {
		return inbound(r.Success)
	}

	return r.Failed(r.Failure)
}

// Tee function is for functionality that primairly for side effects.
func (r *Results[S, E]) Tee(inbound ResultTeeArgs[S]) *Results[S, E] {
	if r.IsSuccess() {
		inbound(r.Success)
	}
	return r
}
