package slices

import (
	"fmt"
	"reflect"
	"testing"
)

const errorFormat = "\ngot: %+v\nexp: %+v\n"

var (
	intSlice         = []int{1, 2, 3, 4, 5}
	intSliceReversed = []int{5, 4, 3, 2, 1}
	floatSlice       = []float64{1, 2, 3, 4, 5}
	stringSlice      = []string{"one", "two", "three", "four", "five"}
)

type testStruct struct {
	field int
}

func (s testStruct) String() string {
	return fmt.Sprintf("%d", s.field)
}

func TestRemoveValue(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		val   int
		exp   []int
	}{
		{name: "empty", input: []int{}, exp: []int{}},
		{name: "3", input: intSlice, val: 3, exp: []int{1, 2, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveValue(tt.input, tt.val); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestRemoveIdx(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		idx   int
		exp   []string
	}{
		{name: "empty", input: []string{}, idx: 0, exp: []string{}},
		{name: "1", input: stringSlice, idx: 1, exp: []string{"one", "three", "four", "five"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveIdx(tt.input, tt.idx); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestHasValue(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		val   string
		exp   bool
	}{
		{name: "empty", input: []string{}, val: "", exp: false},
		{name: "false", input: stringSlice, val: "none", exp: false},
		{name: "true", input: stringSlice, val: "four", exp: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasValue(tt.input, tt.val); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		val   int
		exp   int
	}{
		{name: "empty", input: []int{}, val: 1, exp: -1},
		{name: "-1", input: intSlice, val: 42, exp: -1},
		{name: "0", input: intSlice, val: 1, exp: 0},
		{name: "1", input: intSlice, val: 2, exp: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.input, tt.val); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   []int
	}{
		{name: "empty", input: []int{}, exp: []int{}},
		{name: "123", input: intSlice, exp: intSlice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy(tt.input)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
			if len(tt.input) == 0 {
				return
			}
			got[0] = 0
			if tt.exp[0] == 0 {
				t.Error()
			}
		})
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		others [][]string
		exp    []string
	}{
		{name: "empty", input: []string{}, others: [][]string{}, exp: []string{}},
		{name: "empty_diff", input: stringSlice, others: [][]string{}, exp: stringSlice},
		{name: "one", input: stringSlice, others: [][]string{{"two", "three"}, {"four", "five"}}, exp: []string{"one"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.input, tt.others...); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		others [][]string
		exp    []string
	}{
		{name: "empty", input: []string{}, others: [][]string{}, exp: []string{}},
		{name: "empty_others", input: stringSlice, others: [][]string{}, exp: []string{}},
		{name: "one_empty_other", input: stringSlice, others: [][]string{{"nine", "one"}, {}}, exp: []string{}},
		{name: "no_intersect", input: stringSlice, others: [][]string{{"nine", "ten"}}, exp: []string{}},
		{name: "one", input: stringSlice, others: [][]string{{"one", "none"}, {"nine", "one"}}, exp: []string{"one"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.input, tt.others...); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSafeSlice(t *testing.T) {
	arr := Copy(stringSlice)

	r := SafeSlice(arr, 1, 3)

	arr[1] = "new"

	if r[0] != "two" {
		t.Errorf(errorFormat, r, []string{"two", "three"})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		chunkSize int
		exp       [][]string
	}{
		{name: "empty", input: []string{}, chunkSize: 0, exp: [][]string{}},
		{name: "1", input: stringSlice, chunkSize: 5, exp: [][]string{stringSlice}},
		{name: "2", input: stringSlice, chunkSize: 2, exp: [][]string{{"one", "two"}, {"three", "four"}, {"five"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.input, tt.chunkSize); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		others [][]int
		exp    []int
	}{
		{name: "empty", input: []int{}, others: [][]int{}, exp: []int{}},
		{name: "123", input: []int{1}, others: [][]int{{2, 3}, {1, 3}}, exp: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.input, tt.others...); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   []int
	}{
		{name: "empty", input: []int{}, exp: []int{}},
		{name: "123", input: []int{1, 1, 2, 3, 2, 3}, exp: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.input); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestFill(t *testing.T) {
	tests := []struct {
		name  string
		value int
		count int
		exp   []int
	}{
		{name: "empty", value: 1, count: 0, exp: []int{}},
		{name: "111", value: 1, count: 3, exp: []int{1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fill(tt.value, tt.count); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		f     func(int) bool
		exp   []int
	}{
		{name: "empty", input: []int{}, f: nil, exp: []int{}},
		{name: "123", input: intSlice, f: func(i int) bool { return i < 4 }, exp: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.input, tt.f); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		f     func(string) string
		exp   []string
	}{
		{name: "empty", input: []string{}, f: nil, exp: []string{}},
		{name: "one!", input: []string{"one"}, f: func(s string) string { return s + "!" }, exp: []string{"one!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.input, tt.f); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   []float64
	}{
		{name: "empty", input: []int{}, exp: []float64{}},
		{name: "123", input: intSlice, exp: floatSlice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert[int, float64](tt.input); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		exp   float64
	}{
		{name: "empty", input: []float64{}, exp: float64(0)},
		{name: "1", input: []float64{3, 4, 1, 5, 2}, exp: float64(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.input); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   int
	}{
		{name: "empty", input: []int{}, exp: 0},
		{name: "5", input: intSlice, exp: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.input); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   int
	}{
		{name: "empty", input: []int{}, exp: 0},
		{name: "15", input: intSlice, exp: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.input); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   []int
	}{
		{name: "empty", input: []int{}, exp: []int{}},
		{name: "123", input: intSlice, exp: intSliceReversed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.input); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name   string
		input  []float64
		format string
		exp    []string
	}{
		{name: "empty", input: []float64{}, exp: []string{}},
		{
			name:   "123",
			input:  []float64{0.000001, 0.02, 0.300000003},
			format: "%.2f",
			exp:    []string{"0.00", "0.02", "0.30"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.input, tt.format); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestStringify(t *testing.T) {
	tests := []struct {
		name  string
		input []testStruct
		exp   []string
	}{
		{name: "empty", input: []testStruct{}, exp: []string{}},
		{name: "123", input: []testStruct{{1}, {2}, {3}}, exp: []string{"1", "2", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stringify(tt.input); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestRange(t *testing.T) {
	tests := []struct {
		name  string
		start int
		stop  int
		step  int
		exp   []int
	}{
		{name: "empty", exp: []int{}},
		{name: "impossible", start: 3, stop: 0, step: 1, exp: []int{}},
		{name: "forward", start: 0, stop: 3, step: 1, exp: []int{0, 1, 2}},
		{name: "backward", start: 2, stop: -1, step: -1, exp: []int{2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Range(tt.start, tt.stop, tt.step); !reflect.DeepEqual(got, tt.exp) {
				t.Errorf(errorFormat, got, tt.exp)
			}
		})
	}
}

func TestSequenceGenerator(t *testing.T) {
	tests := []struct {
		name  string
		start int
		step  int
		count int
		exp   int
	}{
		{name: "empty"},
		{name: "3", start: 0, step: 1, count: 3, exp: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := SequenceGenerator(tt.start, tt.step)
			var x int
			for i := 0; i < tt.count; i++ {
				x = gen()
			}
			if x != tt.exp {
				t.Errorf(errorFormat, x, tt.exp)
			}
		})
	}
}
