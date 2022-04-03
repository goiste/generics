package sets

import (
	"reflect"
	"sort"
	"testing"
)

const errorFormat = "\ngot: %+v\nexp: %+v\n"

var (
	stringSet = Set[string]{
		"one":   {},
		"two":   {},
		"three": {},
		"four":  {},
		"five":  {},
	}
	intSet = Set[int]{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
	}
)

func TestMake(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		exp    Set[int]
	}{
		{name: "empty", exp: Set[int]{}},
		{name: "values", values: []int{1, 2, 3, 4, 5}, exp: intSet},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make[int](tt.values...); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	tests := []struct {
		name   string
		set    Set[int]
		values []int
		exp    Set[int]
	}{
		{name: "empty"},
		{name: "123", set: Make[int](), values: []int{1, 2, 3}, exp: Make[int](1, 2, 3)},
		{name: "no_change", set: intSet, values: []int{1, 2, 3}, exp: intSet},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Add(tt.values...)
			if got := tt.set; !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Delete(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		del  []int
		exp  Set[int]
	}{
		{name: "empty"},
		{name: "del_nothing", set: intSet, del: []int{}, exp: intSet},
		{name: "123", set: Make[int](1, 2, 3, 4, 5), del: []int{4, 5, 6, 7}, exp: Make[int](1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Delete(tt.del...)
			if got := tt.set; !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Truncate(t *testing.T) {
	tests := []struct {
		name string
		set  Set[string]
		exp  Set[string]
	}{
		{name: "empty", set: Make[string](), exp: Make[string]()},
		{name: "123", set: Make[string]("1", "2", "3"), exp: Set[string]{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Truncate()
			if got := tt.set; !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Has(t *testing.T) {
	tests := []struct {
		name string
		set  Set[string]
		val  string
		exp  bool
	}{
		{name: "empty"},
		{name: "none", set: stringSet, val: "none", exp: false},
		{name: "one", set: stringSet, val: "one", exp: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Has(tt.val); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Len(t *testing.T) {
	tests := []struct {
		name string
		set  Set[string]
		exp  int
	}{
		{name: "empty"},
		{name: "1", set: Make[string]("one"), exp: 1},
		{name: "5", set: stringSet, exp: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Len(); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Values(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		exp  []int
	}{
		{name: "empty", set: Set[int]{}, exp: []int{}},
		{name: "123", set: Make[int](1, 2, 3), exp: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Values()
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Merge(t *testing.T) {
	tests := []struct {
		name   string
		set    Set[int]
		others []Set[int]
		exp    Set[int]
	}{
		{name: "empty"},
		{name: "123", set: Make[int](1, 2), others: []Set[int]{Make[int](2, 3, 4), Make[int](3, 4, 5)}, exp: intSet},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Merge(tt.others...)
			if got := tt.set; !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Diff(t *testing.T) {
	tests := []struct {
		name   string
		set    Set[int]
		others []Set[int]
		exp    Set[int]
	}{
		{name: "empty"},
		{name: "empty_others", set: intSet, others: []Set[int]{}, exp: intSet},
		{name: "123", set: Make[int](1, 2, 3, 4, 5), others: []Set[int]{{4: {}}, {5: {}}}, exp: Make[int](1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Diff(tt.others...)
			if got := tt.set; !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Intersect(t *testing.T) {
	tests := []struct {
		name   string
		set    Set[int]
		others []Set[int]
		exp    Set[int]
	}{
		{name: "empty", set: Make[int](), exp: Make[int]()},
		{name: "empty_others", set: Make[int](1, 2, 3, 4, 5), others: []Set[int]{}, exp: Make[int]()},
		{name: "123", set: Make[int](1, 2, 3, 4, 5), others: []Set[int]{Make[int](1, 2, 3, 4), Make[int](1, 2, 3, 5)}, exp: Make[int](1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Intersect(tt.others...)
			if got := tt.set; !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Equals(t *testing.T) {
	tests := []struct {
		name  string
		set   Set[string]
		other Set[string]
		exp   bool
	}{
		{name: "empty", exp: true},
		{name: "equals", set: Make[string]("1", "2"), other: Make[string]("1", "2"), exp: true},
		{name: "not_equals", set: stringSet, other: Make[string]("1", "2", "3", "4", "5"), exp: false},
		{name: "not_equals_len", set: stringSet, other: Set[string]{}, exp: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Equals(tt.other); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Filter(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		f    func(int) bool
		exp  Set[int]
	}{
		{name: "empty"},
		{name: "123", set: Make[int](1, 2, 3, 4, 5), f: func(i int) bool { return i < 4 }, exp: Make[int](1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Filter(tt.f)
			if got := tt.set; !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Map(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		f    func(int) int
		exp  Set[int]
	}{
		{name: "empty", set: Make[int](), exp: Make[int]()},
		{name: "+1", set: Make[int](1, 2, 3), f: func(i int) int { return i + 1 }, exp: Make[int](2, 3, 4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Map(tt.f)
			if got := tt.set; !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSet_Copy(t *testing.T) {
	tests := []struct {
		name string
		set  Set[string]
		exp  Set[string]
	}{
		{name: "empty", set: Set[string]{}, exp: Set[string]{}},
		{name: "string", set: stringSet, exp: stringSet},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Copy(); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}
