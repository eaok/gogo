package numbercrunching

import (
	"testing"
)

func Test_divide1(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantRes    int
		wantRemain int
	}{
		// TODO: Add test cases.
		{
			name:       "case 5/2",
			args:       args{5, 2},
			wantRes:    2,
			wantRemain: 1,
		},
		{
			name:       "case 4/2",
			args:       args{4, 2},
			wantRes:    2,
			wantRemain: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := divide1(tt.args.m, tt.args.n)
			if got != tt.wantRes {
				t.Errorf("divide1() got = %v, want %v", got, tt.wantRes)
			}
			if got1 != tt.wantRemain {
				t.Errorf("divide1() got1 = %v, want %v", got1, tt.wantRemain)
			}
		})
	}
}

func Test_divide2(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantRes    int
		wantRemain int
	}{
		// TODO: Add test cases.
		{
			name:       "case 5/2",
			args:       args{5, 2},
			wantRes:    2,
			wantRemain: 1,
		},
		{
			name:       "case 4/2",
			args:       args{4, 2},
			wantRes:    2,
			wantRemain: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotRemain := divide2(tt.args.m, tt.args.n)
			if gotRes != tt.wantRes {
				t.Errorf("divide2() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotRemain != tt.wantRemain {
				t.Errorf("divide2() gotRemain = %v, want %v", gotRemain, tt.wantRemain)
			}
		})
	}
}
