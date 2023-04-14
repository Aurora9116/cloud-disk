package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

// 分片大小
const chunkSize = 100 * 1024 * 1024 // 100M

// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("")
	if err != nil {
		t.Fatal(err)
	}
	// 分片的个数
	chunkNum := math.Ceil(float64(fileInfo.Size() / chunkSize)) // 向上取整
	myFile, err := os.OpenFile("", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		/*
			switch whence {
			case 0:
				w = FILE_BEGIN 从文件开始读
			case 1:
				w = FILE_CURRENT
			case 2:
				w = FILE_END
			}
		*/
		myFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
}

// 分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat("")
	if err != nil {
		t.Fatal(err)
	}
	// 分片的个数
	chunkNum := math.Ceil(float64(fileInfo.Size() / chunkSize)) // 向上取整
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 文件一致性校验
func TestCheck(t *testing.T) {
	f, err := os.OpenFile("", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := os.OpenFile("", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := io.ReadAll(f2)
	if err != nil {
		t.Fatal(err)
	}
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
	fmt.Println(s1 == s2)
}
