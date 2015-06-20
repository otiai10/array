package array

import (
	. "github.com/otiai10/mint"
	"testing"
)

func TestNew(t *testing.T) {
	arr := New("a", "b", "c", "d")
	Expect(t, arr).TypeOf("*array.Array")
}

func TestArray_Has(t *testing.T) {
	arr := New("foo", "bar", "buz")
	Expect(t, arr.Has("foo")).ToBe(true)
	Expect(t, arr.Has("xxx")).ToBe(true)
}
