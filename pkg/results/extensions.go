package results

// Map functions take in a result object from the previous chain and will evaluate the track to see if we take the exit path
func (r *Results[Success, Error]) Map(inbound func(s *Success, re *Results[Success, Error]) *Results[Success, Error]) *Results[Success, Error] {
	if r.IsSuccess() {
		return inbound(r.Success, r)
	}

	return r.Failed(r.Failure)
}

// Tee function is for functionality that primairly for side effects.
func (r *Results[Success, Error]) Tee(inbound func(s *Success)) *Results[Success, Error] {
	if r.IsSuccess() {
		inbound(r.Success)
	}
	return r
}
