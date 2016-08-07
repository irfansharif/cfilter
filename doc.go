/*
Copyright (c) 2016 Irfan Sharif
The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package cfilter is an implementation of the Cuckoo filter, a Bloom filter
// replacement for approximated set-membership queries. Cuckoo filters support
// adding and removing items dynamically while achieving even higher performance
// than Bloom filters.
//
// As documented in the original implementation:
//      Cuckoo filters provide the flexibility to add and remove items dynamically. A
//      cuckoo filter is based on cuckoo hashing (and therefore named as cuckoo
//      filter). It is essentially a cuckoo hash table storing each key's fingerprint.
//      Cuckoo hash tables can be highly compact, thus a cuckoo filter could use less
//      space than conventional Bloom filters, for applications that require low false
//      positive rates (< 3%).
//
// For details about the algorithm and citations please refer to the original
// research paper, "Cuckoo Filter: Better Than Bloom" by Bin Fan, Dave Andersen
// and Michael Kaminsky (https://www.cs.cmu.edu/~dga/papers/cuckoo-conext2014.pdf).
package cfilter
