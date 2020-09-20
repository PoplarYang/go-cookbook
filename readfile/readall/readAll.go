package readall

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// one
func ReadByReadall() []byte {
	file, err := os.Open("test.txt")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	// os.File.Read(), io.ReadFull() 和 io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但 ioutil.ReadAll() 会读取 reader (这个例子中是file)的每一个字节，然后把字节slice返回。
	// 读取文件到byte slice中
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// two
func ReadByRead() []byte {
	file, err := os.Open("test.txt")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	// buffer size is filesize
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	return buffer
}

// three
func ReadFile() []byte {
	// 读取文件到byte slice中
	// ReadFile(filename string)方法之所以速度快的原因就是先计算出file文件的size，在初始化对应size大小的buff，传入ReadRead(p []byte) 来读取字节流
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	return data
}