package reverse_string

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple String",
			args: args{
				input: "Hello World!",
			},
			want: "!dlroW olleH",
		},
		{
			name: "Empty String",
			args: args{
				input: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
