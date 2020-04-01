package dependency_injection

import (
	"bytes"
	"testing"
)

func Test_greet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"non-empty string",
			args{"foo"},
			"Hello, foo!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := bytes.Buffer{}
			greet(&buffer, tt.args.name)
			got := buffer.String()
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
