// Package bloom_filter implementation taken from github.com/bsushmith/filters
package bloom_filter

import (
	"github.com/cespare/xxhash"
	"github.com/spaolacci/murmur3"
	"hash"
	"time"
)

type BloomFilter struct {
	bits    *BitArray
	size    uint64
	hashFn1 hash.Hash64
	hashFn2 hash.Hash64
}

func NewBloomFilter(size uint64) *BloomFilter {
	return &BloomFilter{
		bits:    NewBitArray(size),
		size:    size,
		hashFn1: murmur3.New64WithSeed(uint32(time.Now().Unix())),
		hashFn2: xxhash.New(),
	}
}

func (b *BloomFilter) applyHash(data string, size uint64) (uint64, uint64) {
	b.hashFn1.Write([]byte(data))
	h1 := b.hashFn1.Sum64() % size
	b.hashFn1.Reset()

	b.hashFn2.Write([]byte(data))
	h2 := b.hashFn2.Sum64() % size
	b.hashFn2.Reset()

	return h1, h2
}

func (b *BloomFilter) Add(element string) {
	h1, h2 := b.applyHash(element, b.size)
	b.bits.Set(h1)
	b.bits.Set(h2)
}

func (b *BloomFilter) Exists(element string) bool {
	h1, h2 := b.applyHash(element, b.size)
	return b.bits.IsSet(h1) && b.bits.IsSet(h2)
}

func (b *BloomFilter) ClearAll() {
	b.bits = NewBitArray(b.size)
}
