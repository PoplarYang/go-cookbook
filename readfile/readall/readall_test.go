package readall

import "testing"

func TestReadByReadall(t *testing.T) {
	ReadByReadall()
}

func TestReadByRead(t *testing.T) {
	ReadByRead()
}

func TestReadFile(t *testing.T) {
	ReadFile()
}

func BenchmarkReadByReadall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadByReadall()
	}
}

func BenchmarkReadByRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadByRead()
	}
}

func BenchmarkReadFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadFile()
	}
}