package slices

// 2つのsliceについて、共通するものを返す。
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

// sliceの各要素の数を数える
func CountValueSlice[T comparable](slice []T) map[T]int {
	resultMap := map[T]int{}
	for _, v := range slice {
		resultMap[v] += 1
	}
	return resultMap
}

// 2つのsliceの差分を算出する
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

// 対象となるsliceの要素を削除する
func DeleteValue[T comparable](slice []T, target T) []T {
	resultSlice := []T{}
	for _, v := range slice {
		if v != target {
			resultSlice = append(resultSlice, v)
		}
	}
	return resultSlice
}

// 複数のsliceをmergeする
func MergeSlice[T any](slices ...[]T) []T {
	resultSlice := []T{}
	for _, v := range slices {
		resultSlice = append(resultSlice, v...)
	}
	return resultSlice
}

// sliceの要素をフィルタリングする
func FilterSlice[T any](slice []T, fn func(t T) bool) []T {
	resultSlice := []T{}
	for _, v := range slice {
		if fn(v) {
			resultSlice = append(resultSlice, v)
		}
	}
	return resultSlice
}

// sliceの要素を一括で更新する
func BatchSlice[T any](slice []T, fn func(t T) (T, error)) ([]T, error) {
	resultSlice := make([]T, len(slice))
	for k, v := range slice {
		u, err := fn(v)
		if err != nil {
			return nil, err
		}
		resultSlice[k] = u
	}
	return resultSlice, nil
}
