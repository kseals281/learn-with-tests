package structs_methods_and_interfaces

import "testing"

type shapes interface {
	area() float64
}

func Test_perimeter(t *testing.T) {
	tests := []struct {
		name      string
		rectangle Rectangle
		wantP     float64
	}{
		{
			"Square",
			Rectangle{5.0, 5.0},
			20.0,
		}, {
			"Non-square",
			Rectangle{5.0, 10.0},
			30.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := tt.rectangle.perimeter(); gotP != tt.wantP {
				t.Errorf("perimeter() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func Test_area(t *testing.T) {
	tests := []struct {
		name  string
		shape shapes
		wantA float64
	}{
		{
			"Square",
			Rectangle{5.0, 5.0},
			25.0,
		}, {
			"Non-square",
			Rectangle{5.0, 10.0},
			50.0,
		}, {
			"Non-whole",
			Rectangle{.5, 2},
			1,
		}, {
			"Standard circle",
			Circle{10.0},
			314.1592653589793,
		}, {
			"Standard triangle",
			Triangle{12, 6},
			36.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := tt.shape.area(); gotA != tt.wantA {
				t.Errorf("%#v got %.2f, want %.2f", tt.shape, gotA, tt.wantA)
			}
		})
	}
}
