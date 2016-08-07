package cfilter

import (
	"bufio"
	"os"
	"testing"
)

func TestMultipleInsertions(t *testing.T) {
	cf := New()

	fd, err := os.Open("/usr/share/dict/words")
	if err != nil {
		t.Errorf(err.Error())
	}

	scanner := bufio.NewScanner(fd)
	var words [][]byte
	var wordCount uint
	for scanner.Scan() {
		word := []byte(scanner.Text())

		if !cf.Lookup(word) && cf.Insert(word) {
			wordCount++
		}
		words = append(words, word)
	}

	size := cf.Size()
	if size != wordCount {
		t.Errorf("Expected word count = %d, not %d", wordCount, size)
	}

	for _, word := range words {
		cf.Delete(word)
	}

	size = cf.Size()
	if size != 0 {
		t.Errorf("Expected word count = 0, not %d", size)
	}
}

func TestBasicInsertion(t *testing.T) {
	cf := New()
	if !cf.Insert([]byte("bongiorno")) {
		t.Errorf("Wasn't able to insert very first word, 'bongiorno'")
	}

	size := cf.Size()
	if size != 1 {
		t.Errorf("Expected size after insertion to be 1, not %d", size)
	}

	if !cf.Lookup([]byte("bongiorno")) {
		t.Errorf("Expected to find 'bongiorno' in filter set membership query")
	}

	if !cf.Delete([]byte("bongiorno")) {
		t.Errorf("Expected to be able to delete 'bongiorno' in filter")
	}

	if cf.Lookup([]byte("bongiorno")) {
		t.Errorf("Did not expect to find 'bongiorno' in filter after deletion")
	}

	size = cf.Size()
	if size != 0 {
		t.Errorf("Expected size after deletion to be 0, not %d", size)
	}
}

func TestInitialization(t *testing.T) {
	cf := New()
	size := cf.Size()
	if size != 0 {
		t.Errorf("Expected initial size to be 0, not %d", size)
	}
}

func BenchmarkInsertionAndDeletion(b *testing.B) {
	cf := New()
	for n := 0; n < b.N; n++ {
		cf.Insert([]byte("bongiorno"))
		cf.Delete([]byte("bongiorno"))
	}
}
