package array

import (
	"testing"
)

func Test_findDupByHash(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				[]int{1, 3, 4, 2, 5, 3},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				[]int{1, 2, 4, 2, 5, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDupByHash(tt.args.arr); got != tt.want {
				t.Errorf("findDupByHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDupByXOR(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				[]int{1, 3, 4, 2, 5, 3},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				[]int{1, 2, 4, 2, 5, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDupByXOR(tt.args.arr); got != tt.want {
				t.Errorf("findDupByXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDupByMap(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				[]int{1, 3, 4, 2, 5, 3},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				[]int{1, 2, 4, 2, 5, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDupByMap(tt.args.arr); got != tt.want {
				t.Errorf("findDupByMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDupByLoop(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				[]int{1, 3, 4, 2, 5, 3},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				[]int{1, 2, 4, 2, 5, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDupByLoop(tt.args.arr); got != tt.want {
				t.Errorf("findDupByLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
