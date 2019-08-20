package numbercrunching

import (
	"testing"
)

func Test_isPower1(t *testing.T) {
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
			name: "case1",
			args: args{9},
			want: true,
		},
		{
			name: "case2",
			args: args{64},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPower1(tt.args.n); got != tt.want {
				t.Errorf("isPower1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPower2(t *testing.T) {
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
			name: "9",
			args: args{9},
			want: true,
		},
		{
			name: "64",
			args: args{64},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPower2(tt.args.n); got != tt.want {
				t.Errorf("isPower2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPower3(t *testing.T) {
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
			name: "9",
			args: args{9},
			want: true,
		},
		{
			name: "64",
			args: args{64},
			want: true,
		},
		// {
		// 	name: "66",
		// 	args: args{66},
		// 	want: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPower3(tt.args.n); got != tt.want {
				t.Errorf("isPower3() = %v, want %v", got, tt.want)
			}
		})
	}
}
