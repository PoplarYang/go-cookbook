package readline

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadLineByScanner() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// 缺省的分隔函数是bufio.ScanLines
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		scanner.Text()
	}
}

func ReadLineByReadBytes() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	for {
		//遇到\n结束读取
		_, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		//log.Print(string(b))
	}
}

func ReadLineByReadString() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	for {
		//遇到\n结束读取
		_, errR := buf.ReadString('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		//log.Print(s)
	}
}