package results

type SuccessIFace[T int | string | float32 | interface{}] interface {
	SetValue(T)
	GetValue() T
}

type ErrorIFace[T int | string | float32 | interface{}] interface {
	SetValue(T)
	GetValue() T
}
