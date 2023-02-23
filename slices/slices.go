package slices

import ()

// 2つの配列の共通するものを返す。
func CommonSlice[T comparable](slice1, slice2 []T) []T {
	commonSlice := []T{}
	cpmMap := map[T]int{}
	for _, v := range slice1 {
		_, ok := cpmMap[v]
		// 要素の重複は無視
		if !ok {
			cpmMap[v] += 1
		}
	}

	// slice1の要素と一致しているslice2の要素については、返り値のsliceに加える
	for _, v := range slice2 {
		i, ok := cpmMap[v]
		if ok {
			// 要素の重複は無視
			if i == 1 {
				commonSlice = append(commonSlice, v)
			}
			cpmMap[v] += 1
		}
	}
	return commonSlice
}
