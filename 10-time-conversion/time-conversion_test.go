package timeconversion

import "testing"

func TestTimeConversion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TimeConversionTest",
			args: args{
				s: "07:05:45PM",
			},
			want: "19:05:45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeConversion(tt.args.s); got != tt.want {
				t.Errorf("TimeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
