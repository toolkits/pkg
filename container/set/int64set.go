package set

type Int64Set struct {
	M map[int64]struct{}
}

func NewInt64Set() *Int64Set {
	return &Int64Set{
		M: make(map[int64]struct{}),
	}
}

func (this *Int64Set) Add(elt int64) *Int64Set {
	this.M[elt] = struct{}{}
	return this
}

func (this *Int64Set) Exists(elt int64) bool {
	_, exists := this.M[elt]
	return exists
}

func (this *Int64Set) Delete(elt int64) {
	delete(this.M, elt)
}

func (this *Int64Set) Clear() {
	this.M = make(map[int64]struct{})
}

func (this *Int64Set) ToSlice() []int64 {
	count := len(this.M)
	if count == 0 {
		return []int64{}
	}

	r := make([]int64, count)

	i := 0
	for elt := range this.M {
		r[i] = elt
		i++
	}

	return r
}
