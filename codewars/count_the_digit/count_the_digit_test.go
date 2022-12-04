package countthedigit

import "testing"

type args struct {
	n int
	d int
}

func TestNbDig(t *testing.T) {

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				n: 5750,
				d: 0,
			},
			want: 4700,
		},
		{
			name: "test",
			args: args{
				n: 11011,
				d: 2,
			},
			want: 9481,
		},
		{
			name: "test",
			args: args{
				n: 550,
				d: 5,
			},
			want: 213,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NbDig(tt.args.n, tt.args.d); got != tt.want {
				t.Errorf("NbDig() = %v, want %v", got, tt.want)
			}
		})
	}
}
