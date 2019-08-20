package numbercrunching

import (
	"testing"
)

func Test_isNPower1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 4",
			args: args{4},
			want: true,
		},
		{
			name: "case 16",
			args: args{16},
			want: true,
		},
		// {
		// 	name: "case 22",
		// 	args: args{22},
		// 	want: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNPower1(tt.args.n); got != tt.want {
				t.Errorf("isNPower1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNPower2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 4",
			args: args{4},
			want: true,
		},
		{
			name: "case 16",
			args: args{16},
			want: true,
		},
		// {
		// 	name: "case 22",
		// 	args: args{22},
		// 	want: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNPower2(tt.args.n); got != tt.want {
				t.Errorf("isNPower2() = %v, want %v", got, tt.want)
			}
		})
	}
}
