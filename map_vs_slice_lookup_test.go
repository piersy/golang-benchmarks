package main

import (
	"crypto/rand"

	mr "math/rand"
	"strconv"
	"testing"
)

func BenchmarkMapVsSliceLookup(b *testing.B) {
	maxLen := 10000
	values := make(map[string]string)
	r := make([]byte, 20)
	for i := 0; i < maxLen; i++ {

		_, err := rand.Read(r)
		if err != nil {
			b.Fatal(err)
		}
		v := string(r)
		values[v] = v
	}

	// Double the test size each iteration
	for i := 8; i < maxLen; i *= 2 {

		// Set up a map and a slice for this test
		m := make(map[string]string, i)
		s := make([]string, i)

		// Fill the map and slice, the iteration order of values is randomized.
		j := 0
		for k, _ := range values {
			m[k] = k
			s[j] = k
			j++
			if j == i {
				break
			}
		}

		// Choose a random value.
		key := s[mr.Intn(len(s))]

		b.Run("Map__"+strconv.Itoa(i), func(b *testing.B) { mapLookup(b, m, key) })
		b.Run("Slice"+strconv.Itoa(i), func(b *testing.B) { sliceLookup(b, s, key) })
	}

}

func mapLookup(b *testing.B, m map[string]string, key string) {
	for i := 0; i < b.N; i++ {
		if v, ok := m[key]; !ok {
			b.Fatalf("failed to lookup key %q, instead got %q", key, v)
		}
	}
}

func sliceLookup(b *testing.B, s []string, key string) {
	for i := 0; i < b.N; i++ {
		for _, v := range s {
			if key == v {
				goto FOUND
			}
		}
		b.Fatalf("failed to lookup key %q", key)
	FOUND:
	}
}
