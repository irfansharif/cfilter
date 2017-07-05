package cfilter

import (
	"bytes"
	"hash"
)

type fingerprint []byte

func fprint(item []byte, fpSize uint8, hashfn hash.Hash) fingerprint {
	hashfn.Reset()
	hashfn.Write(item)
	h := hashfn.Sum(nil)

	fp := make(fingerprint, fpSize)
	copy(fp, h)

	return fp
}

func hashfp(f fingerprint) uint {
	var h uint = 5381
	for i := range f {
		h = ((h << 5) + h) + uint(f[i])
	}

	return h
}

func match(a, b fingerprint) bool {
	return bytes.Equal(a, b)
}
