package string

import "testing"

func Test_permutationStr(t *testing.T) {
	type args struct {
		str   []rune
		start int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{[]rune("abc"), 0},
			want: "abc acb bac bca cba cab ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutationStr(tt.args.str, tt.args.start); got != tt.want {
				t.Errorf("permutationStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
