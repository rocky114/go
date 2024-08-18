package algorithm

import "testing"

func Test_partition(t *testing.T) {
	type args struct {
		list  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.list, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuicksort(t *testing.T) {
	type args struct {
		list  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				list:  []int{1, 2, 4, 2, 6, 9, 3, 5},
				left:  0,
				right: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Quicksort(tt.args.list, tt.args.left, tt.args.right)
		})
	}
}
