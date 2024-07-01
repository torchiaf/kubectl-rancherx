package log

import (
	"fmt"
	"testing"
)

func Test_trimColorLabels(t *testing.T) {
	type args struct {
		src string
	}

	type test struct {
		name string
		args args
		want string
	}

	tests := []test{}

	for k, v := range customLevels {
		tests = append(tests, test{
			name: fmt.Sprintf("test trim %s color string", k),
			args: args{
				src: fmt.Sprintf("%s%s: foo-bar", v.color, v.label),
			},
			want: fmt.Sprintf("%s: foo-bar", v.label),
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimColorLabels(tt.args.src); got != tt.want {
				t.Errorf("trimColorLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
