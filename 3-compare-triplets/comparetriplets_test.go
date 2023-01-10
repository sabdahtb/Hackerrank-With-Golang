package comparetriplets

import (
	"reflect"
	"testing"
)

func TestCompareTriplets(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "CompareTripletsTest",
			args: args{
				a: []int32{17, 28, 30},
				b: []int32{99, 16, 8},
			},
			want: []int32{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTriplets(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
