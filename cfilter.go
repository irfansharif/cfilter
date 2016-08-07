package cfilter

import (
	"hash"
	"hash/fnv"
	"math/rand"
)

// The maximum number of times we kick down items/displace from their buckets
const maxCuckooCount = 500

// The number of buckets in the filter
const cfSize = (1 << 18) / bSize

// CFilter represents a Cuckoo Filter, a probabilistic data store
// for approximated set membership queries
type CFilter struct {
	size    uint
	hashfn  hash.Hash64
	buckets []bucket
}

// New returns a new CFilter object. It's Insert, Lookup, Delete and
// Size behave as their names suggest.
func New() *CFilter {
	cf := new(CFilter)

	cf.size = 0
	cf.hashfn = fnv.New64()
	cf.buckets = make([]bucket, cfSize, cfSize)
	for i := range cf.buckets {
		cf.buckets[i] = make([]fingerprint, bSize, bSize)
	}

	return cf
}

// Insert adds an element (in byte-array form) to the Cuckoo filter,
// returns true if successful and false otherwise.
func (cf *CFilter) Insert(item []byte) bool {
	f := fprint(item, cf.hashfn)
	j := hashfp(item) % cfSize
	k := (j ^ hashfp(f)) % cfSize

	if cf.buckets[j].insert(f) || cf.buckets[k].insert(f) {
		cf.size++
		return true
	}

	i := [2]uint{j, k}[rand.Intn(2)]
	for n := 0; n < maxCuckooCount; n++ {
		f = cf.buckets[i].swap(f)
		i ^= hashfp(f) % cfSize

		if cf.buckets[i].insert(f) {
			cf.size++
			return true
		}
	}

	return false
}

// Lookup checks if an element (in byte-array form) exists in the Cuckoo
// Filter, returns true if found and false otherwise.
func (cf *CFilter) Lookup(item []byte) bool {
	f := fprint(item, cf.hashfn)
	j := hashfp(item) % cfSize
	k := (j ^ hashfp(f)) % cfSize

	return cf.buckets[j].lookup(f) || cf.buckets[k].lookup(f)
}

// Delete removes an element (in byte-array form) from the Cuckoo Filter,
// returns true if element existed prior and false otherwise.
func (cf *CFilter) Delete(item []byte) bool {
	f := fprint(item, cf.hashfn)
	j := hashfp(item) % cfSize
	k := (j ^ hashfp(f)) % cfSize

	if cf.buckets[j].remove(f) || cf.buckets[k].remove(f) {
		cf.size--
		return true
	}

	return false
}

// Size returns the total number of elements added to the Cuckoo Filter.
func (cf *CFilter) Size() uint {
	return cf.size
}
