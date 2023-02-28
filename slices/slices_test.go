package slices

import (
	"strings"
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
				t.Errorf("CommonSlice Error got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestDiffSlice(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{
			name:   "slice1がslice2に包含されている",
			slice1: []int{1, 3, 5},
			slice2: []int{1, 2, 3, 4, 5},
			want:   []int{},
		},
		{
			name:   "slice1がslice2を包含している",
			slice1: []int{1, 2, 3, 4, 5},
			slice2: []int{1, 3, 5},
			want:   []int{2, 4},
		},
		{
			name:   "共通している要素がない",
			slice1: []int{1, 2, 3, 4, 5},
			slice2: []int{0, 6, 8},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "1つのsliceに共通している要素が複数ある",
			slice1: []int{1, 2, 3, 3, 4, 5},
			slice2: []int{1, 1, 6, 8},
			want:   []int{2, 3, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := DiffSlice(tt.slice1, tt.slice2)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("DiffSlice Error got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSlice(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		target int
		want   []int
	}{
		{
			name:   "sliceに削除対象がある時は、その要素が削除される",
			slice:  []int{1, 3, 5},
			target: 1,
			want:   []int{3, 5},
		},
		{
			name:   "sliceに削除対象が複数ある時は、対象を全て削除する",
			slice:  []int{1, 2, 3, 3},
			target: 3,
			want:   []int{1, 2},
		},
		{
			name:   "sliceに削除対象がない時は、何もしない",
			slice:  []int{1, 2, 3, 4, 5},
			target: 6,
			want:   []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteValue(tt.slice, tt.target)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("DeleteSlice Error got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestMergeSlice(t *testing.T) {
	tests := []struct {
		name   string
		slices [][]int
		want   []int
	}{
		{
			name: "全てのsliceが連携できる",
			slices: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "全てのsliceが連携できる",
			slices: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
				{9, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "空のsliceがあった時は無視される",
			slices: [][]int{
				{1, 2},
				{},
				{5, 6},
			},
			want: []int{1, 2, 5, 6},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSlice(tt.slices...)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("MergeSlice Error got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestFilterSlice(t *testing.T) {
	type Test struct {
		ID   int
		Name string
	}
	tests := []struct {
		name  string
		slice []Test
		fn    func(i Test) bool
		want  []Test
	}{
		{
			name: "work",
			slice: []Test{
				{
					ID:   1,
					Name: "test1",
				},
				{
					ID:   2,
					Name: "test2",
				},
				{
					ID:   3,
					Name: "error1",
				},
				{
					ID:   4,
					Name: "test3",
				},
			},
			fn: func(t Test) bool {
				return t.ID%2 == 1 && strings.Contains(t.Name, "test")
			},
			want: []Test{
				{
					ID:   1,
					Name: "test1",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := FilterSlice(tt.slice, tt.fn)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("FilterSlice Error got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestBatchSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		fn    func(t int) (int, error)
		want  []int
	}{
		{
			name:  "work",
			slice: []int{1, 2, 3, 4},
			fn: func(t int) (int, error) {
				return t * 2, nil
			},
			want: []int{2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := BatchSlice(tt.slice, tt.fn)
			diff := cmp.Diff(tt.want, got)
			if diff != "" || err != nil {
				t.Errorf("BatchSlice Error got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestCountValueSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  map[int]int
	}{
		{
			name:  "各要素の数を取得できる",
			slice: []int{1, 2, 3, 4},
			want: map[int]int{
				1: 1,
				2: 1,
				3: 1,
				4: 1,
			},
		},
		{
			name:  "被っている要素の数を取得できる",
			slice: []int{1, 2, 3, 4, 1, 3, 1},
			want: map[int]int{
				1: 3,
				2: 1,
				3: 2,
				4: 1,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := CountValueSlice(tt.slice)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("FilterSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
