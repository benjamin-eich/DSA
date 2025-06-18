package divide_and_conquer

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Basic Test",
			args: args{arr: []int{38, 27, 43, 3, 9, 82, 10}},
			want: []int{3, 9, 10, 27, 38, 43, 82},
		},
		{
			name: "Basic Test 2",
			args: args{arr: []int{38, 27, 43, 3, 9, 82}},
			want: []int{3, 9, 27, 38, 43, 82},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
