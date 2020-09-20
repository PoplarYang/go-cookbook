package readline

import "testing"

func TestReadLineByScanner(t *testing.T) {
	ReadLineByScanner()
}

func TestReadLineByReadBytes(t *testing.T) {
	ReadLineByReadBytes()
}

func TestReadLineByReadString(t *testing.T) {
	ReadLineByReadString()
}

func BenchmarkReadLineByScanner(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadLineByScanner()
	}
}

func BenchmarkReadLineByReadBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadLineByReadBytes()
	}
}

func BenchmarkReadLineByReadString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadLineByReadString()
	}
}