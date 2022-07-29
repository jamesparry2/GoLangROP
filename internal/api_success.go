package internal

type SuccessModel struct {
	Value int
}

func NewSuccessModel() *SuccessModel {
	return &SuccessModel{
		Value: 0,
	}
}

func (v *SuccessModel) SetValue(value int) {
	v.Value = value
}

func (v *SuccessModel) GetValue() int {
	return v.Value
}
