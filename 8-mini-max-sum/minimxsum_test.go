package minimaxsum

import (
	"reflect"
	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "MinMaxSumTest",
			args: args{
				arr: []int32{1, 3, 5, 7, 9},
			},
			want: []int32{24, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiniMaxSum(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiniMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
