package search

import (
	"path"
	"testing"
)

var dir = path.Join("testdata", "artifacts")
var answer = "error ERC721UnsupportedToken(address token)"
var selector = [...]byte{0xc7, 0xd8, 0x37, 0xc6}

func TestSearchDirectory(t *testing.T) {
	res, err := SearchDirectory(dir, selector)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if res.String() != answer {
		t.Errorf("expected %s, got %s", answer, res)
	}
}

func BenchmarkSearchDirectory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchDirectory(dir, selector)
	}
}
