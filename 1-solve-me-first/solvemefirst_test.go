package solvemefirst

import "testing"

func Test_solveMeFirst(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "TestSolveMeFirst",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveMeFirst(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solveMeFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
