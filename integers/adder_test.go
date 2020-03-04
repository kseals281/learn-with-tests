package integers

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"positive ints",
			args{1, 2},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleAdd() {
	sum := add(1, 4)
	fmt.Println(sum)
	// Output: 5
}
