package sort

import (
	"fmt"
	"testing"
)

func Test_insertSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				[]int{2, 1, 5, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertSort(tt.args.array)
			fmt.Println(tt.args.array)
		})
	}
}
