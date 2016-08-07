package cfilter

import (
	"math/rand"
)

const kMaxCuckooCount = 500
const cfSize = (1 << 15) / bSize

type CFilter struct {
	buckets []bucket
	size    uint
}

func NewCFilter() *CFilter {
	buckets := make([]bucket, cfSize, cfSize)

	for b := range buckets {
		buckets[b] = make([]fingerprint, bSize, bSize)
	}

	return &CFilter{
		buckets: buckets,
		size:    0,
	}
}

func (cf *CFilter) Insert(item []byte) bool {
	f := fprint(item)
	j := hash(item) % cfSize
	k := (j ^ hash(f)) % cfSize

	if cf.buckets[j].insert(f) || cf.buckets[k].insert(f) {
		cf.size++
		return true
	}

	i := [2]uint{j, k}[rand.Intn(1)]
	for n := 0; n < kMaxCuckooCount; n++ {
		f := cf.buckets[i].swap(f)
		i ^= hash(f)

		if cf.buckets[i].insert(f) {
			cf.size++
			return true
		}
	}

	return false
}

func (cf *CFilter) Lookup(item []byte) bool {
	f := fprint(item)
	j := hash(item) % cfSize
	k := (j ^ hash(f)) % cfSize

	return cf.buckets[j].lookup(f) || cf.buckets[k].lookup(f)
}

func (cf *CFilter) Delete(item []byte) bool {
	f := fprint(item)
	j := hash(item) % cfSize
	k := (j ^ hash(f)) % cfSize

	if cf.buckets[j].remove(f) || cf.buckets[k].remove(f) {
		cf.size--
		return true
	}

	return false
}

func (cf *CFilter) Size() uint {
	return cf.size
}
