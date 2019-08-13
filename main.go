package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"os"
	"runtime"
	"strings"
)

var (
	genHasher func() hash.Hash
	coreNum   = runtime.NumCPU() * 3 / 4

	hashAlgorithms = map[string]func() hash.Hash{
		"md5": func() hash.Hash {
			return md5.New()
		},
		"sha1": func() hash.Hash {
			return sha1.New()
		},
		"sha256": func() hash.Hash {
			return sha256.New()
		},
		"sha512": func() hash.Hash {
			return sha512.New()
		},
	}
)

func hashFunc(target string, prefix []byte, suffix []byte, result chan<- string, length int, id int) {
	var msg []byte
	var src []byte
	fn := genStrFunc(length, id)

	for {
		src = fn()
		msg = append(prefix, src...)
		msg = append(msg, suffix...)
		hasher := genHasher()
		hasher.Write(msg)

		if fmt.Sprintf("%x", string(hasher.Sum(nil)))[:length] == target {
			result <- string(src)
			break
		}
	}
}

var elements = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func genStrFunc(length int, id int) func() []byte {
	ptrs := []int{}
	for i := 0; i < length; i++ {
		ptrs = append(ptrs, 0)
	}
	ptrs[0] = (62 / coreNum) * id

	return func() []byte {
		var str []byte
		for i := 0; i < len(ptrs); i++ {
			str = append(str, elements[ptrs[i]])
		}

		for i := len(ptrs) - 1; i > -1; i-- {
			if ptrs[i] < 61 {
				ptrs[i]++
				break
			} else {
				ptrs[i] = 0
			}
		}
		return str
	}
}

func parseCmd() (target string, prefix string, suffix string, err error) {
	if len(os.Args) < 3 {
		return "", "", "", fmt.Errorf("no enough parameters")
	}
	var ok bool
	genHasher, ok = hashAlgorithms[os.Args[1]]
	if !ok {
		return "", "", "", fmt.Errorf("unrecognized hash algorithm")
	}

	target = strings.ToLower(os.Args[2])

	if len(os.Args) > 3 {
		prefix = os.Args[3]
	}

	if len(os.Args) > 4 {
		suffix = os.Args[4]
	}
	return
}

func banner() {
	fmt.Println("Hash Proofer (md5, sha1, sha256, sha512)\n\n" +
		"Usage:\n" +
		"proof algorithm target [prefix] [suffix]\n")
}

func brute(target string, prefix []byte, suffix []byte, length int) chan string {
	result := make(chan string)
	for i := 0; i < coreNum; i++ {
		go hashFunc(target, prefix, suffix, result, length, i)
	}
	return result
}

func main() {
	target, prefix, suffix, err := parseCmd()
	if err != nil {
		banner()
		fmt.Println(err.Error())
		return
	}

	result := brute(target, []byte(prefix), []byte(suffix), len(target))
	fmt.Println(<-result)
}
