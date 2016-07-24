package cfilter

import "math/rand"

const bSize = 4

type bucket []fingerprint

func (b bucket) insert(f fingerprint) bool {
	for i, fp := range b {
		if fp == nil {
			b[i] = f
			return true
		}
	}

	return false
}

func (b bucket) swap(f fingerprint) fingerprint {
	i := rand.Intn(bSize - 1)
	b[i], f = f, b[i]
	return f
}
