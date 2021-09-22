package algorithmThought

import "testing"

func TestFibonacciDP(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "data1",
			args: args{3},
			want: 2,
		},
		{
			name: "data2",
			args: args{4},
			want: 3,
		},
		{
			name: "data3",
			args: args{6},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciDP(tt.args.n); got != tt.want {
				t.Errorf("FibonacciDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciDPOptimize(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "data1",
			args: args{3},
			want: 2,
		},
		{
			name: "data2",
			args: args{4},
			want: 3,
		},
		{
			name: "data3",
			args: args{6},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciDPOptimize(tt.args.n); got != tt.want {
				t.Errorf("FibonacciDPOptimize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciRE(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "data1",
			args: args{3},
			want: 2,
		},
		{
			name: "data2",
			args: args{4},
			want: 3,
		},
		{
			name: "data3",
			args: args{6},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciRE(tt.args.n); got != tt.want {
				t.Errorf("FibonacciRE() = %v, want %v", got, tt.want)
			}
		})
	}
}
