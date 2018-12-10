package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"flag"
	"github.com/cespare/xxhash"
	"github.com/xrkhill/hashmod/hashmod"
	"hash"
	"hash/crc64"
	"hash/fnv"
	"io"
	"log"
	"os"
)

func getLinesFromStdin(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println("reading standard input:", err)
	}

	return list
}

func main() {
	hashTypePtr := flag.String("hashalg", "md5", "[xxhash64|fnv|crc64|md5|sha1|sha512]: hash algorithim")
	inputTypePtr := flag.String("input", "hostname", "[hostname|stdin]: hash hostname of system or import list from stdin")
	bucketSizePtr := flag.Int("buckets", 2, "how many buckets to divide hosts into")

	flag.Parse()

	hashers := map[string]hash.Hash{
		"xxhash64": xxhash.New(),
		"fnv":      fnv.New128a(),
		"crc64":    crc64.New(crc64.MakeTable(crc64.ISO)),
		"md5":      md5.New(),
		"sha1":     sha1.New(),
		"sha512":   sha512.New512_256(),
	}

	hasher, present := hashers[*hashTypePtr]

	if !present {
		hasher = hashers["xxhash64"]
	}

	enabledCount, enabled := 0, ""
	buckets := uint64(*bucketSizePtr)
	h := hashmod.New(buckets, hasher)

	if *inputTypePtr == "stdin" {
		lines := getLinesFromStdin(os.Stdin)

		for i, line := range lines {
			if h.IsHostEnabled(line, buckets) == true {
				enabled = "true"
				enabledCount++
			} else {
				enabled = "false"
			}

			log.Println("index:", i)
			log.Println("line:", line)
			log.Println("enabled:", enabled)
			log.Println("----------")
		}

		totalCount := len(lines)
		percent := float64(enabledCount) / float64(totalCount) * 100

		log.Println("buckets:", buckets)
		log.Println("enabledCount:", enabledCount)
		log.Println("totalCount:", totalCount)
		log.Printf("percent: %.2f%%\n", percent)
	} else {
		hostname, err := os.Hostname()

		if err != nil {
			log.Fatalln("Unexpected error when looking up hostname: ", err)
		}

		if h.IsHostEnabled(hostname, buckets) {
			enabled = "true"
		} else {
			enabled = "false"
		}

		log.Print(enabled)
	}
}
