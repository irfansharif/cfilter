package cfilter

import (
	"hash"
	"hash/fnv"
)

type option func(*CFilter)

// Size sets the number of buckets in the filter.
// Defaults to ((1 << 18) / BucketSize).
func Size(s uint) option {
	return func(cf *CFilter) {
		cf.size = s
	}
}

// BucketSize sets the size of each bucket in the filter. Defaults to 4.
func BucketSize(s uint8) option {
	return func(cf *CFilter) {
		cf.bSize = s
	}
}

// FingerprintSize sets the size of the fingerprint. Defaults to 2.
func FingerprintSize(s uint8) option {
	return func(cf *CFilter) {
		cf.fpSize = s
	}
}

// MaximumKicks sets the maximum number of times we kick down items/displace
// from their buckets. Defaults to 500.
func MaximumKicks(k uint) option {
	return func(cf *CFilter) {
		cf.kicks = k
	}
}

// HashFn sets the hashing function to be used for fingerprinting. Defaults to
// a 64-bit FNV-1 hash.Hash.
func HashFn(hashfn hash.Hash) option {
	return func(cf *CFilter) {
		cf.hashfn = hashfn
	}
}

func configure(cf *CFilter) {
	if cf.hashfn == nil {
		cf.hashfn = fnv.New64()
	}
	if cf.bSize == 0 {
		cf.bSize = 4
	}
	if cf.fpSize == 0 {
		cf.fpSize = 2
	}
	if cf.kicks == 0 {
		cf.kicks = 500
	}
	if cf.size == 0 {
		cf.size = (1 << 18) / uint(cf.bSize)
	}
}
