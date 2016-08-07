package cfilter

import "hash/fnv"

const fpSize = 1

type fingerprint []byte

var hashera = fnv.New64()

func fprint(item []byte) fingerprint {
	hashera.Reset()
	hashera.Write(item)
	hash := hashera.Sum(nil)

	fp := fingerprint{}
	for i := 0; i < fpSize; i++ {
		fp = append(fp, hash[i])
	}

	if fp == nil {
		fp[0] += 7
	}

	return fp
}
