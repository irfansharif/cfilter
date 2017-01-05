package cfilter

import "math/rand"

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

func (b bucket) lookup(f fingerprint) bool {
	for _, fp := range b {
		if match(fp, f) {
			return true
		}
	}

	return false
}

func (b bucket) remove(f fingerprint) bool {
	for i, fp := range b {
		if match(fp, f) {
			b[i] = nil
			return true
		}
	}

	return false
}

func (b bucket) swap(f fingerprint) fingerprint {
	i := rand.Intn(len(b))
	b[i], f = f, b[i]

	return f
}
