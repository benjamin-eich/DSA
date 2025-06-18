package dynamic_programming

import "testing"

func TestJumpBottomUp(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{n: 5},
			want: 8,
		},
		{
			name: "45",
			args: args{n: 45},
			want: 1836311903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JumpBottomUp(tt.args.n); got != tt.want {
				t.Errorf("JumpBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJumpRecursive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{n: 5},
			want: 8,
		},
		{
			name: "45",
			args: args{n: 45},
			want: 1836311903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JumpRecursive(tt.args.n); got != tt.want {
				t.Errorf("JumpRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJumpRecursiveMemo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{n: 5},
			want: 8,
		},
		{
			name: "45",
			args: args{n: 45},
			want: 1836311903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memo := make(map[int]int)
			if got := JumpRecursiveMemo(tt.args.n, memo); got != tt.want {
				t.Errorf("JumpRecursiveMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
