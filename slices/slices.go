package slices

import ()

// 2つの配列の共通するものを返す。
func CommonSlice[T comparable](slice1, slice2 []T) []T {
	commonSlice := []T{}
	cmpMap := map[T]int{}
	for _, v := range slice1 {
		_, ok := cmpMap[v]
		// 要素の重複は無視
		if !ok {
			cmpMap[v] += 1
		}
	}

	// slice1の要素と一致しているslice2の要素については、返り値のsliceに加える
	for _, v := range slice2 {
		i, ok := cmpMap[v]
		if ok {
			// 要素の重複は無視
			if i == 1 {
				commonSlice = append(commonSlice, v)
			}
			cmpMap[v] += 1
		}
	}
	return commonSlice
}

func DiffSlice[T comparable](slice1, slice2 []T) []T {
	diffSlice := []T{}
	cmpMap := map[T]int{}

	// slice2が各要素が何個あるのかmapに格納
	for _, v := range slice2 {
		cmpMap[v] += 1
	}

	// slice2にある要素をslice2にあるか確認して、なければdiffSliceに格納
	for _, v := range slice1 {
		t, ok := cmpMap[v]
		if ok {
			cmpMap[v] -= 1
			if t == 1 {
				delete(cmpMap, v)
			}
			continue
		}
		diffSlice = append(diffSlice, v)
	}
	return diffSlice
}

func DeleteValue() {

}
