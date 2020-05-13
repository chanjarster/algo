package heap

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{[]int{1, 2, 3, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			args: args{[]int{4, 3, 2, 1}},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.data)
			if !reflect.DeepEqual(tt.args.data, tt.want) {
				t.Errorf("data = %v, want %v", tt.args.data, tt.want)
			}
		})
	}
}
