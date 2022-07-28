package results

type Success struct {
	value interface{}
}

func NewSuccess(value interface{}) *Success {
	return &Success{value: value}
}

func (s *Success) GetValue() interface{} {
	return s.value
}

type Error struct {
	Value interface{}
}

func NewError(value interface{}) *Error {
	return &Error{Value: value}
}

func (e *Error) GetValue() interface {
}
