package readbyte

import (
	"io"
	"log"
	"os"
)

func ByteRead(length int) (numBytesRead int, byteSliceBuffer []byte) {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 返回io.EOF的error的触发条件
	// 读到空文件
	// 读到文件末尾后继续读取，如果文件有内容，即使文件长度不够，也不会报错，[]byte中有默认零值，直接结束。
	byteSliceBuffer = make([]byte, length)
	numBytesRead, err = file.Read(byteSliceBuffer)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ByteReadAtLeast(minBytes int) (numBytesRead int, byteSliceBuffer []byte) {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	byteSliceBuffer = make([]byte, minBytes)
	// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误(unexpected EOF)，但会把已读的文件保留
	numBytesRead, err = io.ReadAtLeast(file, byteSliceBuffer, minBytes)
	if err != nil {
		log.Fatal(err)
	}
	return
}