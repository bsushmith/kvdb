package sstable

import "github.com/bsushmith/kvdb/bloom_filter"

type SSTEntry struct {
	key     string
	value   []byte
	deleted bool
}

type SSTFile struct {
	entries []SSTEntry
}

type SSTManager struct {
	bloom *bloom_filter.BloomFilter
	files map[string]SSTFile
}

func NewSSTManager() *SSTManager {
	return &SSTManager{
		bloom: bloom_filter.NewBloomFilter(1000),
		files: make(map[string]SSTFile),
	}
}

func (s *SSTManager) AddFile(fileName string, entries []SSTEntry) {
	s.files[fileName] = SSTFile{
		entries: entries,
	}
	for _, entry := range entries {
		s.bloom.Add(entry.key)
	}
}

func (s *SSTManager) Get(key string) ([]byte, bool) {
	if !s.bloom.Exists(key) {
		return nil, false
	}
	for _, file := range s.files {
		for _, entry := range file.entries {
			if entry.key == key && !entry.deleted {
				return entry.value, true
			}
		}
	}
	return nil, false
}
