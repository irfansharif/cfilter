package cfilter

import (
	"bytes"
	"hash"
)

const fpSize = 2

type fingerprint []byte

func fprint(item []byte, hashfn hash.Hash64) fingerprint {
	hashfn.Reset()
	hashfn.Write(item)
	h := hashfn.Sum(nil)

	fp := fingerprint{}
	for i := 0; i < fpSize; i++ {
		fp = append(fp, h[i])
	}

	if fp == nil {
		fp[0] += 7
	}

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
