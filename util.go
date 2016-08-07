package cfilter

func hash(item []byte) uint {
	var h uint = 5381
	for i := 0; i < len(item); i++ {
		h = ((h << 5) + h) + uint(item[i])
	}

	return h
}

func equal(a, b fingerprint) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
