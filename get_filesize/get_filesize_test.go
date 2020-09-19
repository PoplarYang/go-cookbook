package filesize

import "testing"

func BenchmarkGetFileSizeByRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetFileSizeByRead()
	}
}

func BenchmarkGetFileSizeByIoutilReadFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetFileSizeByIoutilReadFile()
	}
}

func BenchmarkGetFileSizeByOsStat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetFileSizeByOsStat()
	}
}