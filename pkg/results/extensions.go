package results

// Map functions take in a result object from the previous chain and will evaluate the track to see if we take the exit path
func (r *Results[SuccessIFace, ErrorIFace]) Map(inbound func(s *SuccessIFace, f *ErrorIFace, re *Results[SuccessIFace, ErrorIFace]) *Results[SuccessIFace, ErrorIFace]) *Results[SuccessIFace, ErrorIFace] {
	if r.IsSuccess() {
		return inbound(r.Success, r.Failure, r)
	}

	return r.Failed(r.Failure)
}

// Tee function is for functionality that primairly for side effects.
func (r *Results[SuccessIFace, ErrorIFace]) Tee(inbound func(s *SuccessIFace)) *Results[SuccessIFace, ErrorIFace] {
	if r.IsSuccess() {
		inbound(r.Success)
	}
	return r
}
