package util

func NilInt64(i int64) *int64 {
	if i != 0 {
		return &i
	}

	return nil
}

func NilStringArray(arr []string) *[]string {
	if len(arr) > 0 {
		return &arr
	}

	return nil
}

func NilString(str string) *string {
	if len(str) > 0 {
		return &str
	}

	return nil
}

func NilUint32(i uint32) *uint32 {
	if i != 0 {
		return &i
	}

	return nil
}

func NilUint64(i uint64) *uint64 {
	if i != 0 {
		return &i
	}

	return nil
}