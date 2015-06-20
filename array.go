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

type IntArray []int
type UintArray []uint
type StringArray []string

func Ints(arr ...int) IntArray {
	return IntArray(arr)
}

func (arr IntArray) Has(i int) bool {
	for _, e := range arr {
		if i == e {
			return true
		}
	}
	return false
}

func Uints(arr ...uint) UintArray {
	return UintArray(arr)
}

func (arr UintArray) Has(u uint) bool {
	for _, e := range arr {
		if u == e {
			return true
		}
	}
	return false
}

func Strings(arr ...string) StringArray {
	return StringArray(arr)
}

func (arr StringArray) Has(s string) bool {
	for _, e := range arr {
		if s == e {
			return true
		}
	}
	return false
}
