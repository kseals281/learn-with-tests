package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"Empty slice",
			[]int{},
			0,
		}, {
			"Single num",
			[]int{1},
			1,
		}, {
			"Multiple nums",
			[]int{1, 2, 3, 4, 5},
			15,
		}, {
			"Negative nums",
			[]int{-3, 4},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.nums); got != tt.want {
				t.Errorf("sum gave '%d', want '%d', started with %v", got, tt.want, tt.nums)
			}
		})
	}
}

func ExampleSum() {
	nums := []int{1, 2, 3}
	got := Sum(nums)
	fmt.Println(got)
	// Output: 6
}

func TestSumAll(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			"Two lists of 2",
			[][]int{{1, 2}, {3, 4}},
			[]int{3, 7},
		}, {
			"Empty slices",
			[][]int{},
			nil,
		}, {
			"One empty slice",
			[][]int{{}, {1, 2, 3}},
			[]int{0, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAll(tt.lists...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSumAll() {
	nums := [][]int{{0, 1, 2}, {3, 4, 5}}
	got := SumAll(nums...)
	fmt.Println(got)
	// Output: [3 12]
}

func TestSumAllTails(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{
			"Two elements each",
			[][]int{{0, 1}, {2, 3}},
			[]int{1, 3},
		}, {
			"Empty slices",
			[][]int{},
			nil,
		}, {
			"One empty slice",
			[][]int{{}, {2, 4, 6}},
			[]int{0, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAllTails(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAllTails() = %v, want %v, gave %v", got, tt.want, tt.args)
			}
		})
	}
}

func ExampleSumAllTails() {
	nums := [][]int{{0, 1, 2}, {3, 4, 5}}
	got := SumAllTails(nums...)
	fmt.Println(got)
	// Output: [3 9]
}
