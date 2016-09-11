package cfilter

import (
	"hash"
	"math/rand"
)

// CFilter represents a Cuckoo Filter, a probabilistic data store
// for approximated set membership queries.
type CFilter struct {
	hashfn  hash.Hash // Hash function used for fingerprinting
	buckets []bucket  // Buckets where fingerprints are stored
	count   uint

	bSize  uint8 // Bucket size
	fpSize uint8 // Fingerprint size
	size   uint  // Number of buckets in the filter
	kicks  uint  // Maximum number of times we kick down items from buckets
}

// New returns a new CFilter object. It's Insert, Lookup, Delete and
// Size behave as their names suggest.
// Takes zero or more of the following option functions and applies them in
// order to the Filter:
//      - cfilter.Size(uint) sets the number of buckets in the filter
//      - cfilter.BucketSize(uint8) sets the size of each bucket
//      - cfilter.FingerprintSize(uint8) sets the size of the fingerprint
//      - cfilter.MaximumKicks(uint) sets the maximum number of bucket kicks
//      - cfilter.HashFn(hash.Hash) sets the fingerprinting hashing function
func New(opts ...option) *CFilter {
	cf := new(CFilter)
	for _, opt := range opts {
		opt(cf)
	}
	configure(cf)

	cf.buckets = make([]bucket, cf.size, cf.size)
	for i := range cf.buckets {
		cf.buckets[i] = make([]fingerprint, cf.bSize, cf.bSize)
	}

	return cf
}

// Insert adds an element (in byte-array form) to the Cuckoo filter,
// returns true if successful and false otherwise.
func (cf *CFilter) Insert(item []byte) bool {
	f := fprint(item, cf.fpSize, cf.hashfn)
	j := hashfp(item) % cf.size
	k := (j ^ hashfp(f)) % cf.size

	if cf.buckets[j].insert(f) || cf.buckets[k].insert(f) {
		cf.count++
		return true
	}

	i := [2]uint{j, k}[rand.Intn(2)]
	for n := uint(0); n < cf.kicks; n++ {
		f = cf.buckets[i].swap(f)
		i ^= hashfp(f) % cf.size

		if cf.buckets[i].insert(f) {
			cf.count++
			return true
		}
	}

	return false
}

// Lookup checks if an element (in byte-array form) exists in the Cuckoo
// Filter, returns true if found and false otherwise.
func (cf *CFilter) Lookup(item []byte) bool {
	f := fprint(item, cf.fpSize, cf.hashfn)
	j := hashfp(item) % cf.size
	k := (j ^ hashfp(f)) % cf.size

	return cf.buckets[j].lookup(f) || cf.buckets[k].lookup(f)
}

// Delete removes an element (in byte-array form) from the Cuckoo Filter,
// returns true if element existed prior and false otherwise.
func (cf *CFilter) Delete(item []byte) bool {
	f := fprint(item, cf.fpSize, cf.hashfn)
	j := hashfp(item) % cf.size
	k := (j ^ hashfp(f)) % cf.size

	if cf.buckets[j].remove(f) || cf.buckets[k].remove(f) {
		cf.count--
		return true
	}

	return false
}

// Count returns the total number of elements currently in the Cuckoo Filter.
func (cf *CFilter) Count() uint {
	return cf.count
}
