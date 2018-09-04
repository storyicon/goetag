/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package etag is used to calculate the hash value of the file.
The algorithm is based on the qetag of Qi Niuyun.
Qetag address is
    https://github.com/qiniu/qetag
This package extends two export functions based on Qetag,
named GetEtagByString, GetEtagByBytes
And re-implemented GetEtagByPath
*/

package goetag

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"os"
	"runtime"
)

const (
	BlockSize int64 = 4194304
)

type Reader interface {
	Read([]byte) (int, error)
	ReadAt([]byte, int64) (int, error)
}

//GetEtagByString calculates the hash value from the string
func GetEtagByString(str string) (string, error) {
	return GetEtagByBytes([]byte(str))
}

//GetEtagByString calculates the hash value from file path
func GetEtagByPath(filepath string) (string, error) {
	var stat os.FileInfo

	file, err := os.Open(filepath)
	defer file.Close()
	if err == nil {
		if stat, err = file.Stat(); err == nil {
			size := stat.Size()
			return getEtagByReader(file, size), nil
		}
	}
	return "", err
}

//GetEtagByString calculates the hash value from the byte array
func GetEtagByBytes(content []byte) (string, error) {
	reader := bytes.NewReader(content)
	size := reader.Size()
	return getEtagByReader(reader, size), nil
}

func getEtagByReader(reader Reader, size int64) string {
	buffer := make([]byte, 0, 21)
	if count := blockCount(size); count > 1 {
		buffer = getHugeEtag(reader, count)
	} else {
		buffer = getTinyEtag(reader, buffer)
	}
	return base64.URLEncoding.EncodeToString(buffer)
}

func getTinyEtag(reader Reader, buffer []byte) []byte {
	buffer = append(buffer, 0x16)
	buffer = getSha1ByReader(buffer, reader)
	return buffer
}

func doEtagWork(reader Reader, offsetChan <-chan int, conseqChan chan<- map[int][]byte) {
	for offset := range offsetChan {
		data := io.NewSectionReader(reader, int64(offset)*BlockSize, BlockSize)
		sha1 := getSha1ByReader(nil, data)
		conseqChan <- map[int][]byte{
			offset: sha1,
		}
	}
}

func getHugeEtag(reader Reader, count int64) []byte {
	conseqChan := make(chan map[int][]byte, count)
	offsetChan := make(chan int, count)

	for i := 1; i <= runtime.NumCPU(); i++ {
		go doEtagWork(reader, offsetChan, conseqChan)
	}

	for offset := 0; offset < int(count); offset++ {
		offsetChan <- offset
	}

	close(offsetChan)

	return getSha1ByConseqChan(conseqChan, count)
}

func getSha1ByConseqChan(conseqChan chan map[int][]byte, count int64) (conseq []byte) {
	sha1Map := make(map[int][]byte, 0)
	for i := 0; i < int(count); i++ {
		eachChan := <-conseqChan
		for k, v := range eachChan {
			sha1Map[k] = v
		}
	}
	blockSha1 := make([]byte, 0, count*20)
	for i := 0; int64(i) < count; i++ {
		blockSha1 = append(blockSha1, sha1Map[i]...)
	}
	conseq = make([]byte, 0, 21)
	conseq = append(conseq, 0x96)
	conseq = getSha1ByReader(conseq, bytes.NewReader(blockSha1))
	return
}

func getSha1ByReader(buffer []byte, reader Reader) []byte {
	hash := sha1.New()
	io.Copy(hash, reader)
	return hash.Sum(buffer)
}

func blockCount(size int64) int64 {
	if size > BlockSize {
		count := size / BlockSize
		if size&BlockSize == 0 {
			return count
		}
		return count + 1
	}
	return 1
}
