package slices

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCommonSlice(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{
			name:   "一方が他方を包含",
			slice1: []int{1, 3, 5},
			slice2: []int{1, 2, 3, 4, 5},
			want:   []int{1, 3, 5},
		},
		{
			name:   "sliceが一致している",
			slice1: []int{1, 3, 5},
			slice2: []int{1, 3, 5},
			want:   []int{1, 3, 5},
		},
		{
			name:   "共通している要素がない",
			slice1: []int{1, 2, 3, 4, 5},
			slice2: []int{0, 6, 8},
			want:   []int{},
		},
		{
			name:   "1つのsliceに共通している要素が複数ある",
			slice1: []int{1, 2, 3, 3, 4, 5},
			slice2: []int{1, 1, 6, 8},
			want:   []int{1},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := CommonSlice(tt.slice1, tt.slice2)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("DiffSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
