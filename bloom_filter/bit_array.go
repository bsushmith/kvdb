package bloom_filter

type BitArray struct {
	bits []uint8
}

func NewBitArray(size uint64) *BitArray {
	return &BitArray{
		bits: make([]uint8, (size+7)/8), // +7 to round up to the nearest byte
	}
}

func (b *BitArray) Set(index uint64) {
	b.bits[index/8] |= 1 << (index % 8)
}

func (b *BitArray) Clear(index uint64) {
	b.bits[index/8] &= ^(1 << (index % 8))
}

func (b *BitArray) IsSet(index uint64) bool {
	return (b.bits[index/8] & (1 << (index % 8))) != 0
}
