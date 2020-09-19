package filesize

import (
	"io"
	"os"
)

func GetFileSizeByRead() (fs int) {
	file, err := os.Open("abc.txt")
	if err == nil {
		sum := 0
		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			sum += n
			if err == io.EOF {
				break
			}
		}
	}
	return
}

func GetFileSizeByIoutilReadFile() (fs int64) {
	file,err := os.Open("abc.txt")

	if err == nil {
		fi, _ := file.Stat()
		fs = fi.Size()
	}
	return
}

func GetFileSizeByOsStat() (fs int64) {
	fi, err := os.Stat("abc.txt")
	if err == nil {
		fs = fi.Size()
	}
	return
}