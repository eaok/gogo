package sort

import (
	"reflect"
	"testing"
)

func Test_shellSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{[]int{5, 4, 9, 8, 7, 6, 0, 1, 3, 2}},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
