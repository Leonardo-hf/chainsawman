package common

func SliceRepeat[T any](v T, size int) []T {
	retval := make([]T, 0, size)
	for i := 0; i < size; i++ {
		retval = append(retval, v)
	}
	return retval
}

func Bool2Int64(t bool) int64 {
	if t {
		return 1
	}
	return 0
}

func Int642Bool(i int64) bool {
	if i == 1 {
		return true
	}
	return false
}
