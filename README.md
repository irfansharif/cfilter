# cfilter: Cuckoo Filter implementation in Go

Cuckoo filter is a Bloom filter replacement for approximated set-membership
queries. Cuckoo filters support adding and removing items dynamically while
achieving even higher performance than Bloom filters. For applications that
store many items and target moderately low false positive rates, cuckoo filters
have lower space overhead than space-optimized Bloom filters.
Some possible use-cases that depend on approximated set-membership queries
would be databases, caches, routers, and storage systems where it is used to
decide if a given item is in a (usually large) set, with some small false
positive probability. Alternatively, given it is designed to be a viable
replacement to Bloom filters, it can also be used to reduce the space required
in probabilistic routing tables, speed longest-prefix matching for IP
addresses, improve network state management and monitoring, and encode
multicast forwarding information in packets, among many other applications.

Cuckoo filters provide the flexibility to add and remove items dynamically. A
cuckoo filter is based on cuckoo hashing (and therefore named as cuckoo
filter).  It is essentially a cuckoo hash table storing each key's fingerprint.
Cuckoo hash tables can be highly compact, thus a cuckoo filter could use less
space than conventional Bloom filters, for applications that require low false
positive rates (< 3%).

For details about the algorithm and citations please refer to the original
research paper, ["Cuckoo Filter: Better Than Bloom" by Bin Fan, Dave Andersen
and Michael Kaminsky](https://www.cs.cmu.edu/~dga/papers/cuckoo-conext2014.pdf).

## Interface
A cuckoo filter supports following operations:

*  `Insert(item)`: insert an item to the filter
*  `Lookup(item)`: return if item is already in the filter (may return false
   positive results like Bloom filters)
*  `Delete(item)`: delete the given item from the filter. Note that to use this
   method, it must be ensured that this item is in the filter (e.g., based on
   records on external storage); otherwise, a false item may be deleted.
*  `Size()`: return the total number of items currently in the filter

## Example Usage
```go
import "github.com/irfansharif/cfilter"

cf := cfilter.NewCFilter()

cf.Insert([]byte("bongiorno"))  // inserts 'bongiorno' to the filter
cf.Lookup([]byte("hola"))       // looks up 'hola' in the filter, may return false positive
cf.Size()                       // returns 1 (given only 'bongiorno' was added)
cf.Delete([]byte("bonjour"))    // tries deleting 'bonjour' from filter, may delete false item
```

## Authors
Irfan Sharif <irfanmahmoudsharif@gmail.com>
