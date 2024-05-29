// Package bloom_filter implementation taken from github.com/bsushmith/filters
package bloom_filter

import (
	"github.com/cespare/xxhash"
	"github.com/spaolacci/murmur3"
	"math"
	"time"
)

type BloomFilter struct {
	bits      *BitArray
	size      uint64
	hashFuncs []func(string, uint64) uint64
}

// NewBloomFilter creates a new bloom filter with the given capacity and false positive probability
// The capacity is the expected number of elements to be added to the filter
// The false positive probability is the probability of the filter returning true for an element that was not added
// Ref: https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
func NewBloomFilter(capacity uint64, falsePositiveProbability float64) *BloomFilter {
	bitArraySize := -1 * (float64(capacity) * math.Log(falsePositiveProbability)) / (math.Pow(math.Log(2), 2))
	numHashFunctions := int((bitArraySize / float64(capacity)) * math.Log(2))
	bloomFilter := BloomFilter{
		bits:      NewBitArray(uint64(bitArraySize)),
		size:      uint64(bitArraySize),
		hashFuncs: getHashFunctions(numHashFunctions, uint64(bitArraySize)),
	}
	return &bloomFilter
}

func getHashFunctions(numHashFunctions int, size uint64) []func(string, uint64) uint64 {
	hashFn1 := murmur3.New64WithSeed(uint32(time.Now().Unix()))
	hashFn2 := xxhash.New()
	hashFunctions := make([]func(string, uint64) uint64, numHashFunctions)
	//hash functions of the form ah1 + bh2
	for i := 0; i < numHashFunctions; i++ {
		hashFunctions[i] = func(data string, size uint64) uint64 {
			hashFn1.Write([]byte(data))
			hashFn2.Write([]byte(data))
			hash1 := (i * int(hashFn1.Sum64())) % int(size)
			hash2 := (i * i * int(hashFn2.Sum64())) % int(size)
			hashFn1.Reset()
			hashFn2.Reset()
			return uint64(hash1+hash2) % size
		}
	}

	return hashFunctions
}

func (b *BloomFilter) applyHash(data string, size uint64) []uint64 {
	hashes := make([]uint64, len(b.hashFuncs))
	for i, hashFn := range b.hashFuncs {
		hashes[i] = hashFn(data, size)
	}
	return hashes
}

func (b *BloomFilter) Add(element string) {
	hashes := b.applyHash(element, b.size)
	for _, hash := range hashes {
		b.bits.Set(hash)
	}
}

func (b *BloomFilter) Exists(element string) bool {
	hashes := b.applyHash(element, b.size)
	exists := true
	for _, hash := range hashes {
		exists = exists && b.bits.IsSet(hash)
		if !exists {
			return false
		}
	}
	return true
}

func (b *BloomFilter) ClearAll() {
	b.bits = NewBitArray(b.size)
}
