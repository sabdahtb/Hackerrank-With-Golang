package arraysum

import "testing"

func TestArraySum(t *testing.T) {
	type args struct {
		ar []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "ArraySumTest",
			args: args{
				ar: []int32{1, 2, 3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArraySum(tt.args.ar); got != tt.want {
				t.Errorf("ArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
