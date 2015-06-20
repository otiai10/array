package array

type Array struct {
	elements []interface{}
	Equals   func(a, b interface{}) bool
}

var (
	defaultEqual = func(a, b interface{}) bool {
		return (a == b)
	}
)

func New(elements ...interface{}) *Array {
	return &Array{
		elements: elements,
		Equals:   defaultEqual,
	}
}

func (arr *Array) Has(v interface{}) bool {
	for _, e := range arr.elements {
		if arr.Equals(v, e) {
			return true
		}
	}
	return false
}
