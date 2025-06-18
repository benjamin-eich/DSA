package dynamic_programming

import "testing"

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Fib(0)",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "Fib(1)",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Fib(2)",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "Fib(3)",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "Fib(4)",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "Fib(5)",
			args: args{n: 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibMemo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "FibMemo(0)",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "FibMemo(1)",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "FibMemo(2)",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "FibMemo(3)",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "FibMemo(4)",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "FibMemo(5)",
			args: args{n: 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memo := make(map[int]int)
			if got := FibMemo(tt.args.n, memo); got != tt.want {
				t.Errorf("FibMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibBottomUp(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "FibBottomUp(0)",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "FibBottomUp(1)",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "FibBottomUp(2)",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "FibBottomUp(3)",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "FibBottomUp(4)",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "FibBottomUp(5)",
			args: args{n: 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibBottomUp(tt.args.n); got != tt.want {
				t.Errorf("FibBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
