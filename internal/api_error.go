package internal

type ErrorModel struct {
	Value string
}

func NewErrorModel(msg string) *ErrorModel {
	return &ErrorModel{
		Value: msg,
	}
}

func (v *ErrorModel) SetValue(value string) {
	v.Value = value
}

func (v *ErrorModel) GetValue() string {
	return v.Value
}
