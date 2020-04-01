package test_demo

import "testing"

func Test_byAdd(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "byAdd1",
			args: args{
				s1: "hello",
				s2: "world",
			},
			want: "hello world",
		},
		{
			name: "byAdd2",
			args: args{
				s1: "hello",
				s2: "kitty",
			},
			want: "hello kitty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byAdd(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("byAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bySprintf(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "bySprintf1",
			args: args{
				s1: "hello",
				s2: "world",
			},
			want: "hello world",
		},
		{
			name: "bySprintf2",
			args: args{
				s1: "hello",
				s2: "kitty",
			},
			want: "hello kitty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bySprintf(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("bySprintf() = %v, want %v", got, tt.want)
			}
		})
	}
}